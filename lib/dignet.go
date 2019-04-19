package lib

import (
	"fmt"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
	"math"
	"net"
	"strconv"
	"strings"
)

type (
	// AWSCredentials contains information to allow accessing AWS resources
	// Keys have precedence over Profile and Region overrides that specified
	// in the profile
	AWSCredentials struct {
		Profile         string
		Region          string
		AccessKeyID     string
		SecretAccessKey string
	}

	// Config contains the configuration that needs to be supplied to dignet
	Config struct {
		VpcID       string
		SubnetSize  uint64
		Credentials *AWSCredentials
	}
)

func createSession(c *AWSCredentials) (*session.Session, error) {
	awsConf := aws.Config{}
	if c.Region != "" {
		awsConf.Region = aws.String(c.Region)
	}

	if c.AccessKeyID != "" {
		awsConf.Credentials = credentials.NewStaticCredentials(c.AccessKeyID, c.SecretAccessKey, "")
		return session.Must(session.NewSession(&awsConf)), nil
	}

	if c.Profile != "" {
		return session.Must(
			session.NewSessionWithOptions(session.Options{Config: awsConf, Profile: c.Profile})), nil
	}

	return nil, errors.New("No AWS credentials provided")
}

func getVPCCidr(sess *session.Session, c *Config) (string, error) {
	result, err := ec2.New(sess).DescribeVpcs(&ec2.DescribeVpcsInput{VpcIds: []*string{aws.String(c.VpcID)}})
	if err != nil {
		return "", err
	}
	return *result.Vpcs[0].CidrBlock, nil
}

func getExistingSubnetsFromVPC(sess *session.Session, c *Config) (*ec2.DescribeSubnetsOutput, error) {
	input := &ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			{Name: aws.String("vpc-id"), Values: []*string{aws.String(c.VpcID)}},
		},
	}
	return ec2.New(sess).DescribeSubnets(input)
}

func extractCidrs(subnets *ec2.DescribeSubnetsOutput) []*net.IPNet {
	cidrBlocks := make([]*net.IPNet, len(subnets.Subnets))

	for i, s := range subnets.Subnets {
		_, subnetCidr, _ := net.ParseCIDR(*s.CidrBlock)
		cidrBlocks[i] = subnetCidr
	}

	return cidrBlocks
}

// FindAvailableSubnets queries the target VPC for existing subnets and returns
// CIDRs of subnets of the specified size that are available for use
func FindAvailableSubnets(c *Config) ([]*net.IPNet, error) {
	var result []*net.IPNet

	sess, err := createSession(c.Credentials)
	if err != nil {
		return nil, err
	}

	if c.VpcID == "" {
		return nil, errors.New("VPC ID not provided")
	}
	if c.SubnetSize < 1 || float64(c.SubnetSize) > math.Pow(float64(2), float64(32)) {
		return nil, errors.New(fmt.Sprintf("Invalid subnet size %d", c.SubnetSize))
	}

	vpcCidr, err := getVPCCidr(sess, c)
	if err != nil {
		return nil, err
	}

	subnets, err := getExistingSubnetsFromVPC(sess, c)
	if err != nil {
		return nil, err
	}

	existingSubnetCidrs := extractCidrs(subnets)
	frozenBits := 32 - uint64(math.Ceil(math.Log2(float64(c.SubnetSize))))
	vpcFrozenBits, _ := strconv.Atoi(strings.Split(vpcCidr, "/")[1])
	vpcSize := uint64(math.Pow(2, float64(vpcFrozenBits)))
	_, parsedVpcCidr, _ := net.ParseCIDR(vpcCidr)

	// This validation guarantees that frozenBits > vpcFrozenBits
	if c.SubnetSize > vpcSize {
		return nil, errors.New(
			fmt.Sprintf("Subnet size cannot be greater than VPC size (%d | %s)", vpcSize, vpcCidr))
	}

	newBits := frozenBits - uint64(vpcFrozenBits)
	numOfSubnets := math.Pow(2, float64(frozenBits-uint64(vpcFrozenBits)))

	for netNum := 0; netNum < int(numOfSubnets); netNum++ {
		candidateSubnetCidr, err := cidr.Subnet(parsedVpcCidr, int(newBits), netNum)
		if err != nil {
			return nil, err
		}

		if cidr.VerifyNoOverlap(append(existingSubnetCidrs, candidateSubnetCidr), parsedVpcCidr) == nil {
			result = append(result, candidateSubnetCidr)
		}
	}

	return result, nil
}
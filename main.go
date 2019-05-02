package main

import (
	"fmt"
	dignet "github.com/duaraghav8/dignet/lib"
	"github.com/duaraghav8/dignet/version"
	"github.com/urfave/cli"
	"os"
	"strings"
)

var cmdListAvailableSubnets = cli.Command{
	Name: "list-available-subnets",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "vpc-id",
			Usage: "ID of the target VPC",
		},
		cli.Uint64Flag{
			Name:  "subnet-size",
			Usage: "Desired subnet size",
			Value: 128,
		},
	},
	Action: listAvailableSubnets,
	Usage:  "List CIDRs of available IPv4 subnets of given size in target VPC",
}

func prettyPrint(res *dignet.FindAvailableSubnetsResponse) {
	sep := strings.Repeat("=", 25)

	fmt.Println()
	fmt.Printf("Region:   %s\n", res.Region)
	fmt.Printf("VPC ID:   %s\n", res.VpcID)
	fmt.Printf("VPC CIDR: %s\n", res.VpcCidr)

	fmt.Printf("%s\n%d Available Subnets\n%s\n", sep, len(res.AvailableSubnets), sep)
	for _, cidr := range res.AvailableSubnets {
		fmt.Println(cidr.String())
	}
	fmt.Println()
}

func listAvailableSubnets(c *cli.Context) error {
	config := &dignet.Config{
		VpcID: c.String("vpc-id"),
		Credentials: &dignet.AWSCredentials{
			Profile:         c.GlobalString("profile"),
			Region:          c.GlobalString("region"),
			AccessKeyID:     c.GlobalString("access-key-id"),
			SecretAccessKey: c.GlobalString("secret-access-key"),
		},
		SubnetSize: c.Uint64("subnet-size"),
	}

	res, err := dignet.FindAvailableSubnets(config)
	if err != nil {
		return err
	}

	prettyPrint(res)
	return nil
}

func main() {

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(version.HumanVersion)
	}

	app := cli.NewApp()
	app.Name = version.Name
	app.Version = version.Version
	app.Usage = version.Description
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{cmdListAvailableSubnets}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "profile",
			Usage:  "The profile to use from AWS credentials file",
			EnvVar: "AWS_PROFILE",
		},
		cli.StringFlag{
			Name:   "region",
			Usage:  "AWS Region",
			EnvVar: "AWS_REGION",
			Value:  "us-east-1",
		},
		cli.StringFlag{
			Name:   "access-key-id",
			Usage:  "AWS Access Key ID",
			EnvVar: "AWS_ACCESS_KEY_ID",
		},
		cli.StringFlag{
			Name:   "secret-access-key",
			Usage:  "AWS Secret Access Key",
			EnvVar: "AWS_SECRET_ACCESS_KEY",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s failed to start: %s\n", version.Name, err.Error())
		os.Exit(1)
	}
}

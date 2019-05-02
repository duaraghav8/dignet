package main

import (
	"fmt"
	dignet "github.com/duaraghav8/dignet/lib"
	"log"
	"os"
)

func prettyPrint(res *dignet.FindAvailableSubnetsResponse) {
	fmt.Printf("VPC CIDR: %s\n", res.VpcCidr)
	fmt.Println("Available Subnets:")
	for _, cidr := range res.AvailableSubnets {
		fmt.Println(cidr)
	}
}

func main() {
	config := &dignet.Config{
		SubnetSize: 2000, // x.x.x.x/21
		VpcID:      "vpc-7c872910",
		Credentials: &dignet.AWSCredentials{
			Region:          "us-east-1",
			AccessKeyID:     "abcdefghijklmnop",
			SecretAccessKey: "xxxxxxxxxxxxxxxxx",
		},
	}

	res, err := dignet.FindAvailableSubnets(config)
	if err != nil {
		log.Fatal(err)
	}

	prettyPrint(res)
}

package main

import (
	"fmt"
	"github.com/duaraghav8/netter/version"
	"github.com/urfave/cli"
	"os"
)

func listAvailableSubnets(c *cli.Context) error {
	return nil
}

var cmdListAvailableSubnets = cli.Command{
	Name:  "list-available-subnets",
	Usage: "List available subnet CIDRs of given size in target VPC",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "vpc-id",
			Usage: "ID of the target VPC",
		},
		cli.Uint64Flag{
			Name:  "subnet-size",
			Usage: "Desired subnet size",
		},
	},
	Action: listAvailableSubnets,
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
		fmt.Fprintf(os.Stderr, "Failed to start %s: %s\n", version.Name, err.Error())
		os.Exit(1)
	}
}

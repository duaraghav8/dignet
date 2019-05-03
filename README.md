# Dignet
Dignet is a Go library and a CLI tool to find available IPv4 subnets in Amazon [Virtual Private Clouds](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).

It is useful when you want to determine one or more subnet CIDRs available for use inside a VPC. Specify the size of the subnet(s) required and `dignet` outputs all available subnet CIDRs that do not overlap with any of the existing ones in the target VPC.

## CLI usage
To use a pre-compiled binary, download the package appropriate for your platform from [Releases](https://github.com/duaraghav8/dignet/releases). Unzip the binary into any directory. The recommended approach is to put it in `/usr/local/bin` or any directory in your `PATH`.

Use the `help` command to see the complete list of commands and options available.
```bash
$ dignet help
NAME:
   Dignet - Find available Subnets in Amazon VPC

USAGE:
   dignet [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     list-available-subnets  List CIDRs of available IPv4 subnets of given size in target VPC
     help, h                 Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --profile value            The profile to use from AWS credentials file [$AWS_PROFILE]
   --region value             AWS Region [$AWS_REGION]
   --access-key-id value      AWS Access Key ID [$AWS_ACCESS_KEY_ID]
   --secret-access-key value  AWS Secret Access Key [$AWS_SECRET_ACCESS_KEY]
   --help, -h                 show help
   --version, -v              print the version
```

If you've configured a [Named Profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) for your AWS account, you can use it with dignet. Alternatively, you can provide the region and API credentials.

Note that the region specified via the commandline overrides the region specified in profile.

Use `help <command>` to get help on a particular command.
```bash
$ dignet help list-available-subnets
NAME:
   dignet list-available-subnets - List CIDRs of available IPv4 subnets of given size in target VPC

USAGE:
   dignet list-available-subnets [command options] [arguments...]

OPTIONS:
   --vpc-id value       ID of the target VPC
   --subnet-size value  Desired subnet size (default: 128)
```

To get the list of Subnets available, use the `list-available-subnets` command.
```bash
# Below command shows the list of subnets available that comprise of 4000 IPv4 addresses.
# Since the smallest subnet with at least 4K IPs is /20, dignet looks for all /20 subnets.

$ dignet -profile production list-available-subnets -vpc-id vpc-07h38a995398gg203 -subnet-size 4000

Region:   us-east-1
VPC ID:   vpc-02h68a099382ff017
VPC CIDR: 10.108.0.0/16
===================================
7 Available Subnet(s) of size 4096
===================================
10.108.64.0/20
10.108.80.0/20
10.108.96.0/20
10.108.112.0/20
10.108.128.0/20
...
```

## Library usage
Download dignet using `go get github.com/duaraghav8/dignet`.

See [examples/](https://github.com/duaraghav8/dignet/tree/master/examples) on how to use the library in your application.

## Credentials
The AWS credentials you supply must have the appropriate IAM permissions to [read VPC and subnet information](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_IAM.html#readonlyvpciam).

## License
MIT

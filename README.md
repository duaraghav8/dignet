# Dignet
Dignet is a Go library and a CLI tool to find available IPv4 subnets in Amazon [Virtual Private Clouds](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).

It is useful when you want to determine one or more subnet CIDRs available for use inside a VPC. The user specifies the size of the subnet(s) required and `dignet` outputs all available subnet CIDRs that do not overlap with any of the existing ones in the target VPC.

## CLI usage
Download the latest compiled binary for your platform from the [Releases](https://github.com/duaraghav8/dignet/releases) page.

## Library usage
Download dignet using `go get https://github.com/duaraghav8/dignet`. Then import the library into your project and use it as described below:

```go

```

## Credential Permissions
The AWS credentials provided to Dignet should at least have the following IAM Permissions:

```json

```
# Netter
Netter is a Go library and a CLI tool to find available subnets in Amazon [Virtual Private Clouds](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).

It is useful when a human or a Go application wants to determine one or more subnet CIDRs available for use inside a VPC. The user specifies the size of the subnet(s) required and `netter` outputs all available subnet CIDRs that do not overlap with any of the existing ones in the target VPC.

## CLI usage
Download the latest compiled binary for your platform from the [Releases](https://github.com/duaraghav8/netter/releases) page.

## Library usage
Download netter using `go get https://github.com/duaraghav8/netter`. Then import the library into your project and use it as described below:

```go

```

## Credential Permissions
The AWS credentials provided to Netter should at least have the following IAM Permissions:

```json

```
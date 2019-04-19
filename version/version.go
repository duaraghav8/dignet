package version

import "fmt"

const (
	Version     = "0.1.0"
	Description = "Find available Subnets in Amazon VPC"
)

var (
	Name         = "Netter"
	HumanVersion = fmt.Sprintf("%s version %s", Name, Version)
)

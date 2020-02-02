package main

import (
	"github.com/raskad/aws-resource-list/aws"
)

var gitTag string
var gitCommit string

func main() {
	aws.Start(gitTag, gitCommit)
}

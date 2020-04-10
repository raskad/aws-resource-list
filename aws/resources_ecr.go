package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func getEcr(config aws.Config) (resources awsResourceMap) {
	client := ecr.New(config)

	ecrRepositoryNames := getEcrRepositoryNames(client)

	resources = awsResourceMap{
		ecrRepository: ecrRepositoryNames,
	}
	return
}

func getEcrRepositoryNames(client *ecr.Client) (resources []string) {
	req := client.DescribeRepositoriesRequest(&ecr.DescribeRepositoriesInput{})
	p := ecr.NewDescribeRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Repositories {
			resources = append(resources, *resource.RepositoryName)
		}
	}
	return
}

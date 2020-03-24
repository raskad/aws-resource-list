package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func getEcr(config aws.Config) (resources resourceMap) {
	client := ecr.New(config)

	ecrRepositoryNames := getEcrRepositoryNames(client)

	resources = resourceMap{
		ecrRepository: ecrRepositoryNames,
	}
	return
}

func getEcrRepositoryNames(client *ecr.Client) (resources []string) {
	req := client.DescribeRepositoriesRequest(&ecr.DescribeRepositoriesInput{})
	p := ecr.NewDescribeRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Repositories {
			resources = append(resources, *resource.RepositoryName)
		}
	}
	return
}

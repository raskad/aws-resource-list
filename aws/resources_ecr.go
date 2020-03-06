package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func getEcr(config aws.Config) (resources resourceMap) {
	client := ecr.New(config)
	resources = reduce(
		getEcrRepository(client).unwrap(ecrRepository),
	)
	return
}

func getEcrRepository(client *ecr.Client) (r resourceSliceError) {
	req := client.DescribeRepositoriesRequest(&ecr.DescribeRepositoriesInput{})
	p := ecr.NewDescribeRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Repositories {
			r.resources = append(r.resources, *resource.RepositoryName)
		}
	}
	r.err = p.Err()
	return
}

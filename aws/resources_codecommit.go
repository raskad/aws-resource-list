package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

func getCodeCommit(config aws.Config) (resources resourceMap) {
	client := codecommit.New(config)

	codeCommitRepositoryIDs := getCodeCommitRepositoryIDs(client)

	resources = resourceMap{
		codeCommitRepository: codeCommitRepositoryIDs,
	}
	return
}

func getCodeCommitRepositoryIDs(client *codecommit.Client) (resources []string) {
	req := client.ListRepositoriesRequest(&codecommit.ListRepositoriesInput{})
	p := codecommit.NewListRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Repositories {
			resources = append(resources, *resource.RepositoryId)
		}
	}
	return
}

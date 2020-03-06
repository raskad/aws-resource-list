package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

func getCodeCommit(config aws.Config) (resources resourceMap) {
	client := codecommit.New(config)
	resources = reduce(
		getCodeCommitRepository(client).unwrap(codeCommitRepository),
	)
	return
}

func getCodeCommitRepository(client *codecommit.Client) (r resourceSliceError) {
	req := client.ListRepositoriesRequest(&codecommit.ListRepositoriesInput{})
	p := codecommit.NewListRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Repositories {
			r.resources = append(r.resources, *resource.RepositoryId)
		}
	}
	r.err = p.Err()
	return
}

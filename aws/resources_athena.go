package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
)

func getAthena(config aws.Config) (resources resourceMap) {
	client := athena.New(config)
	resources = reduce(
		getAthenaNamedQuery(client).unwrap(athenaNamedQuery),
	)
	return
}

func getAthenaNamedQuery(client *athena.Client) (r resourceSliceError) {
	req := client.ListNamedQueriesRequest(&athena.ListNamedQueriesInput{})
	p := athena.NewListNamedQueriesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.NamedQueryIds {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

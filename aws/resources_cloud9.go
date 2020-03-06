package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
)

func getCloud9(config aws.Config) (resources resourceMap) {
	client := cloud9.New(config)
	resources = reduce(
		getCloud9EnvironmentEC2(client).unwrap(cloud9EnvironmentEC2),
	)
	return
}

func getCloud9EnvironmentEC2(client *cloud9.Client) (r resourceSliceError) {
	req := client.ListEnvironmentsRequest(&cloud9.ListEnvironmentsInput{})
	p := cloud9.NewListEnvironmentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.EnvironmentIds {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

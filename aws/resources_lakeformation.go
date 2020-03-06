package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
)

func getLakeFormation(config aws.Config) (resources resourceMap) {
	client := lakeformation.New(config)
	resources = reduce(
		getLakeFormationResource(client).unwrap(lakeFormationResource),
	)
	return
}

func getLakeFormationResource(client *lakeformation.Client) (r resourceSliceError) {
	req := client.ListResourcesRequest(&lakeformation.ListResourcesInput{})
	p := lakeformation.NewListResourcesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ResourceInfoList {
			r.resources = append(r.resources, *resource.ResourceArn)
		}
	}
	r.err = p.Err()
	return
}

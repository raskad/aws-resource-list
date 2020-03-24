package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
)

func getLakeFormation(config aws.Config) (resources resourceMap) {
	client := lakeformation.New(config)

	getLakeFormationResourceARNs := getLakeFormationResourceARNs(client)

	resources = resourceMap{
		lakeFormationResource: getLakeFormationResourceARNs,
	}
	return
}

func getLakeFormationResourceARNs(client *lakeformation.Client) (resources []string) {
	req := client.ListResourcesRequest(&lakeformation.ListResourcesInput{})
	p := lakeformation.NewListResourcesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ResourceInfoList {
			resources = append(resources, *resource.ResourceArn)
		}
	}
	return
}

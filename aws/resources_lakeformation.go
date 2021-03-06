package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
)

func getLakeFormation(config aws.Config) (resources awsResourceMap) {
	client := lakeformation.New(config)

	lakeFormationResourceARNs := getLakeFormationResourceARNs(client)

	resources = awsResourceMap{
		lakeFormationResource: lakeFormationResourceARNs,
	}
	return
}

func getLakeFormationResourceARNs(client *lakeformation.Client) (resources []string) {
	req := client.ListResourcesRequest(&lakeformation.ListResourcesInput{})
	p := lakeformation.NewListResourcesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ResourceInfoList {
			resources = append(resources, *resource.ResourceArn)
		}
	}
	return
}

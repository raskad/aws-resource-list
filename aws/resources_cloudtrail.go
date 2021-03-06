package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

func getCloudTrail(config aws.Config) (resources awsResourceMap) {
	client := cloudtrail.New(config)

	cloudTrailTrailNames := getCloudTrailTrailNames(client)

	resources = awsResourceMap{
		cloudTrailTrail: cloudTrailTrailNames,
	}
	return
}

func getCloudTrailTrailNames(client *cloudtrail.Client) (resources []string) {
	req := client.ListTrailsRequest(&cloudtrail.ListTrailsInput{})
	p := cloudtrail.NewListTrailsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Trails {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

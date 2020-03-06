package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

func getCloudTrail(config aws.Config) (resources resourceMap) {
	client := cloudtrail.New(config)
	resources = reduce(
		getCloudTrailTrail(client).unwrap(cloudTrailTrail),
	)
	return
}

func getCloudTrailTrail(client *cloudtrail.Client) (r resourceSliceError) {
	req := client.ListTrailsRequest(&cloudtrail.ListTrailsInput{})
	p := cloudtrail.NewListTrailsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Trails {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

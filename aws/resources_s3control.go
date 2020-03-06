package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

func getS3Control(config aws.Config) (resources resourceMap) {
	client := s3control.New(config)
	resources = reduce(
		getS3AccessPoint(client).unwrap(s3AccessPoint),
	)
	return
}

func getS3AccessPoint(client *s3control.Client) (r resourceSliceError) {
	req := client.ListAccessPointsRequest(&s3control.ListAccessPointsInput{})
	p := s3control.NewListAccessPointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.AccessPointList {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

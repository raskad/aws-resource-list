package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

func getS3Control(config aws.Config) (resources resourceMap) {
	client := s3control.New(config)

	s3AccessPointNames := getS3AccessPointNames(client)

	resources = resourceMap{
		s3AccessPoint: s3AccessPointNames,
	}
	return
}

func getS3AccessPointNames(client *s3control.Client) (resources []string) {
	req := client.ListAccessPointsRequest(&s3control.ListAccessPointsInput{})
	p := s3control.NewListAccessPointsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.AccessPointList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

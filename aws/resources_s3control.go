package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

func getS3Control(config aws.Config) (resources awsResourceMap) {
	client := s3control.New(config)

	s3ControlAccessPointNames := getS3ControlAccessPointNames(client)

	resources = awsResourceMap{
		s3ControlAccessPoint: s3ControlAccessPointNames,
	}
	return
}

func getS3ControlAccessPointNames(client *s3control.Client) (resources []string) {
	req := client.ListAccessPointsRequest(&s3control.ListAccessPointsInput{})
	p := s3control.NewListAccessPointsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.AccessPointList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

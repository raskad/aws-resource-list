package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
)

func getKinesisVideo(config aws.Config) (resources awsResourceMap) {
	client := kinesisvideo.New(config)

	kinesisVideoStreamNames := getKinesisVideoStreamNames(client)

	resources = awsResourceMap{
		kinesisVideoStream: kinesisVideoStreamNames,
	}
	return
}

func getKinesisVideoStreamNames(client *kinesisvideo.Client) (resources []string) {
	req := client.ListStreamsRequest(&kinesisvideo.ListStreamsInput{})
	p := kinesisvideo.NewListStreamsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.StreamInfoList {
			resources = append(resources, *resource.StreamName)
		}
	}
	return
}

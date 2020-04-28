package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
)

func getFirehose(config aws.Config) (resources awsResourceMap) {
	client := firehose.New(config)

	firehoseDeliveryStreamNames := getFirehoseDeliveryStreamNames(client)

	resources = awsResourceMap{
		firehoseDeliveryStream: firehoseDeliveryStreamNames,
	}
	return
}

func getFirehoseDeliveryStreamNames(client *firehose.Client) (resources []string) {
	page, err := client.ListDeliveryStreamsRequest(&firehose.ListDeliveryStreamsInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	resources = append(resources, page.DeliveryStreamNames...)
	return
}

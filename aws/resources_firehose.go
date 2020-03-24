package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
)

func getFirehose(config aws.Config) (resources resourceMap) {
	client := firehose.New(config)

	kinesisFirehoseDeliveryStreamNames := getKinesisFirehoseDeliveryStreamNames(client)

	resources = resourceMap{
		kinesisFirehoseDeliveryStream: kinesisFirehoseDeliveryStreamNames,
	}
	return
}

func getKinesisFirehoseDeliveryStreamNames(client *firehose.Client) (resources []string) {
	page, err := client.ListDeliveryStreamsRequest(&firehose.ListDeliveryStreamsInput{}).Send(context.Background())
	logErr(err)
	resources = append(resources, page.DeliveryStreamNames...)
	return
}

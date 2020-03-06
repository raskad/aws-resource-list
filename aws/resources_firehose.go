package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
)

func getFirehose(config aws.Config) (resources resourceMap) {
	client := firehose.New(config)
	resources = reduce(
		getKinesisFirehoseDeliveryStream(client).unwrap(kinesisFirehoseDeliveryStream),
	)
	return
}

func getKinesisFirehoseDeliveryStream(client *firehose.Client) (r resourceSliceError) {
	page, err := client.ListDeliveryStreamsRequest(&firehose.ListDeliveryStreamsInput{}).Send(context.Background())
	for _, resource := range page.DeliveryStreamNames {
		r.resources = append(r.resources, resource)
	}
	r.err = err
	return
}

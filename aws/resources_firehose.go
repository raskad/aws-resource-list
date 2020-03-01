package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

func getFirehose(session *session.Session) (resources resourceMap) {
	client := firehose.New(session)
	resources = reduce(
		getKinesisFirehoseDeliveryStream(client).unwrap(kinesisFirehoseDeliveryStream),
	)
	return
}

func getKinesisFirehoseDeliveryStream(client *firehose.Firehose) (r resourceSliceError) {
	page, err := client.ListDeliveryStreams(&firehose.ListDeliveryStreamsInput{})
	for _, resource := range page.DeliveryStreamNames {
		r.resources = append(r.resources, *resource)
	}
	r.err = err
	return
}

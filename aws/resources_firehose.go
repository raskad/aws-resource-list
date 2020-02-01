package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

func getFirehose(session *session.Session) (resources resourceMap) {
	client := firehose.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		kinesisFirehoseDeliveryStream: getKinesisFirehoseDeliveryStream(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getKinesisFirehoseDeliveryStream(client *firehose.Firehose) (r resourceSliceError) {
	logDebug("Listing KinesisFirehoseDeliveryStream resources")
	page, err := client.ListDeliveryStreams(&firehose.ListDeliveryStreamsInput{})
	for _, resource := range page.DeliveryStreamNames {
		logDebug("Got KinesisFirehoseDeliveryStream resource with PhysicalResourceId", *resource)
		r.resources = append(r.resources, *resource)
	}
	r.err = err
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func getKinesis(session *session.Session) (resources resourceMap) {
	client := kinesis.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		kinesisStream:         getKinesisStream(client),
		kinesisStreamConsumer: getKinesisStreamConsumer(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getKinesisStream(client *kinesis.Kinesis) (r resourceSliceError) {
	r.err = client.ListStreamsPages(&kinesis.ListStreamsInput{}, func(page *kinesis.ListStreamsOutput, lastPage bool) bool {
		logDebug("Listing KinesisStream resources page")
		for _, resource := range page.StreamNames {
			logDebug("Got KinesisStream resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getKinesisStreamConsumer(client *kinesis.Kinesis) (r resourceSliceError) {
	r.err = client.ListStreamConsumersPages(&kinesis.ListStreamConsumersInput{}, func(page *kinesis.ListStreamConsumersOutput, lastPage bool) bool {
		logDebug("Listing KinesisStreamConsumer resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Consumers {
			logDebug("Got KinesisStreamConsumer resource with PhysicalResourceId", *resource.ConsumerName)
			r.resources = append(r.resources, *resource.ConsumerName)
		}
		return true
	})
	return
}

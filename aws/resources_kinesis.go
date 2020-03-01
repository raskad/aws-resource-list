package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func getKinesis(session *session.Session) (resources resourceMap) {
	client := kinesis.New(session)
	resources = reduce(
		getKinesisStream(client).unwrap(kinesisStream),
		getKinesisStreamConsumer(client).unwrap(kinesisStreamConsumer),
	)
	return
}

func getKinesisStream(client *kinesis.Kinesis) (r resourceSliceError) {
	r.err = client.ListStreamsPages(&kinesis.ListStreamsInput{}, func(page *kinesis.ListStreamsOutput, lastPage bool) bool {
		for _, resource := range page.StreamNames {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getKinesisStreamConsumer(client *kinesis.Kinesis) (r resourceSliceError) {
	r.err = client.ListStreamConsumersPages(&kinesis.ListStreamConsumersInput{}, func(page *kinesis.ListStreamConsumersOutput, lastPage bool) bool {
		for _, resource := range page.Consumers {
			r.resources = append(r.resources, *resource.ConsumerName)
		}
		return true
	})
	return
}

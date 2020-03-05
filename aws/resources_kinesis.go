package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func getKinesis(session *session.Session) (resources resourceMap) {
	client := kinesis.New(session)

	kinesisStreamResourceMap := getKinesisStream(client).unwrap(kinesisStream)
	kinesisStreamARNs := kinesisStreamResourceMap[kinesisStream]

	resources = reduce(
		kinesisStreamResourceMap,
		getKinesisStreamConsumer(client, kinesisStreamARNs).unwrap(kinesisStreamConsumer),
	)
	return
}

func getKinesisStream(client *kinesis.Kinesis) (r resourceSliceError) {
	r.err = client.ListStreamsPages(&kinesis.ListStreamsInput{}, func(page *kinesis.ListStreamsOutput, lastPage bool) bool {
		for _, resource := range page.StreamNames {
			r.err = client.DescribeStreamPages(&kinesis.DescribeStreamInput{
				StreamName: resource,
			}, func(page *kinesis.DescribeStreamOutput, lastPage bool) bool {
				r.resources = append(r.resources, *page.StreamDescription.StreamARN)
				return true
			})
		}
		return true
	})
	return
}

func getKinesisStreamConsumer(client *kinesis.Kinesis, streamARNs []string) (r resourceSliceError) {
	for _, streamARN := range streamARNs {
		r.err = client.ListStreamConsumersPages(&kinesis.ListStreamConsumersInput{
			StreamARN: aws.String(streamARN),
		}, func(page *kinesis.ListStreamConsumersOutput, lastPage bool) bool {
			for _, resource := range page.Consumers {
				r.resources = append(r.resources, *resource.ConsumerName)
			}
			return true
		})
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

func getKinesis(config aws.Config) (resources resourceMap) {
	client := kinesis.New(config)

	kinesisStreamResourceMap := getKinesisStream(client).unwrap(kinesisStream)
	kinesisStreamARNs := kinesisStreamResourceMap[kinesisStream]

	resources = reduce(
		kinesisStreamResourceMap,
		getKinesisStreamConsumer(client, kinesisStreamARNs).unwrap(kinesisStreamConsumer),
	)
	return
}

func getKinesisStream(client *kinesis.Client) (r resourceSliceError) {
	req := client.ListStreamsRequest(&kinesis.ListStreamsInput{})
	p := kinesis.NewListStreamsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.StreamNames {
			req := client.DescribeStreamRequest(&kinesis.DescribeStreamInput{
				StreamName: aws.String(resource),
			})
			p := kinesis.NewDescribeStreamPaginator(req)
			for p.Next(context.Background()) {
				page := p.CurrentPage()
				r.resources = append(r.resources, *page.StreamDescription.StreamARN)
			}
			r.err = p.Err()
			return
		}
	}
	r.err = p.Err()
	return
}

func getKinesisStreamConsumer(client *kinesis.Client, streamARNs []string) (r resourceSliceError) {
	for _, streamARN := range streamARNs {
		req := client.ListStreamConsumersRequest(&kinesis.ListStreamConsumersInput{
			StreamARN: aws.String(streamARN),
		})
		p := kinesis.NewListStreamConsumersPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.Consumers {
				r.resources = append(r.resources, *resource.ConsumerName)
			}
		}
		r.err = p.Err()
		return
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

func getKinesis(config aws.Config) (resources awsResourceMap) {
	client := kinesis.New(config)

	kinesisStreamARNs := getKinesisStreamARNs(client)
	kinesisStreamConsumerNames := getKinesisStreamConsumerNames(client, kinesisStreamARNs)

	resources = awsResourceMap{
		kinesisStream:         kinesisStreamARNs,
		kinesisStreamConsumer: kinesisStreamConsumerNames,
	}
	return
}

func getKinesisStreamARNs(client *kinesis.Client) (resources []string) {
	req := client.ListStreamsRequest(&kinesis.ListStreamsInput{})
	p := kinesis.NewListStreamsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.StreamNames {
			req := client.DescribeStreamRequest(&kinesis.DescribeStreamInput{
				StreamName: aws.String(resource),
			})
			p := kinesis.NewDescribeStreamPaginator(req)
			for p.Next(context.Background()) {
				if p.Err() != nil {
					logErr(p.Err())
					return
				}
				page := p.CurrentPage()
				resources = append(resources, *page.StreamDescription.StreamARN)
			}
		}
	}
	return
}

func getKinesisStreamConsumerNames(client *kinesis.Client, streamARNs []string) (resources []string) {
	for _, streamARN := range streamARNs {
		req := client.ListStreamConsumersRequest(&kinesis.ListStreamConsumersInput{
			StreamARN: aws.String(streamARN),
		})
		p := kinesis.NewListStreamConsumersPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.Consumers {
				resources = append(resources, *resource.ConsumerName)
			}
		}
	}
	return
}

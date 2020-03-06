package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func getSqs(config aws.Config) (resources resourceMap) {
	client := sqs.New(config)
	resources = reduce(
		getSqsQueue(client).unwrap(sqsQueue),
	)
	return
}

func getSqsQueue(client *sqs.Client) (r resourceSliceError) {
	page, err := client.ListQueuesRequest(&sqs.ListQueuesInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.QueueUrls {
		r.resources = append(r.resources, resource)
	}
	return
}

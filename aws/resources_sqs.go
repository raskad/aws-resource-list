package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func getSqs(config aws.Config) (resources awsResourceMap) {
	client := sqs.New(config)

	sqsQueueNames := getSqsQueueNames(client)

	resources = awsResourceMap{
		sqsQueue: sqsQueueNames,
	}
	return
}

func getSqsQueueNames(client *sqs.Client) (resources []string) {
	page, err := client.ListQueuesRequest(&sqs.ListQueuesInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	resources = append(resources, page.QueueUrls...)
	return
}

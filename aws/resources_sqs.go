package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func getSqs(session *session.Session) (resources resourceMap) {
	client := sqs.New(session)
	resources = reduce(
		getSqsQueue(client).unwrap(sqsQueue),
	)
	return
}

func getSqsQueue(client *sqs.SQS) (r resourceSliceError) {
	page, err := client.ListQueues(&sqs.ListQueuesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.QueueUrls {
		r.resources = append(r.resources, *resource)
	}
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func getSqs(session *session.Session) (resources resourceMap) {
	client := sqs.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		sqsQueue: getSqsQueue(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getSqsQueue(client *sqs.SQS) (r resourceSliceError) {
	logDebug("Listing SqsQueue resources")
	page, err := client.ListQueues(&sqs.ListQueuesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.QueueUrls {
		logDebug("Got SqsQueue resource with PhysicalResourceId", *resource)
		r.resources = append(r.resources, *resource)
	}
	return
}

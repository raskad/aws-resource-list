package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
)

func getBatch(session *session.Session) (resources resourceMap) {
	client := batch.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		batchComputeEnvironment: getBatchComputeEnvironment(client),
		batchJobDefinition:      getBatchJobDefinition(client),
		batchJobQueue:           getBatchJobQueue(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getBatchComputeEnvironment(client *batch.Batch) (r resourceSliceError) {
	logDebug("Listing BatchComputeEnvironment resources")
	r.err = client.DescribeComputeEnvironmentsPages(&batch.DescribeComputeEnvironmentsInput{}, func(page *batch.DescribeComputeEnvironmentsOutput, lastPage bool) bool {
		for _, resource := range page.ComputeEnvironments {
			logDebug("Got BatchComputeEnvironment resource with PhysicalResourceId", *resource.ComputeEnvironmentName)
			r.resources = append(r.resources, *resource.ComputeEnvironmentName)
		}
		return true
	})
	return
}

func getBatchJobDefinition(client *batch.Batch) (r resourceSliceError) {
	logDebug("Listing BatchJobDefinition resources")
	r.err = client.DescribeJobDefinitionsPages(&batch.DescribeJobDefinitionsInput{}, func(page *batch.DescribeJobDefinitionsOutput, lastPage bool) bool {
		for _, resource := range page.JobDefinitions {
			logDebug("Got BatchJobDefinition resource with PhysicalResourceId", *resource.JobDefinitionName)
			r.resources = append(r.resources, *resource.JobDefinitionName)
		}
		return true
	})
	return
}

func getBatchJobQueue(client *batch.Batch) (r resourceSliceError) {
	logDebug("Listing BatchJobQueue resources")
	r.err = client.DescribeJobQueuesPages(&batch.DescribeJobQueuesInput{}, func(page *batch.DescribeJobQueuesOutput, lastPage bool) bool {
		for _, resource := range page.JobQueues {
			logDebug("Got BatchJobQueue resource with PhysicalResourceId", *resource.JobQueueName)
			r.resources = append(r.resources, *resource.JobQueueName)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
)

func getBatch(session *session.Session) (resources resourceMap) {
	client := batch.New(session)
	resources = reduce(
		getBatchComputeEnvironment(client).unwrap(batchComputeEnvironment),
		getBatchJobDefinition(client).unwrap(batchJobDefinition),
		getBatchJobQueue(client).unwrap(batchJobQueue),
	)
	return
}

func getBatchComputeEnvironment(client *batch.Batch) (r resourceSliceError) {
	r.err = client.DescribeComputeEnvironmentsPages(&batch.DescribeComputeEnvironmentsInput{}, func(page *batch.DescribeComputeEnvironmentsOutput, lastPage bool) bool {
		for _, resource := range page.ComputeEnvironments {
			r.resources = append(r.resources, *resource.ComputeEnvironmentName)
		}
		return true
	})
	return
}

func getBatchJobDefinition(client *batch.Batch) (r resourceSliceError) {
	r.err = client.DescribeJobDefinitionsPages(&batch.DescribeJobDefinitionsInput{}, func(page *batch.DescribeJobDefinitionsOutput, lastPage bool) bool {
		for _, resource := range page.JobDefinitions {
			r.resources = append(r.resources, *resource.JobDefinitionName)
		}
		return true
	})
	return
}

func getBatchJobQueue(client *batch.Batch) (r resourceSliceError) {
	r.err = client.DescribeJobQueuesPages(&batch.DescribeJobQueuesInput{}, func(page *batch.DescribeJobQueuesOutput, lastPage bool) bool {
		for _, resource := range page.JobQueues {
			r.resources = append(r.resources, *resource.JobQueueName)
		}
		return true
	})
	return
}

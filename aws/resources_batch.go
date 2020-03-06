package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func getBatch(config aws.Config) (resources resourceMap) {
	client := batch.New(config)
	resources = reduce(
		getBatchComputeEnvironment(client).unwrap(batchComputeEnvironment),
		getBatchJobDefinition(client).unwrap(batchJobDefinition),
		getBatchJobQueue(client).unwrap(batchJobQueue),
	)
	return
}

func getBatchComputeEnvironment(client *batch.Client) (r resourceSliceError) {
	req := client.DescribeComputeEnvironmentsRequest(&batch.DescribeComputeEnvironmentsInput{})
	p := batch.NewDescribeComputeEnvironmentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ComputeEnvironments {
			r.resources = append(r.resources, *resource.ComputeEnvironmentName)
		}
	}
	r.err = p.Err()
	return
}

func getBatchJobDefinition(client *batch.Client) (r resourceSliceError) {
	req := client.DescribeJobDefinitionsRequest(&batch.DescribeJobDefinitionsInput{})
	p := batch.NewDescribeJobDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.JobDefinitions {
			r.resources = append(r.resources, *resource.JobDefinitionName)
		}
	}
	r.err = p.Err()
	return
}

func getBatchJobQueue(client *batch.Client) (r resourceSliceError) {
	req := client.DescribeJobQueuesRequest(&batch.DescribeJobQueuesInput{})
	p := batch.NewDescribeJobQueuesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.JobQueues {
			r.resources = append(r.resources, *resource.JobQueueName)
		}
	}
	r.err = p.Err()
	return
}

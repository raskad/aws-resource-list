package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func getBatch(config aws.Config) (resources resourceMap) {
	client := batch.New(config)

	batchComputeEnvironmentNames := getBatchComputeEnvironmentNames(client)
	batchJobDefinitionARNs := getBatchJobDefinitionARNs(client)
	batchJobQueueARNs := getBatchJobQueueARNs(client)

	resources = resourceMap{
		batchComputeEnvironment: batchComputeEnvironmentNames,
		batchJobDefinition:      batchJobDefinitionARNs,
		batchJobQueue:           batchJobQueueARNs,
	}
	return
}

func getBatchComputeEnvironmentNames(client *batch.Client) (resources []string) {
	req := client.DescribeComputeEnvironmentsRequest(&batch.DescribeComputeEnvironmentsInput{})
	p := batch.NewDescribeComputeEnvironmentsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ComputeEnvironments {
			resources = append(resources, *resource.ComputeEnvironmentName)
		}
	}
	return
}

func getBatchJobDefinitionARNs(client *batch.Client) (resources []string) {
	req := client.DescribeJobDefinitionsRequest(&batch.DescribeJobDefinitionsInput{})
	p := batch.NewDescribeJobDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.JobDefinitions {
			resources = append(resources, *resource.JobDefinitionArn)
		}
	}
	return
}

func getBatchJobQueueARNs(client *batch.Client) (resources []string) {
	req := client.DescribeJobQueuesRequest(&batch.DescribeJobQueuesInput{})
	p := batch.NewDescribeJobQueuesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.JobQueues {
			resources = append(resources, *resource.JobQueueArn)
		}
	}
	return
}

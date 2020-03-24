package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

func getCodePipeline(config aws.Config) (resources resourceMap) {
	client := codepipeline.New(config)

	codePipelinePipelineNames := getCodePipelinePipelineNames(client)
	codePipelineWebhookARNs := getCodePipelineWebhookARNs(client)

	resources = resourceMap{
		codePipelinePipeline: codePipelinePipelineNames,
		codePipelineWebhook:  codePipelineWebhookARNs,
	}
	return
}

func getCodePipelinePipelineNames(client *codepipeline.Client) (resources []string) {
	req := client.ListPipelinesRequest(&codepipeline.ListPipelinesInput{})
	p := codepipeline.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Pipelines {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getCodePipelineWebhookARNs(client *codepipeline.Client) (resources []string) {
	req := client.ListWebhooksRequest(&codepipeline.ListWebhooksInput{})
	p := codepipeline.NewListWebhooksPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Webhooks {
			resources = append(resources, *resource.Arn)
		}
	}
	return
}

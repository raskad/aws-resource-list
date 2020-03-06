package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

func getCodePipeline(config aws.Config) (resources resourceMap) {
	client := codepipeline.New(config)
	resources = reduce(
		getCodePipelinePipeline(client).unwrap(codePipelinePipeline),
		getCodePipelineWebhook(client).unwrap(codePipelineWebhook),
	)
	return
}

func getCodePipelinePipeline(client *codepipeline.Client) (r resourceSliceError) {
	req := client.ListPipelinesRequest(&codepipeline.ListPipelinesInput{})
	p := codepipeline.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Pipelines {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getCodePipelineWebhook(client *codepipeline.Client) (r resourceSliceError) {
	req := client.ListWebhooksRequest(&codepipeline.ListWebhooksInput{})
	p := codepipeline.NewListWebhooksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Webhooks {
			r.resources = append(r.resources, *resource.Arn)
		}
	}
	r.err = p.Err()
	return
}

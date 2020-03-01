package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"
)

func getCodePipeline(session *session.Session) (resources resourceMap) {
	client := codepipeline.New(session)
	resources = reduce(
		getCodePipelinePipeline(client).unwrap(codePipelinePipeline),
		getCodePipelineWebhook(client).unwrap(codePipelineWebhook),
	)
	return
}

func getCodePipelinePipeline(client *codepipeline.CodePipeline) (r resourceSliceError) {
	r.err = client.ListPipelinesPages(&codepipeline.ListPipelinesInput{}, func(page *codepipeline.ListPipelinesOutput, lastPage bool) bool {
		for _, resource := range page.Pipelines {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getCodePipelineWebhook(client *codepipeline.CodePipeline) (r resourceSliceError) {
	r.err = client.ListWebhooksPages(&codepipeline.ListWebhooksInput{}, func(page *codepipeline.ListWebhooksOutput, lastPage bool) bool {
		for _, resource := range page.Webhooks {
			r.resources = append(r.resources, *resource.Arn)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"
)

func getCodePipeline(session *session.Session) (resources resourceMap) {
	client := codepipeline.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		codePipelinePipeline: getCodePipelinePipeline(client),
		codePipelineWebhook:  getCodePipelineWebhook(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCodePipelinePipeline(client *codepipeline.CodePipeline) (r resourceSliceError) {
	r.err = client.ListPipelinesPages(&codepipeline.ListPipelinesInput{}, func(page *codepipeline.ListPipelinesOutput, lastPage bool) bool {
		logDebug("List CodePipelinePipeline resources page")
		for _, resource := range page.Pipelines {
			logDebug("Got CodePipelinePipeline resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getCodePipelineWebhook(client *codepipeline.CodePipeline) (r resourceSliceError) {
	r.err = client.ListWebhooksPages(&codepipeline.ListWebhooksInput{}, func(page *codepipeline.ListWebhooksOutput, lastPage bool) bool {
		logDebug("List CodePipelineWebhook resources page")
		for _, resource := range page.Webhooks {
			logDebug("Got CodePipelineWebhook resource with PhysicalResourceId", *resource.Arn)
			r.resources = append(r.resources, *resource.Arn)
		}
		return true
	})
	return
}

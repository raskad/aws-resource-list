package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/datapipeline"
)

func getDataPipeline(session *session.Session) (resources resourceMap) {
	client := datapipeline.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		dataPipelinePipeline: getDataPipelinePipeline(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getDataPipelinePipeline(client *datapipeline.DataPipeline) (r resourceSliceError) {
	logDebug("Listing DataPipelinePipeline resources")
	r.err = client.ListPipelinesPages(&datapipeline.ListPipelinesInput{}, func(page *datapipeline.ListPipelinesOutput, lastPage bool) bool {
		for _, resource := range page.PipelineIdList {
			logDebug("Got DataPipelinePipeline resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

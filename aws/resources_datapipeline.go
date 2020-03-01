package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/datapipeline"
)

func getDataPipeline(session *session.Session) (resources resourceMap) {
	client := datapipeline.New(session)
	resources = reduce(
		getDataPipelinePipeline(client).unwrap(dataPipelinePipeline),
	)
	return
}

func getDataPipelinePipeline(client *datapipeline.DataPipeline) (r resourceSliceError) {
	r.err = client.ListPipelinesPages(&datapipeline.ListPipelinesInput{}, func(page *datapipeline.ListPipelinesOutput, lastPage bool) bool {
		for _, resource := range page.PipelineIdList {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

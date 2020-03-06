package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
)

func getDataPipeline(config aws.Config) (resources resourceMap) {
	client := datapipeline.New(config)
	resources = reduce(
		getDataPipelinePipeline(client).unwrap(dataPipelinePipeline),
	)
	return
}

func getDataPipelinePipeline(client *datapipeline.Client) (r resourceSliceError) {
	req := client.ListPipelinesRequest(&datapipeline.ListPipelinesInput{})
	p := datapipeline.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.PipelineIdList {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

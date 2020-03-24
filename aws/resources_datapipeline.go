package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
)

func getDataPipeline(config aws.Config) (resources resourceMap) {
	client := datapipeline.New(config)

	dataPipelinePipelineIDs := getDataPipelinePipelineIDs(client)

	resources = resourceMap{
		dataPipelinePipeline: dataPipelinePipelineIDs,
	}
	return
}

func getDataPipelinePipelineIDs(client *datapipeline.Client) (resources []string) {
	req := client.ListPipelinesRequest(&datapipeline.ListPipelinesInput{})
	p := datapipeline.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.PipelineIdList {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

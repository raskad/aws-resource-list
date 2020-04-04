package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
)

func getElasticTranscoder(config aws.Config) (resources awsResourceMap) {
	client := elastictranscoder.New(config)

	elasticTranscoderPipelineIDs := getElasticTranscoderPipelineIDs(client)
	elasticTranscoderPresetIDs := getElasticTranscoderPresetIDs(client)

	resources = awsResourceMap{
		elasticTranscoderPipeline: elasticTranscoderPipelineIDs,
		elasticTranscoderPreset:   elasticTranscoderPresetIDs,
	}
	return
}

func getElasticTranscoderPipelineIDs(client *elastictranscoder.Client) (resources []string) {
	req := client.ListPipelinesRequest(&elastictranscoder.ListPipelinesInput{})
	p := elastictranscoder.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Pipelines {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getElasticTranscoderPresetIDs(client *elastictranscoder.Client) (resources []string) {
	req := client.ListPresetsRequest(&elastictranscoder.ListPresetsInput{})
	p := elastictranscoder.NewListPresetsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Presets {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
)

func getMediaConvert(config aws.Config) (resources resourceMap) {
	client := mediaconvert.New(config)
	resources = reduce(
		getMediaConvertJobTemplate(client).unwrap(mediaConvertJobTemplate),
		getMediaConvertPreset(client).unwrap(mediaConvertPreset),
		getMediaConvertQueue(client).unwrap(mediaConvertQueue),
	)
	return
}

func getMediaConvertJobTemplate(client *mediaconvert.Client) (r resourceSliceError) {
	req := client.ListJobTemplatesRequest(&mediaconvert.ListJobTemplatesInput{})
	p := mediaconvert.NewListJobTemplatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.JobTemplates {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getMediaConvertPreset(client *mediaconvert.Client) (r resourceSliceError) {
	req := client.ListPresetsRequest(&mediaconvert.ListPresetsInput{})
	p := mediaconvert.NewListPresetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Presets {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getMediaConvertQueue(client *mediaconvert.Client) (r resourceSliceError) {
	req := client.ListQueuesRequest(&mediaconvert.ListQueuesInput{})
	p := mediaconvert.NewListQueuesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Queues {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

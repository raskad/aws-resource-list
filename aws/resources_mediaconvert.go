package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
)

func getMediaConvert(config aws.Config) (resources resourceMap) {
	client := mediaconvert.New(config)

	mediaConvertJobTemplateNames := getMediaConvertJobTemplateNames(client)
	mediaConvertPresetNames := getMediaConvertPresetNames(client)
	mediaConvertQueueNames := getMediaConvertQueueNames(client)

	resources = resourceMap{
		mediaConvertJobTemplate: mediaConvertJobTemplateNames,
		mediaConvertPreset:      mediaConvertPresetNames,
		mediaConvertQueue:       mediaConvertQueueNames,
	}
	return
}

func getMediaConvertJobTemplateNames(client *mediaconvert.Client) (resources []string) {
	req := client.ListJobTemplatesRequest(&mediaconvert.ListJobTemplatesInput{})
	p := mediaconvert.NewListJobTemplatesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.JobTemplates {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getMediaConvertPresetNames(client *mediaconvert.Client) (resources []string) {
	req := client.ListPresetsRequest(&mediaconvert.ListPresetsInput{})
	p := mediaconvert.NewListPresetsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Presets {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getMediaConvertQueueNames(client *mediaconvert.Client) (resources []string) {
	req := client.ListQueuesRequest(&mediaconvert.ListQueuesInput{})
	p := mediaconvert.NewListQueuesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Queues {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
)

func getMediaConvert(session *session.Session) (resources resourceMap) {
	client := mediaconvert.New(session)
	resources = reduce(
		getMediaConvertJobTemplate(client).unwrap(mediaConvertJobTemplate),
		getMediaConvertPreset(client).unwrap(mediaConvertPreset),
		getMediaConvertQueue(client).unwrap(mediaConvertQueue),
	)
	return
}

func getMediaConvertJobTemplate(client *mediaconvert.MediaConvert) (r resourceSliceError) {
	r.err = client.ListJobTemplatesPages(&mediaconvert.ListJobTemplatesInput{}, func(page *mediaconvert.ListJobTemplatesOutput, lastPage bool) bool {
		for _, resource := range page.JobTemplates {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getMediaConvertPreset(client *mediaconvert.MediaConvert) (r resourceSliceError) {
	r.err = client.ListPresetsPages(&mediaconvert.ListPresetsInput{}, func(page *mediaconvert.ListPresetsOutput, lastPage bool) bool {
		for _, resource := range page.Presets {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getMediaConvertQueue(client *mediaconvert.MediaConvert) (r resourceSliceError) {
	r.err = client.ListQueuesPages(&mediaconvert.ListQueuesInput{}, func(page *mediaconvert.ListQueuesOutput, lastPage bool) bool {
		for _, resource := range page.Queues {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

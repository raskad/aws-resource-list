package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
)

func getMediaConvert(session *session.Session) (resources resourceMap) {
	client := mediaconvert.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		mediaConvertJobTemplate: getMediaConvertJobTemplate(client),
		mediaConvertPreset:      getMediaConvertPreset(client),
		mediaConvertQueue:       getMediaConvertQueue(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getMediaConvertJobTemplate(client *mediaconvert.MediaConvert) (r resourceSliceError) {
	logDebug("Listing MediaConvertJobTemplate resources")
	r.err = client.ListJobTemplatesPages(&mediaconvert.ListJobTemplatesInput{}, func(page *mediaconvert.ListJobTemplatesOutput, lastPage bool) bool {
		for _, resource := range page.JobTemplates {
			logDebug("Got MediaConvertJobTemplate resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getMediaConvertPreset(client *mediaconvert.MediaConvert) (r resourceSliceError) {
	logDebug("Listing MediaConvertPreset resources")
	r.err = client.ListPresetsPages(&mediaconvert.ListPresetsInput{}, func(page *mediaconvert.ListPresetsOutput, lastPage bool) bool {
		for _, resource := range page.Presets {
			logDebug("Got MediaConvertPreset resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getMediaConvertQueue(client *mediaconvert.MediaConvert) (r resourceSliceError) {
	logDebug("Listing MediaConvertQueue resources")
	r.err = client.ListQueuesPages(&mediaconvert.ListQueuesInput{}, func(page *mediaconvert.ListQueuesOutput, lastPage bool) bool {
		for _, resource := range page.Queues {
			logDebug("Got MediaConvertQueue resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

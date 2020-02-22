package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/medialive"
)

func getMediaLive(session *session.Session) (resources resourceMap) {
	client := medialive.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		mediaLiveChannel:            getMediaLiveChannel(client),
		mediaLiveInput:              getMediaLiveInput(client),
		mediaLiveInputSecurityGroup: getMediaLiveInputSecurityGroup(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getMediaLiveChannel(client *medialive.MediaLive) (r resourceSliceError) {
	r.err = client.ListChannelsPages(&medialive.ListChannelsInput{}, func(page *medialive.ListChannelsOutput, lastPage bool) bool {
		logDebug("Listing MediaLiveChannel resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Channels {
			logDebug("Got MediaLiveChannel resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getMediaLiveInput(client *medialive.MediaLive) (r resourceSliceError) {
	r.err = client.ListInputsPages(&medialive.ListInputsInput{}, func(page *medialive.ListInputsOutput, lastPage bool) bool {
		logDebug("Listing MediaLiveInput resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Inputs {
			logDebug("Got MediaLiveInput resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getMediaLiveInputSecurityGroup(client *medialive.MediaLive) (r resourceSliceError) {
	r.err = client.ListInputSecurityGroupsPages(&medialive.ListInputSecurityGroupsInput{}, func(page *medialive.ListInputSecurityGroupsOutput, lastPage bool) bool {
		logDebug("Listing MediaLiveInputSecurityGroup resources page. Remaining pages", page.NextToken)
		for _, resource := range page.InputSecurityGroups {
			logDebug("Got MediaLiveInputSecurityGroup resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

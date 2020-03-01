package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/medialive"
)

func getMediaLive(session *session.Session) (resources resourceMap) {
	client := medialive.New(session)
	resources = reduce(
		getMediaLiveChannel(client).unwrap(mediaLiveChannel),
		getMediaLiveInput(client).unwrap(mediaLiveInput),
		getMediaLiveInputSecurityGroup(client).unwrap(mediaLiveInputSecurityGroup),
	)
	return
}

func getMediaLiveChannel(client *medialive.MediaLive) (r resourceSliceError) {
	r.err = client.ListChannelsPages(&medialive.ListChannelsInput{}, func(page *medialive.ListChannelsOutput, lastPage bool) bool {
		for _, resource := range page.Channels {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getMediaLiveInput(client *medialive.MediaLive) (r resourceSliceError) {
	r.err = client.ListInputsPages(&medialive.ListInputsInput{}, func(page *medialive.ListInputsOutput, lastPage bool) bool {
		for _, resource := range page.Inputs {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getMediaLiveInputSecurityGroup(client *medialive.MediaLive) (r resourceSliceError) {
	r.err = client.ListInputSecurityGroupsPages(&medialive.ListInputSecurityGroupsInput{}, func(page *medialive.ListInputSecurityGroupsOutput, lastPage bool) bool {
		for _, resource := range page.InputSecurityGroups {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

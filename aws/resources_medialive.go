package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

func getMediaLive(config aws.Config) (resources resourceMap) {
	client := medialive.New(config)

	mediaLiveChannelNames := getMediaLiveChannelNames(client)
	mediaLiveInputNames := getMediaLiveInputNames(client)
	mediaLiveInputSecurityGroupIDs := getMediaLiveInputSecurityGroupIDs(client)

	resources = resourceMap{
		mediaLiveChannel:            mediaLiveChannelNames,
		mediaLiveInput:              mediaLiveInputNames,
		mediaLiveInputSecurityGroup: mediaLiveInputSecurityGroupIDs,
	}
	return
}

func getMediaLiveChannelNames(client *medialive.Client) (resources []string) {
	req := client.ListChannelsRequest(&medialive.ListChannelsInput{})
	p := medialive.NewListChannelsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Channels {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getMediaLiveInputNames(client *medialive.Client) (resources []string) {
	req := client.ListInputsRequest(&medialive.ListInputsInput{})
	p := medialive.NewListInputsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Inputs {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getMediaLiveInputSecurityGroupIDs(client *medialive.Client) (resources []string) {
	req := client.ListInputSecurityGroupsRequest(&medialive.ListInputSecurityGroupsInput{})
	p := medialive.NewListInputSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.InputSecurityGroups {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

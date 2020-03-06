package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

func getMediaLive(config aws.Config) (resources resourceMap) {
	client := medialive.New(config)
	resources = reduce(
		getMediaLiveChannel(client).unwrap(mediaLiveChannel),
		getMediaLiveInput(client).unwrap(mediaLiveInput),
		getMediaLiveInputSecurityGroup(client).unwrap(mediaLiveInputSecurityGroup),
	)
	return
}

func getMediaLiveChannel(client *medialive.Client) (r resourceSliceError) {
	req := client.ListChannelsRequest(&medialive.ListChannelsInput{})
	p := medialive.NewListChannelsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Channels {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getMediaLiveInput(client *medialive.Client) (r resourceSliceError) {
	req := client.ListInputsRequest(&medialive.ListInputsInput{})
	p := medialive.NewListInputsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Inputs {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getMediaLiveInputSecurityGroup(client *medialive.Client) (r resourceSliceError) {
	req := client.ListInputSecurityGroupsRequest(&medialive.ListInputSecurityGroupsInput{})
	p := medialive.NewListInputSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.InputSecurityGroups {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

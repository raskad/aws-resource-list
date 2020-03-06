package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/groundstation"
)

func getGroundStation(config aws.Config) (resources resourceMap) {
	client := groundstation.New(config)

	resources = reduce(
		getGroundStationConfig(client).unwrap(groundStationConfig),
		getGroundStationDataflowEndpointGroup(client).unwrap(groundStationDataflowEndpointGroup),
		getGroundStationMissionProfile(client).unwrap(groundStationMissionProfile),
	)
	return
}

func getGroundStationConfig(client *groundstation.Client) (r resourceSliceError) {
	req := client.ListConfigsRequest(&groundstation.ListConfigsInput{})
	p := groundstation.NewListConfigsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ConfigList {
			r.resources = append(r.resources, *resource.ConfigId)
		}
	}
	r.err = p.Err()
	return
}

func getGroundStationDataflowEndpointGroup(client *groundstation.Client) (r resourceSliceError) {
	req := client.ListDataflowEndpointGroupsRequest(&groundstation.ListDataflowEndpointGroupsInput{})
	p := groundstation.NewListDataflowEndpointGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DataflowEndpointGroupList {
			r.resources = append(r.resources, *resource.DataflowEndpointGroupId)
		}
	}
	r.err = p.Err()
	return
}

func getGroundStationMissionProfile(client *groundstation.Client) (r resourceSliceError) {
	req := client.ListMissionProfilesRequest(&groundstation.ListMissionProfilesInput{})
	p := groundstation.NewListMissionProfilesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.MissionProfileList {
			r.resources = append(r.resources, *resource.MissionProfileId)
		}
	}
	r.err = p.Err()
	return
}

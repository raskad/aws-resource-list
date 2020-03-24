package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/groundstation"
)

func getGroundStation(config aws.Config) (resources resourceMap) {
	client := groundstation.New(config)

	groundStationConfigIDs := getGroundStationConfigIDs(client)
	groundStationDataflowEndpointGroupIDs := getGroundStationDataflowEndpointGroupIDs(client)
	groundStationMissionProfileIDs := getGroundStationMissionProfileIDs(client)

	resources = resourceMap{
		groundStationConfig:                groundStationConfigIDs,
		groundStationDataflowEndpointGroup: groundStationDataflowEndpointGroupIDs,
		groundStationMissionProfile:        groundStationMissionProfileIDs,
	}
	return
}

func getGroundStationConfigIDs(client *groundstation.Client) (resources []string) {
	req := client.ListConfigsRequest(&groundstation.ListConfigsInput{})
	p := groundstation.NewListConfigsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ConfigList {
			resources = append(resources, *resource.ConfigId)
		}
	}
	return
}

func getGroundStationDataflowEndpointGroupIDs(client *groundstation.Client) (resources []string) {
	req := client.ListDataflowEndpointGroupsRequest(&groundstation.ListDataflowEndpointGroupsInput{})
	p := groundstation.NewListDataflowEndpointGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.DataflowEndpointGroupList {
			resources = append(resources, *resource.DataflowEndpointGroupId)
		}
	}
	return
}

func getGroundStationMissionProfileIDs(client *groundstation.Client) (resources []string) {
	req := client.ListMissionProfilesRequest(&groundstation.ListMissionProfilesInput{})
	p := groundstation.NewListMissionProfilesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.MissionProfileList {
			resources = append(resources, *resource.MissionProfileId)
		}
	}
	return
}

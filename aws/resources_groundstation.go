package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/groundstation"
)

func getGroundStation(session *session.Session) (resources resourceMap) {
	client := groundstation.New(session)

	resources = reduce(
		getGroundStationConfig(client).unwrap(groundStationConfig),
		getGroundStationDataflowEndpointGroup(client).unwrap(groundStationDataflowEndpointGroup),
		getGroundStationMissionProfile(client).unwrap(groundStationMissionProfile),
	)
	return
}

func getGroundStationConfig(client *groundstation.GroundStation) (r resourceSliceError) {
	r.err = client.ListConfigsPages(&groundstation.ListConfigsInput{}, func(page *groundstation.ListConfigsOutput, lastPage bool) bool {
		for _, resource := range page.ConfigList {
			r.resources = append(r.resources, *resource.ConfigId)
		}
		return true
	})
	return
}

func getGroundStationDataflowEndpointGroup(client *groundstation.GroundStation) (r resourceSliceError) {
	r.err = client.ListDataflowEndpointGroupsPages(&groundstation.ListDataflowEndpointGroupsInput{}, func(page *groundstation.ListDataflowEndpointGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DataflowEndpointGroupList {
			r.resources = append(r.resources, *resource.DataflowEndpointGroupId)
		}
		return true
	})
	return
}

func getGroundStationMissionProfile(client *groundstation.GroundStation) (r resourceSliceError) {
	r.err = client.ListMissionProfilesPages(&groundstation.ListMissionProfilesInput{}, func(page *groundstation.ListMissionProfilesOutput, lastPage bool) bool {
		for _, resource := range page.MissionProfileList {
			r.resources = append(r.resources, *resource.MissionProfileId)
		}
		return true
	})
	return
}

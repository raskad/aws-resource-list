package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/robomaker"
)

func getRoboMaker(config aws.Config) (resources resourceMap) {
	client := robomaker.New(config)

	roboMakerFleetNames := getRoboMakerFleetNames(client)
	roboMakerRobotNames := getRoboMakerRobotNames(client)
	roboMakerRobotApplicationNames := getRoboMakerRobotApplicationNames(client)
	roboMakerSimulationApplicationNames := getRoboMakerSimulationApplicationNames(client)

	resources = resourceMap{
		roboMakerFleet:                 roboMakerFleetNames,
		roboMakerRobot:                 roboMakerRobotNames,
		roboMakerRobotApplication:      roboMakerRobotApplicationNames,
		roboMakerSimulationApplication: roboMakerSimulationApplicationNames,
	}
	return
}

func getRoboMakerFleetNames(client *robomaker.Client) (resources []string) {
	req := client.ListFleetsRequest(&robomaker.ListFleetsInput{})
	p := robomaker.NewListFleetsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.FleetDetails {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getRoboMakerRobotNames(client *robomaker.Client) (resources []string) {
	req := client.ListRobotsRequest(&robomaker.ListRobotsInput{})
	p := robomaker.NewListRobotsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Robots {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getRoboMakerRobotApplicationNames(client *robomaker.Client) (resources []string) {
	req := client.ListRobotApplicationsRequest(&robomaker.ListRobotApplicationsInput{})
	p := robomaker.NewListRobotApplicationsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.RobotApplicationSummaries {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getRoboMakerSimulationApplicationNames(client *robomaker.Client) (resources []string) {
	req := client.ListSimulationApplicationsRequest(&robomaker.ListSimulationApplicationsInput{})
	p := robomaker.NewListSimulationApplicationsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.SimulationApplicationSummaries {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

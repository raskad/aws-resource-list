package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/robomaker"
)

func getRoboMaker(config aws.Config) (resources resourceMap) {
	client := robomaker.New(config)
	resources = reduce(
		getRoboMakerFleet(client).unwrap(roboMakerFleet),
		getRoboMakerRobot(client).unwrap(roboMakerRobot),
		getRoboMakerRobotApplication(client).unwrap(roboMakerRobotApplication),
		getRoboMakerSimulationApplication(client).unwrap(roboMakerSimulationApplication),
	)
	return
}

func getRoboMakerFleet(client *robomaker.Client) (r resourceSliceError) {
	req := client.ListFleetsRequest(&robomaker.ListFleetsInput{})
	p := robomaker.NewListFleetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.FleetDetails {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getRoboMakerRobot(client *robomaker.Client) (r resourceSliceError) {
	req := client.ListRobotsRequest(&robomaker.ListRobotsInput{})
	p := robomaker.NewListRobotsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Robots {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getRoboMakerRobotApplication(client *robomaker.Client) (r resourceSliceError) {
	req := client.ListRobotApplicationsRequest(&robomaker.ListRobotApplicationsInput{})
	p := robomaker.NewListRobotApplicationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.RobotApplicationSummaries {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getRoboMakerSimulationApplication(client *robomaker.Client) (r resourceSliceError) {
	req := client.ListSimulationApplicationsRequest(&robomaker.ListSimulationApplicationsInput{})
	p := robomaker.NewListSimulationApplicationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.SimulationApplicationSummaries {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

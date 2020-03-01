package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/robomaker"
)

func getRoboMaker(session *session.Session) (resources resourceMap) {
	client := robomaker.New(session)
	resources = reduce(
		getRoboMakerFleet(client).unwrap(roboMakerFleet),
		getRoboMakerRobot(client).unwrap(roboMakerRobot),
		getRoboMakerRobotApplication(client).unwrap(roboMakerRobotApplication),
		getRoboMakerSimulationApplication(client).unwrap(roboMakerSimulationApplication),
	)
	return
}

func getRoboMakerFleet(client *robomaker.RoboMaker) (r resourceSliceError) {
	r.err = client.ListFleetsPages(&robomaker.ListFleetsInput{}, func(page *robomaker.ListFleetsOutput, lastPage bool) bool {
		for _, resource := range page.FleetDetails {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getRoboMakerRobot(client *robomaker.RoboMaker) (r resourceSliceError) {
	r.err = client.ListRobotsPages(&robomaker.ListRobotsInput{}, func(page *robomaker.ListRobotsOutput, lastPage bool) bool {
		for _, resource := range page.Robots {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getRoboMakerRobotApplication(client *robomaker.RoboMaker) (r resourceSliceError) {
	r.err = client.ListRobotApplicationsPages(&robomaker.ListRobotApplicationsInput{}, func(page *robomaker.ListRobotApplicationsOutput, lastPage bool) bool {
		for _, resource := range page.RobotApplicationSummaries {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getRoboMakerSimulationApplication(client *robomaker.RoboMaker) (r resourceSliceError) {
	r.err = client.ListSimulationApplicationsPages(&robomaker.ListSimulationApplicationsInput{}, func(page *robomaker.ListSimulationApplicationsOutput, lastPage bool) bool {
		for _, resource := range page.SimulationApplicationSummaries {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

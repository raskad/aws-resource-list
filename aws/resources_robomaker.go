package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/robomaker"
)

func getRoboMaker(session *session.Session) (resources resourceMap) {
	client := robomaker.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		roboMakerFleet:                 getRoboMakerFleet(client),
		roboMakerRobot:                 getRoboMakerRobot(client),
		roboMakerRobotApplication:      getRoboMakerRobotApplication(client),
		roboMakerSimulationApplication: getRoboMakerSimulationApplication(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getRoboMakerFleet(client *robomaker.RoboMaker) (r resourceSliceError) {
	r.err = client.ListFleetsPages(&robomaker.ListFleetsInput{}, func(page *robomaker.ListFleetsOutput, lastPage bool) bool {
		logDebug("Listing RoboMakerFleet resources page. Remaining pages", page.NextToken)
		for _, resource := range page.FleetDetails {
			logDebug("Got RoboMakerFleet resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getRoboMakerRobot(client *robomaker.RoboMaker) (r resourceSliceError) {
	r.err = client.ListRobotsPages(&robomaker.ListRobotsInput{}, func(page *robomaker.ListRobotsOutput, lastPage bool) bool {
		logDebug("Listing RoboMakerRobot resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Robots {
			logDebug("Got RoboMakerRobot resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getRoboMakerRobotApplication(client *robomaker.RoboMaker) (r resourceSliceError) {
	r.err = client.ListRobotApplicationsPages(&robomaker.ListRobotApplicationsInput{}, func(page *robomaker.ListRobotApplicationsOutput, lastPage bool) bool {
		logDebug("Listing RoboMakerRobotApplication resources page. Remaining pages", page.NextToken)
		for _, resource := range page.RobotApplicationSummaries {
			logDebug("Got RoboMakerRobotApplication resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getRoboMakerSimulationApplication(client *robomaker.RoboMaker) (r resourceSliceError) {
	r.err = client.ListSimulationApplicationsPages(&robomaker.ListSimulationApplicationsInput{}, func(page *robomaker.ListSimulationApplicationsOutput, lastPage bool) bool {
		logDebug("Listing RoboMakerSimulationApplication resources page. Remaining pages", page.NextToken)
		for _, resource := range page.SimulationApplicationSummaries {
			logDebug("Got RoboMakerSimulationApplication resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

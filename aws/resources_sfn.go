package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

func getSfn(session *session.Session) (resources resourceMap) {
	client := sfn.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		stepFunctionsActivity:     getStepFunctionsActivity(client),
		stepFunctionsStateMachine: getStepFunctionsStateMachine(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getStepFunctionsActivity(client *sfn.SFN) (r resourceSliceError) {
	r.err = client.ListActivitiesPages(&sfn.ListActivitiesInput{}, func(page *sfn.ListActivitiesOutput, lastPage bool) bool {
		logDebug("List StepFunctionsActivity resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Activities {
			logDebug("Got StepFunctionsActivity resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getStepFunctionsStateMachine(client *sfn.SFN) (r resourceSliceError) {
	r.err = client.ListStateMachinesPages(&sfn.ListStateMachinesInput{}, func(page *sfn.ListStateMachinesOutput, lastPage bool) bool {
		logDebug("List StepFunctionsStateMachine resources page. Remaining pages", page.NextToken)
		for _, resource := range page.StateMachines {
			logDebug("Got StepFunctionsStateMachine resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}
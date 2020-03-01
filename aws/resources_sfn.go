package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

func getSfn(session *session.Session) (resources resourceMap) {
	client := sfn.New(session)
	resources = reduce(
		getStepFunctionsActivity(client).unwrap(stepFunctionsActivity),
		getStepFunctionsStateMachine(client).unwrap(stepFunctionsStateMachine),
	)
	return
}

func getStepFunctionsActivity(client *sfn.SFN) (r resourceSliceError) {
	r.err = client.ListActivitiesPages(&sfn.ListActivitiesInput{}, func(page *sfn.ListActivitiesOutput, lastPage bool) bool {
		for _, resource := range page.Activities {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getStepFunctionsStateMachine(client *sfn.SFN) (r resourceSliceError) {
	r.err = client.ListStateMachinesPages(&sfn.ListStateMachinesInput{}, func(page *sfn.ListStateMachinesOutput, lastPage bool) bool {
		for _, resource := range page.StateMachines {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

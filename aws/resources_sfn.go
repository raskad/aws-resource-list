package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func getSfn(config aws.Config) (resources resourceMap) {
	client := sfn.New(config)
	resources = reduce(
		getStepFunctionsActivity(client).unwrap(stepFunctionsActivity),
		getStepFunctionsStateMachine(client).unwrap(stepFunctionsStateMachine),
	)
	return
}

func getStepFunctionsActivity(client *sfn.Client) (r resourceSliceError) {
	req := client.ListActivitiesRequest(&sfn.ListActivitiesInput{})
	p := sfn.NewListActivitiesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Activities {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getStepFunctionsStateMachine(client *sfn.Client) (r resourceSliceError) {
	req := client.ListStateMachinesRequest(&sfn.ListStateMachinesInput{})
	p := sfn.NewListStateMachinesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.StateMachines {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

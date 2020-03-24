package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func getSfn(config aws.Config) (resources resourceMap) {
	client := sfn.New(config)

	stepFunctionsActivityNames := getStepFunctionsActivityNames(client)
	stepFunctionsStateMachineNames := getStepFunctionsStateMachineNames(client)

	resources = resourceMap{
		stepFunctionsActivity:     stepFunctionsActivityNames,
		stepFunctionsStateMachine: stepFunctionsStateMachineNames,
	}
	return
}

func getStepFunctionsActivityNames(client *sfn.Client) (resources []string) {
	req := client.ListActivitiesRequest(&sfn.ListActivitiesInput{})
	p := sfn.NewListActivitiesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Activities {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getStepFunctionsStateMachineNames(client *sfn.Client) (resources []string) {
	req := client.ListStateMachinesRequest(&sfn.ListStateMachinesInput{})
	p := sfn.NewListStateMachinesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.StateMachines {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

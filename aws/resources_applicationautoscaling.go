package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
)

func getApplicationAutoScaling(config aws.Config) (resources awsResourceMap) {
	client := applicationautoscaling.New(config)

	applicationAutoScalingScheduledActionNames := getApplicationAutoScalingScheduledActionNames(client)

	resources = awsResourceMap{
		applicationAutoScalingScheduledAction: applicationAutoScalingScheduledActionNames,
	}
	return
}

func getApplicationAutoScalingScheduledActionNames(client *applicationautoscaling.Client) (resources []string) {
	req := client.DescribeScheduledActionsRequest(&applicationautoscaling.DescribeScheduledActionsInput{})
	p := applicationautoscaling.NewDescribeScheduledActionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ScheduledActions {
			resources = append(resources, *resource.ScheduledActionName)
		}
	}
	return
}

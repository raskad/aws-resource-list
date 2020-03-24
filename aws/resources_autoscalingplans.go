package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
)

func getAutoScalingPlans(config aws.Config) (resources resourceMap) {
	client := autoscalingplans.New(config)

	autoScalingPlansScalingPlanNames := getAutoScalingPlansScalingPlanNames(client)

	resources = resourceMap{
		autoScalingPlansScalingPlan: autoScalingPlansScalingPlanNames,
	}
	return
}

func getAutoScalingPlansScalingPlanNames(client *autoscalingplans.Client) (resources []string) {
	input := autoscalingplans.DescribeScalingPlansInput{}
	for {
		page, err := client.DescribeScalingPlansRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.ScalingPlans {
			resources = append(resources, *resource.ScalingPlanName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

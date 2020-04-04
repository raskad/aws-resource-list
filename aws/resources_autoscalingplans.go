package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
)

func getAutoScalingPlans(config aws.Config) (resources awsResourceMap) {
	client := autoscalingplans.New(config)

	autoScalingPlansScalingPlanNames := getAutoScalingPlansScalingPlanNames(client)

	resources = awsResourceMap{
		autoScalingPlansScalingPlan: autoScalingPlansScalingPlanNames,
	}
	return
}

func getAutoScalingPlansScalingPlanNames(client *autoscalingplans.Client) (resources []string) {
	input := autoscalingplans.DescribeScalingPlansInput{}
	for {
		page, err := client.DescribeScalingPlansRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ScalingPlans {
			resources = append(resources, *resource.ScalingPlanName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

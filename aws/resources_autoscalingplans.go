package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
)

func getAutoScalingPlans(config aws.Config) (resources resourceMap) {
	client := autoscalingplans.New(config)
	resources = reduce(
		getAutoScalingPlansScalingPlan(client).unwrap(autoScalingPlansScalingPlan),
	)
	return
}

func getAutoScalingPlansScalingPlan(client *autoscalingplans.Client) (r resourceSliceError) {
	input := autoscalingplans.DescribeScalingPlansInput{}
	for {
		page, err := client.DescribeScalingPlansRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ScalingPlans {
			r.resources = append(r.resources, *resource.ScalingPlanName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

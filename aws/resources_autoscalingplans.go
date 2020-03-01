package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
)

func getAutoScalingPlans(session *session.Session) (resources resourceMap) {
	client := autoscalingplans.New(session)
	resources = reduce(
		getAutoScalingPlansScalingPlan(client).unwrap(autoScalingPlansScalingPlan),
	)
	return
}

func getAutoScalingPlansScalingPlan(client *autoscalingplans.AutoScalingPlans) (r resourceSliceError) {
	input := autoscalingplans.DescribeScalingPlansInput{}
	for {
		page, err := client.DescribeScalingPlans(&input)
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

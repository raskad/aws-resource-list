package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
)

func getAutoScalingPlans(session *session.Session) (resources resourceMap) {
	client := autoscalingplans.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		autoScalingPlansScalingPlan: getAutoScalingPlansScalingPlan(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAutoScalingPlansScalingPlan(client *autoscalingplans.AutoScalingPlans) (r resourceSliceError) {
	logDebug("Listing AutoScalingPlansScalingPlan resources")
	input := autoscalingplans.DescribeScalingPlansInput{}
	for {
		page, err := client.DescribeScalingPlans(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ScalingPlans {
			logDebug("Got AutoScalingPlansScalingPlan resource with PhysicalResourceId", *resource.ScalingPlanName)
			r.resources = append(r.resources, *resource.ScalingPlanName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

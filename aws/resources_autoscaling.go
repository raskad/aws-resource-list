package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func getAutoScaling(session *session.Session) (resources resourceMap) {
	client := autoscaling.New(session)
	resources = reduce(
		getAutoScalingAutoScalingGroup(client).unwrap(autoScalingAutoScalingGroup),
		getAutoScalingLaunchConfiguration(client).unwrap(autoScalingLaunchConfiguration),
		getAutoScalingScalingPolicy(client).unwrap(autoScalingScalingPolicy),
		getAutoScalingScheduledAction(client).unwrap(autoScalingScheduledAction),
	)
	return
}

func getAutoScalingAutoScalingGroup(client *autoscaling.AutoScaling) (r resourceSliceError) {
	r.err = client.DescribeAutoScalingGroupsPages(&autoscaling.DescribeAutoScalingGroupsInput{}, func(page *autoscaling.DescribeAutoScalingGroupsOutput, lastPage bool) bool {
		for _, resource := range page.AutoScalingGroups {
			r.resources = append(r.resources, *resource.AutoScalingGroupName)
		}
		return true
	})
	return
}

func getAutoScalingLaunchConfiguration(client *autoscaling.AutoScaling) (r resourceSliceError) {
	r.err = client.DescribeLaunchConfigurationsPages(&autoscaling.DescribeLaunchConfigurationsInput{}, func(page *autoscaling.DescribeLaunchConfigurationsOutput, lastPage bool) bool {
		for _, resource := range page.LaunchConfigurations {
			r.resources = append(r.resources, *resource.LaunchConfigurationName)
		}
		return true
	})
	return
}

func getAutoScalingScalingPolicy(client *autoscaling.AutoScaling) (r resourceSliceError) {
	r.err = client.DescribePoliciesPages(&autoscaling.DescribePoliciesInput{}, func(page *autoscaling.DescribePoliciesOutput, lastPage bool) bool {
		for _, resource := range page.ScalingPolicies {
			r.resources = append(r.resources, *resource.PolicyName)
		}
		return true
	})
	return
}

func getAutoScalingScheduledAction(client *autoscaling.AutoScaling) (r resourceSliceError) {
	r.err = client.DescribeScheduledActionsPages(&autoscaling.DescribeScheduledActionsInput{}, func(page *autoscaling.DescribeScheduledActionsOutput, lastPage bool) bool {
		for _, resource := range page.ScheduledUpdateGroupActions {
			r.resources = append(r.resources, *resource.ScheduledActionName)
		}
		return true
	})
	return
}

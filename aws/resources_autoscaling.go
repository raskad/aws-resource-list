package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func getAutoScaling(session *session.Session) (resources resourceMap) {
	client := autoscaling.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		autoScalingAutoScalingGroup:    getAutoScalingAutoScalingGroup(client),
		autoScalingLaunchConfiguration: getAutoScalingLaunchConfiguration(client),
		autoScalingScalingPolicy:       getAutoScalingScalingPolicy(client),
		autoScalingScheduledAction:     getAutoScalingScheduledAction(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAutoScalingAutoScalingGroup(client *autoscaling.AutoScaling) (r resourceSliceError) {
	r.err = client.DescribeAutoScalingGroupsPages(&autoscaling.DescribeAutoScalingGroupsInput{}, func(page *autoscaling.DescribeAutoScalingGroupsOutput, lastPage bool) bool {
		logDebug("List AutoScalingAutoScalingGroup resources page. Remaining pages", page.NextToken)
		for _, resource := range page.AutoScalingGroups {
			logDebug("Got AutoScalingAutoScalingGroup resource with PhysicalResourceId", *resource.AutoScalingGroupName)
			r.resources = append(r.resources, *resource.AutoScalingGroupName)
		}
		return true
	})
	return
}

func getAutoScalingLaunchConfiguration(client *autoscaling.AutoScaling) (r resourceSliceError) {
	r.err = client.DescribeLaunchConfigurationsPages(&autoscaling.DescribeLaunchConfigurationsInput{}, func(page *autoscaling.DescribeLaunchConfigurationsOutput, lastPage bool) bool {
		logDebug("List AutoScalingLaunchConfiguration resources page. Remaining pages", page.NextToken)
		for _, resource := range page.LaunchConfigurations {
			logDebug("Got AutoScalingLaunchConfiguration resource with PhysicalResourceId", *resource.LaunchConfigurationName)
			r.resources = append(r.resources, *resource.LaunchConfigurationName)
		}
		return true
	})
	return
}

func getAutoScalingScalingPolicy(client *autoscaling.AutoScaling) (r resourceSliceError) {
	r.err = client.DescribePoliciesPages(&autoscaling.DescribePoliciesInput{}, func(page *autoscaling.DescribePoliciesOutput, lastPage bool) bool {
		logDebug("List AutoScalingScalingPolicy resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ScalingPolicies {
			logDebug("Got AutoScalingScalingPolicy resource with PhysicalResourceId", *resource.PolicyName)
			r.resources = append(r.resources, *resource.PolicyName)
		}
		return true
	})
	return
}

func getAutoScalingScheduledAction(client *autoscaling.AutoScaling) (r resourceSliceError) {
	r.err = client.DescribeScheduledActionsPages(&autoscaling.DescribeScheduledActionsInput{}, func(page *autoscaling.DescribeScheduledActionsOutput, lastPage bool) bool {
		logDebug("List AutoScalingScheduledAction resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ScheduledUpdateGroupActions {
			logDebug("Got AutoScalingScheduledAction resource with PhysicalResourceId", *resource.ScheduledActionName)
			r.resources = append(r.resources, *resource.ScheduledActionName)
		}
		return true
	})
	return
}

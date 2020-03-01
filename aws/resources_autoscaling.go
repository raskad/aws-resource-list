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
	logDebug("Listing AutoScalingAutoScalingGroup resources")
	r.err = client.DescribeAutoScalingGroupsPages(&autoscaling.DescribeAutoScalingGroupsInput{}, func(page *autoscaling.DescribeAutoScalingGroupsOutput, lastPage bool) bool {
		for _, resource := range page.AutoScalingGroups {
			logDebug("Got AutoScalingAutoScalingGroup resource with PhysicalResourceId", *resource.AutoScalingGroupName)
			r.resources = append(r.resources, *resource.AutoScalingGroupName)
		}
		return true
	})
	return
}

func getAutoScalingLaunchConfiguration(client *autoscaling.AutoScaling) (r resourceSliceError) {
	logDebug("Listing AutoScalingLaunchConfiguration resources")
	r.err = client.DescribeLaunchConfigurationsPages(&autoscaling.DescribeLaunchConfigurationsInput{}, func(page *autoscaling.DescribeLaunchConfigurationsOutput, lastPage bool) bool {
		for _, resource := range page.LaunchConfigurations {
			logDebug("Got AutoScalingLaunchConfiguration resource with PhysicalResourceId", *resource.LaunchConfigurationName)
			r.resources = append(r.resources, *resource.LaunchConfigurationName)
		}
		return true
	})
	return
}

func getAutoScalingScalingPolicy(client *autoscaling.AutoScaling) (r resourceSliceError) {
	logDebug("Listing AutoScalingScalingPolicy resources")
	r.err = client.DescribePoliciesPages(&autoscaling.DescribePoliciesInput{}, func(page *autoscaling.DescribePoliciesOutput, lastPage bool) bool {
		for _, resource := range page.ScalingPolicies {
			logDebug("Got AutoScalingScalingPolicy resource with PhysicalResourceId", *resource.PolicyName)
			r.resources = append(r.resources, *resource.PolicyName)
		}
		return true
	})
	return
}

func getAutoScalingScheduledAction(client *autoscaling.AutoScaling) (r resourceSliceError) {
	logDebug("Listing AutoScalingScheduledAction resources")
	r.err = client.DescribeScheduledActionsPages(&autoscaling.DescribeScheduledActionsInput{}, func(page *autoscaling.DescribeScheduledActionsOutput, lastPage bool) bool {
		for _, resource := range page.ScheduledUpdateGroupActions {
			logDebug("Got AutoScalingScheduledAction resource with PhysicalResourceId", *resource.ScheduledActionName)
			r.resources = append(r.resources, *resource.ScheduledActionName)
		}
		return true
	})
	return
}

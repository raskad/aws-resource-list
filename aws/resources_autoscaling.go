package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func getAutoScaling(config aws.Config) (resources resourceMap) {
	client := autoscaling.New(config)

	autoScalingAutoScalingGroupNames := getAutoScalingAutoScalingGroupNames(client)
	autoScalingLaunchConfigurationNames := getAutoScalingLaunchConfigurationNames(client)
	autoScalingScalingPolicyNames := getAutoScalingScalingPolicyNames(client)
	autoScalingScheduledActionNames := getAutoScalingScheduledActionNames(client)

	resources = resourceMap{
		autoScalingAutoScalingGroup:    autoScalingAutoScalingGroupNames,
		autoScalingLaunchConfiguration: autoScalingLaunchConfigurationNames,
		autoScalingScalingPolicy:       autoScalingScalingPolicyNames,
		autoScalingScheduledAction:     autoScalingScheduledActionNames,
	}
	return
}

func getAutoScalingAutoScalingGroupNames(client *autoscaling.Client) (resources []string) {
	req := client.DescribeAutoScalingGroupsRequest(&autoscaling.DescribeAutoScalingGroupsInput{})
	p := autoscaling.NewDescribeAutoScalingGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.AutoScalingGroups {
			resources = append(resources, *resource.AutoScalingGroupName)
		}
	}
	return
}

func getAutoScalingLaunchConfigurationNames(client *autoscaling.Client) (resources []string) {
	req := client.DescribeLaunchConfigurationsRequest(&autoscaling.DescribeLaunchConfigurationsInput{})
	p := autoscaling.NewDescribeLaunchConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.LaunchConfigurations {
			resources = append(resources, *resource.LaunchConfigurationName)
		}
	}
	return
}

func getAutoScalingScalingPolicyNames(client *autoscaling.Client) (resources []string) {
	req := client.DescribePoliciesRequest(&autoscaling.DescribePoliciesInput{})
	p := autoscaling.NewDescribePoliciesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ScalingPolicies {
			resources = append(resources, *resource.PolicyName)
		}
	}
	return
}

func getAutoScalingScheduledActionNames(client *autoscaling.Client) (resources []string) {
	req := client.DescribeScheduledActionsRequest(&autoscaling.DescribeScheduledActionsInput{})
	p := autoscaling.NewDescribeScheduledActionsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ScheduledUpdateGroupActions {
			resources = append(resources, *resource.ScheduledActionName)
		}
	}
	return
}

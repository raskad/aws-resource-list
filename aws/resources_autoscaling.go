package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func getAutoScaling(config aws.Config) (resources resourceMap) {
	client := autoscaling.New(config)
	resources = reduce(
		getAutoScalingAutoScalingGroup(client).unwrap(autoScalingAutoScalingGroup),
		getAutoScalingLaunchConfiguration(client).unwrap(autoScalingLaunchConfiguration),
		getAutoScalingScalingPolicy(client).unwrap(autoScalingScalingPolicy),
		getAutoScalingScheduledAction(client).unwrap(autoScalingScheduledAction),
	)
	return
}

func getAutoScalingAutoScalingGroup(client *autoscaling.Client) (r resourceSliceError) {
	req := client.DescribeAutoScalingGroupsRequest(&autoscaling.DescribeAutoScalingGroupsInput{})
	p := autoscaling.NewDescribeAutoScalingGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.AutoScalingGroups {
			r.resources = append(r.resources, *resource.AutoScalingGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getAutoScalingLaunchConfiguration(client *autoscaling.Client) (r resourceSliceError) {
	req := client.DescribeLaunchConfigurationsRequest(&autoscaling.DescribeLaunchConfigurationsInput{})
	p := autoscaling.NewDescribeLaunchConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.LaunchConfigurations {
			r.resources = append(r.resources, *resource.LaunchConfigurationName)
		}
	}
	r.err = p.Err()
	return
}

func getAutoScalingScalingPolicy(client *autoscaling.Client) (r resourceSliceError) {
	req := client.DescribePoliciesRequest(&autoscaling.DescribePoliciesInput{})
	p := autoscaling.NewDescribePoliciesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ScalingPolicies {
			r.resources = append(r.resources, *resource.PolicyName)
		}
	}
	r.err = p.Err()
	return
}

func getAutoScalingScheduledAction(client *autoscaling.Client) (r resourceSliceError) {
	req := client.DescribeScheduledActionsRequest(&autoscaling.DescribeScheduledActionsInput{})
	p := autoscaling.NewDescribeScheduledActionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ScheduledUpdateGroupActions {
			r.resources = append(r.resources, *resource.ScheduledActionName)
		}
	}
	r.err = p.Err()
	return
}

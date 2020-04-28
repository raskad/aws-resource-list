package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func getCloudwatchLogs(config aws.Config) (resources awsResourceMap) {
	client := cloudwatchlogs.New(config)

	cloudwatchLogsLogGroupNames := getCloudwatchLogsLogGroupNames(client)
	cloudwatchLogsDestinationNames := getCloudwatchLogsDestinationNames(client)
	cloudwatchLogsMetricFilterNames := getCloudwatchLogsMetricFilterNames(client)
	cloudwatchLogsSubscriptionFilterNames := getCloudwatchLogsSubscriptionFilterNames(client, cloudwatchLogsLogGroupNames)
	cloudwatchLogsResourcePolicyNames := getCloudwatchLogsResourcePolicyNames(client)

	resources = awsResourceMap{
		cloudwatchLogsLogGroup:           cloudwatchLogsLogGroupNames,
		cloudwatchLogsDestination:        cloudwatchLogsDestinationNames,
		cloudwatchLogsMetricFilter:       cloudwatchLogsMetricFilterNames,
		cloudwatchLogsSubscriptionFilter: cloudwatchLogsSubscriptionFilterNames,
		cloudwatchLogsResourcePolicy:     cloudwatchLogsResourcePolicyNames,
	}
	return
}

func getCloudwatchLogsDestinationNames(client *cloudwatchlogs.Client) (resources []string) {
	req := client.DescribeDestinationsRequest(&cloudwatchlogs.DescribeDestinationsInput{})
	p := cloudwatchlogs.NewDescribeDestinationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Destinations {
			resources = append(resources, *resource.DestinationName)
		}
	}
	return
}

func getCloudwatchLogsLogGroupNames(client *cloudwatchlogs.Client) (resources []string) {
	req := client.DescribeLogGroupsRequest(&cloudwatchlogs.DescribeLogGroupsInput{})
	p := cloudwatchlogs.NewDescribeLogGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.LogGroups {
			resources = append(resources, *resource.LogGroupName)
		}
	}
	return
}

func getCloudwatchLogsMetricFilterNames(client *cloudwatchlogs.Client) (resources []string) {
	req := client.DescribeMetricFiltersRequest(&cloudwatchlogs.DescribeMetricFiltersInput{})
	p := cloudwatchlogs.NewDescribeMetricFiltersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.MetricFilters {
			resources = append(resources, *resource.FilterName)
		}
	}
	return
}

func getCloudwatchLogsSubscriptionFilterNames(client *cloudwatchlogs.Client, logGroupNames []string) (resources []string) {
	for _, logGroupName := range logGroupNames {
		req := client.DescribeSubscriptionFiltersRequest(&cloudwatchlogs.DescribeSubscriptionFiltersInput{
			LogGroupName: aws.String(logGroupName),
		})
		p := cloudwatchlogs.NewDescribeSubscriptionFiltersPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.SubscriptionFilters {
				resources = append(resources, *resource.FilterName)
			}
		}
	}
	return
}

func getCloudwatchLogsResourcePolicyNames(client *cloudwatchlogs.Client) (resources []string) {
	input := cloudwatchlogs.DescribeResourcePoliciesInput{}
	for {
		page, err := client.DescribeResourcePoliciesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ResourcePolicies {
			resources = append(resources, *resource.PolicyName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

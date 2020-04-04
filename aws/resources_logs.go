package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func getCloudwatchLogs(config aws.Config) (resources awsResourceMap) {
	client := cloudwatchlogs.New(config)

	logsLogGroupNames := getLogsLogGroupNames(client)
	logsDestinationNames := getLogsDestinationNames(client)
	logsMetricFilterNames := getLogsMetricFilterNames(client)
	logsSubscriptionFilterNames := getLogsSubscriptionFilterNames(client, logsLogGroupNames)
	logsResourcePolicyNames := getLogsResourcePolicyNames(client)

	resources = awsResourceMap{
		logsLogGroup:           logsLogGroupNames,
		logsDestination:        logsDestinationNames,
		logsMetricFilter:       logsMetricFilterNames,
		logsSubscriptionFilter: logsSubscriptionFilterNames,
		logsResourcePolicy:     logsResourcePolicyNames,
	}
	return
}

func getLogsDestinationNames(client *cloudwatchlogs.Client) (resources []string) {
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

func getLogsLogGroupNames(client *cloudwatchlogs.Client) (resources []string) {
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

func getLogsMetricFilterNames(client *cloudwatchlogs.Client) (resources []string) {
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

func getLogsSubscriptionFilterNames(client *cloudwatchlogs.Client, logGroupNames []string) (resources []string) {
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

func getLogsResourcePolicyNames(client *cloudwatchlogs.Client) (resources []string) {
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

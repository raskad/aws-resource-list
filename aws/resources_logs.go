package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func getCloudwatchLogs(config aws.Config) (resources resourceMap) {
	client := cloudwatchlogs.New(config)

	logsLogGroupResourceMap := getLogsLogGroup(client).unwrap(logsLogGroup)
	logsLogGroupNames := logsLogGroupResourceMap[logsLogGroup]

	resources = reduce(
		getLogsDestination(client).unwrap(logsDestination),
		logsLogGroupResourceMap,
		getLogsMetricFilter(client).unwrap(logsMetricFilter),
		getLogsSubscriptionFilter(client, logsLogGroupNames).unwrap(logsSubscriptionFilter),
	)
	return
}

func getLogsDestination(client *cloudwatchlogs.Client) (r resourceSliceError) {
	req := client.DescribeDestinationsRequest(&cloudwatchlogs.DescribeDestinationsInput{})
	p := cloudwatchlogs.NewDescribeDestinationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Destinations {
			r.resources = append(r.resources, *resource.DestinationName)
		}
	}
	r.err = p.Err()
	return
}

func getLogsLogGroup(client *cloudwatchlogs.Client) (r resourceSliceError) {
	req := client.DescribeLogGroupsRequest(&cloudwatchlogs.DescribeLogGroupsInput{})
	p := cloudwatchlogs.NewDescribeLogGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.LogGroups {
			r.resources = append(r.resources, *resource.LogGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getLogsMetricFilter(client *cloudwatchlogs.Client) (r resourceSliceError) {
	req := client.DescribeMetricFiltersRequest(&cloudwatchlogs.DescribeMetricFiltersInput{})
	p := cloudwatchlogs.NewDescribeMetricFiltersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.MetricFilters {
			r.resources = append(r.resources, *resource.FilterName)
		}
	}
	r.err = p.Err()
	return
}

func getLogsSubscriptionFilter(client *cloudwatchlogs.Client, logGroupNames []string) (r resourceSliceError) {
	for _, logGroupName := range logGroupNames {
		req := client.DescribeSubscriptionFiltersRequest(&cloudwatchlogs.DescribeSubscriptionFiltersInput{
			LogGroupName: aws.String(logGroupName),
		})
		p := cloudwatchlogs.NewDescribeSubscriptionFiltersPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.SubscriptionFilters {
				r.resources = append(r.resources, *resource.FilterName)
			}
		}
		r.err = p.Err()
		return
	}
	return
}

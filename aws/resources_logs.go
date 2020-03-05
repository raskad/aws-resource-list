package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func getCloudwatchLogs(session *session.Session) (resources resourceMap) {
	client := cloudwatchlogs.New(session)

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

func getLogsDestination(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	r.err = client.DescribeDestinationsPages(&cloudwatchlogs.DescribeDestinationsInput{}, func(page *cloudwatchlogs.DescribeDestinationsOutput, lastPage bool) bool {
		for _, resource := range page.Destinations {
			r.resources = append(r.resources, *resource.DestinationName)
		}
		return true
	})
	return
}

func getLogsLogGroup(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	r.err = client.DescribeLogGroupsPages(&cloudwatchlogs.DescribeLogGroupsInput{}, func(page *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {
		for _, resource := range page.LogGroups {
			r.resources = append(r.resources, *resource.LogGroupName)
		}
		return true
	})
	return
}

func getLogsMetricFilter(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	r.err = client.DescribeMetricFiltersPages(&cloudwatchlogs.DescribeMetricFiltersInput{}, func(page *cloudwatchlogs.DescribeMetricFiltersOutput, lastPage bool) bool {
		for _, resource := range page.MetricFilters {
			r.resources = append(r.resources, *resource.FilterName)
		}
		return true
	})
	return
}

func getLogsSubscriptionFilter(client *cloudwatchlogs.CloudWatchLogs, logGroupNames []string) (r resourceSliceError) {
	for _, logGroupName := range logGroupNames {
		r.err = client.DescribeSubscriptionFiltersPages(&cloudwatchlogs.DescribeSubscriptionFiltersInput{
			LogGroupName: aws.String(logGroupName),
		}, func(page *cloudwatchlogs.DescribeSubscriptionFiltersOutput, lastPage bool) bool {
			for _, resource := range page.SubscriptionFilters {
				r.resources = append(r.resources, *resource.FilterName)
			}
			return true
		})
	}
	return
}

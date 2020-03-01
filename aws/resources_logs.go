package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func getCloudwatchLogs(session *session.Session) (resources resourceMap) {
	client := cloudwatchlogs.New(session)
	resources = reduce(
		getLogsDestination(client).unwrap(logsDestination),
		getLogsLogGroup(client).unwrap(logsLogGroup),
		getLogsLogStream(client).unwrap(logsLogStream),
		getLogsMetricFilter(client).unwrap(logsMetricFilter),
		getLogsSubscriptionFilter(client).unwrap(logsSubscriptionFilter),
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

func getLogsLogStream(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	r.err = client.DescribeLogStreamsPages(&cloudwatchlogs.DescribeLogStreamsInput{}, func(page *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
		for _, resource := range page.LogStreams {
			r.resources = append(r.resources, *resource.LogStreamName)
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

func getLogsSubscriptionFilter(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	r.err = client.DescribeSubscriptionFiltersPages(&cloudwatchlogs.DescribeSubscriptionFiltersInput{}, func(page *cloudwatchlogs.DescribeSubscriptionFiltersOutput, lastPage bool) bool {
		for _, resource := range page.SubscriptionFilters {
			r.resources = append(r.resources, *resource.FilterName)
		}
		return true
	})
	return
}

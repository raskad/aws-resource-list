package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func getCloudwatchLogs(session *session.Session) (resources resourceMap) {
	client := cloudwatchlogs.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		logsDestination:        getLogsDestination(client),
		logsLogGroup:           getLogsLogGroup(client),
		logsLogStream:          getLogsLogStream(client),
		logsMetricFilter:       getLogsMetricFilter(client),
		logsSubscriptionFilter: getLogsSubscriptionFilter(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getLogsDestination(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	logDebug("Listing LogsDestination resources")
	r.err = client.DescribeDestinationsPages(&cloudwatchlogs.DescribeDestinationsInput{}, func(page *cloudwatchlogs.DescribeDestinationsOutput, lastPage bool) bool {
		for _, resource := range page.Destinations {
			logDebug("Got LogsDestination resource with PhysicalResourceId", *resource.DestinationName)
			r.resources = append(r.resources, *resource.DestinationName)
		}
		return true
	})
	return
}

func getLogsLogGroup(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	logDebug("Listing LogsLogGroup resources")
	r.err = client.DescribeLogGroupsPages(&cloudwatchlogs.DescribeLogGroupsInput{}, func(page *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {
		for _, resource := range page.LogGroups {
			logDebug("Got LogsLogGroup resource with PhysicalResourceId", *resource.LogGroupName)
			r.resources = append(r.resources, *resource.LogGroupName)
		}
		return true
	})
	return
}

func getLogsLogStream(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	logDebug("Listing LogsLogStream resources")
	r.err = client.DescribeLogStreamsPages(&cloudwatchlogs.DescribeLogStreamsInput{}, func(page *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
		for _, resource := range page.LogStreams {
			logDebug("Got LogsLogStream resource with PhysicalResourceId", *resource.LogStreamName)
			r.resources = append(r.resources, *resource.LogStreamName)
		}
		return true
	})
	return
}

func getLogsMetricFilter(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	logDebug("Listing LogsMetricFilter resources")
	r.err = client.DescribeMetricFiltersPages(&cloudwatchlogs.DescribeMetricFiltersInput{}, func(page *cloudwatchlogs.DescribeMetricFiltersOutput, lastPage bool) bool {
		for _, resource := range page.MetricFilters {
			logDebug("Got LogsMetricFilter resource with PhysicalResourceId", *resource.FilterName)
			r.resources = append(r.resources, *resource.FilterName)
		}
		return true
	})
	return
}

func getLogsSubscriptionFilter(client *cloudwatchlogs.CloudWatchLogs) (r resourceSliceError) {
	logDebug("Listing LogsSubscriptionFilter resources")
	r.err = client.DescribeSubscriptionFiltersPages(&cloudwatchlogs.DescribeSubscriptionFiltersInput{}, func(page *cloudwatchlogs.DescribeSubscriptionFiltersOutput, lastPage bool) bool {
		for _, resource := range page.SubscriptionFilters {
			logDebug("Got LogsSubscriptionFilter resource with PhysicalResourceId", *resource.FilterName)
			r.resources = append(r.resources, *resource.FilterName)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func getCloudWatch(session *session.Session) (resources resourceMap) {
	client := cloudwatch.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		cloudWatchAlarm:       getCloudWatchAlarm(client),
		cloudWatchDashboard:   getCloudWatchDashboard(client),
		cloudWatchInsightRule: getCloudWatchInsightRule(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCloudWatchAlarm(client *cloudwatch.CloudWatch) (r resourceSliceError) {
	logDebug("Listing CloudWatchAlarm resources")
	r.err = client.DescribeAlarmsPages(&cloudwatch.DescribeAlarmsInput{}, func(page *cloudwatch.DescribeAlarmsOutput, lastPage bool) bool {
		for _, resource := range page.MetricAlarms {
			logDebug("Got CloudWatchAlarm resource with PhysicalResourceId", *resource.AlarmName)
			r.resources = append(r.resources, *resource.AlarmName)
		}
		return true
	})
	return
}

func getCloudWatchDashboard(client *cloudwatch.CloudWatch) (r resourceSliceError) {
	logDebug("Listing CloudWatchDashboard resources")
	r.err = client.ListDashboardsPages(&cloudwatch.ListDashboardsInput{}, func(page *cloudwatch.ListDashboardsOutput, lastPage bool) bool {
		for _, resource := range page.DashboardEntries {
			logDebug("Got CloudWatchDashboard resource with PhysicalResourceId", *resource.DashboardName)
			r.resources = append(r.resources, *resource.DashboardName)
		}
		return true
	})
	return
}

func getCloudWatchInsightRule(client *cloudwatch.CloudWatch) (r resourceSliceError) {
	logDebug("Listing CloudWatchInsightRule resources")
	r.err = client.DescribeInsightRulesPages(&cloudwatch.DescribeInsightRulesInput{}, func(page *cloudwatch.DescribeInsightRulesOutput, lastPage bool) bool {
		for _, resource := range page.InsightRules {
			logDebug("Got CloudWatchInsightRule resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

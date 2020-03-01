package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func getCloudWatch(session *session.Session) (resources resourceMap) {
	client := cloudwatch.New(session)
	resources = reduce(
		getCloudWatchAlarm(client).unwrap(cloudWatchAlarm),
		getCloudWatchDashboard(client).unwrap(cloudWatchDashboard),
		getCloudWatchInsightRule(client).unwrap(cloudWatchInsightRule),
	)
	return
}

func getCloudWatchAlarm(client *cloudwatch.CloudWatch) (r resourceSliceError) {
	r.err = client.DescribeAlarmsPages(&cloudwatch.DescribeAlarmsInput{}, func(page *cloudwatch.DescribeAlarmsOutput, lastPage bool) bool {
		for _, resource := range page.MetricAlarms {
			r.resources = append(r.resources, *resource.AlarmName)
		}
		return true
	})
	return
}

func getCloudWatchDashboard(client *cloudwatch.CloudWatch) (r resourceSliceError) {
	r.err = client.ListDashboardsPages(&cloudwatch.ListDashboardsInput{}, func(page *cloudwatch.ListDashboardsOutput, lastPage bool) bool {
		for _, resource := range page.DashboardEntries {
			r.resources = append(r.resources, *resource.DashboardName)
		}
		return true
	})
	return
}

func getCloudWatchInsightRule(client *cloudwatch.CloudWatch) (r resourceSliceError) {
	r.err = client.DescribeInsightRulesPages(&cloudwatch.DescribeInsightRulesInput{}, func(page *cloudwatch.DescribeInsightRulesOutput, lastPage bool) bool {
		for _, resource := range page.InsightRules {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

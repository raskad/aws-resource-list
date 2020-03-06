package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func getCloudWatch(config aws.Config) (resources resourceMap) {
	client := cloudwatch.New(config)
	resources = reduce(
		getCloudWatchAlarm(client).unwrap(cloudWatchAlarm),
		getCloudWatchDashboard(client).unwrap(cloudWatchDashboard),
		getCloudWatchInsightRule(client).unwrap(cloudWatchInsightRule),
	)
	return
}

func getCloudWatchAlarm(client *cloudwatch.Client) (r resourceSliceError) {
	req := client.DescribeAlarmsRequest(&cloudwatch.DescribeAlarmsInput{})
	p := cloudwatch.NewDescribeAlarmsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.MetricAlarms {
			r.resources = append(r.resources, *resource.AlarmName)
		}
	}
	r.err = p.Err()
	return
}

func getCloudWatchDashboard(client *cloudwatch.Client) (r resourceSliceError) {
	req := client.ListDashboardsRequest(&cloudwatch.ListDashboardsInput{})
	p := cloudwatch.NewListDashboardsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DashboardEntries {
			r.resources = append(r.resources, *resource.DashboardName)
		}
	}
	r.err = p.Err()
	return
}

func getCloudWatchInsightRule(client *cloudwatch.Client) (r resourceSliceError) {
	req := client.DescribeInsightRulesRequest(&cloudwatch.DescribeInsightRulesInput{})
	p := cloudwatch.NewDescribeInsightRulesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.InsightRules {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

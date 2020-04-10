package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func getCloudWatch(config aws.Config) (resources awsResourceMap) {
	client := cloudwatch.New(config)

	cloudWatchAlarmNames := getCloudWatchAlarmNames(client)
	cloudWatchDashboardNames := getCloudWatchDashboardNames(client)
	cloudWatchInsightRuleNames := getCloudWatchInsightRuleNames(client)

	resources = awsResourceMap{
		cloudWatchAlarm:       cloudWatchAlarmNames,
		cloudWatchDashboard:   cloudWatchDashboardNames,
		cloudWatchInsightRule: cloudWatchInsightRuleNames,
	}
	return
}

func getCloudWatchAlarmNames(client *cloudwatch.Client) (resources []string) {
	req := client.DescribeAlarmsRequest(&cloudwatch.DescribeAlarmsInput{})
	p := cloudwatch.NewDescribeAlarmsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.MetricAlarms {
			resources = append(resources, *resource.AlarmName)
		}
	}
	return
}

func getCloudWatchDashboardNames(client *cloudwatch.Client) (resources []string) {
	req := client.ListDashboardsRequest(&cloudwatch.ListDashboardsInput{})
	p := cloudwatch.NewListDashboardsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DashboardEntries {
			resources = append(resources, *resource.DashboardName)
		}
	}
	return
}

func getCloudWatchInsightRuleNames(client *cloudwatch.Client) (resources []string) {
	req := client.DescribeInsightRulesRequest(&cloudwatch.DescribeInsightRulesInput{})
	p := cloudwatch.NewDescribeInsightRulesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.InsightRules {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

func getCloudWatchEvents(config aws.Config) (resources awsResourceMap) {
	client := cloudwatchevents.New(config)

	eventsEventBusNames := getEventsEventBusNames(client)
	eventsRuleNames := getEventsRuleNames(client)
	eventsTargetIDs := getEventsTargetIDs(client, eventsRuleNames)

	resources = awsResourceMap{
		eventsEventBus: eventsEventBusNames,
		eventsRule:     eventsRuleNames,
		eventsTarget:   eventsTargetIDs,
	}
	return
}

func getEventsEventBusNames(client *cloudwatchevents.Client) (resources []string) {
	input := cloudwatchevents.ListEventBusesInput{}
	for {
		page, err := client.ListEventBusesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.EventBuses {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getEventsRuleNames(client *cloudwatchevents.Client) (resources []string) {
	input := cloudwatchevents.ListRulesInput{}
	for {
		page, err := client.ListRulesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Rules {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getEventsTargetIDs(client *cloudwatchevents.Client, eventsRuleNames []string) (resources []string) {
	for _, eventsRuleName := range eventsRuleNames {
		input := cloudwatchevents.ListTargetsByRuleInput{
			Rule: &eventsRuleName,
		}
		for {
			page, err := client.ListTargetsByRuleRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.Targets {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

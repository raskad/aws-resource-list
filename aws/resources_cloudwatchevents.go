package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

func getCloudWatchEvents(config aws.Config) (resources resourceMap) {
	client := cloudwatchevents.New(config)

	eventsEventBusNames := getEventsEventBusNames(client)
	eventsRuleNames := getEventsRuleNames(client)

	resources = resourceMap{
		eventsEventBus: eventsEventBusNames,
		eventsRule:     eventsRuleNames,
	}
	return
}

func getEventsEventBusNames(client *cloudwatchevents.Client) (resources []string) {
	input := cloudwatchevents.ListEventBusesInput{}
	for {
		page, err := client.ListEventBusesRequest(&input).Send(context.Background())
		logErr(err)
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
		logErr(err)
		for _, resource := range page.Rules {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

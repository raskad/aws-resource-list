package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

func getCloudWatchEvents(config aws.Config) (resources resourceMap) {
	client := cloudwatchevents.New(config)
	resources = reduce(
		getEventsEventBus(client).unwrap(eventsEventBus),
		getEventsRule(client).unwrap(eventsRule),
	)
	return
}

func getEventsEventBus(client *cloudwatchevents.Client) (r resourceSliceError) {
	input := cloudwatchevents.ListEventBusesInput{}
	for {
		page, err := client.ListEventBusesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.EventBuses {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getEventsRule(client *cloudwatchevents.Client) (r resourceSliceError) {
	input := cloudwatchevents.ListRulesInput{}
	for {
		page, err := client.ListRulesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Rules {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

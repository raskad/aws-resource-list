package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func getCloudWatchEvents(session *session.Session) (resources resourceMap) {
	client := cloudwatchevents.New(session)
	resources = reduce(
		getEventsEventBus(client).unwrap(eventsEventBus),
		getEventsRule(client).unwrap(eventsRule),
	)
	return
}

func getEventsEventBus(client *cloudwatchevents.CloudWatchEvents) (r resourceSliceError) {
	input := cloudwatchevents.ListEventBusesInput{}
	for {
		page, err := client.ListEventBuses(&input)
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

func getEventsRule(client *cloudwatchevents.CloudWatchEvents) (r resourceSliceError) {
	input := cloudwatchevents.ListRulesInput{}
	for {
		page, err := client.ListRules(&input)
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

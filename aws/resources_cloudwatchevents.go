package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func getCloudWatchEvents(session *session.Session) (resources resourceMap) {
	client := cloudwatchevents.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		eventsEventBus: getEventsEventBus(client),
		eventsRule:     getEventsRule(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
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
		logDebug("Listing EventsEventBus resources page. Remaining pages", page.NextToken)
		for _, resource := range page.EventBuses {
			logDebug("Got EventsEventBus resource with PhysicalResourceId", *resource.Name)
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
		logDebug("Listing EventsRule resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Rules {
			logDebug("Got EventsRule resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

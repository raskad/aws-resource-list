package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func getSns(config aws.Config) (resources awsResourceMap) {
	client := sns.New(config)

	snsPlatformApplicationARNs := getSnsPlatformApplicationARNs(client)
	snsSubscriptionARNs := getSnsSubscriptionARNs(client)
	snsTopicARNs := getSnsTopicARNs(client)

	resources = awsResourceMap{
		snsPlatformApplication: snsPlatformApplicationARNs,
		snsSubscription:        snsSubscriptionARNs,
		snsTopic:               snsTopicARNs,
	}
	return
}

func getSnsPlatformApplicationARNs(client *sns.Client) (resources []string) {
	req := client.ListPlatformApplicationsRequest(&sns.ListPlatformApplicationsInput{})
	p := sns.NewListPlatformApplicationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.PlatformApplications {
			resources = append(resources, *resource.PlatformApplicationArn)
		}
	}
	return
}

func getSnsSubscriptionARNs(client *sns.Client) (resources []string) {
	req := client.ListSubscriptionsRequest(&sns.ListSubscriptionsInput{})
	p := sns.NewListSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Subscriptions {
			resources = append(resources, *resource.SubscriptionArn)
		}
	}
	return
}

func getSnsTopicARNs(client *sns.Client) (resources []string) {
	req := client.ListTopicsRequest(&sns.ListTopicsInput{})
	p := sns.NewListTopicsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Topics {
			resources = append(resources, *resource.TopicArn)
		}
	}
	return
}

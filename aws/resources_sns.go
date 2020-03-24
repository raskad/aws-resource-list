package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func getSns(config aws.Config) (resources resourceMap) {
	client := sns.New(config)

	snsSubscriptionARNs := getSnsSubscriptionARNs(client)
	snsTopicARNs := getSnsTopicARNs(client)

	resources = resourceMap{
		snsSubscription: snsSubscriptionARNs,
		snsTopic:        snsTopicARNs,
	}
	return
}

func getSnsSubscriptionARNs(client *sns.Client) (resources []string) {
	req := client.ListSubscriptionsRequest(&sns.ListSubscriptionsInput{})
	p := sns.NewListSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
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
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Topics {
			resources = append(resources, *resource.TopicArn)
		}
	}
	return
}

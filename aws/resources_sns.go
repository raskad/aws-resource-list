package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func getSns(config aws.Config) (resources resourceMap) {
	client := sns.New(config)
	resources = reduce(
		getSnsSubscription(client).unwrap(snsSubscription),
		getSnsTopic(client).unwrap(snsTopic),
	)
	return
}

func getSnsSubscription(client *sns.Client) (r resourceSliceError) {
	req := client.ListSubscriptionsRequest(&sns.ListSubscriptionsInput{})
	p := sns.NewListSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Subscriptions {
			r.resources = append(r.resources, *resource.SubscriptionArn)
		}
	}
	r.err = p.Err()
	return
}

func getSnsTopic(client *sns.Client) (r resourceSliceError) {
	req := client.ListTopicsRequest(&sns.ListTopicsInput{})
	p := sns.NewListTopicsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Topics {
			r.resources = append(r.resources, *resource.TopicArn)
		}
	}
	r.err = p.Err()
	return
}

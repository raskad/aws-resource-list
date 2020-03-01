package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func getSns(session *session.Session) (resources resourceMap) {
	client := sns.New(session)
	resources = reduce(
		getSnsSubscription(client).unwrap(snsSubscription),
		getSnsTopic(client).unwrap(snsTopic),
	)
	return
}

func getSnsSubscription(client *sns.SNS) (r resourceSliceError) {
	r.err = client.ListSubscriptionsPages(&sns.ListSubscriptionsInput{}, func(page *sns.ListSubscriptionsOutput, lastPage bool) bool {
		for _, resource := range page.Subscriptions {
			r.resources = append(r.resources, *resource.SubscriptionArn)
		}
		return true
	})
	return
}

func getSnsTopic(client *sns.SNS) (r resourceSliceError) {
	r.err = client.ListTopicsPages(&sns.ListTopicsInput{}, func(page *sns.ListTopicsOutput, lastPage bool) bool {
		for _, resource := range page.Topics {
			r.resources = append(r.resources, *resource.TopicArn)
		}
		return true
	})
	return
}

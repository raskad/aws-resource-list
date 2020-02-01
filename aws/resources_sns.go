package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func getSns(session *session.Session) (resources resourceMap) {
	client := sns.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		snsSubscription: getSnsSubscription(client),
		snsTopic:        getSnsTopic(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getSnsSubscription(client *sns.SNS) (r resourceSliceError) {
	r.err = client.ListSubscriptionsPages(&sns.ListSubscriptionsInput{}, func(page *sns.ListSubscriptionsOutput, lastPage bool) bool {
		logDebug("Listing SnsSubscription resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Subscriptions {
			logDebug("Got SnsSubscription resource with PhysicalResourceId", *resource.SubscriptionArn)
			r.resources = append(r.resources, *resource.SubscriptionArn)
		}
		return true
	})
	return
}

func getSnsTopic(client *sns.SNS) (r resourceSliceError) {
	r.err = client.ListTopicsPages(&sns.ListTopicsInput{}, func(page *sns.ListTopicsOutput, lastPage bool) bool {
		logDebug("Listing SnsTopic resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Topics {
			logDebug("Got SnsTopic resource with PhysicalResourceId", *resource.TopicArn)
			r.resources = append(r.resources, *resource.TopicArn)
		}
		return true
	})
	return
}

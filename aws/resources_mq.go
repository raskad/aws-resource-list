package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mq"
)

func getMq(session *session.Session) (resources resourceMap) {
	client := mq.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		amazonMQBroker:        getAmazonMQBroker(client),
		amazonMQConfiguration: getAmazonMQConfiguration(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAmazonMQBroker(client *mq.MQ) (r resourceSliceError) {
	input := mq.ListBrokersInput{}
	for {
		page, err := client.ListBrokers(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing AmazonMQBroker resources page. Remaining pages", page.NextToken)
		for _, resource := range page.BrokerSummaries {
			logDebug("Got AmazonMQBroker resource with PhysicalResourceId", *resource.BrokerName)
			r.resources = append(r.resources, *resource.BrokerName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAmazonMQConfiguration(client *mq.MQ) (r resourceSliceError) {
	input := mq.ListConfigurationsInput{}
	for {
		page, err := client.ListConfigurations(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing AmazonMQConfiguration resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Configurations {
			logDebug("Got AmazonMQConfiguration resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

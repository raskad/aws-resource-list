package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mq"
)

func getMq(session *session.Session) (resources resourceMap) {
	client := mq.New(session)
	resources = reduce(
		getAmazonMQBroker(client).unwrap(amazonMQBroker),
		getAmazonMQConfiguration(client).unwrap(amazonMQConfiguration),
	)
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
		for _, resource := range page.BrokerSummaries {
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
		for _, resource := range page.Configurations {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

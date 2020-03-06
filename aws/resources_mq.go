package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func getMq(config aws.Config) (resources resourceMap) {
	client := mq.New(config)
	resources = reduce(
		getAmazonMQBroker(client).unwrap(amazonMQBroker),
		getAmazonMQConfiguration(client).unwrap(amazonMQConfiguration),
	)
	return
}

func getAmazonMQBroker(client *mq.Client) (r resourceSliceError) {
	input := mq.ListBrokersInput{}
	for {
		page, err := client.ListBrokersRequest(&input).Send(context.Background())
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

func getAmazonMQConfiguration(client *mq.Client) (r resourceSliceError) {
	input := mq.ListConfigurationsInput{}
	for {
		page, err := client.ListConfigurationsRequest(&input).Send(context.Background())
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

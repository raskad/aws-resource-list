package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func getMq(config aws.Config) (resources resourceMap) {
	client := mq.New(config)

	amazonMQBrokerIDs := getAmazonMQBrokerIDs(client)
	amazonMQConfigurationIDs := getAmazonMQConfigurationIDs(client)

	resources = resourceMap{
		amazonMQBroker:        amazonMQBrokerIDs,
		amazonMQConfiguration: amazonMQConfigurationIDs,
	}
	return
}

func getAmazonMQBrokerIDs(client *mq.Client) (resources []string) {
	input := mq.ListBrokersInput{}
	for {
		page, err := client.ListBrokersRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.BrokerSummaries {
			resources = append(resources, *resource.BrokerId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAmazonMQConfigurationIDs(client *mq.Client) (resources []string) {
	input := mq.ListConfigurationsInput{}
	for {
		page, err := client.ListConfigurationsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Configurations {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func getMQ(config aws.Config) (resources awsResourceMap) {
	client := mq.New(config)

	mqBrokerIDs := getMQBrokerIDs(client)
	mqConfigurationIDs := getMQConfigurationIDs(client)

	resources = awsResourceMap{
		mqBroker:        mqBrokerIDs,
		mqConfiguration: mqConfigurationIDs,
	}
	return
}

func getMQBrokerIDs(client *mq.Client) (resources []string) {
	input := mq.ListBrokersInput{}
	for {
		page, err := client.ListBrokersRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.BrokerSummaries {
			resources = append(resources, *resource.BrokerId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getMQConfigurationIDs(client *mq.Client) (resources []string) {
	input := mq.ListConfigurationsInput{}
	for {
		page, err := client.ListConfigurationsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Configurations {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

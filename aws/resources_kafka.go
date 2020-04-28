package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

func getKafka(config aws.Config) (resources awsResourceMap) {
	client := kafka.New(config)

	kafkaClusterNames := getKafkaClusterNames(client)
	kafkaConfigurationNames := getKafkaConfigurationNames(client)

	resources = awsResourceMap{
		kafkaCluster:       kafkaClusterNames,
		kafkaConfiguration: kafkaConfigurationNames,
	}
	return
}

func getKafkaClusterNames(client *kafka.Client) (resources []string) {
	req := client.ListClustersRequest(&kafka.ListClustersInput{})
	p := kafka.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ClusterInfoList {
			resources = append(resources, *resource.ClusterName)
		}
	}
	return
}

func getKafkaConfigurationNames(client *kafka.Client) (resources []string) {
	req := client.ListConfigurationsRequest(&kafka.ListConfigurationsInput{})
	p := kafka.NewListConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Configurations {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

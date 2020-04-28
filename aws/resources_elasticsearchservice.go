package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

func getElasticSearchService(config aws.Config) (resources awsResourceMap) {
	client := elasticsearchservice.New(config)

	elasticSearchServiceDomainNames := getElasticSearchServiceDomainNames(client)

	resources = awsResourceMap{
		elasticSearchServiceDomain: elasticSearchServiceDomainNames,
	}
	return
}

func getElasticSearchServiceDomainNames(client *elasticsearchservice.Client) (resources []string) {
	page, err := client.ListDomainNamesRequest(&elasticsearchservice.ListDomainNamesInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.DomainNames {
		resources = append(resources, *resource.DomainName)
	}
	return
}

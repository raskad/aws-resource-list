package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

func getElasticsearch(config aws.Config) (resources resourceMap) {
	client := elasticsearchservice.New(config)

	elasticsearchDomainNames := getElasticsearchDomainNames(client)

	resources = resourceMap{
		elasticsearchDomain: elasticsearchDomainNames,
	}
	return
}

func getElasticsearchDomainNames(client *elasticsearchservice.Client) (resources []string) {
	page, err := client.ListDomainNamesRequest(&elasticsearchservice.ListDomainNamesInput{}).Send(context.Background())
	logErr(err)
	for _, resource := range page.DomainNames {
		resources = append(resources, *resource.DomainName)
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

func getElasticsearch(config aws.Config) (resources resourceMap) {
	client := elasticsearchservice.New(config)
	resources = reduce(
		getElasticsearchDomain(client).unwrap(elasticsearchDomain),
	)
	return
}

func getElasticsearchDomain(client *elasticsearchservice.Client) (r resourceSliceError) {
	page, err := client.ListDomainNamesRequest(&elasticsearchservice.ListDomainNamesInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.DomainNames {
		r.resources = append(r.resources, *resource.DomainName)
	}
	return
}

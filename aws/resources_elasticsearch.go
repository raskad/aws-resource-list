package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
)

func getElasticsearch(session *session.Session) (resources resourceMap) {
	client := elasticsearchservice.New(session)
	resources = reduce(
		getElasticsearchDomain(client).unwrap(elasticsearchDomain),
	)
	return
}

func getElasticsearchDomain(client *elasticsearchservice.ElasticsearchService) (r resourceSliceError) {
	page, err := client.ListDomainNames(&elasticsearchservice.ListDomainNamesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.DomainNames {
		r.resources = append(r.resources, *resource.DomainName)
	}
	return
}

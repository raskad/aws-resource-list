package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
)

func getElasticsearch(session *session.Session) (resources resourceMap) {
	client := elasticsearchservice.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		elasticsearchDomain: getElasticsearchDomain(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getElasticsearchDomain(client *elasticsearchservice.ElasticsearchService) (r resourceSliceError) {
	page, err := client.ListDomainNames(&elasticsearchservice.ListDomainNamesInput{})
	if err != nil {
		r.err = err
		return
	}
	logDebug("Listing ElasticsearchDomain resources page.")
	for _, resource := range page.DomainNames {
		logDebug("Got ElasticsearchDomain resource with PhysicalResourceId", *resource.DomainName)
		r.resources = append(r.resources, *resource.DomainName)
	}
	return
}

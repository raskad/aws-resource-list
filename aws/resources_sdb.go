package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/simpledb"
)

func getSdb(config aws.Config) (resources resourceMap) {
	client := simpledb.New(config)

	sdbDomainNames := getSdbDomainNames(client)

	resources = resourceMap{
		sdbDomain: sdbDomainNames,
	}
	return
}

func getSdbDomainNames(client *simpledb.Client) (resources []string) {
	req := client.ListDomainsRequest(&simpledb.ListDomainsInput{})
	p := simpledb.NewListDomainsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.DomainNames {
			resources = append(resources, resource)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/simpledb"
)

func getSdb(config aws.Config) (resources resourceMap) {
	client := simpledb.New(config)
	resources = reduce(
		getSdbDomain(client).unwrap(sdbDomain),
	)
	return
}

func getSdbDomain(client *simpledb.Client) (r resourceSliceError) {
	req := client.ListDomainsRequest(&simpledb.ListDomainsInput{})
	p := simpledb.NewListDomainsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DomainNames {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

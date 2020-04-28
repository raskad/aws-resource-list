package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/simpledb"
)

func getSimpleDB(config aws.Config) (resources awsResourceMap) {
	client := simpledb.New(config)

	simpleDBDomainNames := getSimpleDBDomainNames(client)

	resources = awsResourceMap{
		simpleDBDomain: simpleDBDomainNames,
	}
	return
}

func getSimpleDBDomainNames(client *simpledb.Client) (resources []string) {
	req := client.ListDomainsRequest(&simpledb.ListDomainsInput{})
	p := simpledb.NewListDomainsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.DomainNames...)
	}
	return
}

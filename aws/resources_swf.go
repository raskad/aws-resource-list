package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/swf"
)

func getSWF(config aws.Config) (resources awsResourceMap) {
	client := swf.New(config)

	swfDomainNames := getSWFDomainNames(client)

	resources = awsResourceMap{
		swfDomain: swfDomainNames,
	}
	return
}

func getSWFDomainNames(client *swf.Client) (resources []string) {
	req := client.ListDomainsRequest(&swf.ListDomainsInput{})
	p := swf.NewListDomainsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DomainInfos {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

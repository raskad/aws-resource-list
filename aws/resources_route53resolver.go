package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

func getRoute53Resolver(config aws.Config) (resources resourceMap) {
	client := route53resolver.New(config)
	resources = reduce(
		getRoute53ResolverResolverEndpoint(client).unwrap(route53ResolverResolverEndpoint),
		getRoute53ResolverResolverRule(client).unwrap(route53ResolverResolverRule),
		getRoute53ResolverResolverRuleAssociation(client).unwrap(route53ResolverResolverRuleAssociation),
	)
	return
}

func getRoute53ResolverResolverEndpoint(client *route53resolver.Client) (r resourceSliceError) {
	req := client.ListResolverEndpointsRequest(&route53resolver.ListResolverEndpointsInput{})
	p := route53resolver.NewListResolverEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ResolverEndpoints {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getRoute53ResolverResolverRule(client *route53resolver.Client) (r resourceSliceError) {
	req := client.ListResolverRulesRequest(&route53resolver.ListResolverRulesInput{})
	p := route53resolver.NewListResolverRulesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ResolverRules {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getRoute53ResolverResolverRuleAssociation(client *route53resolver.Client) (r resourceSliceError) {
	req := client.ListResolverRuleAssociationsRequest(&route53resolver.ListResolverRuleAssociationsInput{})
	p := route53resolver.NewListResolverRuleAssociationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ResolverRuleAssociations {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

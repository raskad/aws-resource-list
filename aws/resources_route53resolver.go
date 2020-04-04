package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

func getRoute53Resolver(config aws.Config) (resources awsResourceMap) {
	client := route53resolver.New(config)

	route53ResolverResolverEndpointIDs := getRoute53ResolverResolverEndpointIDs(client)
	route53ResolverResolverRuleIDs := getRoute53ResolverResolverRuleIDs(client)
	route53ResolverResolverRuleAssociationIDs := getRoute53ResolverResolverRuleAssociationIDs(client)

	resources = awsResourceMap{
		route53ResolverResolverEndpoint:        route53ResolverResolverEndpointIDs,
		route53ResolverResolverRule:            route53ResolverResolverRuleIDs,
		route53ResolverResolverRuleAssociation: route53ResolverResolverRuleAssociationIDs,
	}
	return
}

func getRoute53ResolverResolverEndpointIDs(client *route53resolver.Client) (resources []string) {
	req := client.ListResolverEndpointsRequest(&route53resolver.ListResolverEndpointsInput{})
	p := route53resolver.NewListResolverEndpointsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ResolverEndpoints {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getRoute53ResolverResolverRuleIDs(client *route53resolver.Client) (resources []string) {
	req := client.ListResolverRulesRequest(&route53resolver.ListResolverRulesInput{})
	p := route53resolver.NewListResolverRulesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ResolverRules {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getRoute53ResolverResolverRuleAssociationIDs(client *route53resolver.Client) (resources []string) {
	req := client.ListResolverRuleAssociationsRequest(&route53resolver.ListResolverRuleAssociationsInput{})
	p := route53resolver.NewListResolverRuleAssociationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ResolverRuleAssociations {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

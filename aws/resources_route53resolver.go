package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53resolver"
)

func getRoute53Resolver(session *session.Session) (resources resourceMap) {
	client := route53resolver.New(session)
	resources = reduce(
		getRoute53ResolverResolverEndpoint(client).unwrap(route53ResolverResolverEndpoint),
		getRoute53ResolverResolverRule(client).unwrap(route53ResolverResolverRule),
		getRoute53ResolverResolverRuleAssociation(client).unwrap(route53ResolverResolverRuleAssociation),
	)
	return
}

func getRoute53ResolverResolverEndpoint(client *route53resolver.Route53Resolver) (r resourceSliceError) {
	r.err = client.ListResolverEndpointsPages(&route53resolver.ListResolverEndpointsInput{}, func(page *route53resolver.ListResolverEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.ResolverEndpoints {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getRoute53ResolverResolverRule(client *route53resolver.Route53Resolver) (r resourceSliceError) {
	r.err = client.ListResolverRulesPages(&route53resolver.ListResolverRulesInput{}, func(page *route53resolver.ListResolverRulesOutput, lastPage bool) bool {
		for _, resource := range page.ResolverRules {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getRoute53ResolverResolverRuleAssociation(client *route53resolver.Route53Resolver) (r resourceSliceError) {
	r.err = client.ListResolverRuleAssociationsPages(&route53resolver.ListResolverRuleAssociationsInput{}, func(page *route53resolver.ListResolverRuleAssociationsOutput, lastPage bool) bool {
		for _, resource := range page.ResolverRuleAssociations {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

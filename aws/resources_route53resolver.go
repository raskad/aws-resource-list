package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53resolver"
)

func getRoute53Resolver(session *session.Session) (resources resourceMap) {
	client := route53resolver.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		route53ResolverResolverEndpoint:        getRoute53ResolverResolverEndpoint(client),
		route53ResolverResolverRule:            getRoute53ResolverResolverRule(client),
		route53ResolverResolverRuleAssociation: getRoute53ResolverResolverRuleAssociation(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getRoute53ResolverResolverEndpoint(client *route53resolver.Route53Resolver) (r resourceSliceError) {
	r.err = client.ListResolverEndpointsPages(&route53resolver.ListResolverEndpointsInput{}, func(page *route53resolver.ListResolverEndpointsOutput, lastPage bool) bool {
		logDebug("Listing Route53ResolverResolverEndpoint resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ResolverEndpoints {
			logDebug("Got Route53ResolverResolverEndpoint resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getRoute53ResolverResolverRule(client *route53resolver.Route53Resolver) (r resourceSliceError) {
	r.err = client.ListResolverRulesPages(&route53resolver.ListResolverRulesInput{}, func(page *route53resolver.ListResolverRulesOutput, lastPage bool) bool {
		logDebug("Listing Route53ResolverResolverRule resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ResolverRules {
			logDebug("Got Route53ResolverResolverRule resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getRoute53ResolverResolverRuleAssociation(client *route53resolver.Route53Resolver) (r resourceSliceError) {
	r.err = client.ListResolverRuleAssociationsPages(&route53resolver.ListResolverRuleAssociationsInput{}, func(page *route53resolver.ListResolverRuleAssociationsOutput, lastPage bool) bool {
		logDebug("Listing Route53ResolverResolverRuleAssociation resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ResolverRuleAssociations {
			logDebug("Got Route53ResolverResolverRuleAssociation resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

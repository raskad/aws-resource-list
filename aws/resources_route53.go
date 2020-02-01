package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

func getRoute53(session *session.Session) (resources resourceMap) {
	client := route53.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		route53HealthCheck: getRoute53HealthCheck(client),
		route53HostedZone:  getRoute53HostedZone(client),
		route53RecordSet:   getRoute53RecordSet(client),
		//route53ResolverResolverEndpoint: getRoute53ResolverResolverEndpoint(client),
		//route53ResolverResolverRule: getRoute53ResolverResolverRule(client),
		//route53ResolverResolverRuleAssociation: getRoute53ResolverResolverRuleAssociation(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getRoute53HealthCheck(client *route53.Route53) (r resourceSliceError) {
	r.err = client.ListHealthChecksPages(&route53.ListHealthChecksInput{}, func(page *route53.ListHealthChecksOutput, lastPage bool) bool {
		logDebug("Listing Route53HealthCheck resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.HealthChecks {
			logDebug("Got Route53HealthCheck resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getRoute53HostedZone(client *route53.Route53) (r resourceSliceError) {
	r.err = client.ListHostedZonesPages(&route53.ListHostedZonesInput{}, func(page *route53.ListHostedZonesOutput, lastPage bool) bool {
		logDebug("Listing Route53HostedZone resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.HostedZones {
			logDebug("Got Route53HostedZone resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getRoute53RecordSet(client *route53.Route53) (r resourceSliceError) {
	r.err = client.ListResourceRecordSetsPages(&route53.ListResourceRecordSetsInput{}, func(page *route53.ListResourceRecordSetsOutput, lastPage bool) bool {
		logDebug("Listing Route53RecordSet resources page")
		for _, resource := range page.ResourceRecordSets {
			logDebug("Got Route53RecordSet resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

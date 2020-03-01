package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

func getRoute53(session *session.Session) (resources resourceMap) {
	client := route53.New(session)
	resources = reduce(
		getRoute53HealthCheck(client).unwrap(route53HealthCheck),
		getRoute53HostedZone(client).unwrap(route53HostedZone),
		getRoute53RecordSet(client).unwrap(route53RecordSet),
	)
	return
}

func getRoute53HealthCheck(client *route53.Route53) (r resourceSliceError) {
	r.err = client.ListHealthChecksPages(&route53.ListHealthChecksInput{}, func(page *route53.ListHealthChecksOutput, lastPage bool) bool {
		for _, resource := range page.HealthChecks {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getRoute53HostedZone(client *route53.Route53) (r resourceSliceError) {
	r.err = client.ListHostedZonesPages(&route53.ListHostedZonesInput{}, func(page *route53.ListHostedZonesOutput, lastPage bool) bool {
		for _, resource := range page.HostedZones {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getRoute53RecordSet(client *route53.Route53) (r resourceSliceError) {
	r.err = client.ListResourceRecordSetsPages(&route53.ListResourceRecordSetsInput{}, func(page *route53.ListResourceRecordSetsOutput, lastPage bool) bool {
		for _, resource := range page.ResourceRecordSets {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

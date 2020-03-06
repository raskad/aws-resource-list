package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func getRoute53(config aws.Config) (resources resourceMap) {
	client := route53.New(config)

	route53HostedZoneResourceMap := getRoute53HostedZone(client).unwrap(route53HostedZone)
	route53HostedZoneIDs := route53HostedZoneResourceMap[route53HostedZone]

	resources = reduce(
		getRoute53HealthCheck(client).unwrap(route53HealthCheck),
		route53HostedZoneResourceMap,
		getRoute53RecordSet(client, route53HostedZoneIDs).unwrap(route53RecordSet),
	)
	return
}

func getRoute53HealthCheck(client *route53.Client) (r resourceSliceError) {
	req := client.ListHealthChecksRequest(&route53.ListHealthChecksInput{})
	p := route53.NewListHealthChecksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.HealthChecks {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getRoute53HostedZone(client *route53.Client) (r resourceSliceError) {
	req := client.ListHostedZonesRequest(&route53.ListHostedZonesInput{})
	p := route53.NewListHostedZonesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.HostedZones {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getRoute53RecordSet(client *route53.Client, hostedZoneIDs []string) (r resourceSliceError) {
	for _, hostedZoneID := range hostedZoneIDs {
		req := client.ListResourceRecordSetsRequest(&route53.ListResourceRecordSetsInput{
			HostedZoneId: aws.String(hostedZoneID),
		})
		p := route53.NewListResourceRecordSetsPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.ResourceRecordSets {
				r.resources = append(r.resources, *resource.Name)
			}
		}
		r.err = p.Err()
		return
	}
	return
}

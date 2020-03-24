package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func getRoute53(config aws.Config) (resources resourceMap) {
	client := route53.New(config)

	route53HostedZoneIDs := getRoute53HostedZoneIDs(client)
	route53HealthCheckIDs := getRoute53HealthCheckIDs(client)
	route53RecordSetNames := getRoute53RecordSetNames(client, route53HostedZoneIDs)

	resources = resourceMap{
		route53HostedZone:  route53HostedZoneIDs,
		route53HealthCheck: route53HealthCheckIDs,
		route53RecordSet:   route53RecordSetNames,
	}
	return
}

func getRoute53HealthCheckIDs(client *route53.Client) (resources []string) {
	req := client.ListHealthChecksRequest(&route53.ListHealthChecksInput{})
	p := route53.NewListHealthChecksPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.HealthChecks {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getRoute53HostedZoneIDs(client *route53.Client) (resources []string) {
	req := client.ListHostedZonesRequest(&route53.ListHostedZonesInput{})
	p := route53.NewListHostedZonesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.HostedZones {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getRoute53RecordSetNames(client *route53.Client, hostedZoneIDs []string) (resources []string) {
	for _, hostedZoneID := range hostedZoneIDs {
		req := client.ListResourceRecordSetsRequest(&route53.ListResourceRecordSetsInput{
			HostedZoneId: aws.String(hostedZoneID),
		})
		p := route53.NewListResourceRecordSetsPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.ResourceRecordSets {
				resources = append(resources, *resource.Name)
			}
		}
	}
	return
}

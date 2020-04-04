package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func getRoute53(config aws.Config) (resources awsResourceMap) {
	client := route53.New(config)

	route53DelegationSetIDs := getRoute53DelegationSetIDs(client)
	route53HostedZoneIDs := getRoute53HostedZoneIDs(client)
	route53HealthCheckIDs := getRoute53HealthCheckIDs(client)
	route53RecordSetNames := getRoute53RecordSetNames(client, route53HostedZoneIDs)
	route53QueryLogIDs := getRoute53QueryLogIDs(client)

	resources = awsResourceMap{
		route53DelegationSet: route53DelegationSetIDs,
		route53HealthCheck:   route53HealthCheckIDs,
		route53HostedZone:    route53HostedZoneIDs,
		route53RecordSet:     route53RecordSetNames,
		route53QueryLog:      route53QueryLogIDs,
	}
	return
}

func getRoute53DelegationSetIDs(client *route53.Client) (resources []string) {
	input := route53.ListReusableDelegationSetsInput{}
	for {
		page, err := client.ListReusableDelegationSetsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.DelegationSets {
			resources = append(resources, *resource.Id)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getRoute53HealthCheckIDs(client *route53.Client) (resources []string) {
	req := client.ListHealthChecksRequest(&route53.ListHealthChecksInput{})
	p := route53.NewListHealthChecksPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.ResourceRecordSets {
				resources = append(resources, *resource.Name)
			}
		}
	}
	return
}

func getRoute53QueryLogIDs(client *route53.Client) (resources []string) {
	input := route53.ListQueryLoggingConfigsInput{}
	for {
		page, err := client.ListQueryLoggingConfigsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.QueryLoggingConfigs {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

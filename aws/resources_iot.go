package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func getIoT(config aws.Config) (resources resourceMap) {
	client := iot.New(config)
	resources = reduce(
		getIoTCertificate(client).unwrap(ioTCertificate),
		getIoTPolicy(client).unwrap(ioTPolicy),
		getIoTThing(client).unwrap(ioTThing),
		getIoTTopicRule(client).unwrap(ioTTopicRule),
	)
	return
}

func getIoTCertificate(client *iot.Client) (r resourceSliceError) {
	input := iot.ListCertificatesInput{}
	for {
		page, err := client.ListCertificatesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Certificates {
			r.resources = append(r.resources, *resource.CertificateId)
		}
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

func getIoTPolicy(client *iot.Client) (r resourceSliceError) {
	input := iot.ListPoliciesInput{}
	for {
		page, err := client.ListPoliciesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Policies {
			r.resources = append(r.resources, *resource.PolicyName)
		}
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

func getIoTThing(client *iot.Client) (r resourceSliceError) {
	input := iot.ListThingsInput{}
	for {
		page, err := client.ListThingsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Things {
			r.resources = append(r.resources, *resource.ThingName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getIoTTopicRule(client *iot.Client) (r resourceSliceError) {
	input := iot.ListTopicRulesInput{}
	for {
		page, err := client.ListTopicRulesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Rules {
			r.resources = append(r.resources, *resource.RuleName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

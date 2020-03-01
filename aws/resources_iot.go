package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
)

func getIoT(session *session.Session) (resources resourceMap) {
	client := iot.New(session)
	resources = reduce(
		getIoTCertificate(client).unwrap(ioTCertificate),
		getIoTPolicy(client).unwrap(ioTPolicy),
		getIoTThing(client).unwrap(ioTThing),
		getIoTTopicRule(client).unwrap(ioTTopicRule),
	)
	return
}

func getIoTCertificate(client *iot.IoT) (r resourceSliceError) {
	input := iot.ListCertificatesInput{}
	for {
		page, err := client.ListCertificates(&input)
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

func getIoTPolicy(client *iot.IoT) (r resourceSliceError) {
	input := iot.ListPoliciesInput{}
	for {
		page, err := client.ListPolicies(&input)
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

func getIoTThing(client *iot.IoT) (r resourceSliceError) {
	input := iot.ListThingsInput{}
	for {
		page, err := client.ListThings(&input)
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

func getIoTTopicRule(client *iot.IoT) (r resourceSliceError) {
	input := iot.ListTopicRulesInput{}
	for {
		page, err := client.ListTopicRules(&input)
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

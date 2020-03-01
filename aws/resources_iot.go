package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
)

func getIoT(session *session.Session) (resources resourceMap) {
	client := iot.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		ioTCertificate: getIoTCertificate(client),
		ioTPolicy:      getIoTPolicy(client),
		ioTThing:       getIoTThing(client),
		ioTTopicRule:   getIoTTopicRule(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getIoTCertificate(client *iot.IoT) (r resourceSliceError) {
	logDebug("Listing IoTCertificate resources")
	input := iot.ListCertificatesInput{}
	for {
		page, err := client.ListCertificates(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Certificates {
			logDebug("Got IoTCertificate resource with PhysicalResourceId", *resource.CertificateId)
			r.resources = append(r.resources, *resource.CertificateId)
		}
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

func getIoTPolicy(client *iot.IoT) (r resourceSliceError) {
	logDebug("Listing IoTPolicy resources")
	input := iot.ListPoliciesInput{}
	for {
		page, err := client.ListPolicies(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Policies {
			logDebug("Got IoTPolicy resource with PhysicalResourceId", *resource.PolicyName)
			r.resources = append(r.resources, *resource.PolicyName)
		}
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

func getIoTThing(client *iot.IoT) (r resourceSliceError) {
	logDebug("Listing IoTThing resources")
	input := iot.ListThingsInput{}
	for {
		page, err := client.ListThings(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Things {
			logDebug("Got IoTThing resource with PhysicalResourceId", *resource.ThingName)
			r.resources = append(r.resources, *resource.ThingName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getIoTTopicRule(client *iot.IoT) (r resourceSliceError) {
	logDebug("Listing IoTTopicRule resources")
	input := iot.ListTopicRulesInput{}
	for {
		page, err := client.ListTopicRules(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Rules {
			logDebug("Got IoTTopicRule resource with PhysicalResourceId", *resource.RuleName)
			r.resources = append(r.resources, *resource.RuleName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

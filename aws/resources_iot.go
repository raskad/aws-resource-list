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
	input := iot.ListCertificatesInput{}
	for {
		page, err := client.ListCertificates(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing IoTCertificate resources page. Remaining pages", page.NextMarker)
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
	input := iot.ListPoliciesInput{}
	for {
		page, err := client.ListPolicies(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing IoTPolicy resources page. Remaining pages", page.NextMarker)
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
	input := iot.ListThingsInput{}
	for {
		page, err := client.ListThings(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing IoTThing resources page. Remaining pages", page.NextToken)
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
	input := iot.ListTopicRulesInput{}
	for {
		page, err := client.ListTopicRules(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing IoTTopicRule resources page. Remaining pages", page.NextToken)
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

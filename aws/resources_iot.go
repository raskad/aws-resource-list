package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func getIoT(config aws.Config) (resources awsResourceMap) {
	client := iot.New(config)

	ioTCertificateIDs := getIoTCertificateIDs(client)
	ioTPolicyNames := getIoTPolicyNames(client)
	iotRoleAliasNames := getIotRoleAliasNames(client)
	ioTThingNames := getIoTThingNames(client)
	iotThingTypeNames := getIotThingTypeNames(client)
	ioTTopicRuleNames := getIoTTopicRuleNames(client)

	resources = awsResourceMap{
		ioTCertificate: ioTCertificateIDs,
		ioTPolicy:      ioTPolicyNames,
		iotRoleAlias:   iotRoleAliasNames,
		ioTThing:       ioTThingNames,
		iotThingType:   iotThingTypeNames,
		ioTTopicRule:   ioTTopicRuleNames,
	}
	return
}

func getIoTCertificateIDs(client *iot.Client) (resources []string) {
	input := iot.ListCertificatesInput{}
	for {
		page, err := client.ListCertificatesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Certificates {
			resources = append(resources, *resource.CertificateId)
		}
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

func getIoTPolicyNames(client *iot.Client) (resources []string) {
	input := iot.ListPoliciesInput{}
	for {
		page, err := client.ListPoliciesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Policies {
			resources = append(resources, *resource.PolicyName)
		}
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

func getIotRoleAliasNames(client *iot.Client) (resources []string) {
	input := iot.ListRoleAliasesInput{}
	for {
		page, err := client.ListRoleAliasesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		resources = append(resources, page.RoleAliases...)
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

func getIoTThingNames(client *iot.Client) (resources []string) {
	input := iot.ListThingsInput{}
	for {
		page, err := client.ListThingsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Things {
			resources = append(resources, *resource.ThingName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getIoTTopicRuleNames(client *iot.Client) (resources []string) {
	input := iot.ListTopicRulesInput{}
	for {
		page, err := client.ListTopicRulesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Rules {
			resources = append(resources, *resource.RuleName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getIotThingTypeNames(client *iot.Client) (resources []string) {
	input := iot.ListThingTypesInput{}
	for {
		page, err := client.ListThingTypesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ThingTypes {
			resources = append(resources, *resource.ThingTypeName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

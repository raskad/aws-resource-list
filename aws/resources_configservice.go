package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func getConfigService(config aws.Config) (resources awsResourceMap) {
	client := configservice.New(config)

	configServiceConfigRuleNames := getConfigServiceConfigRuleNames(client)
	configServiceAggregationAuthorizationARNs := getConfigServiceAggregationAuthorizationARNs(client)
	configServiceConfigurationAggregatorNames := getConfigServiceConfigurationAggregatorNames(client)
	configServiceConfigurationRecorderNames := getConfigServiceConfigurationRecorderNames(client)
	configServiceConformancePackNames := getConfigServiceConformancePackNames(client)
	configServiceDeliveryChannelNames := getConfigServiceDeliveryChannelNames(client)
	configServiceOrganizationConfigRuleNames := getConfigServiceOrganizationConfigRuleNames(client)
	configServiceOrganizationConformancePackNames := getConfigServiceOrganizationConformancePackNames(client)
	configServiceRemediationConfigurationNames := getConfigServiceRemediationConfigurationNames(client, configServiceConfigRuleNames)

	resources = awsResourceMap{
		configServiceConfigRule:                  configServiceConfigRuleNames,
		configServiceAggregationAuthorization:    configServiceAggregationAuthorizationARNs,
		configServiceConfigurationAggregator:     configServiceConfigurationAggregatorNames,
		configServiceConfigurationRecorder:       configServiceConfigurationRecorderNames,
		configServiceConformancePack:             configServiceConformancePackNames,
		configServiceDeliveryChannel:             configServiceDeliveryChannelNames,
		configServiceOrganizationConfigRule:      configServiceOrganizationConfigRuleNames,
		configServiceOrganizationConformancePack: configServiceOrganizationConformancePackNames,
		configServiceRemediationConfiguration:    configServiceRemediationConfigurationNames,
	}
	return
}

func getConfigServiceAggregationAuthorizationARNs(client *configservice.Client) (resources []string) {
	input := configservice.DescribeAggregationAuthorizationsInput{}
	for {
		page, err := client.DescribeAggregationAuthorizationsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.AggregationAuthorizations {
			resources = append(resources, *resource.AggregationAuthorizationArn)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigServiceConfigRuleNames(client *configservice.Client) (resources []string) {
	input := configservice.DescribeConfigRulesInput{}
	for {
		page, err := client.DescribeConfigRulesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ConfigRules {
			resources = append(resources, *resource.ConfigRuleName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigServiceConfigurationAggregatorNames(client *configservice.Client) (resources []string) {
	input := configservice.DescribeConfigurationAggregatorsInput{}
	for {
		page, err := client.DescribeConfigurationAggregatorsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ConfigurationAggregators {
			resources = append(resources, *resource.ConfigurationAggregatorName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigServiceConfigurationRecorderNames(client *configservice.Client) (resources []string) {
	page, err := client.DescribeConfigurationRecordersRequest(&configservice.DescribeConfigurationRecordersInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.ConfigurationRecorders {
		resources = append(resources, *resource.Name)
	}
	return
}

func getConfigServiceConformancePackNames(client *configservice.Client) (resources []string) {
	input := configservice.DescribeConformancePacksInput{}
	for {
		page, err := client.DescribeConformancePacksRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ConformancePackDetails {
			resources = append(resources, *resource.ConformancePackName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigServiceDeliveryChannelNames(client *configservice.Client) (resources []string) {
	page, err := client.DescribeDeliveryChannelsRequest(&configservice.DescribeDeliveryChannelsInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.DeliveryChannels {
		resources = append(resources, *resource.Name)
	}
	return
}

func getConfigServiceOrganizationConfigRuleNames(client *configservice.Client) (resources []string) {
	input := configservice.DescribeOrganizationConfigRulesInput{}
	for {
		page, err := client.DescribeOrganizationConfigRulesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.OrganizationConfigRules {
			resources = append(resources, *resource.OrganizationConfigRuleName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigServiceOrganizationConformancePackNames(client *configservice.Client) (resources []string) {
	input := configservice.DescribeOrganizationConformancePacksInput{}
	for {
		page, err := client.DescribeOrganizationConformancePacksRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.OrganizationConformancePacks {
			resources = append(resources, *resource.OrganizationConformancePackName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigServiceRemediationConfigurationNames(client *configservice.Client, configRuleNames []string) (resources []string) {
	page, err := client.DescribeRemediationConfigurationsRequest(&configservice.DescribeRemediationConfigurationsInput{
		ConfigRuleNames: configRuleNames,
	}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.RemediationConfigurations {
		resources = append(resources, *resource.ConfigRuleName)
	}
	return
}

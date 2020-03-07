package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func getConfig(config aws.Config) (resources resourceMap) {
	client := configservice.New(config)

	configConfigRuleResourceMap := getConfigConfigRule(client).unwrap(configConfigRule)
	configConfigRuleNames := configConfigRuleResourceMap[configConfigRule]

	resources = reduce(
		getConfigAggregationAuthorization(client).unwrap(configAggregationAuthorization),
		configConfigRuleResourceMap,
		getConfigConfigurationAggregator(client).unwrap(configConfigurationAggregator),
		getConfigConfigurationRecorder(client).unwrap(configConfigurationRecorder),
		getConfigConformancePack(client).unwrap(configConformancePack),
		getConfigDeliveryChannel(client).unwrap(configDeliveryChannel),
		getConfigOrganizationConfigRule(client).unwrap(configOrganizationConfigRule),
		getConfigOrganizationConformancePack(client).unwrap(configOrganizationConformancePack),
		getConfigRemediationConfiguration(client, configConfigRuleNames).unwrap(configRemediationConfiguration),
	)
	return
}

func getConfigAggregationAuthorization(client *configservice.Client) (r resourceSliceError) {
	input := configservice.DescribeAggregationAuthorizationsInput{}
	for {
		page, err := client.DescribeAggregationAuthorizationsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.AggregationAuthorizations {
			r.resources = append(r.resources, *resource.AggregationAuthorizationArn)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigConfigRule(client *configservice.Client) (r resourceSliceError) {
	input := configservice.DescribeConfigRulesInput{}
	for {
		page, err := client.DescribeConfigRulesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ConfigRules {
			r.resources = append(r.resources, *resource.ConfigRuleName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigConfigurationAggregator(client *configservice.Client) (r resourceSliceError) {
	input := configservice.DescribeConfigurationAggregatorsInput{}
	for {
		page, err := client.DescribeConfigurationAggregatorsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ConfigurationAggregators {
			r.resources = append(r.resources, *resource.ConfigurationAggregatorName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigConfigurationRecorder(client *configservice.Client) (r resourceSliceError) {
	page, err := client.DescribeConfigurationRecordersRequest(&configservice.DescribeConfigurationRecordersInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.ConfigurationRecorders {
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getConfigConformancePack(client *configservice.Client) (r resourceSliceError) {
	input := configservice.DescribeConformancePacksInput{}
	for {
		page, err := client.DescribeConformancePacksRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ConformancePackDetails {
			r.resources = append(r.resources, *resource.ConformancePackName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigDeliveryChannel(client *configservice.Client) (r resourceSliceError) {
	page, err := client.DescribeDeliveryChannelsRequest(&configservice.DescribeDeliveryChannelsInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.DeliveryChannels {
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getConfigOrganizationConfigRule(client *configservice.Client) (r resourceSliceError) {
	input := configservice.DescribeOrganizationConfigRulesInput{}
	for {
		page, err := client.DescribeOrganizationConfigRulesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.OrganizationConfigRules {
			r.resources = append(r.resources, *resource.OrganizationConfigRuleName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigOrganizationConformancePack(client *configservice.Client) (r resourceSliceError) {
	input := configservice.DescribeOrganizationConformancePacksInput{}
	for {
		page, err := client.DescribeOrganizationConformancePacksRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.OrganizationConformancePacks {
			r.resources = append(r.resources, *resource.OrganizationConformancePackName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigRemediationConfiguration(client *configservice.Client, configRuleNames []string) (r resourceSliceError) {
	page, err := client.DescribeRemediationConfigurationsRequest(&configservice.DescribeRemediationConfigurationsInput{
		ConfigRuleNames: configRuleNames,
	}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.RemediationConfigurations {
		r.resources = append(r.resources, *resource.ConfigRuleName)
	}
	return
}

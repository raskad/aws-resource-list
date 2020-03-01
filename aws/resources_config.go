package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
)

func getConfig(session *session.Session) (resources resourceMap) {
	client := configservice.New(session)
	resources = reduce(
		getConfigAggregationAuthorization(client).unwrap(configAggregationAuthorization),
		getConfigConfigRule(client).unwrap(configConfigRule),
		getConfigConfigurationAggregator(client).unwrap(configConfigurationAggregator),
		getConfigConfigurationRecorder(client).unwrap(configConfigurationRecorder),
		getConfigConformancePack(client).unwrap(configConformancePack),
		getConfigDeliveryChannel(client).unwrap(configDeliveryChannel),
		getConfigOrganizationConfigRule(client).unwrap(configOrganizationConfigRule),
		getConfigOrganizationConformancePack(client).unwrap(configOrganizationConformancePack),
		getConfigRemediationConfiguration(client).unwrap(configRemediationConfiguration),
	)
	return
}

func getConfigAggregationAuthorization(client *configservice.ConfigService) (r resourceSliceError) {
	input := configservice.DescribeAggregationAuthorizationsInput{}
	for {
		page, err := client.DescribeAggregationAuthorizations(&input)
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

func getConfigConfigRule(client *configservice.ConfigService) (r resourceSliceError) {
	input := configservice.DescribeConfigRulesInput{}
	for {
		page, err := client.DescribeConfigRules(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ConfigRules {
			r.resources = append(r.resources, *resource.ConfigRuleId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigConfigurationAggregator(client *configservice.ConfigService) (r resourceSliceError) {
	input := configservice.DescribeConfigurationAggregatorsInput{}
	for {
		page, err := client.DescribeConfigurationAggregators(&input)
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

func getConfigConfigurationRecorder(client *configservice.ConfigService) (r resourceSliceError) {
	page, err := client.DescribeConfigurationRecorders(&configservice.DescribeConfigurationRecordersInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.ConfigurationRecorders {
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getConfigConformancePack(client *configservice.ConfigService) (r resourceSliceError) {
	input := configservice.DescribeConformancePacksInput{}
	for {
		page, err := client.DescribeConformancePacks(&input)
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

func getConfigDeliveryChannel(client *configservice.ConfigService) (r resourceSliceError) {
	page, err := client.DescribeDeliveryChannels(&configservice.DescribeDeliveryChannelsInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.DeliveryChannels {
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getConfigOrganizationConfigRule(client *configservice.ConfigService) (r resourceSliceError) {
	input := configservice.DescribeOrganizationConfigRulesInput{}
	for {
		page, err := client.DescribeOrganizationConfigRules(&input)
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

func getConfigOrganizationConformancePack(client *configservice.ConfigService) (r resourceSliceError) {
	input := configservice.DescribeOrganizationConformancePacksInput{}
	for {
		page, err := client.DescribeOrganizationConformancePacks(&input)
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

func getConfigRemediationConfiguration(client *configservice.ConfigService) (r resourceSliceError) {
	page, err := client.DescribeRemediationConfigurations(&configservice.DescribeRemediationConfigurationsInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.RemediationConfigurations {
		r.resources = append(r.resources, *resource.Arn)
	}
	return
}

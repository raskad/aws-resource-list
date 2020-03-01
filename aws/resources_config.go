package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
)

func getConfig(session *session.Session) (resources resourceMap) {
	client := configservice.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		configAggregationAuthorization:    getConfigAggregationAuthorization(client),
		configConfigRule:                  getConfigConfigRule(client),
		configConfigurationAggregator:     getConfigConfigurationAggregator(client),
		configConfigurationRecorder:       getConfigConfigurationRecorder(client),
		configConformancePack:             getConfigConformancePack(client),
		configDeliveryChannel:             getConfigDeliveryChannel(client),
		configOrganizationConfigRule:      getConfigOrganizationConfigRule(client),
		configOrganizationConformancePack: getConfigOrganizationConformancePack(client),
		configRemediationConfiguration:    getConfigRemediationConfiguration(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getConfigAggregationAuthorization(client *configservice.ConfigService) (r resourceSliceError) {
	logDebug("Listing ConfigAggregationAuthorization resources")
	input := configservice.DescribeAggregationAuthorizationsInput{}
	for {
		page, err := client.DescribeAggregationAuthorizations(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.AggregationAuthorizations {
			logDebug("Got ConfigAggregationAuthorization resource with PhysicalResourceId", *resource.AggregationAuthorizationArn)
			r.resources = append(r.resources, *resource.AggregationAuthorizationArn)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigConfigRule(client *configservice.ConfigService) (r resourceSliceError) {
	logDebug("Listing ConfigConfigRule resources")
	input := configservice.DescribeConfigRulesInput{}
	for {
		page, err := client.DescribeConfigRules(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ConfigRules {
			logDebug("Got ConfigConfigRule resource with PhysicalResourceId", *resource.ConfigRuleId)
			r.resources = append(r.resources, *resource.ConfigRuleId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigConfigurationAggregator(client *configservice.ConfigService) (r resourceSliceError) {
	logDebug("Listing ConfigConfigurationAggregator resources")
	input := configservice.DescribeConfigurationAggregatorsInput{}
	for {
		page, err := client.DescribeConfigurationAggregators(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ConfigurationAggregators {
			logDebug("Got ConfigConfigurationAggregator resource with PhysicalResourceId", *resource.ConfigurationAggregatorName)
			r.resources = append(r.resources, *resource.ConfigurationAggregatorName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigConfigurationRecorder(client *configservice.ConfigService) (r resourceSliceError) {
	logDebug("Listing ConfigConfigurationRecorder resources")
	page, err := client.DescribeConfigurationRecorders(&configservice.DescribeConfigurationRecordersInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.ConfigurationRecorders {
		logDebug("Got ConfigConfigurationRecorder resource with PhysicalResourceId", *resource.Name)
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getConfigConformancePack(client *configservice.ConfigService) (r resourceSliceError) {
	logDebug("Listing ConfigConformancePack resources")
	input := configservice.DescribeConformancePacksInput{}
	for {
		page, err := client.DescribeConformancePacks(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ConformancePackDetails {
			logDebug("Got ConfigConformancePack resource with PhysicalResourceId", *resource.ConformancePackName)
			r.resources = append(r.resources, *resource.ConformancePackName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigDeliveryChannel(client *configservice.ConfigService) (r resourceSliceError) {
	logDebug("Listing ConfigDeliveryChannel resources")
	page, err := client.DescribeDeliveryChannels(&configservice.DescribeDeliveryChannelsInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.DeliveryChannels {
		logDebug("Got ConfigDeliveryChannel resource with PhysicalResourceId", *resource.Name)
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getConfigOrganizationConfigRule(client *configservice.ConfigService) (r resourceSliceError) {
	logDebug("Listing ConfigOrganizationConfigRule resources")
	input := configservice.DescribeOrganizationConfigRulesInput{}
	for {
		page, err := client.DescribeOrganizationConfigRules(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.OrganizationConfigRules {
			logDebug("Got ConfigOrganizationConfigRule resource with PhysicalResourceId", *resource.OrganizationConfigRuleName)
			r.resources = append(r.resources, *resource.OrganizationConfigRuleName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigOrganizationConformancePack(client *configservice.ConfigService) (r resourceSliceError) {
	logDebug("Listing ConfigOrganizationConformancePack resources")
	input := configservice.DescribeOrganizationConformancePacksInput{}
	for {
		page, err := client.DescribeOrganizationConformancePacks(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.OrganizationConformancePacks {
			logDebug("Got ConfigOrganizationConformancePack resource with PhysicalResourceId", *resource.OrganizationConformancePackName)
			r.resources = append(r.resources, *resource.OrganizationConformancePackName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getConfigRemediationConfiguration(client *configservice.ConfigService) (r resourceSliceError) {
	logDebug("Listing ConfigRemediationConfiguration resources")
	page, err := client.DescribeRemediationConfigurations(&configservice.DescribeRemediationConfigurationsInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.RemediationConfigurations {
		logDebug("Got ConfigRemediationConfiguration resource with PhysicalResourceId", *resource.Arn)
		r.resources = append(r.resources, *resource.Arn)
	}
	return
}

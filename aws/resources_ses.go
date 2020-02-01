package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func getSes(session *session.Session) (resources resourceMap) {
	client := ses.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		sesConfigurationSet: getSesConfigurationSet(client),
		sesReceiptFilter:    getSesReceiptFilter(client),
		sesReceiptRuleSet:   getSesReceiptRuleSet(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getSesConfigurationSet(client *ses.SES) (r resourceSliceError) {
	input := ses.ListConfigurationSetsInput{}
	for {
		page, err := client.ListConfigurationSets(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing SesConfigurationSet resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ConfigurationSets {
			logDebug("Got SesConfigurationSet resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSesReceiptFilter(client *ses.SES) (r resourceSliceError) {
	logDebug("Listing SesReceiptFilter resources")
	page, err := client.ListReceiptFilters(&ses.ListReceiptFiltersInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.Filters {
		logDebug("Got SesReceiptFilter resource with PhysicalResourceId", *resource.Name)
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getSesReceiptRuleSet(client *ses.SES) (r resourceSliceError) {
	input := ses.ListReceiptRuleSetsInput{}
	for {
		page, err := client.ListReceiptRuleSets(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing SesReceiptRuleSet resources page. Remaining pages", page.NextToken)
		for _, resource := range page.RuleSets {
			logDebug("Got SesReceiptRuleSet resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSesTemplate(client *ses.SES) (r resourceSliceError) {
	input := ses.ListTemplatesInput{}
	for {
		page, err := client.ListTemplates(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing SesTemplate resources page. Remaining pages", page.NextToken)
		for _, resource := range page.TemplatesMetadata {
			logDebug("Got SesTemplate resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

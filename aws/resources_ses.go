package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func getSes(session *session.Session) (resources resourceMap) {
	client := ses.New(session)
	resources = reduce(
		getSesConfigurationSet(client).unwrap(sesConfigurationSet),
		getSesReceiptFilter(client).unwrap(sesReceiptFilter),
		getSesReceiptRuleSet(client).unwrap(sesReceiptRuleSet),
		getSesTemplate(client).unwrap(sesTemplate),
	)
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
		for _, resource := range page.ConfigurationSets {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSesReceiptFilter(client *ses.SES) (r resourceSliceError) {
	page, err := client.ListReceiptFilters(&ses.ListReceiptFiltersInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.Filters {
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
		for _, resource := range page.RuleSets {
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
		for _, resource := range page.TemplatesMetadata {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

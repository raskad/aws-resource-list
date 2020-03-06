package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func getSes(config aws.Config) (resources resourceMap) {
	client := ses.New(config)
	resources = reduce(
		getSesConfigurationSet(client).unwrap(sesConfigurationSet),
		getSesReceiptFilter(client).unwrap(sesReceiptFilter),
		getSesReceiptRuleSet(client).unwrap(sesReceiptRuleSet),
		getSesTemplate(client).unwrap(sesTemplate),
	)
	return
}

func getSesConfigurationSet(client *ses.Client) (r resourceSliceError) {
	input := ses.ListConfigurationSetsInput{}
	for {
		page, err := client.ListConfigurationSetsRequest(&input).Send(context.Background())
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

func getSesReceiptFilter(client *ses.Client) (r resourceSliceError) {
	page, err := client.ListReceiptFiltersRequest(&ses.ListReceiptFiltersInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.Filters {
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getSesReceiptRuleSet(client *ses.Client) (r resourceSliceError) {
	input := ses.ListReceiptRuleSetsInput{}
	for {
		page, err := client.ListReceiptRuleSetsRequest(&input).Send(context.Background())
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

func getSesTemplate(client *ses.Client) (r resourceSliceError) {
	input := ses.ListTemplatesInput{}
	for {
		page, err := client.ListTemplatesRequest(&input).Send(context.Background())
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

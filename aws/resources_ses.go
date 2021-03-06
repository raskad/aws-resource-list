package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func getSes(config aws.Config) (resources awsResourceMap) {
	client := ses.New(config)

	sesConfigurationSetNames := getSesConfigurationSetNames(client)
	sesDomainIdentityNames := getSesDomainIdentityNames(client)
	sesEmailIdentityNames := getSesEmailIdentityNames(client)
	sesReceiptFilterNames := getSesReceiptFilterNames(client)
	sesReceiptRuleSetNames := getSesReceiptRuleSetNames(client)
	sesTemplateNames := getSesTemplateNames(client)

	resources = awsResourceMap{
		sesConfigurationSet: sesConfigurationSetNames,
		sesDomainIdentity:   sesDomainIdentityNames,
		sesEmailIdentity:    sesEmailIdentityNames,
		sesReceiptFilter:    sesReceiptFilterNames,
		sesReceiptRuleSet:   sesReceiptRuleSetNames,
		sesTemplate:         sesTemplateNames,
	}
	return
}

func getSesConfigurationSetNames(client *ses.Client) (resources []string) {
	input := ses.ListConfigurationSetsInput{}
	for {
		page, err := client.ListConfigurationSetsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ConfigurationSets {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSesDomainIdentityNames(client *ses.Client) (resources []string) {
	input := ses.ListIdentitiesInput{
		IdentityType: ses.IdentityTypeDomain,
	}
	for {
		page, err := client.ListIdentitiesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		resources = append(resources, page.Identities...)
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSesEmailIdentityNames(client *ses.Client) (resources []string) {
	input := ses.ListIdentitiesInput{
		IdentityType: ses.IdentityTypeEmailAddress,
	}
	for {
		page, err := client.ListIdentitiesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		resources = append(resources, page.Identities...)
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSesReceiptFilterNames(client *ses.Client) (resources []string) {
	page, err := client.ListReceiptFiltersRequest(&ses.ListReceiptFiltersInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.Filters {
		resources = append(resources, *resource.Name)
	}
	return
}

func getSesReceiptRuleSetNames(client *ses.Client) (resources []string) {
	input := ses.ListReceiptRuleSetsInput{}
	for {
		page, err := client.ListReceiptRuleSetsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.RuleSets {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSesTemplateNames(client *ses.Client) (resources []string) {
	input := ses.ListTemplatesInput{}
	for {
		page, err := client.ListTemplatesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.TemplatesMetadata {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

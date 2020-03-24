package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
)

func getWafv2(config aws.Config) (resources resourceMap) {
	client := wafv2.New(config)

	wafv2IPSetIDs := getWafv2IPSetIDs(client)
	wafv2RegexPatternSetIDs := getWafv2RegexPatternSetIDs(client)
	wafv2RuleGroupIDs := getWafv2RuleGroupIDs(client)
	wafv2WebACLIDs := getWafv2WebACLIDs(client)

	resources = resourceMap{
		wafv2IPSet:           wafv2IPSetIDs,
		wafv2RegexPatternSet: wafv2RegexPatternSetIDs,
		wafv2RuleGroup:       wafv2RuleGroupIDs,
		wafv2WebACL:          wafv2WebACLIDs,
	}
	return
}

func getWafv2IPSetIDs(client *wafv2.Client) (resources []string) {
	input := wafv2.ListIPSetsInput{
		Scope: wafv2.ScopeRegional,
	}
	for {
		page, err := client.ListIPSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.IPSets {
			resources = append(resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafv2RegexPatternSetIDs(client *wafv2.Client) (resources []string) {
	input := wafv2.ListRegexPatternSetsInput{
		Scope: wafv2.ScopeRegional,
	}
	for {
		page, err := client.ListRegexPatternSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.RegexPatternSets {
			resources = append(resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafv2RuleGroupIDs(client *wafv2.Client) (resources []string) {
	input := wafv2.ListRuleGroupsInput{
		Scope: wafv2.ScopeRegional,
	}
	for {
		page, err := client.ListRuleGroupsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.RuleGroups {
			resources = append(resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafv2WebACLIDs(client *wafv2.Client) (resources []string) {
	input := wafv2.ListWebACLsInput{
		Scope: wafv2.ScopeRegional,
	}
	for {
		page, err := client.ListWebACLsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.WebACLs {
			resources = append(resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
)

func getWafv2(config aws.Config) (resources resourceMap) {
	client := wafv2.New(config)
	resources = reduce(
		getWafv2IPSet(client).unwrap(wafv2IPSet),
		getWafv2RegexPatternSet(client).unwrap(wafv2RegexPatternSet),
		getWafv2RuleGroup(client).unwrap(wafv2RuleGroup),
		getWafv2WebACL(client).unwrap(wafv2WebACL),
	)
	return
}

func getWafv2IPSet(client *wafv2.Client) (r resourceSliceError) {
	input := wafv2.ListIPSetsInput{
		Scope: wafv2.ScopeRegional,
	}
	for {
		page, err := client.ListIPSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.IPSets {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafv2RegexPatternSet(client *wafv2.Client) (r resourceSliceError) {
	input := wafv2.ListRegexPatternSetsInput{
		Scope: wafv2.ScopeRegional,
	}
	for {
		page, err := client.ListRegexPatternSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.RegexPatternSets {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafv2RuleGroup(client *wafv2.Client) (r resourceSliceError) {
	input := wafv2.ListRuleGroupsInput{
		Scope: wafv2.ScopeRegional,
	}
	for {
		page, err := client.ListRuleGroupsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.RuleGroups {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafv2WebACL(client *wafv2.Client) (r resourceSliceError) {
	input := wafv2.ListWebACLsInput{
		Scope: wafv2.ScopeRegional,
	}
	for {
		page, err := client.ListWebACLsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.WebACLs {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

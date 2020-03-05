package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/wafv2"
)

func getWafv2(session *session.Session) (resources resourceMap) {
	client := wafv2.New(session)
	resources = reduce(
		getWafv2IPSet(client).unwrap(wafv2IPSet),
		getWafv2RegexPatternSet(client).unwrap(wafv2RegexPatternSet),
		getWafv2RuleGroup(client).unwrap(wafv2RuleGroup),
		getWafv2WebACL(client).unwrap(wafv2WebACL),
	)
	return
}

func getWafv2IPSet(client *wafv2.WAFV2) (r resourceSliceError) {
	input := wafv2.ListIPSetsInput{
		Scope: aws.String(wafv2.ScopeRegional),
	}
	for {
		page, err := client.ListIPSets(&input)
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

func getWafv2RegexPatternSet(client *wafv2.WAFV2) (r resourceSliceError) {
	input := wafv2.ListRegexPatternSetsInput{
		Scope: aws.String(wafv2.ScopeRegional),
	}
	for {
		page, err := client.ListRegexPatternSets(&input)
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

func getWafv2RuleGroup(client *wafv2.WAFV2) (r resourceSliceError) {
	input := wafv2.ListRuleGroupsInput{
		Scope: aws.String(wafv2.ScopeRegional),
	}
	for {
		page, err := client.ListRuleGroups(&input)
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

func getWafv2WebACL(client *wafv2.WAFV2) (r resourceSliceError) {
	input := wafv2.ListWebACLsInput{
		Scope: aws.String(wafv2.ScopeRegional),
	}
	for {
		page, err := client.ListWebACLs(&input)
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

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func getWafRegional(config aws.Config) (resources resourceMap) {
	client := wafregional.New(config)

	wafRegionalByteMatchSetIDs := getWafRegionalByteMatchSetIDs(client)
	wafRegionalGeoMatchSetIDs := getWafRegionalGeoMatchSetIDs(client)
	wafRegionalIPSetIDs := getWafRegionalIPSetIDs(client)
	wafRegionalRateBasedRuleIDs := getWafRegionalRateBasedRuleIDs(client)
	wafRegionalRegexPatternSetIDs := getWafRegionalRegexPatternSetIDs(client)
	wafRegionalRuleIDs := getWafRegionalRuleIDs(client)
	wafRegionalSizeConstraintSetIDs := getWafRegionalSizeConstraintSetIDs(client)
	wafRegionalSQLInjectionMatchSetIDs := getWafRegionalSQLInjectionMatchSetIDs(client)
	wafRegionalWebACLIDs := getWafRegionalWebACLIDs(client)
	wafRegionalXSSMatchSetIDs := getWafRegionalXSSMatchSetIDs(client)

	resources = resourceMap{
		wafRegionalByteMatchSet:         wafRegionalByteMatchSetIDs,
		wafRegionalGeoMatchSet:          wafRegionalGeoMatchSetIDs,
		wafRegionalIPSet:                wafRegionalIPSetIDs,
		wafRegionalRateBasedRule:        wafRegionalRateBasedRuleIDs,
		wafRegionalRegexPatternSet:      wafRegionalRegexPatternSetIDs,
		wafRegionalRule:                 wafRegionalRuleIDs,
		wafRegionalSizeConstraintSet:    wafRegionalSizeConstraintSetIDs,
		wafRegionalSQLInjectionMatchSet: wafRegionalSQLInjectionMatchSetIDs,
		wafRegionalWebACL:               wafRegionalWebACLIDs,
		wafRegionalXSSMatchSet:          wafRegionalXSSMatchSetIDs,
	}
	return
}

func getWafRegionalByteMatchSetIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListByteMatchSetsInput{}
	for {
		page, err := client.ListByteMatchSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.ByteMatchSets {
			resources = append(resources, *resource.ByteMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalGeoMatchSetIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListGeoMatchSetsInput{}
	for {
		page, err := client.ListGeoMatchSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.GeoMatchSets {
			resources = append(resources, *resource.GeoMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalIPSetIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListIPSetsInput{}
	for {
		page, err := client.ListIPSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.IPSets {
			resources = append(resources, *resource.IPSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalRateBasedRuleIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListRateBasedRulesInput{}
	for {
		page, err := client.ListRateBasedRulesRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Rules {
			resources = append(resources, *resource.RuleId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalRegexPatternSetIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListRegexPatternSetsInput{}
	for {
		page, err := client.ListRegexPatternSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.RegexPatternSets {
			resources = append(resources, *resource.RegexPatternSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalRuleIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListRulesInput{}
	for {
		page, err := client.ListRulesRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Rules {
			resources = append(resources, *resource.RuleId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalSizeConstraintSetIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListSizeConstraintSetsInput{}
	for {
		page, err := client.ListSizeConstraintSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.SizeConstraintSets {
			resources = append(resources, *resource.SizeConstraintSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalSQLInjectionMatchSetIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListSqlInjectionMatchSetsInput{}
	for {
		page, err := client.ListSqlInjectionMatchSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.SqlInjectionMatchSets {
			resources = append(resources, *resource.SqlInjectionMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalWebACLIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListWebACLsInput{}
	for {
		page, err := client.ListWebACLsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.WebACLs {
			resources = append(resources, *resource.WebACLId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalXSSMatchSetIDs(client *wafregional.Client) (resources []string) {
	input := wafregional.ListXssMatchSetsInput{}
	for {
		page, err := client.ListXssMatchSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.XssMatchSets {
			resources = append(resources, *resource.XssMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

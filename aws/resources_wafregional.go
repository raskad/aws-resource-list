package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func getWafRegional(config aws.Config) (resources resourceMap) {
	client := wafregional.New(config)
	resources = reduce(
		getWafRegionalByteMatchSet(client).unwrap(wafRegionalByteMatchSet),
		getWafRegionalGeoMatchSet(client).unwrap(wafRegionalGeoMatchSet),
		getWafRegionalIPSet(client).unwrap(wafRegionalIPSet),
		getWafRegionalRateBasedRule(client).unwrap(wafRegionalRateBasedRule),
		getWafRegionalRegexPatternSet(client).unwrap(wafRegionalRegexPatternSet),
		getWafRegionalRule(client).unwrap(wafRegionalRule),
		getWafRegionalSizeConstraintSet(client).unwrap(wafRegionalSizeConstraintSet),
		getWafRegionalSQLInjectionMatchSet(client).unwrap(wafRegionalSQLInjectionMatchSet),
		getWafRegionalWebACL(client).unwrap(wafRegionalWebACL),
		getWafRegionalXSSMatchSet(client).unwrap(wafRegionalXSSMatchSet),
	)
	return
}

func getWafRegionalByteMatchSet(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListByteMatchSetsInput{}
	for {
		page, err := client.ListByteMatchSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ByteMatchSets {
			r.resources = append(r.resources, *resource.ByteMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalGeoMatchSet(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListGeoMatchSetsInput{}
	for {
		page, err := client.ListGeoMatchSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.GeoMatchSets {
			r.resources = append(r.resources, *resource.GeoMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalIPSet(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListIPSetsInput{}
	for {
		page, err := client.ListIPSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.IPSets {
			r.resources = append(r.resources, *resource.IPSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalRateBasedRule(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListRateBasedRulesInput{}
	for {
		page, err := client.ListRateBasedRulesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Rules {
			r.resources = append(r.resources, *resource.RuleId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalRegexPatternSet(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListRegexPatternSetsInput{}
	for {
		page, err := client.ListRegexPatternSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.RegexPatternSets {
			r.resources = append(r.resources, *resource.RegexPatternSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalRule(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListRulesInput{}
	for {
		page, err := client.ListRulesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Rules {
			r.resources = append(r.resources, *resource.RuleId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalSizeConstraintSet(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListSizeConstraintSetsInput{}
	for {
		page, err := client.ListSizeConstraintSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.SizeConstraintSets {
			r.resources = append(r.resources, *resource.SizeConstraintSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalSQLInjectionMatchSet(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListSqlInjectionMatchSetsInput{}
	for {
		page, err := client.ListSqlInjectionMatchSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.SqlInjectionMatchSets {
			r.resources = append(r.resources, *resource.SqlInjectionMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalWebACL(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListWebACLsInput{}
	for {
		page, err := client.ListWebACLsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.WebACLs {
			r.resources = append(r.resources, *resource.WebACLId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRegionalXSSMatchSet(client *wafregional.Client) (r resourceSliceError) {
	input := wafregional.ListXssMatchSetsInput{}
	for {
		page, err := client.ListXssMatchSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.XssMatchSets {
			r.resources = append(r.resources, *resource.XssMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
)

func getWafRegional(session *session.Session) (resources resourceMap) {
	client := wafregional.New(session)
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

func getWafRegionalByteMatchSet(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListByteMatchSetsInput{}
	for {
		page, err := client.ListByteMatchSets(&input)
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

func getWafRegionalGeoMatchSet(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListGeoMatchSetsInput{}
	for {
		page, err := client.ListGeoMatchSets(&input)
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

func getWafRegionalIPSet(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListIPSetsInput{}
	for {
		page, err := client.ListIPSets(&input)
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

func getWafRegionalRateBasedRule(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListRateBasedRulesInput{}
	for {
		page, err := client.ListRateBasedRules(&input)
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

func getWafRegionalRegexPatternSet(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListRegexPatternSetsInput{}
	for {
		page, err := client.ListRegexPatternSets(&input)
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

func getWafRegionalRule(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListRulesInput{}
	for {
		page, err := client.ListRules(&input)
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

func getWafRegionalSizeConstraintSet(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListSizeConstraintSetsInput{}
	for {
		page, err := client.ListSizeConstraintSets(&input)
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

func getWafRegionalSQLInjectionMatchSet(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListSqlInjectionMatchSetsInput{}
	for {
		page, err := client.ListSqlInjectionMatchSets(&input)
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

func getWafRegionalWebACL(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListWebACLsInput{}
	for {
		page, err := client.ListWebACLs(&input)
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

func getWafRegionalXSSMatchSet(client *wafregional.WAFRegional) (r resourceSliceError) {
	input := waf.ListXssMatchSetsInput{}
	for {
		page, err := client.ListXssMatchSets(&input)
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

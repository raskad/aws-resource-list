package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
)

func getWafRegional(session *session.Session) (resources resourceMap) {
	client := wafregional.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		wafRegionalByteMatchSet:         getWafRegionalByteMatchSet(client),
		wafRegionalGeoMatchSet:          getWafRegionalGeoMatchSet(client),
		wafRegionalIPSet:                getWafRegionalIPSet(client),
		wafRegionalRateBasedRule:        getWafRegionalRateBasedRule(client),
		wafRegionalRegexPatternSet:      getWafRegionalRegexPatternSet(client),
		wafRegionalRule:                 getWafRegionalRule(client),
		wafRegionalSizeConstraintSet:    getWafRegionalSizeConstraintSet(client),
		wafRegionalSQLInjectionMatchSet: getWafRegionalSQLInjectionMatchSet(client),
		wafRegionalWebACL:               getWafRegionalWebACL(client),
		wafRegionalXSSMatchSet:          getWafRegionalXSSMatchSet(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
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
		logDebug("Listing WafRegionalByteMatchSet resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.ByteMatchSets {
			logDebug("Got WafRegionalByteMatchSet resource with PhysicalResourceId", *resource.ByteMatchSetId)
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
		logDebug("Listing WafRegionalGeoMatchSet resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.GeoMatchSets {
			logDebug("Got WafRegionalGeoMatchSet resource with PhysicalResourceId", *resource.GeoMatchSetId)
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
		logDebug("Listing WafRegionalIPSet resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.IPSets {
			logDebug("Got WafRegionalIPSet resource with PhysicalResourceId", *resource.IPSetId)
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
		logDebug("Listing WafRegionalRateBasedRule resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.Rules {
			logDebug("Got WafRegionalRateBasedRule resource with PhysicalResourceId", *resource.RuleId)
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
		logDebug("Listing WafRegionalRegexPatternSet resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.RegexPatternSets {
			logDebug("Got WafRegionalRegexPatternSet resource with PhysicalResourceId", *resource.RegexPatternSetId)
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
		logDebug("Listing WafRegionalRule resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.Rules {
			logDebug("Got WafRegionalRule resource with PhysicalResourceId", *resource.RuleId)
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
		logDebug("Listing WafRegionalSizeConstraintSet resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.SizeConstraintSets {
			logDebug("Got WafRegionalSizeConstraintSet resource with PhysicalResourceId", *resource.SizeConstraintSetId)
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
		logDebug("Listing WafRegionalSQLInjectionMatchSet resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.SqlInjectionMatchSets {
			logDebug("Got WafRegionalSQLInjectionMatchSet resource with PhysicalResourceId", *resource.SqlInjectionMatchSetId)
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
		logDebug("Listing WafRegionalWebACL resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.WebACLs {
			logDebug("Got WafRegionalWebACL resource with PhysicalResourceId", *resource.WebACLId)
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
		logDebug("Listing WafRegionalXSSMatchSet resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.XssMatchSets {
			logDebug("Got WafRegionalXSSMatchSet resource with PhysicalResourceId", *resource.XssMatchSetId)
			r.resources = append(r.resources, *resource.XssMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

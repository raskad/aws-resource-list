package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
)

func getWaf(session *session.Session) (resources resourceMap) {
	client := waf.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		wafByteMatchSet:         getWafByteMatchSet(client),
		wafIPSet:                getWafIPSet(client),
		wafRule:                 getWafRule(client),
		wafSizeConstraintSet:    getWafSizeConstraintSet(client),
		wafSQLInjectionMatchSet: getWafSQLInjectionMatchSet(client),
		wafWebACL:               getWafWebACL(client),
		wafXSSMatchSet:          getWafXSSMatchSet(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getWafByteMatchSet(client *waf.WAF) (r resourceSliceError) {
	logDebug("Listing WafByteMatchSet resources")
	input := waf.ListByteMatchSetsInput{}
	for {
		page, err := client.ListByteMatchSets(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ByteMatchSets {
			logDebug("Got WafByteMatchSet resource with PhysicalResourceId", *resource.ByteMatchSetId)
			r.resources = append(r.resources, *resource.ByteMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafIPSet(client *waf.WAF) (r resourceSliceError) {
	logDebug("Listing WafIPSet resources")
	input := waf.ListIPSetsInput{}
	for {
		page, err := client.ListIPSets(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.IPSets {
			logDebug("Got WafIPSet resource with PhysicalResourceId", *resource.IPSetId)
			r.resources = append(r.resources, *resource.IPSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafRule(client *waf.WAF) (r resourceSliceError) {
	logDebug("Listing WafRule resources")
	input := waf.ListRulesInput{}
	for {
		page, err := client.ListRules(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Rules {
			logDebug("Got WafRule resource with PhysicalResourceId", *resource.RuleId)
			r.resources = append(r.resources, *resource.RuleId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafSizeConstraintSet(client *waf.WAF) (r resourceSliceError) {
	logDebug("Listing WafSizeConstraintSet resources")
	input := waf.ListSizeConstraintSetsInput{}
	for {
		page, err := client.ListSizeConstraintSets(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.SizeConstraintSets {
			logDebug("Got WafSizeConstraintSet resource with PhysicalResourceId", *resource.SizeConstraintSetId)
			r.resources = append(r.resources, *resource.SizeConstraintSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafSQLInjectionMatchSet(client *waf.WAF) (r resourceSliceError) {
	logDebug("Listing WafSQLInjectionMatchSet resources")
	input := waf.ListSqlInjectionMatchSetsInput{}
	for {
		page, err := client.ListSqlInjectionMatchSets(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.SqlInjectionMatchSets {
			logDebug("Got WafSQLInjectionMatchSet resource with PhysicalResourceId", *resource.SqlInjectionMatchSetId)
			r.resources = append(r.resources, *resource.SqlInjectionMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafWebACL(client *waf.WAF) (r resourceSliceError) {
	logDebug("Listing WafWebACL resources")
	input := waf.ListWebACLsInput{}
	for {
		page, err := client.ListWebACLs(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.WebACLs {
			logDebug("Got WafWebACL resource with PhysicalResourceId", *resource.WebACLId)
			r.resources = append(r.resources, *resource.WebACLId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafXSSMatchSet(client *waf.WAF) (r resourceSliceError) {
	logDebug("Listing WafXSSMatchSet resources")
	input := waf.ListXssMatchSetsInput{}
	for {
		page, err := client.ListXssMatchSets(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.XssMatchSets {
			logDebug("Got WafXSSMatchSet resource with PhysicalResourceId", *resource.XssMatchSetId)
			r.resources = append(r.resources, *resource.XssMatchSetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

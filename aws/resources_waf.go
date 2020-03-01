package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
)

func getWaf(session *session.Session) (resources resourceMap) {
	client := waf.New(session)
	resources = reduce(
		getWafByteMatchSet(client).unwrap(wafByteMatchSet),
		getWafIPSet(client).unwrap(wafIPSet),
		getWafRule(client).unwrap(wafRule),
		getWafSizeConstraintSet(client).unwrap(wafSizeConstraintSet),
		getWafSQLInjectionMatchSet(client).unwrap(wafSQLInjectionMatchSet),
		getWafWebACL(client).unwrap(wafWebACL),
		getWafXSSMatchSet(client).unwrap(wafXSSMatchSet),
	)
	return
}

func getWafByteMatchSet(client *waf.WAF) (r resourceSliceError) {
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

func getWafIPSet(client *waf.WAF) (r resourceSliceError) {
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

func getWafRule(client *waf.WAF) (r resourceSliceError) {
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

func getWafSizeConstraintSet(client *waf.WAF) (r resourceSliceError) {
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

func getWafSQLInjectionMatchSet(client *waf.WAF) (r resourceSliceError) {
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

func getWafWebACL(client *waf.WAF) (r resourceSliceError) {
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

func getWafXSSMatchSet(client *waf.WAF) (r resourceSliceError) {
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

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/wafv2"
)

func getWafv2(session *session.Session) (resources resourceMap) {
	client := wafv2.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		wafv2IPSet:           getWafv2IPSet(client),
		wafv2RegexPatternSet: getWafv2RegexPatternSet(client),
		wafv2RuleGroup:       getWafv2RuleGroup(client),
		wafv2WebACL:          getWafv2WebACL(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getWafv2IPSet(client *wafv2.WAFV2) (r resourceSliceError) {
	input := wafv2.ListIPSetsInput{}
	for {
		page, err := client.ListIPSets(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing Wafv2IPSet resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.IPSets {
			logDebug("Got Wafv2IPSet resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafv2RegexPatternSet(client *wafv2.WAFV2) (r resourceSliceError) {
	input := wafv2.ListRegexPatternSetsInput{}
	for {
		page, err := client.ListRegexPatternSets(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing Wafv2RegexPatternSet resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.RegexPatternSets {
			logDebug("Got Wafv2RegexPatternSet resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafv2RuleGroup(client *wafv2.WAFV2) (r resourceSliceError) {
	input := wafv2.ListRuleGroupsInput{}
	for {
		page, err := client.ListRuleGroups(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing Wafv2RuleGroup resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.RuleGroups {
			logDebug("Got Wafv2RuleGroup resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

func getWafv2WebACL(client *wafv2.WAFV2) (r resourceSliceError) {
	input := wafv2.ListWebACLsInput{}
	for {
		page, err := client.ListWebACLs(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing Wafv2WebACL resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.WebACLs {
			logDebug("Got Wafv2WebACL resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextMarker == nil {
			return
		}
		input.NextMarker = page.NextMarker
	}
}

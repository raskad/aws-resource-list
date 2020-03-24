package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func getWaf(config aws.Config) (resources resourceMap) {
	client := waf.New(config)

	wafByteMatchSetIDs := getWafByteMatchSetIDs(client)
	wafIPSetIDs := getWafIPSetIDs(client)
	wafRuleIDs := getWafRuleIDs(client)
	wafSizeConstraintSetIDs := getWafSizeConstraintSetIDs(client)
	wafSQLInjectionMatchSetIDs := getWafSQLInjectionMatchSetIDs(client)
	wafWebACLIDs := getWafWebACLIDs(client)
	wafXSSMatchSetIDs := getWafXSSMatchSetIDs(client)

	resources = resourceMap{
		wafByteMatchSet:         wafByteMatchSetIDs,
		wafIPSet:                wafIPSetIDs,
		wafRule:                 wafRuleIDs,
		wafSizeConstraintSet:    wafSizeConstraintSetIDs,
		wafSQLInjectionMatchSet: wafSQLInjectionMatchSetIDs,
		wafWebACL:               wafWebACLIDs,
		wafXSSMatchSet:          wafXSSMatchSetIDs,
	}
	return
}

func getWafByteMatchSetIDs(client *waf.Client) (resources []string) {
	input := waf.ListByteMatchSetsInput{}
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

func getWafIPSetIDs(client *waf.Client) (resources []string) {
	input := waf.ListIPSetsInput{}
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

func getWafRuleIDs(client *waf.Client) (resources []string) {
	input := waf.ListRulesInput{}
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

func getWafSizeConstraintSetIDs(client *waf.Client) (resources []string) {
	input := waf.ListSizeConstraintSetsInput{}
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

func getWafSQLInjectionMatchSetIDs(client *waf.Client) (resources []string) {
	input := waf.ListSqlInjectionMatchSetsInput{}
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

func getWafWebACLIDs(client *waf.Client) (resources []string) {
	input := waf.ListWebACLsInput{}
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

func getWafXSSMatchSetIDs(client *waf.Client) (resources []string) {
	input := waf.ListXssMatchSetsInput{}
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

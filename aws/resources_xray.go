package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

func getXray(config aws.Config) (resources awsResourceMap) {
	client := xray.New(config)

	xraySamplingRuleNames := getXraySamplingRuleNames(client)

	resources = awsResourceMap{
		xraySamplingRule: xraySamplingRuleNames,
	}

	return
}

func getXraySamplingRuleNames(client *xray.Client) (resources []string) {
	req := client.GetSamplingRulesRequest(&xray.GetSamplingRulesInput{})
	p := xray.NewGetSamplingRulesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.SamplingRuleRecords {
			resources = append(resources, *resource.SamplingRule.RuleName)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

func getAccessAnalyzer(config aws.Config) (resources resourceMap) {
	client := accessanalyzer.New(config)

	accessAnalyzerAnalyzerNames := getAccessAnalyzerAnalyzerNames(client)

	resources = resourceMap{
		accessAnalyzerAnalyzer: accessAnalyzerAnalyzerNames,
	}
	return
}

func getAccessAnalyzerAnalyzerNames(client *accessanalyzer.Client) (resources []string) {
	req := client.ListAnalyzersRequest(&accessanalyzer.ListAnalyzersInput{})
	p := accessanalyzer.NewListAnalyzersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Analyzers {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

func getAccessAnalyzer(config aws.Config) (resources resourceMap) {
	client := accessanalyzer.New(config)
	resources = reduce(
		getAccessAnalyzerAnalyzer(client).unwrap(accessAnalyzerAnalyzer),
	)
	return
}

func getAccessAnalyzerAnalyzer(client *accessanalyzer.Client) (r resourceSliceError) {
	req := client.ListAnalyzersRequest(&accessanalyzer.ListAnalyzersInput{})
	p := accessanalyzer.NewListAnalyzersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Analyzers {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

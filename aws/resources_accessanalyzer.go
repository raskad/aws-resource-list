package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
)

func getAccessAnalyzer(session *session.Session) (resources resourceMap) {
	client := accessanalyzer.New(session)
	resources = reduce(
		getAccessAnalyzerAnalyzer(client).unwrap(accessAnalyzerAnalyzer),
	)
	return
}

func getAccessAnalyzerAnalyzer(client *accessanalyzer.AccessAnalyzer) (r resourceSliceError) {
	r.err = client.ListAnalyzersPages(&accessanalyzer.ListAnalyzersInput{}, func(page *accessanalyzer.ListAnalyzersOutput, lastPage bool) bool {
		for _, resource := range page.Analyzers {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

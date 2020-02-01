package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
)

func getAccessAnalyzer(session *session.Session) (resources resourceMap) {
	client := accessanalyzer.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		accessAnalyzerAnalyzer: getAccessAnalyzerAnalyzer(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAccessAnalyzerAnalyzer(client *accessanalyzer.AccessAnalyzer) (r resourceSliceError) {
	r.err = client.ListAnalyzersPages(&accessanalyzer.ListAnalyzersInput{}, func(page *accessanalyzer.ListAnalyzersOutput, lastPage bool) bool {
		logDebug("List accessAnalyzerAnalyzer resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Analyzers {
			logDebug("Got accessAnalyzerAnalyzer resource with PhysicalResourceId", *resource.Arn)
			r.resources = append(r.resources, *resource.Arn)
		}
		return true
	})
	return
}

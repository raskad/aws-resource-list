package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
)

func getInspector(session *session.Session) (resources resourceMap) {
	client := inspector.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		inspectorAssessmentTarget:   getInspectorAssessmentTarget(client),
		inspectorAssessmentTemplate: getInspectorAssessmentTemplate(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getInspectorAssessmentTarget(client *inspector.Inspector) (r resourceSliceError) {
	r.err = client.ListAssessmentTargetsPages(&inspector.ListAssessmentTargetsInput{}, func(page *inspector.ListAssessmentTargetsOutput, lastPage bool) bool {
		logDebug("Listing InspectorAssessmentTarget resources page. Remaining pages", page.NextToken)
		for _, resource := range page.AssessmentTargetArns {
			logDebug("Got InspectorAssessmentTarget resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getInspectorAssessmentTemplate(client *inspector.Inspector) (r resourceSliceError) {
	r.err = client.ListAssessmentTemplatesPages(&inspector.ListAssessmentTemplatesInput{}, func(page *inspector.ListAssessmentTemplatesOutput, lastPage bool) bool {
		logDebug("Listing InspectorAssessmentTemplate resources page. Remaining pages", page.NextToken)
		for _, resource := range page.AssessmentTemplateArns {
			logDebug("Got InspectorAssessmentTemplate resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

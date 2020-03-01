package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
)

func getInspector(session *session.Session) (resources resourceMap) {
	client := inspector.New(session)
	resources = reduce(
		getInspectorAssessmentTarget(client).unwrap(inspectorAssessmentTarget),
		getInspectorAssessmentTemplate(client).unwrap(inspectorAssessmentTemplate),
	)
	return
}

func getInspectorAssessmentTarget(client *inspector.Inspector) (r resourceSliceError) {
	r.err = client.ListAssessmentTargetsPages(&inspector.ListAssessmentTargetsInput{}, func(page *inspector.ListAssessmentTargetsOutput, lastPage bool) bool {
		for _, resource := range page.AssessmentTargetArns {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getInspectorAssessmentTemplate(client *inspector.Inspector) (r resourceSliceError) {
	r.err = client.ListAssessmentTemplatesPages(&inspector.ListAssessmentTemplatesInput{}, func(page *inspector.ListAssessmentTemplatesOutput, lastPage bool) bool {
		for _, resource := range page.AssessmentTemplateArns {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

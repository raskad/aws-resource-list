package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
)

func getInspector(config aws.Config) (resources resourceMap) {
	client := inspector.New(config)
	resources = reduce(
		getInspectorAssessmentTarget(client).unwrap(inspectorAssessmentTarget),
		getInspectorAssessmentTemplate(client).unwrap(inspectorAssessmentTemplate),
	)
	return
}

func getInspectorAssessmentTarget(client *inspector.Client) (r resourceSliceError) {
	req := client.ListAssessmentTargetsRequest(&inspector.ListAssessmentTargetsInput{})
	p := inspector.NewListAssessmentTargetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.AssessmentTargetArns {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

func getInspectorAssessmentTemplate(client *inspector.Client) (r resourceSliceError) {
	req := client.ListAssessmentTemplatesRequest(&inspector.ListAssessmentTemplatesInput{})
	p := inspector.NewListAssessmentTemplatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.AssessmentTemplateArns {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

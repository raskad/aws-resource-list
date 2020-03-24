package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
)

func getInspector(config aws.Config) (resources resourceMap) {
	client := inspector.New(config)

	inspectorAssessmentTargetARNs := getInspectorAssessmentTargetARNs(client)
	inspectorAssessmentTemplateARNs := getInspectorAssessmentTemplateARNs(client)

	resources = resourceMap{
		inspectorAssessmentTarget:   inspectorAssessmentTargetARNs,
		inspectorAssessmentTemplate: inspectorAssessmentTemplateARNs,
	}
	return
}

func getInspectorAssessmentTargetARNs(client *inspector.Client) (resources []string) {
	req := client.ListAssessmentTargetsRequest(&inspector.ListAssessmentTargetsInput{})
	p := inspector.NewListAssessmentTargetsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.AssessmentTargetArns {
			resources = append(resources, resource)
		}
	}
	return
}

func getInspectorAssessmentTemplateARNs(client *inspector.Client) (resources []string) {
	req := client.ListAssessmentTemplatesRequest(&inspector.ListAssessmentTemplatesInput{})
	p := inspector.NewListAssessmentTemplatesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.AssessmentTemplateArns {
			resources = append(resources, resource)
		}
	}
	return
}

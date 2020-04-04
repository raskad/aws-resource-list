package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
)

func getInspector(config aws.Config) (resources awsResourceMap) {
	client := inspector.New(config)

	inspectorAssessmentTargetARNs := getInspectorAssessmentTargetARNs(client)
	inspectorAssessmentTemplateARNs := getInspectorAssessmentTemplateARNs(client)

	resources = awsResourceMap{
		inspectorAssessmentTarget:   inspectorAssessmentTargetARNs,
		inspectorAssessmentTemplate: inspectorAssessmentTemplateARNs,
	}
	return
}

func getInspectorAssessmentTargetARNs(client *inspector.Client) (resources []string) {
	req := client.ListAssessmentTargetsRequest(&inspector.ListAssessmentTargetsInput{})
	p := inspector.NewListAssessmentTargetsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.AssessmentTargetArns...)
	}
	return
}

func getInspectorAssessmentTemplateARNs(client *inspector.Client) (resources []string) {
	req := client.ListAssessmentTemplatesRequest(&inspector.ListAssessmentTemplatesInput{})
	p := inspector.NewListAssessmentTemplatesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.AssessmentTemplateArns...)
	}
	return
}

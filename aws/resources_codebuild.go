package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

func getCodeBuild(session *session.Session) (resources resourceMap) {
	client := codebuild.New(session)
	resources = reduce(
		getCodeBuildProject(client).unwrap(codeBuildProject),
		getCodeBuildReportGroup(client).unwrap(codeBuildReportGroup),
		getCodeBuildSourceCredential(client).unwrap(codeBuildSourceCredential),
	)
	return
}

func getCodeBuildProject(client *codebuild.CodeBuild) (r resourceSliceError) {
	input := codebuild.ListProjectsInput{}
	for {
		page, err := client.ListProjects(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Projects {
			r.resources = append(r.resources, *resource)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getCodeBuildReportGroup(client *codebuild.CodeBuild) (r resourceSliceError) {
	input := codebuild.ListReportGroupsInput{}
	for {
		page, err := client.ListReportGroups(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ReportGroups {
			r.resources = append(r.resources, *resource)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getCodeBuildSourceCredential(client *codebuild.CodeBuild) (r resourceSliceError) {
	input := codebuild.ListSourceCredentialsInput{}
	page, err := client.ListSourceCredentials(&input)
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.SourceCredentialsInfos {
		r.resources = append(r.resources, *resource.Arn)
	}
	return
}

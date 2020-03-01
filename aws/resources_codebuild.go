package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

func getCodeBuild(session *session.Session) (resources resourceMap) {
	client := codebuild.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		codeBuildProject:          getCodeBuildProject(client),
		codeBuildReportGroup:      getCodeBuildReportGroup(client),
		codeBuildSourceCredential: getCodeBuildSourceCredential(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCodeBuildProject(client *codebuild.CodeBuild) (r resourceSliceError) {
	logDebug("Listing CodeBuildProject resources")
	input := codebuild.ListProjectsInput{}
	for {
		page, err := client.ListProjects(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Projects {
			logDebug("Got CodeBuildProject resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getCodeBuildReportGroup(client *codebuild.CodeBuild) (r resourceSliceError) {
	logDebug("Listing CodeBuildReportGroup resources")
	input := codebuild.ListReportGroupsInput{}
	for {
		page, err := client.ListReportGroups(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ReportGroups {
			logDebug("Got CodeBuildReportGroup resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getCodeBuildSourceCredential(client *codebuild.CodeBuild) (r resourceSliceError) {
	logDebug("Listing CodeBuildSourceCredential resources")
	input := codebuild.ListSourceCredentialsInput{}
	page, err := client.ListSourceCredentials(&input)
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.SourceCredentialsInfos {
		logDebug("Got CodeBuildSourceCredential resource with PhysicalResourceId", *resource.Arn)
		r.resources = append(r.resources, *resource.Arn)
	}
	return
}

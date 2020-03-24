package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

func getCodeBuild(config aws.Config) (resources resourceMap) {
	client := codebuild.New(config)

	codeBuildProjectNames := getCodeBuildProjectNames(client)
	codeBuildReportGroupARNs := getCodeBuildReportGroupARNs(client)
	codeBuildSourceCredentialARNs := getCodeBuildSourceCredentialARNs(client)

	resources = resourceMap{
		codeBuildProject:          codeBuildProjectNames,
		codeBuildReportGroup:      codeBuildReportGroupARNs,
		codeBuildSourceCredential: codeBuildSourceCredentialARNs,
	}
	return
}

func getCodeBuildProjectNames(client *codebuild.Client) (resources []string) {
	input := codebuild.ListProjectsInput{}
	for {
		page, err := client.ListProjectsRequest(&input).Send(context.Background())
		logErr(err)
		resources = append(resources, page.Projects...)
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getCodeBuildReportGroupARNs(client *codebuild.Client) (resources []string) {
	input := codebuild.ListReportGroupsInput{}
	for {
		page, err := client.ListReportGroupsRequest(&input).Send(context.Background())
		logErr(err)
		resources = append(resources, page.ReportGroups...)
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getCodeBuildSourceCredentialARNs(client *codebuild.Client) (resources []string) {
	input := codebuild.ListSourceCredentialsInput{}
	page, err := client.ListSourceCredentialsRequest(&input).Send(context.Background())
	logErr(err)
	for _, resource := range page.SourceCredentialsInfos {
		resources = append(resources, *resource.Arn)
	}
	return
}

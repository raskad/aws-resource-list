package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

func getCodeBuild(config aws.Config) (resources resourceMap) {
	client := codebuild.New(config)
	resources = reduce(
		getCodeBuildProject(client).unwrap(codeBuildProject),
		getCodeBuildReportGroup(client).unwrap(codeBuildReportGroup),
		getCodeBuildSourceCredential(client).unwrap(codeBuildSourceCredential),
	)
	return
}

func getCodeBuildProject(client *codebuild.Client) (r resourceSliceError) {
	input := codebuild.ListProjectsInput{}
	for {
		page, err := client.ListProjectsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Projects {
			r.resources = append(r.resources, resource)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getCodeBuildReportGroup(client *codebuild.Client) (r resourceSliceError) {
	input := codebuild.ListReportGroupsInput{}
	for {
		page, err := client.ListReportGroupsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ReportGroups {
			r.resources = append(r.resources, resource)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getCodeBuildSourceCredential(client *codebuild.Client) (r resourceSliceError) {
	input := codebuild.ListSourceCredentialsInput{}
	page, err := client.ListSourceCredentialsRequest(&input).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.SourceCredentialsInfos {
		r.resources = append(r.resources, *resource.Arn)
	}
	return
}

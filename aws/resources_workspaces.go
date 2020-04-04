package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

func getWorkSpaces(config aws.Config) (resources awsResourceMap) {
	client := workspaces.New(config)

	workSpacesDirectoryIDs := getWorkSpacesDirectoryIDs(client)
	workspacesIPGroupIDs := getWorkspacesIPGroupIDs(client)
	workSpacesWorkspaceIDs := getWorkSpacesWorkspaceIDs(client)

	resources = awsResourceMap{
		workSpacesDirectory: workSpacesDirectoryIDs,
		workspacesIPGroup:   workspacesIPGroupIDs,
		workSpacesWorkspace: workSpacesWorkspaceIDs,
	}

	return
}

func getWorkSpacesDirectoryIDs(client *workspaces.Client) (resources []string) {
	req := client.DescribeWorkspaceDirectoriesRequest(&workspaces.DescribeWorkspaceDirectoriesInput{})
	p := workspaces.NewDescribeWorkspaceDirectoriesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Directories {
			resources = append(resources, *resource.DirectoryId)
		}
	}
	return
}

func getWorkspacesIPGroupIDs(client *workspaces.Client) (resources []string) {
	input := workspaces.DescribeIpGroupsInput{}
	for {
		page, err := client.DescribeIpGroupsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Result {
			resources = append(resources, *resource.GroupId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getWorkSpacesWorkspaceIDs(client *workspaces.Client) (resources []string) {
	req := client.DescribeWorkspacesRequest(&workspaces.DescribeWorkspacesInput{})
	p := workspaces.NewDescribeWorkspacesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Workspaces {
			resources = append(resources, *resource.WorkspaceId)
		}
	}
	return
}

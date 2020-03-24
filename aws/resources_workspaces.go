package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

func getWorkSpaces(config aws.Config) (resources resourceMap) {
	client := workspaces.New(config)

	workSpacesWorkspaceIDs := getWorkSpacesWorkspaceIDs(client)

	resources = resourceMap{
		workSpacesWorkspace: workSpacesWorkspaceIDs,
	}

	return
}

func getWorkSpacesWorkspaceIDs(client *workspaces.Client) (resources []string) {
	req := client.DescribeWorkspacesRequest(&workspaces.DescribeWorkspacesInput{})
	p := workspaces.NewDescribeWorkspacesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Workspaces {
			resources = append(resources, *resource.WorkspaceId)
		}
	}
	return
}

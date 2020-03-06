package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

func getWorkSpaces(config aws.Config) (resources resourceMap) {
	client := workspaces.New(config)
	resources = reduce(
		getWorkSpacesWorkspace(client).unwrap(workSpacesWorkspace),
	)
	return
}

func getWorkSpacesWorkspace(client *workspaces.Client) (r resourceSliceError) {
	req := client.DescribeWorkspacesRequest(&workspaces.DescribeWorkspacesInput{})
	p := workspaces.NewDescribeWorkspacesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Workspaces {
			r.resources = append(r.resources, *resource.WorkspaceId)
		}
	}
	r.err = p.Err()
	return
}

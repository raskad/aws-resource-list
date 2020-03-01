package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

func getWorkSpaces(session *session.Session) (resources resourceMap) {
	client := workspaces.New(session)
	resources = reduce(
		getWorkSpacesWorkspace(client).unwrap(workSpacesWorkspace),
	)
	return
}

func getWorkSpacesWorkspace(client *workspaces.WorkSpaces) (r resourceSliceError) {
	r.err = client.DescribeWorkspacesPages(&workspaces.DescribeWorkspacesInput{}, func(page *workspaces.DescribeWorkspacesOutput, lastPage bool) bool {
		for _, resource := range page.Workspaces {
			r.resources = append(r.resources, *resource.WorkspaceId)
		}
		return true
	})
	return
}

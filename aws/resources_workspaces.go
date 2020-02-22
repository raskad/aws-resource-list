package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

func getWorkSpaces(session *session.Session) (resources resourceMap) {
	client := workspaces.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		workSpacesWorkspace: getWorkSpacesWorkspace(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getWorkSpacesWorkspace(client *workspaces.WorkSpaces) (r resourceSliceError) {
	r.err = client.DescribeWorkspacesPages(&workspaces.DescribeWorkspacesInput{}, func(page *workspaces.DescribeWorkspacesOutput, lastPage bool) bool {
		logDebug("List WorkSpacesWorkspace resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Workspaces {
			logDebug("Got WorkSpacesWorkspace resource with PhysicalResourceId", *resource.WorkspaceId)
			r.resources = append(r.resources, *resource.WorkspaceId)
		}
		return true
	})
	return
}

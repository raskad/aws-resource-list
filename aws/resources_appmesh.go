package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appmesh"
)

func getAppMesh(session *session.Session) (resources resourceMap) {
	client := appmesh.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		appMeshMesh: getAppMeshMesh(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAppMeshMesh(client *appmesh.AppMesh) (r resourceSliceError) {
	r.err = client.ListMeshesPages(&appmesh.ListMeshesInput{}, func(page *appmesh.ListMeshesOutput, lastPage bool) bool {
		logDebug("Listing AppMeshMesh resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Meshes {
			logDebug("Got AppMeshMesh resource with PhysicalResourceId", *resource.MeshName)
			r.resources = append(r.resources, *resource.MeshName)
		}
		return true
	})
	return
}

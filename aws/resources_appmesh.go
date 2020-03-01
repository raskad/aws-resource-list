package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appmesh"
)

func getAppMesh(session *session.Session) (resources resourceMap) {
	client := appmesh.New(session)
	resources = reduce(
		getAppMeshMesh(client).unwrap(appMeshMesh),
	)
	return
}

func getAppMeshMesh(client *appmesh.AppMesh) (r resourceSliceError) {
	r.err = client.ListMeshesPages(&appmesh.ListMeshesInput{}, func(page *appmesh.ListMeshesOutput, lastPage bool) bool {
		for _, resource := range page.Meshes {
			r.resources = append(r.resources, *resource.MeshName)
		}
		return true
	})
	return
}

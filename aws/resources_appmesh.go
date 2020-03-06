package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

func getAppMesh(config aws.Config) (resources resourceMap) {
	client := appmesh.New(config)
	resources = reduce(
		getAppMeshMesh(client).unwrap(appMeshMesh),
	)
	return
}

func getAppMeshMesh(client *appmesh.Client) (r resourceSliceError) {
	req := client.ListMeshesRequest(&appmesh.ListMeshesInput{})
	p := appmesh.NewListMeshesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Meshes {
			r.resources = append(r.resources, *resource.MeshName)
		}
	}
	r.err = p.Err()
	return
}

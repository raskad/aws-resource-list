package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

func getAppMesh(config aws.Config) (resources awsResourceMap) {
	client := appmesh.New(config)

	appMeshMeshNames := getAppMeshMeshNames(client)

	resources = awsResourceMap{
		appMeshMesh: appMeshMeshNames,
	}
	return
}

func getAppMeshMeshNames(client *appmesh.Client) (resources []string) {
	req := client.ListMeshesRequest(&appmesh.ListMeshesInput{})
	p := appmesh.NewListMeshesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Meshes {
			resources = append(resources, *resource.MeshName)
		}
	}
	return
}

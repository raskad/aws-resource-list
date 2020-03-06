package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func getEcs(config aws.Config) (resources resourceMap) {
	client := ecs.New(config)

	ecsClusterResourceMap := getEcsCluster(client).unwrap(ecsCluster)
	ecsClusterNames := ecsClusterResourceMap[ecsCluster]

	resources = reduce(
		ecsClusterResourceMap,
		getEcsService(client, ecsClusterNames).unwrap(ecsService),
		getEcsTaskDefinition(client).unwrap(ecsTaskDefinition),
	)
	return
}

func getEcsCluster(client *ecs.Client) (r resourceSliceError) {
	req := client.ListClustersRequest(&ecs.ListClustersInput{})
	p := ecs.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ClusterArns {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

func getEcsService(client *ecs.Client, clusterNames []string) (r resourceSliceError) {
	for _, clusterName := range clusterNames {
		req := client.ListServicesRequest(&ecs.ListServicesInput{
			Cluster: aws.String(clusterName),
		})
		p := ecs.NewListServicesPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.ServiceArns {
				r.resources = append(r.resources, resource)
			}
		}
		r.err = p.Err()
	}
	return
}

func getEcsTaskDefinition(client *ecs.Client) (r resourceSliceError) {
	req := client.ListTaskDefinitionsRequest(&ecs.ListTaskDefinitionsInput{})
	p := ecs.NewListTaskDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TaskDefinitionArns {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

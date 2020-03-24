package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func getEcs(config aws.Config) (resources resourceMap) {
	client := ecs.New(config)

	ecsClusterARNs := getEcsClusterARNs(client)
	ecsServiceARNs := getEcsServiceARNs(client, ecsClusterARNs)
	ecsTaskDefinitionARNs := getEcsTaskDefinitionARNs(client)

	resources = resourceMap{
		ecsCluster:        ecsClusterARNs,
		ecsService:        ecsServiceARNs,
		ecsTaskDefinition: ecsTaskDefinitionARNs,
	}
	return
}

func getEcsClusterARNs(client *ecs.Client) (resources []string) {
	req := client.ListClustersRequest(&ecs.ListClustersInput{})
	p := ecs.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ClusterArns {
			resources = append(resources, resource)
		}
	}
	return
}

func getEcsServiceARNs(client *ecs.Client, clusterARNs []string) (resources []string) {
	for _, clusterARN := range clusterARNs {
		req := client.ListServicesRequest(&ecs.ListServicesInput{
			Cluster: aws.String(clusterARN),
		})
		p := ecs.NewListServicesPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.ServiceArns {
				resources = append(resources, resource)
			}
		}
	}
	return
}

func getEcsTaskDefinitionARNs(client *ecs.Client) (resources []string) {
	req := client.ListTaskDefinitionsRequest(&ecs.ListTaskDefinitionsInput{})
	p := ecs.NewListTaskDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.TaskDefinitionArns {
			resources = append(resources, resource)
		}
	}
	return
}

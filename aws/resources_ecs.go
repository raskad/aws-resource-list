package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func getEcs(config aws.Config) (resources awsResourceMap) {
	client := ecs.New(config)

	ecsClusterARNs := getEcsClusterARNs(client)
	ecsServiceARNs := getEcsServiceARNs(client, ecsClusterARNs)
	ecsTaskDefinitionARNs := getEcsTaskDefinitionARNs(client)
	ecsCapacityProviderARNs := getEcsCapacityProviderARNs(client)

	resources = awsResourceMap{
		ecsCluster:          ecsClusterARNs,
		ecsService:          ecsServiceARNs,
		ecsTaskDefinition:   ecsTaskDefinitionARNs,
		ecsCapacityProvider: ecsCapacityProviderARNs,
	}
	return
}

func getEcsClusterARNs(client *ecs.Client) (resources []string) {
	req := client.ListClustersRequest(&ecs.ListClustersInput{})
	p := ecs.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.ClusterArns...)
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
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			resources = append(resources, page.ServiceArns...)
		}
	}
	return
}

func getEcsTaskDefinitionARNs(client *ecs.Client) (resources []string) {
	req := client.ListTaskDefinitionsRequest(&ecs.ListTaskDefinitionsInput{})
	p := ecs.NewListTaskDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.TaskDefinitionArns...)
	}
	return
}

func getEcsCapacityProviderARNs(client *ecs.Client) (resources []string) {
	input := ecs.DescribeCapacityProvidersInput{}
	for {
		page, err := client.DescribeCapacityProvidersRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.CapacityProviders {
			resources = append(resources, *resource.CapacityProviderArn)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

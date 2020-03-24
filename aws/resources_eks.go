package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func getEks(config aws.Config) (resources resourceMap) {
	client := eks.New(config)

	eksClusterNames := getEksClusterNames(client)
	eksNodegroupNames := getEksNodegroupNames(client, eksClusterNames)

	resources = resourceMap{
		eksCluster:   eksClusterNames,
		eksNodegroup: eksNodegroupNames,
	}
	return
}

func getEksClusterNames(client *eks.Client) (resources []string) {
	req := client.ListClustersRequest(&eks.ListClustersInput{})
	p := eks.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		resources = append(resources, page.Clusters...)
	}
	return
}

func getEksNodegroupNames(client *eks.Client, clusterNames []string) (resources []string) {
	for _, clusterName := range clusterNames {
		req := client.ListNodegroupsRequest(&eks.ListNodegroupsInput{
			ClusterName: aws.String(clusterName),
		})
		p := eks.NewListNodegroupsPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			resources = append(resources, page.Nodegroups...)
		}
	}
	return
}

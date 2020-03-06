package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func getEks(config aws.Config) (resources resourceMap) {
	client := eks.New(config)

	eksClusterResourceMap := getEksCluster(client).unwrap(eksCluster)
	eksClusterNames := eksClusterResourceMap[eksCluster]

	resources = reduce(
		eksClusterResourceMap,
		getEksNodegroup(client, eksClusterNames).unwrap(eksNodegroup),
	)
	return
}

func getEksCluster(client *eks.Client) (r resourceSliceError) {
	req := client.ListClustersRequest(&eks.ListClustersInput{})
	p := eks.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Clusters {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

func getEksNodegroup(client *eks.Client, clusterNames []string) (r resourceSliceError) {
	for _, clusterName := range clusterNames {
		req := client.ListNodegroupsRequest(&eks.ListNodegroupsInput{
			ClusterName: aws.String(clusterName),
		})
		p := eks.NewListNodegroupsPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.Nodegroups {
				r.resources = append(r.resources, resource)
			}
		}
		r.err = p.Err()
		return
	}
	return
}

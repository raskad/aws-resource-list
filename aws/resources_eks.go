package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func getEks(config aws.Config) (resources awsResourceMap) {
	client := eks.New(config)

	eksClusterNames := getEksClusterNames(client)
	eksNodegroupNames := getEksNodegroupNames(client, eksClusterNames)
	eksFargateProfileNames := getEksFargateProfileNames(client, eksClusterNames)

	resources = awsResourceMap{
		eksCluster:        eksClusterNames,
		eksNodegroup:      eksNodegroupNames,
		eksFargateProfile: eksFargateProfileNames,
	}
	return
}

func getEksClusterNames(client *eks.Client) (resources []string) {
	req := client.ListClustersRequest(&eks.ListClustersInput{})
	p := eks.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			resources = append(resources, page.Nodegroups...)
		}
	}
	return
}

func getEksFargateProfileNames(client *eks.Client, clusterNames []string) (resources []string) {
	for _, clusterName := range clusterNames {
		req := client.ListFargateProfilesRequest(&eks.ListFargateProfilesInput{
			ClusterName: aws.String(clusterName),
		})
		p := eks.NewListFargateProfilesPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			resources = append(resources, page.FargateProfileNames...)
		}
	}
	return
}

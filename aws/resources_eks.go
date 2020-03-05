package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

func getEks(session *session.Session) (resources resourceMap) {
	client := eks.New(session)

	eksClusterResourceMap := getEksCluster(client).unwrap(eksCluster)
	eksClusterNames := eksClusterResourceMap[eksCluster]

	resources = reduce(
		eksClusterResourceMap,
		getEksNodegroup(client, eksClusterNames).unwrap(eksNodegroup),
	)
	return
}

func getEksCluster(client *eks.EKS) (r resourceSliceError) {
	r.err = client.ListClustersPages(&eks.ListClustersInput{}, func(page *eks.ListClustersOutput, lastPage bool) bool {
		for _, resource := range page.Clusters {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getEksNodegroup(client *eks.EKS, clusterNames []string) (r resourceSliceError) {
	for _, clusterName := range clusterNames {
		r.err = client.ListNodegroupsPages(&eks.ListNodegroupsInput{
			ClusterName: aws.String(clusterName),
		}, func(page *eks.ListNodegroupsOutput, lastPage bool) bool {
			for _, resource := range page.Nodegroups {
				r.resources = append(r.resources, *resource)
			}
			return true
		})
	}
	return
}

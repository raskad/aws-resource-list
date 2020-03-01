package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

func getEks(session *session.Session) (resources resourceMap) {
	client := eks.New(session)
	resources = reduce(
		getEksCluster(client).unwrap(eksCluster),
		getEksNodegroup(client).unwrap(eksNodegroup),
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

func getEksNodegroup(client *eks.EKS) (r resourceSliceError) {
	r.err = client.ListNodegroupsPages(&eks.ListNodegroupsInput{}, func(page *eks.ListNodegroupsOutput, lastPage bool) bool {
		for _, resource := range page.Nodegroups {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

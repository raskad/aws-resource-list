package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

func getEks(session *session.Session) (resources resourceMap) {
	client := eks.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		eksCluster:   getEksCluster(client),
		eksNodegroup: getEksNodegroup(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getEksCluster(client *eks.EKS) (r resourceSliceError) {
	r.err = client.ListClustersPages(&eks.ListClustersInput{}, func(page *eks.ListClustersOutput, lastPage bool) bool {
		logDebug("Listing EksCluster resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Clusters {
			logDebug("Got EksCluster resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getEksNodegroup(client *eks.EKS) (r resourceSliceError) {
	r.err = client.ListNodegroupsPages(&eks.ListNodegroupsInput{}, func(page *eks.ListNodegroupsOutput, lastPage bool) bool {
		logDebug("Listing EksNodegroup resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Nodegroups {
			logDebug("Got EksNodegroup resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

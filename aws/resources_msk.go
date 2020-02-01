package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kafka"
)

func getMsk(session *session.Session) (resources resourceMap) {
	client := kafka.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		mskCluster: getMskCluster(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getMskCluster(client *kafka.Kafka) (r resourceSliceError) {
	r.err = client.ListClustersPages(&kafka.ListClustersInput{}, func(page *kafka.ListClustersOutput, lastPage bool) bool {
		logDebug("Listing MskCluster resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ClusterInfoList {
			logDebug("Got MskCluster resource with PhysicalResourceId", *resource.ClusterName)
			r.resources = append(r.resources, *resource.ClusterName)
		}
		return true
	})
	return
}

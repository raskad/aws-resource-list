package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kafka"
)

func getMsk(session *session.Session) (resources resourceMap) {
	client := kafka.New(session)
	resources = reduce(
		getMskCluster(client).unwrap(mskCluster),
	)
	return
}

func getMskCluster(client *kafka.Kafka) (r resourceSliceError) {
	r.err = client.ListClustersPages(&kafka.ListClustersInput{}, func(page *kafka.ListClustersOutput, lastPage bool) bool {
		for _, resource := range page.ClusterInfoList {
			r.resources = append(r.resources, *resource.ClusterName)
		}
		return true
	})
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

func getMsk(config aws.Config) (resources resourceMap) {
	client := kafka.New(config)
	resources = reduce(
		getMskCluster(client).unwrap(mskCluster),
	)
	return
}

func getMskCluster(client *kafka.Client) (r resourceSliceError) {
	req := client.ListClustersRequest(&kafka.ListClustersInput{})
	p := kafka.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ClusterInfoList {
			r.resources = append(r.resources, *resource.ClusterName)
		}
	}
	r.err = p.Err()
	return
}

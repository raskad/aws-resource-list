package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

func getMsk(config aws.Config) (resources resourceMap) {
	client := kafka.New(config)

	mskClusterNames := getMskClusterNames(client)

	resources = resourceMap{
		mskCluster: mskClusterNames,
	}
	return
}

func getMskClusterNames(client *kafka.Client) (resources []string) {
	req := client.ListClustersRequest(&kafka.ListClustersInput{})
	p := kafka.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ClusterInfoList {
			resources = append(resources, *resource.ClusterName)
		}
	}
	return
}

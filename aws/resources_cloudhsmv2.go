package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
)

func getCloudHSMV2(config aws.Config) (resources awsResourceMap) {
	client := cloudhsmv2.New(config)

	cloudHSMV2ClusterIDs, cloudHSMV2HSMIDs := getCloudHSMV2ClusterIDs(client)

	resources = awsResourceMap{
		cloudHSMV2Cluster: cloudHSMV2ClusterIDs,
		cloudHSMV2HSM:     cloudHSMV2HSMIDs,
	}
	return
}

func getCloudHSMV2ClusterIDs(client *cloudhsmv2.Client) (cloudHSMV2ClusterIDs []string, cloudHSMV2HSMIDs []string) {
	buckets, err := client.DescribeClustersRequest(&cloudhsmv2.DescribeClustersInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range buckets.Clusters {
		cloudHSMV2ClusterIDs = append(cloudHSMV2ClusterIDs, *resource.ClusterId)
		for _, resource := range resource.Hsms {
			cloudHSMV2HSMIDs = append(cloudHSMV2HSMIDs, *resource.HsmId)
		}
	}
	return
}

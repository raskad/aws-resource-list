package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
)

func getStorageGateway(config aws.Config) (resources awsResourceMap) {
	client := storagegateway.New(config)

	storageGatewayCachedISCSIVolumeIDs := getStorageGatewayCachedISCSIVolumeIDs(client)
	storageGatewayGatewayIDs := getStorageGatewayGatewayIDs(client)
	storageGatewayNFSFileShareIDs, storageGatewaySMBFileShareIDs := getStorageGatewayNFSFileShareIDsAndStorageGatewaySMBFileShareIDs(client)

	resources = awsResourceMap{
		storageGatewayCachedISCSIVolume: storageGatewayCachedISCSIVolumeIDs,
		storageGatewayGateway:           storageGatewayGatewayIDs,
		storageGatewayNFSFileShare:      storageGatewayNFSFileShareIDs,
		storageGatewaySMBFileShare:      storageGatewaySMBFileShareIDs,
	}
	return
}

func getStorageGatewayCachedISCSIVolumeIDs(client *storagegateway.Client) (resources []string) {
	req := client.ListVolumesRequest(&storagegateway.ListVolumesInput{})
	p := storagegateway.NewListVolumesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.VolumeInfos {
			resources = append(resources, *resource.VolumeId)
		}
	}
	return
}

func getStorageGatewayGatewayIDs(client *storagegateway.Client) (resources []string) {
	req := client.ListGatewaysRequest(&storagegateway.ListGatewaysInput{})
	p := storagegateway.NewListGatewaysPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Gateways {
			resources = append(resources, *resource.GatewayId)
		}
	}
	return
}

func getStorageGatewayNFSFileShareIDsAndStorageGatewaySMBFileShareIDs(client *storagegateway.Client) (storageGatewayNFSFileShareIDs []string, storageGatewaySMBFileShareIDs []string) {
	req := client.ListFileSharesRequest(&storagegateway.ListFileSharesInput{})
	p := storagegateway.NewListFileSharesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.FileShareInfoList {
			if resource.FileShareType == storagegateway.FileShareTypeNfs {
				storageGatewayNFSFileShareIDs = append(storageGatewayNFSFileShareIDs, *resource.FileShareId)
			} else if resource.FileShareType == storagegateway.FileShareTypeSmb {
				storageGatewaySMBFileShareIDs = append(storageGatewaySMBFileShareIDs, *resource.FileShareId)
			}
		}
	}
	return
}

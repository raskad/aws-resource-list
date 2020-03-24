package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
)

func getEfs(config aws.Config) (resources resourceMap) {
	client := efs.New(config)

	efsFileSystemIDs := getEfsFileSystemIDs(client)
	efsMountTargetIDs := getEfsMountTargetIDs(client, efsFileSystemIDs)

	resources = resourceMap{
		efsFileSystem:  efsFileSystemIDs,
		efsMountTarget: efsMountTargetIDs,
	}
	return
}

func getEfsFileSystemIDs(client *efs.Client) (resources []string) {
	req := client.DescribeFileSystemsRequest(&efs.DescribeFileSystemsInput{})
	p := efs.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.FileSystems {
			resources = append(resources, *resource.FileSystemId)
		}
	}
	return
}

func getEfsMountTargetIDs(client *efs.Client, fileSystemIDs []string) (resources []string) {
	for _, fileSystemID := range fileSystemIDs {
		input := efs.DescribeMountTargetsInput{
			FileSystemId: aws.String(fileSystemID),
		}
		for {
			page, err := client.DescribeMountTargetsRequest(&input).Send(context.Background())
			logErr(err)
			for _, resource := range page.MountTargets {
				resources = append(resources, *resource.MountTargetId)
			}
			if page.NextMarker == nil {
				return
			}
			input.Marker = page.NextMarker
		}
	}
	return
}

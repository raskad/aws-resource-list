package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
)

func getEfs(config aws.Config) (resources resourceMap) {
	client := efs.New(config)

	efsFileSystemResourceMap := getEfsFileSystem(client).unwrap(efsFileSystem)
	efsFileSystemIDs := efsFileSystemResourceMap[efsFileSystem]

	resources = reduce(
		efsFileSystemResourceMap,
		getEfsMountTarget(client, efsFileSystemIDs).unwrap(efsMountTarget),
	)
	return
}

func getEfsFileSystem(client *efs.Client) (r resourceSliceError) {
	req := client.DescribeFileSystemsRequest(&efs.DescribeFileSystemsInput{})
	p := efs.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.FileSystems {
			r.resources = append(r.resources, *resource.FileSystemId)
		}
	}
	r.err = p.Err()
	return
}

func getEfsMountTarget(client *efs.Client, fileSystemIDs []string) (r resourceSliceError) {
	for _, fileSystemID := range fileSystemIDs {
		input := efs.DescribeMountTargetsInput{
			FileSystemId: aws.String(fileSystemID),
		}
		for {
			page, err := client.DescribeMountTargetsRequest(&input).Send(context.Background())
			if err != nil {
				r.err = err
				return
			}
			for _, resource := range page.MountTargets {
				r.resources = append(r.resources, *resource.MountTargetId)
			}
			if page.NextMarker == nil {
				return
			}
			input.Marker = page.NextMarker
		}
	}
	return
}

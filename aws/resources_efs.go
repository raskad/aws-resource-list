package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"
)

func getEfs(session *session.Session) (resources resourceMap) {
	client := efs.New(session)

	efsFileSystemResourceMap := getEfsFileSystem(client).unwrap(efsFileSystem)
	efsFileSystemIDs := efsFileSystemResourceMap[efsFileSystem]

	resources = reduce(
		efsFileSystemResourceMap,
		getEfsMountTarget(client, efsFileSystemIDs).unwrap(efsMountTarget),
	)
	return
}

func getEfsFileSystem(client *efs.EFS) (r resourceSliceError) {
	r.err = client.DescribeFileSystemsPages(&efs.DescribeFileSystemsInput{}, func(page *efs.DescribeFileSystemsOutput, lastPage bool) bool {
		for _, resource := range page.FileSystems {
			r.resources = append(r.resources, *resource.FileSystemId)
		}
		return true
	})
	return
}

func getEfsMountTarget(client *efs.EFS, fileSystemIDs []string) (r resourceSliceError) {
	for _, fileSystemID := range fileSystemIDs {
		input := efs.DescribeMountTargetsInput{
			FileSystemId: aws.String(fileSystemID),
		}
		for {
			page, err := client.DescribeMountTargets(&input)
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

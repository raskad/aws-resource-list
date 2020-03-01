package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"
)

func getEfs(session *session.Session) (resources resourceMap) {
	client := efs.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		efsFileSystem:  getEfsFileSystem(client),
		efsMountTarget: getEfsMountTarget(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getEfsFileSystem(client *efs.EFS) (r resourceSliceError) {
	logDebug("Listing EfsFileSystem resources")
	r.err = client.DescribeFileSystemsPages(&efs.DescribeFileSystemsInput{}, func(page *efs.DescribeFileSystemsOutput, lastPage bool) bool {
		for _, resource := range page.FileSystems {
			logDebug("Got EfsFileSystem resource with PhysicalResourceId", *resource.FileSystemId)
			r.resources = append(r.resources, *resource.FileSystemId)
		}
		return true
	})
	return
}

func getEfsMountTarget(client *efs.EFS) (r resourceSliceError) {
	logDebug("Listing EfsMountTarget resources")
	input := efs.DescribeMountTargetsInput{}
	for {
		page, err := client.DescribeMountTargets(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.MountTargets {
			logDebug("Got EfsMountTarget resource with PhysicalResourceId", *resource.MountTargetId)
			r.resources = append(r.resources, *resource.MountTargetId)
		}
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

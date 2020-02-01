package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/fsx"
)

func getFsx(session *session.Session) (resources resourceMap) {
	client := fsx.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		fsxFileSystem: getFsxFileSystem(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getFsxFileSystem(client *fsx.FSx) (r resourceSliceError) {
	r.err = client.DescribeFileSystemsPages(&fsx.DescribeFileSystemsInput{}, func(page *fsx.DescribeFileSystemsOutput, lastPage bool) bool {
		logDebug("Listing FsxFileSystem resources page. Remaining pages", page.NextToken)
		for _, resource := range page.FileSystems {
			logDebug("Got FsxFileSystem resource with PhysicalResourceId", *resource.FileSystemId)
			r.resources = append(r.resources, *resource.FileSystemId)
		}
		return true
	})
	return
}

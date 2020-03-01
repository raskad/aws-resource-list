package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/fsx"
)

func getFsx(session *session.Session) (resources resourceMap) {
	client := fsx.New(session)
	resources = reduce(
		getFsxFileSystem(client).unwrap(fsxFileSystem),
	)
	return
}

func getFsxFileSystem(client *fsx.FSx) (r resourceSliceError) {
	r.err = client.DescribeFileSystemsPages(&fsx.DescribeFileSystemsInput{}, func(page *fsx.DescribeFileSystemsOutput, lastPage bool) bool {
		for _, resource := range page.FileSystems {
			r.resources = append(r.resources, *resource.FileSystemId)
		}
		return true
	})
	return
}

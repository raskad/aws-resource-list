package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
)

func getFsx(config aws.Config) (resources resourceMap) {
	client := fsx.New(config)
	resources = reduce(
		getFsxFileSystem(client).unwrap(fsxFileSystem),
	)
	return
}

func getFsxFileSystem(client *fsx.Client) (r resourceSliceError) {
	req := client.DescribeFileSystemsRequest(&fsx.DescribeFileSystemsInput{})
	p := fsx.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.FileSystems {
			r.resources = append(r.resources, *resource.FileSystemId)
		}
	}
	r.err = p.Err()
	return
}

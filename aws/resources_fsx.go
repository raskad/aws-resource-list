package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
)

func getFsx(config aws.Config) (resources resourceMap) {
	client := fsx.New(config)

	fsxFileSystemLustreResourceMap := getFsxFileSystemLustre(client).unwrap(fsxFileSystemLustre)
	fsxFileSystemWindowsResourceMap := getFsxFileSystemWindows(client).unwrap(fsxFileSystemWindows)
	fsxFileSystemLustreIDs := fsxFileSystemLustreResourceMap[fsxFileSystemLustre]
	fsxFileSystemWindowsIDs := fsxFileSystemWindowsResourceMap[fsxFileSystemWindows]

	resources = reduce(
		resourceMap{fsxFileSystem: append(fsxFileSystemLustreIDs, fsxFileSystemWindowsIDs...)},
		fsxFileSystemLustreResourceMap,
		fsxFileSystemWindowsResourceMap,
	)
	return
}

func getFsxFileSystemLustre(client *fsx.Client) (r resourceSliceError) {
	req := client.DescribeFileSystemsRequest(&fsx.DescribeFileSystemsInput{})
	p := fsx.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.FileSystems {
			if resource.FileSystemType == fsx.FileSystemTypeLustre {
				r.resources = append(r.resources, *resource.FileSystemId)
			}
		}
	}
	r.err = p.Err()
	return
}

func getFsxFileSystemWindows(client *fsx.Client) (r resourceSliceError) {
	req := client.DescribeFileSystemsRequest(&fsx.DescribeFileSystemsInput{})
	p := fsx.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.FileSystems {
			if resource.FileSystemType == fsx.FileSystemTypeWindows {
				r.resources = append(r.resources, *resource.FileSystemId)
			}
		}
	}
	r.err = p.Err()
	return
}

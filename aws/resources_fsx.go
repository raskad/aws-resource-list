package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
)

func getFsx(config aws.Config) (resources resourceMap) {
	client := fsx.New(config)

	fsxFileSystemLustreIDs, fsxFileSystemWindowsIDs := getFsxFileSystemIDs(client)

	resources = resourceMap{
		fsxFileSystem:        append(fsxFileSystemLustreIDs, fsxFileSystemWindowsIDs...),
		fsxFileSystemLustre:  fsxFileSystemLustreIDs,
		fsxFileSystemWindows: fsxFileSystemWindowsIDs,
	}
	return
}

func getFsxFileSystemIDs(client *fsx.Client) (fsxFileSystemLustreIDs []string, fsxFileSystemWindowsIDs []string) {
	req := client.DescribeFileSystemsRequest(&fsx.DescribeFileSystemsInput{})
	p := fsx.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.FileSystems {
			if resource.FileSystemType == fsx.FileSystemTypeLustre {
				fsxFileSystemLustreIDs = append(fsxFileSystemLustreIDs, *resource.FileSystemId)
			}
			if resource.FileSystemType == fsx.FileSystemTypeWindows {
				fsxFileSystemWindowsIDs = append(fsxFileSystemWindowsIDs, *resource.FileSystemId)
			}
		}
	}
	return
}

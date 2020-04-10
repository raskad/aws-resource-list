package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
)

func getFsx(config aws.Config) (resources awsResourceMap) {
	client := fsx.New(config)

	resources = awsResourceMap{
		fsxFileSystem: getFsxFileSystemIDs(client),
	}
	return
}

func getFsxFileSystemIDs(client *fsx.Client) (resources []string) {
	req := client.DescribeFileSystemsRequest(&fsx.DescribeFileSystemsInput{})
	p := fsx.NewDescribeFileSystemsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.FileSystems {
			resources = append(resources, *resource.FileSystemId)
		}
	}
	return
}

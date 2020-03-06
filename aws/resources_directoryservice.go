package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
)

func getDirectoryService(config aws.Config) (resources resourceMap) {
	client := directoryservice.New(config)
	resources = reduce(
		getDirectoryServiceMicrosoftAD(client).unwrap(directoryServiceMicrosoftAD),
		getDirectoryServiceSimpleAD(client).unwrap(directoryServiceSimpleAD),
	)
	return
}

func getDirectoryServiceMicrosoftAD(client *directoryservice.Client) (r resourceSliceError) {
	input := directoryservice.DescribeDirectoriesInput{}
	for {
		page, err := client.DescribeDirectoriesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DirectoryDescriptions {
			if resource.Type == directoryservice.DirectoryTypeMicrosoftAd {
				r.resources = append(r.resources, *resource.DirectoryId)
			}
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getDirectoryServiceSimpleAD(client *directoryservice.Client) (r resourceSliceError) {
	input := directoryservice.DescribeDirectoriesInput{}
	for {
		page, err := client.DescribeDirectoriesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DirectoryDescriptions {
			if resource.Type == directoryservice.DirectoryTypeSimpleAd {
				r.resources = append(r.resources, *resource.DirectoryId)
			}
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

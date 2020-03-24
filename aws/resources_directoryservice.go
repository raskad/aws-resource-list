package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
)

func getDirectoryService(config aws.Config) (resources resourceMap) {
	client := directoryservice.New(config)

	directoryServiceMicrosoftADIDs := getDirectoryServiceMicrosoftADIDs(client)
	directoryServiceSimpleADIDs := getDirectoryServiceSimpleADIDs(client)

	resources = resourceMap{
		directoryServiceDirectory:   append(directoryServiceMicrosoftADIDs, directoryServiceSimpleADIDs...),
		directoryServiceMicrosoftAD: directoryServiceMicrosoftADIDs,
		directoryServiceSimpleAD:    directoryServiceSimpleADIDs,
	}
	return
}

func getDirectoryServiceMicrosoftADIDs(client *directoryservice.Client) (resources []string) {
	input := directoryservice.DescribeDirectoriesInput{}
	for {
		page, err := client.DescribeDirectoriesRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.DirectoryDescriptions {
			if resource.Type == directoryservice.DirectoryTypeMicrosoftAd {
				resources = append(resources, *resource.DirectoryId)
			}
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getDirectoryServiceSimpleADIDs(client *directoryservice.Client) (resources []string) {
	input := directoryservice.DescribeDirectoriesInput{}
	for {
		page, err := client.DescribeDirectoriesRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.DirectoryDescriptions {
			if resource.Type == directoryservice.DirectoryTypeSimpleAd {
				resources = append(resources, *resource.DirectoryId)
			}
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

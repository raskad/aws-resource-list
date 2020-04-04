package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
)

func getDirectoryService(config aws.Config) (resources awsResourceMap) {
	client := directoryservice.New(config)

	directoryServiceDirectoryIDs := getDirectoryServiceDirectoryIDs(client)

	resources = awsResourceMap{
		directoryServiceDirectory: directoryServiceDirectoryIDs,
	}
	return
}

func getDirectoryServiceDirectoryIDs(client *directoryservice.Client) (resources []string) {
	input := directoryservice.DescribeDirectoriesInput{}
	for {
		page, err := client.DescribeDirectoriesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.DirectoryDescriptions {
			resources = append(resources, *resource.DirectoryId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

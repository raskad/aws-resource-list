package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directoryservice"
)

func getDirectoryService(session *session.Session) (resources resourceMap) {
	client := directoryservice.New(session)
	resources = reduce(
		getDirectoryServiceMicrosoftAD(client).unwrap(directoryServiceMicrosoftAD),
		getDirectoryServiceSimpleAD(client).unwrap(directoryServiceSimpleAD),
	)
	return
}

func getDirectoryServiceMicrosoftAD(client *directoryservice.DirectoryService) (r resourceSliceError) {
	input := directoryservice.DescribeDirectoriesInput{}
	for {
		page, err := client.DescribeDirectories(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DirectoryDescriptions {
			if *resource.Type == directoryservice.DirectoryTypeMicrosoftAd {
				r.resources = append(r.resources, *resource.DirectoryId)
			}
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getDirectoryServiceSimpleAD(client *directoryservice.DirectoryService) (r resourceSliceError) {
	input := directoryservice.DescribeDirectoriesInput{}
	for {
		page, err := client.DescribeDirectories(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DirectoryDescriptions {
			if *resource.Type == directoryservice.DirectoryTypeSimpleAd {
				r.resources = append(r.resources, *resource.DirectoryId)
			}
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

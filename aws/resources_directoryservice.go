package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directoryservice"
)

func getDirectoryService(session *session.Session) (resources resourceMap) {
	client := directoryservice.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		directoryServiceMicrosoftAD: getDirectoryServiceMicrosoftAD(client),
		directoryServiceSimpleAD:    getDirectoryServiceSimpleAD(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
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
		logDebug("Listing DirectoryServiceMicrosoftAD resources page. Remaining pages", page.NextToken)
		for _, resource := range page.DirectoryDescriptions {
			if *resource.Type == directoryservice.DirectoryTypeMicrosoftAd {
				logDebug("Got DirectoryServiceMicrosoftAD resource with PhysicalResourceId", *resource.DirectoryId)
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
		logDebug("Listing DirectoryServiceMicrosoftAD resources page. Remaining pages", page.NextToken)
		for _, resource := range page.DirectoryDescriptions {
			if *resource.Type == directoryservice.DirectoryTypeSimpleAd {
				logDebug("Got DirectoryServiceMicrosoftAD resource with PhysicalResourceId", *resource.DirectoryId)
				r.resources = append(r.resources, *resource.DirectoryId)
			}
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}
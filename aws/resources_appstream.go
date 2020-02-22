package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appstream"
)

func getAppStream(session *session.Session) (resources resourceMap) {
	client := appstream.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		appStreamDirectoryConfig: getAppStreamDirectoryConfig(client),
		appStreamFleet:           getAppStreamFleet(client),
		appStreamImageBuilder:    getAppStreamImageBuilder(client),
		appStreamStack:           getAppStreamStack(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAppStreamDirectoryConfig(client *appstream.AppStream) (r resourceSliceError) {
	input := appstream.DescribeDirectoryConfigsInput{}
	for {
		page, err := client.DescribeDirectoryConfigs(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing AppStreamDirectoryConfig resources page. Remaining pages", page.NextToken)
		for _, resource := range page.DirectoryConfigs {
			logDebug("Got AppStreamDirectoryConfig resource with PhysicalResourceId", *resource.DirectoryName)
			r.resources = append(r.resources, *resource.DirectoryName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAppStreamFleet(client *appstream.AppStream) (r resourceSliceError) {
	input := appstream.DescribeFleetsInput{}
	for {
		page, err := client.DescribeFleets(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing AppStreamFleet resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Fleets {
			logDebug("Got AppStreamFleet resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAppStreamImageBuilder(client *appstream.AppStream) (r resourceSliceError) {
	input := appstream.DescribeImageBuildersInput{}
	for {
		page, err := client.DescribeImageBuilders(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing AppStreamImageBuilder resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ImageBuilders {
			logDebug("Got AppStreamImageBuilder resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAppStreamStack(client *appstream.AppStream) (r resourceSliceError) {
	input := appstream.DescribeStacksInput{}
	for {
		page, err := client.DescribeStacks(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing AppStreamStack resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Stacks {
			logDebug("Got AppStreamStack resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appstream"
)

func getAppStream(session *session.Session) (resources resourceMap) {
	client := appstream.New(session)
	resources = reduce(
		getAppStreamDirectoryConfig(client).unwrap(appStreamDirectoryConfig),
		getAppStreamFleet(client).unwrap(appStreamFleet),
		getAppStreamImageBuilder(client).unwrap(appStreamImageBuilder),
		getAppStreamStack(client).unwrap(appStreamStack),
	)
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
		for _, resource := range page.DirectoryConfigs {
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
		for _, resource := range page.Fleets {
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
		for _, resource := range page.ImageBuilders {
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
		for _, resource := range page.Stacks {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

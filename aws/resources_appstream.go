package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
)

func getAppStream(config aws.Config) (resources resourceMap) {
	client := appstream.New(config)
	resources = reduce(
		getAppStreamDirectoryConfig(client).unwrap(appStreamDirectoryConfig),
		getAppStreamFleet(client).unwrap(appStreamFleet),
		getAppStreamImageBuilder(client).unwrap(appStreamImageBuilder),
		getAppStreamStack(client).unwrap(appStreamStack),
	)
	return
}

func getAppStreamDirectoryConfig(client *appstream.Client) (r resourceSliceError) {
	input := appstream.DescribeDirectoryConfigsInput{}
	for {
		page, err := client.DescribeDirectoryConfigsRequest(&input).Send(context.Background())
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

func getAppStreamFleet(client *appstream.Client) (r resourceSliceError) {
	input := appstream.DescribeFleetsInput{}
	for {
		page, err := client.DescribeFleetsRequest(&input).Send(context.Background())
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

func getAppStreamImageBuilder(client *appstream.Client) (r resourceSliceError) {
	input := appstream.DescribeImageBuildersInput{}
	for {
		page, err := client.DescribeImageBuildersRequest(&input).Send(context.Background())
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

func getAppStreamStack(client *appstream.Client) (r resourceSliceError) {
	input := appstream.DescribeStacksInput{}
	for {
		page, err := client.DescribeStacksRequest(&input).Send(context.Background())
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

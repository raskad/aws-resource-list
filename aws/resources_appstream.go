package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
)

func getAppStream(config aws.Config) (resources resourceMap) {
	client := appstream.New(config)

	appStreamDirectoryConfigNames := getAppStreamDirectoryConfigNames(client)
	appStreamFleetNames := getAppStreamFleetNames(client)
	appStreamImageBuilderNames := getAppStreamImageBuilderNames(client)
	appStreamStackNames := getAppStreamStackNames(client)

	resources = resourceMap{
		appStreamDirectoryConfig: appStreamDirectoryConfigNames,
		appStreamFleet:           appStreamFleetNames,
		appStreamImageBuilder:    appStreamImageBuilderNames,
		appStreamStack:           appStreamStackNames,
	}
	return
}

func getAppStreamDirectoryConfigNames(client *appstream.Client) (resources []string) {
	input := appstream.DescribeDirectoryConfigsInput{}
	for {
		page, err := client.DescribeDirectoryConfigsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.DirectoryConfigs {
			resources = append(resources, *resource.DirectoryName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAppStreamFleetNames(client *appstream.Client) (resources []string) {
	input := appstream.DescribeFleetsInput{}
	for {
		page, err := client.DescribeFleetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Fleets {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAppStreamImageBuilderNames(client *appstream.Client) (resources []string) {
	input := appstream.DescribeImageBuildersInput{}
	for {
		page, err := client.DescribeImageBuildersRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.ImageBuilders {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAppStreamStackNames(client *appstream.Client) (resources []string) {
	input := appstream.DescribeStacksInput{}
	for {
		page, err := client.DescribeStacksRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Stacks {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

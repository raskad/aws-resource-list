package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass"
)

func getGreengrass(config aws.Config) (resources resourceMap) {
	client := greengrass.New(config)
	resources = reduce(
		getGreengrassConnectorDefinition(client).unwrap(greengrassConnectorDefinition),
		getGreengrassConnectorDefinitionVersion(client).unwrap(greengrassConnectorDefinitionVersion),
		getGreengrassCoreDefinition(client).unwrap(greengrassCoreDefinition),
		getGreengrassCoreDefinitionVersion(client).unwrap(greengrassCoreDefinitionVersion),
		getGreengrassDeviceDefinition(client).unwrap(greengrassDeviceDefinition),
		getGreengrassDeviceDefinitionVersion(client).unwrap(greengrassDeviceDefinitionVersion),
		getGreengrassFunctionDefinition(client).unwrap(greengrassFunctionDefinition),
		getGreengrassFunctionDefinitionVersion(client).unwrap(greengrassFunctionDefinitionVersion),
		getGreengrassGroup(client).unwrap(greengrassGroup),
		getGreengrassGroupVersion(client).unwrap(greengrassGroupVersion),
		getGreengrassLoggerDefinition(client).unwrap(greengrassLoggerDefinition),
		getGreengrassLoggerDefinitionVersion(client).unwrap(greengrassLoggerDefinitionVersion),
		getGreengrassResourceDefinition(client).unwrap(greengrassResourceDefinition),
		getGreengrassResourceDefinitionVersion(client).unwrap(greengrassResourceDefinitionVersion),
		getGreengrassSubscriptionDefinition(client).unwrap(greengrassSubscriptionDefinition),
		getGreengrassSubscriptionDefinitionVersion(client).unwrap(greengrassSubscriptionDefinitionVersion),
	)
	return
}

func getGreengrassConnectorDefinition(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListConnectorDefinitionsInput{}
	for {
		page, err := client.ListConnectorDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Definitions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassConnectorDefinitionVersion(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListConnectorDefinitionVersionsInput{}
	for {
		page, err := client.ListConnectorDefinitionVersionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Versions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassCoreDefinition(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListCoreDefinitionsInput{}
	for {
		page, err := client.ListCoreDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Definitions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassCoreDefinitionVersion(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListCoreDefinitionVersionsInput{}
	for {
		page, err := client.ListCoreDefinitionVersionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Versions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassDeviceDefinition(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListDeviceDefinitionsInput{}
	for {
		page, err := client.ListDeviceDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Definitions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassDeviceDefinitionVersion(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListDeviceDefinitionVersionsInput{}
	for {
		page, err := client.ListDeviceDefinitionVersionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Versions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassFunctionDefinition(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListFunctionDefinitionsInput{}
	for {
		page, err := client.ListFunctionDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Definitions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassFunctionDefinitionVersion(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListFunctionDefinitionVersionsInput{}
	for {
		page, err := client.ListFunctionDefinitionVersionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Versions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassGroup(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListGroupsInput{}
	for {
		page, err := client.ListGroupsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Groups {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassGroupVersion(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListGroupVersionsInput{}
	for {
		page, err := client.ListGroupVersionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Versions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassLoggerDefinition(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListLoggerDefinitionsInput{}
	for {
		page, err := client.ListLoggerDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Definitions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassLoggerDefinitionVersion(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListLoggerDefinitionVersionsInput{}
	for {
		page, err := client.ListLoggerDefinitionVersionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Versions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassResourceDefinition(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListResourceDefinitionsInput{}
	for {
		page, err := client.ListResourceDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Definitions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassResourceDefinitionVersion(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListResourceDefinitionVersionsInput{}
	for {
		page, err := client.ListResourceDefinitionVersionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Versions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassSubscriptionDefinition(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListSubscriptionDefinitionsInput{}
	for {
		page, err := client.ListSubscriptionDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Definitions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassSubscriptionDefinitionVersion(client *greengrass.Client) (r resourceSliceError) {
	input := greengrass.ListSubscriptionDefinitionVersionsInput{}
	for {
		page, err := client.ListSubscriptionDefinitionVersionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Versions {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

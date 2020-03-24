package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass"
)

func getGreengrass(config aws.Config) (resources resourceMap) {
	client := greengrass.New(config)

	greengrassConnectorDefinitionIDs := getGreengrassConnectorDefinitionIDs(client)
	greengrassConnectorDefinitionVersionIDs := getGreengrassConnectorDefinitionVersionIDs(client)
	greengrassCoreDefinitionIDs := getGreengrassCoreDefinitionIDs(client)
	greengrassCoreDefinitionVersionIDs := getGreengrassCoreDefinitionVersionIDs(client)
	greengrassDeviceDefinitionIDs := getGreengrassDeviceDefinitionIDs(client)
	greengrassDeviceDefinitionVersionIDs := getGreengrassDeviceDefinitionVersionIDs(client)
	greengrassFunctionDefinitionIDs := getGreengrassFunctionDefinitionIDs(client)
	greengrassFunctionDefinitionVersionIDs := getGreengrassFunctionDefinitionVersionIDs(client)
	greengrassGroupIDs := getGreengrassGroupIDs(client)
	greengrassGroupVersionIDs := getGreengrassGroupVersionIDs(client)
	greengrassLoggerDefinitionIDs := getGreengrassLoggerDefinitionIDs(client)
	greengrassLoggerDefinitionVersionIDs := getGreengrassLoggerDefinitionVersionIDs(client)
	greengrassResourceDefinitionIDs := getGreengrassResourceDefinitionIDs(client)
	greengrassResourceDefinitionVersionIDs := getGreengrassResourceDefinitionVersionIDs(client)
	greengrassSubscriptionDefinitionIDs := getGreengrassSubscriptionDefinitionIDs(client)
	greengrassSubscriptionDefinitionVersionIDs := getGreengrassSubscriptionDefinitionVersionIDs(client)

	resources = resourceMap{
		greengrassConnectorDefinition:           greengrassConnectorDefinitionIDs,
		greengrassConnectorDefinitionVersion:    greengrassConnectorDefinitionVersionIDs,
		greengrassCoreDefinition:                greengrassCoreDefinitionIDs,
		greengrassCoreDefinitionVersion:         greengrassCoreDefinitionVersionIDs,
		greengrassDeviceDefinition:              greengrassDeviceDefinitionIDs,
		greengrassDeviceDefinitionVersion:       greengrassDeviceDefinitionVersionIDs,
		greengrassFunctionDefinition:            greengrassFunctionDefinitionIDs,
		greengrassFunctionDefinitionVersion:     greengrassFunctionDefinitionVersionIDs,
		greengrassGroup:                         greengrassGroupIDs,
		greengrassGroupVersion:                  greengrassGroupVersionIDs,
		greengrassLoggerDefinition:              greengrassLoggerDefinitionIDs,
		greengrassLoggerDefinitionVersion:       greengrassLoggerDefinitionVersionIDs,
		greengrassResourceDefinition:            greengrassResourceDefinitionIDs,
		greengrassResourceDefinitionVersion:     greengrassResourceDefinitionVersionIDs,
		greengrassSubscriptionDefinition:        greengrassSubscriptionDefinitionIDs,
		greengrassSubscriptionDefinitionVersion: greengrassSubscriptionDefinitionVersionIDs,
	}
	return
}

func getGreengrassConnectorDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListConnectorDefinitionsInput{}
	for {
		page, err := client.ListConnectorDefinitionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassConnectorDefinitionVersionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListConnectorDefinitionVersionsInput{}
	for {
		page, err := client.ListConnectorDefinitionVersionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Versions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassCoreDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListCoreDefinitionsInput{}
	for {
		page, err := client.ListCoreDefinitionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassCoreDefinitionVersionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListCoreDefinitionVersionsInput{}
	for {
		page, err := client.ListCoreDefinitionVersionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Versions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassDeviceDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListDeviceDefinitionsInput{}
	for {
		page, err := client.ListDeviceDefinitionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassDeviceDefinitionVersionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListDeviceDefinitionVersionsInput{}
	for {
		page, err := client.ListDeviceDefinitionVersionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Versions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassFunctionDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListFunctionDefinitionsInput{}
	for {
		page, err := client.ListFunctionDefinitionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassFunctionDefinitionVersionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListFunctionDefinitionVersionsInput{}
	for {
		page, err := client.ListFunctionDefinitionVersionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Versions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassGroupIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListGroupsInput{}
	for {
		page, err := client.ListGroupsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Groups {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassGroupVersionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListGroupVersionsInput{}
	for {
		page, err := client.ListGroupVersionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Versions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassLoggerDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListLoggerDefinitionsInput{}
	for {
		page, err := client.ListLoggerDefinitionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassLoggerDefinitionVersionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListLoggerDefinitionVersionsInput{}
	for {
		page, err := client.ListLoggerDefinitionVersionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Versions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassResourceDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListResourceDefinitionsInput{}
	for {
		page, err := client.ListResourceDefinitionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassResourceDefinitionVersionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListResourceDefinitionVersionsInput{}
	for {
		page, err := client.ListResourceDefinitionVersionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Versions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassSubscriptionDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListSubscriptionDefinitionsInput{}
	for {
		page, err := client.ListSubscriptionDefinitionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassSubscriptionDefinitionVersionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListSubscriptionDefinitionVersionsInput{}
	for {
		page, err := client.ListSubscriptionDefinitionVersionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Versions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

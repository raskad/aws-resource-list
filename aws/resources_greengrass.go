package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass"
)

func getGreengrass(config aws.Config) (resources awsResourceMap) {
	client := greengrass.New(config)

	greengrassConnectorDefinitionIDs := getGreengrassConnectorDefinitionIDs(client)
	greengrassConnectorDefinitionVersionIDs := getGreengrassConnectorDefinitionVersionIDs(client, greengrassConnectorDefinitionIDs)
	greengrassCoreDefinitionIDs := getGreengrassCoreDefinitionIDs(client)
	greengrassCoreDefinitionVersionIDs := getGreengrassCoreDefinitionVersionIDs(client, greengrassCoreDefinitionIDs)
	greengrassDeviceDefinitionIDs := getGreengrassDeviceDefinitionIDs(client)
	greengrassDeviceDefinitionVersionIDs := getGreengrassDeviceDefinitionVersionIDs(client, greengrassDeviceDefinitionIDs)
	greengrassFunctionDefinitionIDs := getGreengrassFunctionDefinitionIDs(client)
	greengrassFunctionDefinitionVersionIDs := getGreengrassFunctionDefinitionVersionIDs(client, greengrassFunctionDefinitionIDs)
	greengrassGroupIDs := getGreengrassGroupIDs(client)
	greengrassGroupVersionIDs := getGreengrassGroupVersionIDs(client, greengrassGroupIDs)
	greengrassLoggerDefinitionIDs := getGreengrassLoggerDefinitionIDs(client)
	greengrassLoggerDefinitionVersionIDs := getGreengrassLoggerDefinitionVersionIDs(client, greengrassLoggerDefinitionIDs)
	greengrassResourceDefinitionIDs := getGreengrassResourceDefinitionIDs(client)
	greengrassResourceDefinitionVersionIDs := getGreengrassResourceDefinitionVersionIDs(client, greengrassResourceDefinitionIDs)
	greengrassSubscriptionDefinitionIDs := getGreengrassSubscriptionDefinitionIDs(client)
	greengrassSubscriptionDefinitionVersionIDs := getGreengrassSubscriptionDefinitionVersionIDs(client, greengrassSubscriptionDefinitionIDs)

	resources = awsResourceMap{
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
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassConnectorDefinitionVersionIDs(client *greengrass.Client, greengrassConnectorDefinitionIDs []string) (resources []string) {
	for _, greengrassConnectorDefinitionID := range greengrassConnectorDefinitionIDs {
		input := greengrass.ListConnectorDefinitionVersionsInput{
			ConnectorDefinitionId: &greengrassConnectorDefinitionID,
		}
		for {
			page, err := client.ListConnectorDefinitionVersionsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.Versions {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getGreengrassCoreDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListCoreDefinitionsInput{}
	for {
		page, err := client.ListCoreDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassCoreDefinitionVersionIDs(client *greengrass.Client, greengrassCoreDefinitionIDs []string) (resources []string) {
	for _, greengrassCoreDefinitionID := range greengrassCoreDefinitionIDs {
		input := greengrass.ListCoreDefinitionVersionsInput{
			CoreDefinitionId: &greengrassCoreDefinitionID,
		}
		for {
			page, err := client.ListCoreDefinitionVersionsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.Versions {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getGreengrassDeviceDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListDeviceDefinitionsInput{}
	for {
		page, err := client.ListDeviceDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassDeviceDefinitionVersionIDs(client *greengrass.Client, greengrassDeviceDefinitionIDs []string) (resources []string) {
	for _, greengrassDeviceDefinitionID := range greengrassDeviceDefinitionIDs {
		input := greengrass.ListDeviceDefinitionVersionsInput{
			DeviceDefinitionId: &greengrassDeviceDefinitionID,
		}
		for {
			page, err := client.ListDeviceDefinitionVersionsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				return
			}
			for _, resource := range page.Versions {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getGreengrassFunctionDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListFunctionDefinitionsInput{}
	for {
		page, err := client.ListFunctionDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassFunctionDefinitionVersionIDs(client *greengrass.Client, greengrassFunctionDefinitionIDs []string) (resources []string) {
	for _, greengrassFunctionDefinitionID := range greengrassFunctionDefinitionIDs {
		input := greengrass.ListFunctionDefinitionVersionsInput{
			FunctionDefinitionId: &greengrassFunctionDefinitionID,
		}
		for {
			page, err := client.ListFunctionDefinitionVersionsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				return
			}
			for _, resource := range page.Versions {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getGreengrassGroupIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListGroupsInput{}
	for {
		page, err := client.ListGroupsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Groups {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassGroupVersionIDs(client *greengrass.Client, greengrassGroupIDs []string) (resources []string) {
	for _, greengrassGroupID := range greengrassGroupIDs {
		input := greengrass.ListGroupVersionsInput{
			GroupId: &greengrassGroupID,
		}
		for {
			page, err := client.ListGroupVersionsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				return
			}
			for _, resource := range page.Versions {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getGreengrassLoggerDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListLoggerDefinitionsInput{}
	for {
		page, err := client.ListLoggerDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassLoggerDefinitionVersionIDs(client *greengrass.Client, greengrassLoggerDefinitionIDs []string) (resources []string) {
	for _, greengrassLoggerDefinitionID := range greengrassLoggerDefinitionIDs {
		input := greengrass.ListLoggerDefinitionVersionsInput{
			LoggerDefinitionId: &greengrassLoggerDefinitionID,
		}
		for {
			page, err := client.ListLoggerDefinitionVersionsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				return
			}
			for _, resource := range page.Versions {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getGreengrassResourceDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListResourceDefinitionsInput{}
	for {
		page, err := client.ListResourceDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassResourceDefinitionVersionIDs(client *greengrass.Client, greengrassResourceDefinitionIDs []string) (resources []string) {
	for _, greengrassResourceDefinitionID := range greengrassResourceDefinitionIDs {
		input := greengrass.ListResourceDefinitionVersionsInput{
			ResourceDefinitionId: &greengrassResourceDefinitionID,
		}
		for {
			page, err := client.ListResourceDefinitionVersionsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				return
			}
			for _, resource := range page.Versions {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getGreengrassSubscriptionDefinitionIDs(client *greengrass.Client) (resources []string) {
	input := greengrass.ListSubscriptionDefinitionsInput{}
	for {
		page, err := client.ListSubscriptionDefinitionsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Definitions {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGreengrassSubscriptionDefinitionVersionIDs(client *greengrass.Client, greengrassSubscriptionDefinitionIDs []string) (resources []string) {
	for _, greengrassSubscriptionDefinitionID := range greengrassSubscriptionDefinitionIDs {
		input := greengrass.ListSubscriptionDefinitionVersionsInput{
			SubscriptionDefinitionId: &greengrassSubscriptionDefinitionID,
		}
		for {
			page, err := client.ListSubscriptionDefinitionVersionsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				return
			}
			for _, resource := range page.Versions {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

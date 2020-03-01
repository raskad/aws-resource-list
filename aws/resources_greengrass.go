package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/greengrass"
)

func getGreengrass(session *session.Session) (resources resourceMap) {
	client := greengrass.New(session)
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

func getGreengrassConnectorDefinition(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListConnectorDefinitionsInput{}
	for {
		page, err := client.ListConnectorDefinitions(&input)
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

func getGreengrassConnectorDefinitionVersion(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListConnectorDefinitionVersionsInput{}
	for {
		page, err := client.ListConnectorDefinitionVersions(&input)
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

func getGreengrassCoreDefinition(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListCoreDefinitionsInput{}
	for {
		page, err := client.ListCoreDefinitions(&input)
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

func getGreengrassCoreDefinitionVersion(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListCoreDefinitionVersionsInput{}
	for {
		page, err := client.ListCoreDefinitionVersions(&input)
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

func getGreengrassDeviceDefinition(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListDeviceDefinitionsInput{}
	for {
		page, err := client.ListDeviceDefinitions(&input)
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

func getGreengrassDeviceDefinitionVersion(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListDeviceDefinitionVersionsInput{}
	for {
		page, err := client.ListDeviceDefinitionVersions(&input)
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

func getGreengrassFunctionDefinition(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListFunctionDefinitionsInput{}
	for {
		page, err := client.ListFunctionDefinitions(&input)
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

func getGreengrassFunctionDefinitionVersion(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListFunctionDefinitionVersionsInput{}
	for {
		page, err := client.ListFunctionDefinitionVersions(&input)
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

func getGreengrassGroup(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListGroupsInput{}
	for {
		page, err := client.ListGroups(&input)
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

func getGreengrassGroupVersion(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListGroupVersionsInput{}
	for {
		page, err := client.ListGroupVersions(&input)
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

func getGreengrassLoggerDefinition(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListLoggerDefinitionsInput{}
	for {
		page, err := client.ListLoggerDefinitions(&input)
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

func getGreengrassLoggerDefinitionVersion(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListLoggerDefinitionVersionsInput{}
	for {
		page, err := client.ListLoggerDefinitionVersions(&input)
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

func getGreengrassResourceDefinition(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListResourceDefinitionsInput{}
	for {
		page, err := client.ListResourceDefinitions(&input)
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

func getGreengrassResourceDefinitionVersion(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListResourceDefinitionVersionsInput{}
	for {
		page, err := client.ListResourceDefinitionVersions(&input)
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

func getGreengrassSubscriptionDefinition(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListSubscriptionDefinitionsInput{}
	for {
		page, err := client.ListSubscriptionDefinitions(&input)
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

func getGreengrassSubscriptionDefinitionVersion(client *greengrass.Greengrass) (r resourceSliceError) {
	input := greengrass.ListSubscriptionDefinitionVersionsInput{}
	for {
		page, err := client.ListSubscriptionDefinitionVersions(&input)
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

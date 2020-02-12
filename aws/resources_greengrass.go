package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/greengrass"
)

func getGreengrass(session *session.Session) (resources resourceMap) {
	client := greengrass.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		greengrassConnectorDefinition:           getGreengrassConnectorDefinition(client),
		greengrassConnectorDefinitionVersion:    getGreengrassConnectorDefinitionVersion(client),
		greengrassCoreDefinition:                getGreengrassCoreDefinition(client),
		greengrassCoreDefinitionVersion:         getGreengrassCoreDefinitionVersion(client),
		greengrassDeviceDefinition:              getGreengrassDeviceDefinition(client),
		greengrassDeviceDefinitionVersion:       getGreengrassDeviceDefinitionVersion(client),
		greengrassFunctionDefinition:            getGreengrassFunctionDefinition(client),
		greengrassFunctionDefinitionVersion:     getGreengrassFunctionDefinitionVersion(client),
		greengrassGroup:                         getGreengrassGroup(client),
		greengrassGroupVersion:                  getGreengrassGroupVersion(client),
		greengrassLoggerDefinition:              getGreengrassLoggerDefinition(client),
		greengrassLoggerDefinitionVersion:       getGreengrassLoggerDefinitionVersion(client),
		greengrassResourceDefinition:            getGreengrassResourceDefinition(client),
		greengrassResourceDefinitionVersion:     getGreengrassResourceDefinitionVersion(client),
		greengrassSubscriptionDefinition:        getGreengrassSubscriptionDefinition(client),
		greengrassSubscriptionDefinitionVersion: getGreengrassSubscriptionDefinitionVersion(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
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
		logDebug("Listing GreengrassConnectorDefinition resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Definitions {
			logDebug("Got GreengrassConnectorDefinition resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassConnectorDefinitionVersion resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Versions {
			logDebug("Got GreengrassConnectorDefinitionVersion resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassCoreDefinition resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Definitions {
			logDebug("Got GreengrassCoreDefinition resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassCoreDefinitionVersion resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Versions {
			logDebug("Got GreengrassCoreDefinitionVersion resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassDeviceDefinition resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Definitions {
			logDebug("Got GreengrassDeviceDefinition resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassDeviceDefinitionVersion resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Versions {
			logDebug("Got GreengrassDeviceDefinitionVersion resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassFunctionDefinition resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Definitions {
			logDebug("Got GreengrassFunctionDefinition resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassFunctionDefinitionVersion resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Versions {
			logDebug("Got GreengrassFunctionDefinitionVersion resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassGroup resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Groups {
			logDebug("Got GreengrassGroup resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassGroupVersion resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Versions {
			logDebug("Got GreengrassGroupVersion resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassLoggerDefinition resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Definitions {
			logDebug("Got GreengrassLoggerDefinition resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassLoggerDefinitionVersion resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Versions {
			logDebug("Got GreengrassLoggerDefinitionVersion resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassResourceDefinition resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Definitions {
			logDebug("Got GreengrassResourceDefinition resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassResourceDefinitionVersion resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Versions {
			logDebug("Got GreengrassResourceDefinitionVersion resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassSubscriptionDefinition resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Definitions {
			logDebug("Got GreengrassSubscriptionDefinition resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing GreengrassSubscriptionDefinitionVersion resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Versions {
			logDebug("Got GreengrassSubscriptionDefinitionVersion resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}
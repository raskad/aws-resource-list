package aws

import (
	"fmt"
	"io/ioutil"
	"os"

	tfjson "github.com/hashicorp/terraform-json"
)

func getTerraformState(path string) (resources resourceMap, err error) {
	state := tfjson.State{}

	stateFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer stateFile.Close()
	byteValue, err := ioutil.ReadAll(stateFile)
	if err != nil {
		return
	}
	err = state.UnmarshalJSON(byteValue)
	if err != nil {
		return
	}

	resources = resourceMap{}
	resources = getTerraformResources(state.Values.RootModule, resources)
	return resources, nil
}

func getTerraformResources(module *tfjson.StateModule, resources resourceMap) resourceMap {
	for _, resource := range module.Resources {
		resourceType, ok := fromTerraformType(resource.Type)
		if ok {
			physicalResourceIDType, ok := resourceType.physicalResourceIDTerraform()
			if ok {
				physicalResourceID := fmt.Sprintf("%v", resource.AttributeValues[physicalResourceIDType])
				logDebug("Got terraform resource with ResourceType", resourceType, "and PhysicalResourceId", physicalResourceID)
				resources[resourceType] = append(resources[resourceType], physicalResourceID)
			} else {
				logError("Terraform resourceType has no physicalResourceIDType")
			}
		}
	}
	for _, module := range module.ChildModules {
		resources = getTerraformResources(module, resources)
	}
	return resources
}

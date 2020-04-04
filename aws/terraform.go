package aws

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
)

func getTerraformState(path string, resources *extResourceMap) (err error) {
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

	getTerraformResources(state.Values.RootModule, resources)
	return nil
}

func getTerraformResources(module *tfjson.StateModule, resources *extResourceMap) {
	for _, resource := range module.Resources {
		if strings.HasPrefix(resource.Address, "data.") {
			continue
		}
		resourceType, ok := fromTerraformType(resource.Type)
		if ok {
			physicalResourceIDType, ok := resourceType.physicalResourceIDTerraform()
			if ok {
				physicalResourceID := fmt.Sprintf("%v", resource.AttributeValues[physicalResourceIDType])
				logDebug("Got terraform resource with ResourceType", resourceType, "and PhysicalResourceId", physicalResourceID)
				resourceExists := false
				for _, resourceID := range (*resources)[resource.Type] {
					if resourceID == physicalResourceID {
						resourceExists = true
						break
					}
				}
				if !resourceExists {
					(*resources)[resource.Type] = append((*resources)[resource.Type], physicalResourceID)
				}
			} else {
				logError("Terraform resourceType has no physicalResourceIDType")
			}
		}
	}
	for _, module := range module.ChildModules {
		getTerraformResources(module, resources)
	}
	return
}

func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

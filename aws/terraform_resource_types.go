package aws

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func getTerraformResourceTypes() (resources []string, err error) {
	resp, err := http.Get("https://www.terraform.io/docs/providers/aws/index.html")
	if err != nil {
		return resources, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resources, err
	}
	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	re := regexp.MustCompile(`.*/docs/providers/aws/r/.*\.html\">(aws_.*)</a>`)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) > 0 {
			resources = append(resources, matches[1])
		}
	}
	return
}

func printMissingTerraformResources() {
	resources, err := getTerraformResourceTypes()
	if err != nil {
		logFatal(err)
	}
	for _, resource := range resources {
		found := false
		for existingResource := range terraformTypeMap {
			if resource == existingResource {
				found = true
				break
			}
		}
		if !found {
		foundResource:
			for _, blacklistedResources := range resourceBlacklistMap {
				for _, blacklistedResource := range blacklistedResources {
					if resource == blacklistedResource {
						found = true
						break foundResource
					}
				}
			}
		}
		if !found {
			fmt.Println(resource)
		}
	}
}

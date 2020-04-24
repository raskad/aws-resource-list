package aws

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type awsCloudFormationTypesJOSN struct {
	PropertyTypes                interface{}
	ResourceTypes                map[string]interface{}
	ResourceSpecificationVersion string
}

func printMissingCloudformationResources() {
	resources, err := getCloudFormationResources()
	if err != nil {
		logFatal(err)
	}
	for _, resource := range resources {
		found := false
		for existingResource := range cloudformationTypeMap {
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

func getCloudFormationResources() (resources []string, err error) {
	cloudFormationResourceTypeJSONUrls, err := getCloudFormationResourceTypeJSONUrls()
	if err != nil {
		return resources, err
	}

	for _, cloudFormationResourceTypeJSONUrl := range cloudFormationResourceTypeJSONUrls {
		types, err := getCloudFormationTypes(cloudFormationResourceTypeJSONUrl)
		if err != nil {
			return resources, err
		}

		for _, resourceType := range types {
			resources = appendIfMissing(resources, resourceType)
		}
	}
	return resources, nil
}

func getCloudFormationTypes(url string) (resourceTypes []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var responseJSON awsCloudFormationTypesJOSN

	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		return
	}

	for resourceType := range responseJSON.ResourceTypes {
		resourceTypes = append(resourceTypes, resourceType)
	}

	return resourceTypes, nil
}

func getCloudFormationResourceTypeJSONUrls() (urls []string, err error) {
	resp, err := http.Get("https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-resource-specification.html")
	if err != nil {
		return urls, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return urls, err
	}
	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	re := regexp.MustCompile(`https\:\/\/(.*)\.cloudfront\.net\/latest\/gzip\/CloudFormationResourceSpecification\.json`)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) > 0 {
			urls = append(urls, fmt.Sprintf("https://%s.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json", matches[1]))
		}
	}
	return
}

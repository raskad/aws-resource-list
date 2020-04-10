package aws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type awsCloudformationDocJSON struct {
	Contents []struct {
		Title    string
		Href     string
		Contents []struct {
			Title    string
			Href     string
			Contents []struct {
				Title string
				Href  string
			}
		}
	}
}

func getCloudFormationResources() (resources []string, err error) {
	cloudformationServices, err := getCloudformationServices()
	if err != nil {
		return resources, err
	}
	for _, service := range cloudformationServices {
		url := fmt.Sprintf("https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/toc-%v.json", service)
		serviceResource, err := getResourceTypes(url)
		if err != nil {
			return resources, err
		}
		resources = append(resources, serviceResource...)
	}
	return resources, nil
}

func getCloudformationServices() (cloudformationServices []string, err error) {
	url := "https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.partial.html"
	resp, err := http.Get(url)
	if err != nil {
		return cloudformationServices, err
	}
	defer resp.Body.Close()
	if err != nil {
		return cloudformationServices, err
	}
	z := html.NewTokenizer(resp.Body)
	for {
		token := z.Next()
		switch {
		case token == html.ErrorToken:
			return cloudformationServices, nil
		case token == html.StartTagToken:
			token := z.Token()
			if token.Data == "li" {
				_ = z.Next()
				token := z.Token()
				if token.Data == "a" {
					service := token.Attr[0].Val
					service = strings.Replace(service, "./", "", -1)
					service = strings.Replace(service, ".html", "", -1)
					if service != "cfn-reference-shared" {
						cloudformationServices = append(cloudformationServices, service)
					}
				}
			}
		}
	}
}

func getResourceTypes(url string) (resources []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return resources, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resources, err
	}
	var doc awsCloudformationDocJSON
	err = json.Unmarshal(body, &doc)
	if err != nil {
		return resources, err
	}
	for _, resource := range doc.Contents[0].Contents {
		resources = append(resources, resource.Title)
	}
	return
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

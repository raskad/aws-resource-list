package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
)

func getIoT1ClickProjects(config aws.Config) (resources resourceMap) {
	client := iot1clickprojects.New(config)

	ioT1ClickProjectNames := getIoT1ClickProjectNames(client)

	resources = resourceMap{
		ioT1ClickProject: ioT1ClickProjectNames,
	}
	return
}

func getIoT1ClickProjectNames(client *iot1clickprojects.Client) (resources []string) {
	req := client.ListProjectsRequest(&iot1clickprojects.ListProjectsInput{})
	p := iot1clickprojects.NewListProjectsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Projects {
			resources = append(resources, *resource.ProjectName)
		}
	}
	return
}

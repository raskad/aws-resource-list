package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
)

func getIoT1ClickProjects(config aws.Config) (resources awsResourceMap) {
	client := iot1clickprojects.New(config)

	ioT1ClickProjectsProjectNames := getIoT1ClickProjectsProjectNames(client)

	resources = awsResourceMap{
		ioT1ClickProjectsProject: ioT1ClickProjectsProjectNames,
	}
	return
}

func getIoT1ClickProjectsProjectNames(client *iot1clickprojects.Client) (resources []string) {
	req := client.ListProjectsRequest(&iot1clickprojects.ListProjectsInput{})
	p := iot1clickprojects.NewListProjectsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Projects {
			resources = append(resources, *resource.ProjectName)
		}
	}
	return
}

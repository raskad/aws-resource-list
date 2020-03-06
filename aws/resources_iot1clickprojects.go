package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
)

func getIoT1ClickProjects(config aws.Config) (resources resourceMap) {
	client := iot1clickprojects.New(config)
	resources = reduce(
		getIoT1ClickProject(client).unwrap(ioT1ClickProject),
	)
	return
}

func getIoT1ClickProject(client *iot1clickprojects.Client) (r resourceSliceError) {
	req := client.ListProjectsRequest(&iot1clickprojects.ListProjectsInput{})
	p := iot1clickprojects.NewListProjectsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Projects {
			r.resources = append(r.resources, *resource.ProjectName)
		}
	}
	r.err = p.Err()
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
)

func getIoT1ClickProjects(session *session.Session) (resources resourceMap) {
	client := iot1clickprojects.New(session)
	resources = reduce(
		getIoT1ClickProject(client).unwrap(ioT1ClickProject),
	)
	return
}

func getIoT1ClickProject(client *iot1clickprojects.IoT1ClickProjects) (r resourceSliceError) {
	r.err = client.ListProjectsPages(&iot1clickprojects.ListProjectsInput{}, func(page *iot1clickprojects.ListProjectsOutput, lastPage bool) bool {
		for _, resource := range page.Projects {
			r.resources = append(r.resources, *resource.ProjectName)
		}
		return true
	})
	return
}

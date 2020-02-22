package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
)

func getIoT1ClickProjects(session *session.Session) (resources resourceMap) {
	client := iot1clickprojects.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		ioT1ClickProject: getIoT1ClickProject(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getIoT1ClickProject(client *iot1clickprojects.IoT1ClickProjects) (r resourceSliceError) {
	r.err = client.ListProjectsPages(&iot1clickprojects.ListProjectsInput{}, func(page *iot1clickprojects.ListProjectsOutput, lastPage bool) bool {
		logDebug("Listing IoT1ClickProject resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Projects {
			logDebug("Got IoT1ClickProject resource with PhysicalResourceId", *resource.ProjectName)
			r.resources = append(r.resources, *resource.ProjectName)
		}
		return true
	})
	return
}

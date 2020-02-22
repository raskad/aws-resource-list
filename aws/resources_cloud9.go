package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloud9"
)

func getCloud9(session *session.Session) (resources resourceMap) {
	client := cloud9.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		cloud9EnvironmentEC2: getCloud9EnvironmentEC2(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCloud9EnvironmentEC2(client *cloud9.Cloud9) (r resourceSliceError) {
	r.err = client.ListEnvironmentsPages(&cloud9.ListEnvironmentsInput{}, func(page *cloud9.ListEnvironmentsOutput, lastPage bool) bool {
		logDebug("List Cloud9EnvironmentEC2 resources page. Remaining pages", page.NextToken)
		for _, resource := range page.EnvironmentIds {
			logDebug("Got Cloud9EnvironmentEC2 resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloud9"
)

func getCloud9(session *session.Session) (resources resourceMap) {
	client := cloud9.New(session)
	resources = reduce(
		getCloud9EnvironmentEC2(client).unwrap(cloud9EnvironmentEC2),
	)
	return
}

func getCloud9EnvironmentEC2(client *cloud9.Cloud9) (r resourceSliceError) {
	r.err = client.ListEnvironmentsPages(&cloud9.ListEnvironmentsInput{}, func(page *cloud9.ListEnvironmentsOutput, lastPage bool) bool {
		for _, resource := range page.EnvironmentIds {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

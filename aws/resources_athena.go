package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

func getAthena(session *session.Session) (resources resourceMap) {
	client := athena.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		athenaNamedQuery: getAthenaNamedQuery(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAthenaNamedQuery(client *athena.Athena) (r resourceSliceError) {
	logDebug("Listing AthenaNamedQuery resources")
	r.err = client.ListNamedQueriesPages(&athena.ListNamedQueriesInput{}, func(page *athena.ListNamedQueriesOutput, lastPage bool) bool {
		for _, resource := range page.NamedQueryIds {
			logDebug("Got AthenaNamedQuery resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

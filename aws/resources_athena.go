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
	r.err = client.ListNamedQueriesPages(&athena.ListNamedQueriesInput{}, func(page *athena.ListNamedQueriesOutput, lastPage bool) bool {
		logDebug("List AthenaNamedQuery resources page. Remaining pages", page.NextToken)
		for _, resource := range page.NamedQueryIds {
			logDebug("Got AthenaNamedQuery resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

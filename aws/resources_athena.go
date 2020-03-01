package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

func getAthena(session *session.Session) (resources resourceMap) {
	client := athena.New(session)
	resources = reduce(
		getAthenaNamedQuery(client).unwrap(athenaNamedQuery),
	)
	return
}

func getAthenaNamedQuery(client *athena.Athena) (r resourceSliceError) {
	r.err = client.ListNamedQueriesPages(&athena.ListNamedQueriesInput{}, func(page *athena.ListNamedQueriesOutput, lastPage bool) bool {
		for _, resource := range page.NamedQueryIds {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

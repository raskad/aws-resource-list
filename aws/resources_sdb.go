package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/simpledb"
)

func getSdb(session *session.Session) (resources resourceMap) {
	client := simpledb.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		sdbDomain: getSdbDomain(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getSdbDomain(client *simpledb.SimpleDB) (r resourceSliceError) {
	r.err = client.ListDomainsPages(&simpledb.ListDomainsInput{}, func(page *simpledb.ListDomainsOutput, lastPage bool) bool {
		logDebug("List SdbDomain resources page. Remaining pages", page.NextToken)
		for _, resource := range page.DomainNames {
			logDebug("Got SdbDomain resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

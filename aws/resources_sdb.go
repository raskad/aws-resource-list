package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/simpledb"
)

func getSdb(session *session.Session) (resources resourceMap) {
	client := simpledb.New(session)
	resources = reduce(
		getSdbDomain(client).unwrap(sdbDomain),
	)
	return
}

func getSdbDomain(client *simpledb.SimpleDB) (r resourceSliceError) {
	r.err = client.ListDomainsPages(&simpledb.ListDomainsInput{}, func(page *simpledb.ListDomainsOutput, lastPage bool) bool {
		for _, resource := range page.DomainNames {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

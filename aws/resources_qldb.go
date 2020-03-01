package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldb"
)

func getQLDB(session *session.Session) (resources resourceMap) {
	client := qldb.New(session)
	resources = reduce(
		getQLDBLedger(client).unwrap(qLDBLedger),
	)
	return
}

func getQLDBLedger(client *qldb.QLDB) (r resourceSliceError) {
	r.err = client.ListLedgersPages(&qldb.ListLedgersInput{}, func(page *qldb.ListLedgersOutput, lastPage bool) bool {
		for _, resource := range page.Ledgers {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

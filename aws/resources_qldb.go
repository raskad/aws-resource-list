package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldb"
)

func getQLDB(session *session.Session) (resources resourceMap) {
	client := qldb.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		qLDBLedger: getQLDBLedger(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getQLDBLedger(client *qldb.QLDB) (r resourceSliceError) {
	logDebug("Listing QLDBLedger resources")
	r.err = client.ListLedgersPages(&qldb.ListLedgersInput{}, func(page *qldb.ListLedgersOutput, lastPage bool) bool {
		for _, resource := range page.Ledgers {
			logDebug("Got QLDBLedger resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transfer"
)

func getTransfer(session *session.Session) (resources resourceMap) {
	client := transfer.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		transferServer: getTransferServer(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getTransferServer(client *transfer.Transfer) (r resourceSliceError) {
	r.err = client.ListServersPages(&transfer.ListServersInput{}, func(page *transfer.ListServersOutput, lastPage bool) bool {
		logDebug("List TransferServer resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Servers {
			logDebug("Got TransferServer resource with PhysicalResourceId", *resource.ServerId)
			r.resources = append(r.resources, *resource.ServerId)
		}
		return true
	})
	return
}

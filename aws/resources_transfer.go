package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transfer"
)

func getTransfer(session *session.Session) (resources resourceMap) {
	client := transfer.New(session)
	resources = reduce(
		getTransferServer(client).unwrap(transferServer),
	)
	return
}

func getTransferServer(client *transfer.Transfer) (r resourceSliceError) {
	r.err = client.ListServersPages(&transfer.ListServersInput{}, func(page *transfer.ListServersOutput, lastPage bool) bool {
		for _, resource := range page.Servers {
			r.resources = append(r.resources, *resource.ServerId)
		}
		return true
	})
	return
}

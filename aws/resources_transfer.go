package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
)

func getTransfer(config aws.Config) (resources resourceMap) {
	client := transfer.New(config)
	resources = reduce(
		getTransferServer(client).unwrap(transferServer),
	)
	return
}

func getTransferServer(client *transfer.Client) (r resourceSliceError) {
	req := client.ListServersRequest(&transfer.ListServersInput{})
	p := transfer.NewListServersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Servers {
			r.resources = append(r.resources, *resource.ServerId)
		}
	}
	r.err = p.Err()
	return
}

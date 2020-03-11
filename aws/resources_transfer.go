package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
)

func getTransfer(config aws.Config) (resources resourceMap) {
	client := transfer.New(config)

	transferServerResourceMap := getTransferServer(client).unwrap(transferServer)
	transferServerIDs := transferServerResourceMap[transferServer]

	resources = reduce(
		transferServerResourceMap,
		getTransferUser(client, transferServerIDs).unwrap(transferUser),
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

func getTransferUser(client *transfer.Client, serverIDs []string) (r resourceSliceError) {
	for _, serverID := range serverIDs {
		req := client.ListUsersRequest(&transfer.ListUsersInput{
			ServerId: aws.String(serverID),
		})
		p := transfer.NewListUsersPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.Users {
				r.resources = append(r.resources, *resource.Arn)
			}
		}
		r.err = p.Err()
	}
	return
}

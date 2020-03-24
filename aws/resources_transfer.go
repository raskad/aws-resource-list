package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
)

func getTransfer(config aws.Config) (resources resourceMap) {
	client := transfer.New(config)

	transferServerIDs := getTransferServerIDs(client)
	transferUserARNs := getTransferUserARNs(client, transferServerIDs)

	resources = resourceMap{
		transferServer: transferServerIDs,
		transferUser:   transferUserARNs,
	}
	return
}

func getTransferServerIDs(client *transfer.Client) (resources []string) {
	req := client.ListServersRequest(&transfer.ListServersInput{})
	p := transfer.NewListServersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Servers {
			resources = append(resources, *resource.ServerId)
		}
	}
	return
}

func getTransferUserARNs(client *transfer.Client, serverIDs []string) (resources []string) {
	for _, serverID := range serverIDs {
		req := client.ListUsersRequest(&transfer.ListUsersInput{
			ServerId: aws.String(serverID),
		})
		p := transfer.NewListUsersPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.Users {
				resources = append(resources, *resource.Arn)
			}
		}
	}
	return
}

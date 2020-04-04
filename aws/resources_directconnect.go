package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func getDirectConnect(config aws.Config) (resources awsResourceMap) {
	client := directconnect.New(config)

	directConnectConnectionIDs := getDirectConnectConnectionIDs(client)
	directConnectGatewayIDs := getDirectConnectGatewayIDs(client)
	directConnectGatewayAssociationIDs := getDirectConnectGatewayAssociationIDs(client)
	directConnectGatewayAssociationProposalIDs := getDirectConnectGatewayAssociationProposalIDs(client)
	directConnectLAGIDs := getDirectConnectLAGIDs(client)
	directConnectVirtualInterfaceIDs := getDirectConnectVirtualInterfaceIDs(client)

	resources = awsResourceMap{
		directConnectConnection:                 directConnectConnectionIDs,
		directConnectGateway:                    directConnectGatewayIDs,
		directConnectGatewayAssociation:         directConnectGatewayAssociationIDs,
		directConnectGatewayAssociationProposal: directConnectGatewayAssociationProposalIDs,
		directConnectLAG:                        directConnectLAGIDs,
		directConnectVirtualInterface:           directConnectVirtualInterfaceIDs,
	}
	return
}

func getDirectConnectConnectionIDs(client *directconnect.Client) (resources []string) {
	input := directconnect.DescribeConnectionsInput{}
	page, err := client.DescribeConnectionsRequest(&input).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.Connections {
		resources = append(resources, *resource.ConnectionId)
	}
	return
}

func getDirectConnectGatewayIDs(client *directconnect.Client) (resources []string) {
	input := directconnect.DescribeDirectConnectGatewaysInput{}
	page, err := client.DescribeDirectConnectGatewaysRequest(&input).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.DirectConnectGateways {
		resources = append(resources, *resource.DirectConnectGatewayId)
	}
	return
}

func getDirectConnectGatewayAssociationIDs(client *directconnect.Client) (resources []string) {
	input := directconnect.DescribeDirectConnectGatewayAssociationsInput{}
	page, err := client.DescribeDirectConnectGatewayAssociationsRequest(&input).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.DirectConnectGatewayAssociations {
		resources = append(resources, *resource.AssociationId)
	}
	return
}

func getDirectConnectGatewayAssociationProposalIDs(client *directconnect.Client) (resources []string) {
	input := directconnect.DescribeDirectConnectGatewayAssociationProposalsInput{}
	page, err := client.DescribeDirectConnectGatewayAssociationProposalsRequest(&input).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.DirectConnectGatewayAssociationProposals {
		resources = append(resources, *resource.ProposalId)
	}
	return
}

func getDirectConnectLAGIDs(client *directconnect.Client) (resources []string) {
	input := directconnect.DescribeLagsInput{}
	page, err := client.DescribeLagsRequest(&input).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.Lags {
		resources = append(resources, *resource.LagId)
	}
	return
}

func getDirectConnectVirtualInterfaceIDs(client *directconnect.Client) (resources []string) {
	input := directconnect.DescribeVirtualInterfacesInput{}
	page, err := client.DescribeVirtualInterfacesRequest(&input).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.VirtualInterfaces {
		resources = append(resources, *resource.VirtualInterfaceId)
	}
	return
}

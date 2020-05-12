package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func getEc2(config aws.Config) (resources awsResourceMap) {
	client := ec2.New(config)

	ec2CapacityReservationIDs := getEc2CapacityReservationIDs(client)
	ec2ClientVpnEndpointIDs := getEc2ClientVpnEndpointIDs(client)
	ec2CustomerGatewayIDs := getEc2CustomerGatewayIDs(client)
	ec2DHCPOptionsIDs := getEc2DHCPOptionsIDs(client)
	ec2FleetIDs := getEc2FleetIDs(client)
	ec2EgressOnlyInternetGatewayIDs := getEc2EgressOnlyInternetGatewayIDs(client)
	ec2EipIDs := getEc2EipIDs(client)
	ec2EipAssociationIDs := getEc2EipAssociationIDs(client)
	ec2FlowLogIDs := getEc2FlowLogIDs(client)
	ec2HostIDs := getEc2HostIDs(client)
	ec2ImageIDs := getEc2ImageIDs(client)
	ec2InstaceIDs := getEc2InstaceIDs(client)
	ec2InternetGatewayIDs := getEc2InternetGatewayIDs(client)
	ec2LaunchTemplateIDs := getEc2LaunchTemplateIDs(client)
	ec2NatGatewayIDs := getEc2NatGatewayIDs(client)
	ec2NetworkACLIDs := getEc2NetworkACLIDs(client)
	ec2NetworkACLSubnetAssociationIDs := getEc2NetworkACLSubnetAssociationIDs(client)
	ec2NetworkInterfaceIDs := getEc2NetworkInterfaceIDs(client)
	ec2NetworkInterfaceAttachmentIDs := getEc2NetworkInterfaceAttachmentIDs(client)
	ec2NetworkInterfacePermissionIDs := getEc2NetworkInterfacePermissionIDs(client)
	ec2PlacementGroupIDs := getEc2PlacementGroupIDs(client)
	ec2RouteTableIDs := getEc2RouteTableIDs(client)
	ec2RouteTableSubnetAssociationIDs := getEc2RouteTableSubnetAssociationIDs(client)
	ec2SecurityGroupIDs := getEc2SecurityGroupIDs(client)
	ec2SnapshotIDs := getEc2SnapshotIDs(client)
	ec2SpotFleetIDs := getEc2SpotFleetIDs(client)
	ec2SubnetIDs := getEc2SubnetIDs(client)
	ec2TrafficMirrorFilterIDs := getEc2TrafficMirrorFilterIDs(client)
	ec2TrafficMirrorFilterRuleIDs := getEc2TrafficMirrorFilterRuleIDs(client)
	ec2TrafficMirrorSessionIDs := getEc2TrafficMirrorSessionIDs(client)
	ec2TrafficMirrorTargetIDs := getEc2TrafficMirrorTargetIDs(client)
	ec2TransitGatewayIDs := getEc2TransitGatewayIDs(client)
	ec2TransitGatewayAttachmentIDs := getEc2TransitGatewayAttachmentIDs(client)
	ec2TransitGatewayRouteTableIDs := getEc2TransitGatewayRouteTableIDs(client)
	ec2TransitGatewayPeeringAttachmentIDs := getEc2TransitGatewayPeeringAttachmentIDs(client)
	ec2VolumeIDs := getEc2VolumeIDs(client)
	ec2VPCIDs := getEc2VPCIDs(client)
	ec2VPCCidrBlockIDs := getEc2VPCCidrBlockIDs(client)
	ec2VPCEndpointIDs := getEc2VPCEndpointIDs(client)
	ec2VPCEndpointConnectionNotificationIDs := getEc2VPCEndpointConnectionNotificationIDs(client)
	ec2VPCEndpointServiceIDs := getEc2VPCEndpointServiceIDs(client)
	ec2VPCPeeringConnectionIDs := getEc2VPCPeeringConnectionIDs(client)
	ec2VPNConnectionIDs := getEc2VPNConnectionIDs(client)
	ec2VPNGatewayIDs := getEc2VPNGatewayIDs(client)
	ec2KeyPairIDs := getEc2KeyPairIDs(client)
	ec2SpotInstanceRequestIDs := getEc2SpotInstanceRequestIDs(client)

	resources = awsResourceMap{
		ec2CapacityReservation:               ec2CapacityReservationIDs,
		ec2ClientVpnEndpoint:                 ec2ClientVpnEndpointIDs,
		ec2CustomerGateway:                   ec2CustomerGatewayIDs,
		ec2DHCPOptions:                       ec2DHCPOptionsIDs,
		ec2EC2Fleet:                          ec2FleetIDs,
		ec2EgressOnlyInternetGateway:         ec2EgressOnlyInternetGatewayIDs,
		ec2EIP:                               ec2EipIDs,
		ec2EIPAssociation:                    ec2EipAssociationIDs,
		ec2FlowLog:                           ec2FlowLogIDs,
		ec2Host:                              ec2HostIDs,
		ec2Image:                             ec2ImageIDs,
		ec2Instance:                          ec2InstaceIDs,
		ec2InternetGateway:                   ec2InternetGatewayIDs,
		ec2LaunchTemplate:                    ec2LaunchTemplateIDs,
		ec2NatGateway:                        ec2NatGatewayIDs,
		ec2NetworkACL:                        ec2NetworkACLIDs,
		ec2NetworkACLSubnetAssociation:       ec2NetworkACLSubnetAssociationIDs,
		ec2NetworkInterface:                  ec2NetworkInterfaceIDs,
		ec2NetworkInterfaceAttachment:        ec2NetworkInterfaceAttachmentIDs,
		ec2NetworkInterfacePermission:        ec2NetworkInterfacePermissionIDs,
		ec2PlacementGroup:                    ec2PlacementGroupIDs,
		ec2RouteTable:                        ec2RouteTableIDs,
		ec2RouteTableSubnetAssociation:       ec2RouteTableSubnetAssociationIDs,
		ec2SecurityGroup:                     ec2SecurityGroupIDs,
		ec2Snapshot:                          ec2SnapshotIDs,
		ec2SpotFleet:                         ec2SpotFleetIDs,
		ec2Subnet:                            ec2SubnetIDs,
		ec2TrafficMirrorFilter:               ec2TrafficMirrorFilterIDs,
		ec2TrafficMirrorFilterRule:           ec2TrafficMirrorFilterRuleIDs,
		ec2TrafficMirrorSession:              ec2TrafficMirrorSessionIDs,
		ec2TrafficMirrorTarget:               ec2TrafficMirrorTargetIDs,
		ec2TransitGateway:                    ec2TransitGatewayIDs,
		ec2TransitGatewayAttachment:          ec2TransitGatewayAttachmentIDs,
		ec2TransitGatewayRouteTable:          ec2TransitGatewayRouteTableIDs,
		ec2TransitGatewayPeeringAttachment:   ec2TransitGatewayPeeringAttachmentIDs,
		ec2Volume:                            ec2VolumeIDs,
		ec2VPC:                               ec2VPCIDs,
		ec2VPCCidrBlock:                      ec2VPCCidrBlockIDs,
		ec2VPCEndpoint:                       ec2VPCEndpointIDs,
		ec2VPCEndpointConnectionNotification: ec2VPCEndpointConnectionNotificationIDs,
		ec2VPCEndpointService:                ec2VPCEndpointServiceIDs,
		ec2VPCPeeringConnection:              ec2VPCPeeringConnectionIDs,
		ec2VPNConnection:                     ec2VPNConnectionIDs,
		ec2VPNGateway:                        ec2VPNGatewayIDs,
		ec2KeyPair:                           ec2KeyPairIDs,
		ec2SpotInstanceRequest:               ec2SpotInstanceRequestIDs,
	}
	return
}

func getEc2SnapshotIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeSnapshotsRequest(&ec2.DescribeSnapshotsInput{
		OwnerIds: []string{accountID},
	})
	p := ec2.NewDescribeSnapshotsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Snapshots {
			resources = append(resources, *resource.SnapshotId)
		}
	}
	return
}

func getEc2CapacityReservationIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeCapacityReservationsRequest(&ec2.DescribeCapacityReservationsInput{})
	p := ec2.NewDescribeCapacityReservationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.CapacityReservations {
			resources = append(resources, *resource.CapacityReservationId)
		}
	}
	return
}

func getEc2ClientVpnEndpointIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeClientVpnEndpointsRequest(&ec2.DescribeClientVpnEndpointsInput{})
	p := ec2.NewDescribeClientVpnEndpointsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ClientVpnEndpoints {
			resources = append(resources, *resource.ClientVpnEndpointId)
		}
	}
	return
}

func getEc2CustomerGatewayIDs(client *ec2.Client) (resources []string) {
	page, err := client.DescribeCustomerGatewaysRequest(&ec2.DescribeCustomerGatewaysInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.CustomerGateways {
		resources = append(resources, *resource.CustomerGatewayId)
	}
	return
}

func getEc2DHCPOptionsIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeDhcpOptionsRequest(&ec2.DescribeDhcpOptionsInput{})
	p := ec2.NewDescribeDhcpOptionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DhcpOptions {
			resources = append(resources, *resource.DhcpOptionsId)
		}
	}
	return
}

func getEc2FleetIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeFleetsRequest(&ec2.DescribeFleetsInput{})
	p := ec2.NewDescribeFleetsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Fleets {
			resources = append(resources, *resource.FleetId)
		}
	}
	return
}

func getEc2EgressOnlyInternetGatewayIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeEgressOnlyInternetGatewaysRequest(&ec2.DescribeEgressOnlyInternetGatewaysInput{})
	p := ec2.NewDescribeEgressOnlyInternetGatewaysPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.EgressOnlyInternetGateways {
			resources = append(resources, *resource.EgressOnlyInternetGatewayId)
		}
	}
	return
}

func getEc2EipIDs(client *ec2.Client) (resources []string) {
	page, err := client.DescribeAddressesRequest(&ec2.DescribeAddressesInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.Addresses {
		resources = append(resources, *resource.AllocationId)
	}
	return
}

func getEc2EipAssociationIDs(client *ec2.Client) (resources []string) {
	page, err := client.DescribeAddressesRequest(&ec2.DescribeAddressesInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.Addresses {
		if resource.AssociationId != nil {
			resources = append(resources, *resource.AssociationId)
		}
	}
	return
}

func getEc2FlowLogIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeFlowLogsRequest(&ec2.DescribeFlowLogsInput{})
	p := ec2.NewDescribeFlowLogsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.FlowLogs {
			resources = append(resources, *resource.FlowLogId)
		}
	}
	return
}

func getEc2HostIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeHostsRequest(&ec2.DescribeHostsInput{})
	p := ec2.NewDescribeHostsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Hosts {
			resources = append(resources, *resource.HostId)
		}
	}
	return
}

func getEc2ImageIDs(client *ec2.Client) (resources []string) {
	page, err := client.DescribeImagesRequest(&ec2.DescribeImagesInput{
		Owners: []string{"self"},
	}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.Images {
		resources = append(resources, *resource.ImageId)
	}
	return
}

func getEc2InstaceIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	p := ec2.NewDescribeInstancesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Reservations {
			for _, resource := range resource.Instances {
				resources = append(resources, *resource.InstanceId)
			}
		}
	}
	return
}

func getEc2InternetGatewayIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeInternetGatewaysRequest(&ec2.DescribeInternetGatewaysInput{})
	p := ec2.NewDescribeInternetGatewaysPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.InternetGateways {
			resources = append(resources, *resource.InternetGatewayId)
		}
	}
	return
}

func getEc2LaunchTemplateIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeLaunchTemplatesRequest(&ec2.DescribeLaunchTemplatesInput{})
	p := ec2.NewDescribeLaunchTemplatesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.LaunchTemplates {
			resources = append(resources, *resource.LaunchTemplateId)
		}
	}
	return
}

func getEc2NatGatewayIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeNatGatewaysRequest(&ec2.DescribeNatGatewaysInput{})
	p := ec2.NewDescribeNatGatewaysPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.NatGateways {
			resources = append(resources, *resource.NatGatewayId)
		}
	}
	return
}

func getEc2NetworkACLIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{})
	p := ec2.NewDescribeNetworkAclsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.NetworkAcls {
			resources = append(resources, *resource.NetworkAclId)
		}
	}
	return
}

func getEc2NetworkACLSubnetAssociationIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{})
	p := ec2.NewDescribeNetworkAclsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.NetworkAcls {
			for _, resource := range resource.Associations {
				resources = append(resources, *resource.NetworkAclAssociationId)
			}
		}
	}
	return
}

func getEc2NetworkInterfaceIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeNetworkInterfacesRequest(&ec2.DescribeNetworkInterfacesInput{})
	p := ec2.NewDescribeNetworkInterfacesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.NetworkInterfaces {
			resources = append(resources, *resource.NetworkInterfaceId)
		}
	}
	return
}

func getEc2NetworkInterfaceAttachmentIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeNetworkInterfacesRequest(&ec2.DescribeNetworkInterfacesInput{})
	p := ec2.NewDescribeNetworkInterfacesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.NetworkInterfaces {
			if resource.Attachment != nil {
				resources = append(resources, *resource.Attachment.AttachmentId)
			}
		}
	}
	return
}

func getEc2NetworkInterfacePermissionIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeNetworkInterfacePermissionsRequest(&ec2.DescribeNetworkInterfacePermissionsInput{})
	p := ec2.NewDescribeNetworkInterfacePermissionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.NetworkInterfacePermissions {
			resources = append(resources, *resource.NetworkInterfacePermissionId)
		}
	}
	return
}

func getEc2PlacementGroupIDs(client *ec2.Client) (resources []string) {
	page, err := client.DescribePlacementGroupsRequest(&ec2.DescribePlacementGroupsInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.PlacementGroups {
		resources = append(resources, *resource.GroupId)
	}
	return
}

func getEc2RouteTableIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeRouteTablesRequest(&ec2.DescribeRouteTablesInput{})
	p := ec2.NewDescribeRouteTablesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.RouteTables {
			resources = append(resources, *resource.RouteTableId)
		}
	}
	return
}

func getEc2RouteTableSubnetAssociationIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeRouteTablesRequest(&ec2.DescribeRouteTablesInput{})
	p := ec2.NewDescribeRouteTablesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.RouteTables {
			for _, resource := range resource.Associations {
				resources = append(resources, *resource.RouteTableAssociationId)
			}
		}
	}
	return
}

func getEc2SecurityGroupIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeSecurityGroupsRequest(&ec2.DescribeSecurityGroupsInput{})
	p := ec2.NewDescribeSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.SecurityGroups {
			resources = append(resources, *resource.GroupId)
		}
	}
	return
}

func getEc2SpotFleetIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeSpotFleetRequestsRequest(&ec2.DescribeSpotFleetRequestsInput{})
	p := ec2.NewDescribeSpotFleetRequestsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.SpotFleetRequestConfigs {
			resources = append(resources, *resource.SpotFleetRequestId)
		}
	}
	return
}

func getEc2SubnetIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeSubnetsRequest(&ec2.DescribeSubnetsInput{})
	p := ec2.NewDescribeSubnetsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Subnets {
			resources = append(resources, *resource.SubnetId)
		}
	}
	return
}

func getEc2TrafficMirrorFilterIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeTrafficMirrorFiltersRequest(&ec2.DescribeTrafficMirrorFiltersInput{})
	p := ec2.NewDescribeTrafficMirrorFiltersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.TrafficMirrorFilters {
			resources = append(resources, *resource.TrafficMirrorFilterId)
		}
	}
	return
}

func getEc2TrafficMirrorFilterRuleIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeTrafficMirrorFiltersRequest(&ec2.DescribeTrafficMirrorFiltersInput{})
	p := ec2.NewDescribeTrafficMirrorFiltersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.TrafficMirrorFilters {
			for _, resource := range resource.EgressFilterRules {
				resources = append(resources, *resource.TrafficMirrorFilterRuleId)
			}
			for _, resource := range resource.IngressFilterRules {
				resources = append(resources, *resource.TrafficMirrorFilterRuleId)
			}
		}
	}
	return
}

func getEc2TrafficMirrorSessionIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeTrafficMirrorSessionsRequest(&ec2.DescribeTrafficMirrorSessionsInput{})
	p := ec2.NewDescribeTrafficMirrorSessionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.TrafficMirrorSessions {
			resources = append(resources, *resource.TrafficMirrorSessionId)
		}
	}
	return
}

func getEc2TrafficMirrorTargetIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeTrafficMirrorTargetsRequest(&ec2.DescribeTrafficMirrorTargetsInput{})
	p := ec2.NewDescribeTrafficMirrorTargetsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.TrafficMirrorTargets {
			resources = append(resources, *resource.TrafficMirrorTargetId)
		}
	}
	return
}

func getEc2TransitGatewayIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeTransitGatewaysRequest(&ec2.DescribeTransitGatewaysInput{})
	p := ec2.NewDescribeTransitGatewaysPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.TransitGateways {
			resources = append(resources, *resource.TransitGatewayId)
		}
	}
	return
}

func getEc2TransitGatewayAttachmentIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeTransitGatewayAttachmentsRequest(&ec2.DescribeTransitGatewayAttachmentsInput{})
	p := ec2.NewDescribeTransitGatewayAttachmentsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.TransitGatewayAttachments {
			resources = append(resources, *resource.TransitGatewayAttachmentId)
		}
	}
	return
}

func getEc2TransitGatewayRouteTableIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeTransitGatewayRouteTablesRequest(&ec2.DescribeTransitGatewayRouteTablesInput{})
	p := ec2.NewDescribeTransitGatewayRouteTablesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.TransitGatewayRouteTables {
			resources = append(resources, *resource.TransitGatewayRouteTableId)
		}
	}
	return
}

func getEc2TransitGatewayPeeringAttachmentIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeTransitGatewayPeeringAttachmentsRequest(&ec2.DescribeTransitGatewayPeeringAttachmentsInput{})
	p := ec2.NewDescribeTransitGatewayPeeringAttachmentsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.TransitGatewayPeeringAttachments {
			resources = append(resources, *resource.TransitGatewayAttachmentId)
		}
	}
	return
}

func getEc2VolumeIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
	p := ec2.NewDescribeVolumesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Volumes {
			resources = append(resources, *resource.VolumeId)
		}
	}
	return
}

func getEc2VPCIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})
	p := ec2.NewDescribeVpcsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Vpcs {
			resources = append(resources, *resource.VpcId)
		}
	}
	return
}

func getEc2VPCCidrBlockIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})
	p := ec2.NewDescribeVpcsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Vpcs {
			for _, resource := range resource.CidrBlockAssociationSet {
				resources = append(resources, *resource.AssociationId)
			}
		}
	}
	return
}

func getEc2VPCEndpointIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeVpcEndpointsRequest(&ec2.DescribeVpcEndpointsInput{})
	p := ec2.NewDescribeVpcEndpointsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.VpcEndpoints {
			resources = append(resources, *resource.VpcEndpointId)
		}
	}
	return
}

func getEc2VPCEndpointConnectionNotificationIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeVpcEndpointConnectionNotificationsRequest(&ec2.DescribeVpcEndpointConnectionNotificationsInput{})
	p := ec2.NewDescribeVpcEndpointConnectionNotificationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ConnectionNotificationSet {
			resources = append(resources, *resource.ConnectionNotificationId)
		}
	}
	return
}

func getEc2VPCEndpointServiceIDs(client *ec2.Client) (resources []string) {
	page, err := client.DescribeVpcEndpointServicesRequest(&ec2.DescribeVpcEndpointServicesInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.ServiceDetails {
		resources = append(resources, *resource.ServiceId)
	}
	return
}

func getEc2VPCPeeringConnectionIDs(client *ec2.Client) (resources []string) {
	req := client.DescribeVpcPeeringConnectionsRequest(&ec2.DescribeVpcPeeringConnectionsInput{})
	p := ec2.NewDescribeVpcPeeringConnectionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.VpcPeeringConnections {
			resources = append(resources, *resource.VpcPeeringConnectionId)
		}
	}
	return
}

func getEc2VPNConnectionIDs(client *ec2.Client) (resources []string) {
	page, err := client.DescribeVpnConnectionsRequest(&ec2.DescribeVpnConnectionsInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.VpnConnections {
		if resource.ConnectionId == nil {
			continue
		}
		resources = append(resources, *resource.ConnectionId)
	}
	return
}

func getEc2VPNGatewayIDs(client *ec2.Client) (resources []string) {
	page, err := client.DescribeVpnGatewaysRequest(&ec2.DescribeVpnGatewaysInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.VpnGateways {
		resources = append(resources, *resource.VpnGatewayId)
	}
	return
}

func getEc2KeyPairIDs(client *ec2.Client) (resources []string) {
	page, err := client.DescribeKeyPairsRequest(&ec2.DescribeKeyPairsInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.KeyPairs {
		resources = append(resources, *resource.KeyPairId)
	}
	return
}

func getEc2SpotInstanceRequestIDs(client *ec2.Client) (resources []string) {
	input := ec2.DescribeSpotInstanceRequestsInput{}
	for {
		page, err := client.DescribeSpotInstanceRequestsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.SpotInstanceRequests {
			resources = append(resources, *resource.SpotInstanceRequestId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

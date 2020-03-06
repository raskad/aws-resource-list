package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func getEc2(config aws.Config) (resources resourceMap) {
	client := ec2.New(config)
	resources = reduce(
		getEc2CapacityReservation(client).unwrap(ec2CapacityReservation),
		getEc2ClientVpnEndpoint(client).unwrap(ec2ClientVpnEndpoint),
		getEc2CustomerGateway(client).unwrap(ec2CustomerGateway),
		getEc2DHCPOptions(client).unwrap(ec2DHCPOptions),
		getEc2Fleet(client).unwrap(ec2EC2Fleet),
		getEc2EgressOnlyInternetGateway(client).unwrap(ec2EgressOnlyInternetGateway),
		getEc2Eip(client).unwrap(ec2EIP),
		getEc2EipAssociation(client).unwrap(ec2EIPAssociation),
		getEc2FlowLog(client).unwrap(ec2FlowLog),
		getEc2Host(client).unwrap(ec2Host),
		getEc2Instace(client).unwrap(ec2Instance),
		getEc2InternetGateway(client).unwrap(ec2InternetGateway),
		getEc2LaunchTemplate(client).unwrap(ec2LaunchTemplate),
		getEc2NatGateway(client).unwrap(ec2NatGateway),
		getEc2NetworkACL(client).unwrap(ec2NetworkACL),
		getEc2NetworkACLSubnetAssociation(client).unwrap(ec2NetworkACLSubnetAssociation),
		getEc2NetworkInterface(client).unwrap(ec2NetworkInterface),
		getEc2NetworkInterfaceAttachment(client).unwrap(ec2NetworkInterfaceAttachment),
		getEc2NetworkInterfacePermission(client).unwrap(ec2NetworkInterfacePermission),
		getEc2PlacementGroup(client).unwrap(ec2PlacementGroup),
		getEc2RouteTable(client).unwrap(ec2RouteTable),
		getEc2RouteTableSubnetAssociation(client).unwrap(ec2RouteTableSubnetAssociation),
		getEc2SecurityGroup(client).unwrap(ec2SecurityGroup),
		getEc2SpotFleet(client).unwrap(ec2SpotFleet),
		getEc2Subnet(client).unwrap(ec2Subnet),
		getEc2TrafficMirrorFilter(client).unwrap(ec2TrafficMirrorFilter),
		getEc2TrafficMirrorFilterRule(client).unwrap(ec2TrafficMirrorFilterRule),
		getEc2TrafficMirrorSession(client).unwrap(ec2TrafficMirrorSession),
		getEc2TrafficMirrorTarget(client).unwrap(ec2TrafficMirrorTarget),
		getEc2TransitGateway(client).unwrap(ec2TransitGateway),
		getEc2TransitGatewayAttachment(client).unwrap(ec2TransitGatewayAttachment),
		getEc2TransitGatewayRouteTable(client).unwrap(ec2TransitGatewayRouteTable),
		getEc2Volume(client).unwrap(ec2Volume),
		getEc2VPC(client).unwrap(ec2VPC),
		getEc2VPCCidrBlock(client).unwrap(ec2VPCCidrBlock),
		getEc2VPCEndpoint(client).unwrap(ec2VPCEndpoint),
		getEc2VPCEndpointConnectionNotification(client).unwrap(ec2VPCEndpointConnectionNotification),
		getEc2VPCEndpointService(client).unwrap(ec2VPCEndpointService),
		getEc2VPCPeeringConnection(client).unwrap(ec2VPCPeeringConnection),
		getEc2VPNConnection(client).unwrap(ec2VPNConnection),
		getEc2VPNGateway(client).unwrap(ec2VPNGateway),
	)
	return
}

func getEc2CapacityReservation(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeCapacityReservationsRequest(&ec2.DescribeCapacityReservationsInput{})
	p := ec2.NewDescribeCapacityReservationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.CapacityReservations {
			r.resources = append(r.resources, *resource.CapacityReservationId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2ClientVpnEndpoint(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeClientVpnEndpointsRequest(&ec2.DescribeClientVpnEndpointsInput{})
	p := ec2.NewDescribeClientVpnEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ClientVpnEndpoints {
			r.resources = append(r.resources, *resource.ClientVpnEndpointId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2CustomerGateway(client *ec2.Client) (r resourceSliceError) {
	output, err := client.DescribeCustomerGatewaysRequest(&ec2.DescribeCustomerGatewaysInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.CustomerGateways {
		r.resources = append(r.resources, *resource.CustomerGatewayId)
	}
	return
}

func getEc2DHCPOptions(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeDhcpOptionsRequest(&ec2.DescribeDhcpOptionsInput{})
	p := ec2.NewDescribeDhcpOptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DhcpOptions {
			r.resources = append(r.resources, *resource.DhcpOptionsId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2Fleet(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeFleetsRequest(&ec2.DescribeFleetsInput{})
	p := ec2.NewDescribeFleetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Fleets {
			r.resources = append(r.resources, *resource.FleetId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2EgressOnlyInternetGateway(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeEgressOnlyInternetGatewaysRequest(&ec2.DescribeEgressOnlyInternetGatewaysInput{})
	p := ec2.NewDescribeEgressOnlyInternetGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.EgressOnlyInternetGateways {
			r.resources = append(r.resources, *resource.EgressOnlyInternetGatewayId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2Eip(client *ec2.Client) (r resourceSliceError) {
	output, err := client.DescribeAddressesRequest(&ec2.DescribeAddressesInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Addresses {
		r.resources = append(r.resources, *resource.AllocationId)
	}
	return
}

func getEc2EipAssociation(client *ec2.Client) (r resourceSliceError) {
	output, err := client.DescribeAddressesRequest(&ec2.DescribeAddressesInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Addresses {
		if resource.AssociationId != nil {
			r.resources = append(r.resources, *resource.AssociationId)
		}
	}
	return
}

func getEc2FlowLog(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeFlowLogsRequest(&ec2.DescribeFlowLogsInput{})
	p := ec2.NewDescribeFlowLogsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.FlowLogs {
			r.resources = append(r.resources, *resource.FlowLogId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2Host(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeHostsRequest(&ec2.DescribeHostsInput{})
	p := ec2.NewDescribeHostsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Hosts {
			r.resources = append(r.resources, *resource.HostId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2Instace(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	p := ec2.NewDescribeInstancesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Reservations {
			for _, resource := range resource.Instances {
				r.resources = append(r.resources, *resource.InstanceId)
			}
		}
	}
	r.err = p.Err()
	return
}

func getEc2InternetGateway(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeInternetGatewaysRequest(&ec2.DescribeInternetGatewaysInput{})
	p := ec2.NewDescribeInternetGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.InternetGateways {
			r.resources = append(r.resources, *resource.InternetGatewayId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2LaunchTemplate(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeLaunchTemplatesRequest(&ec2.DescribeLaunchTemplatesInput{})
	p := ec2.NewDescribeLaunchTemplatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.LaunchTemplates {
			r.resources = append(r.resources, *resource.LaunchTemplateId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2NatGateway(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeNatGatewaysRequest(&ec2.DescribeNatGatewaysInput{})
	p := ec2.NewDescribeNatGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.NatGateways {
			r.resources = append(r.resources, *resource.NatGatewayId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2NetworkACL(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{})
	p := ec2.NewDescribeNetworkAclsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.NetworkAcls {
			r.resources = append(r.resources, *resource.NetworkAclId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2NetworkACLSubnetAssociation(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{})
	p := ec2.NewDescribeNetworkAclsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.NetworkAcls {
			for _, resource := range resource.Associations {
				r.resources = append(r.resources, *resource.NetworkAclAssociationId)
			}
		}
	}
	r.err = p.Err()
	return
}

func getEc2NetworkInterface(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeNetworkInterfacesRequest(&ec2.DescribeNetworkInterfacesInput{})
	p := ec2.NewDescribeNetworkInterfacesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.NetworkInterfaces {
			r.resources = append(r.resources, *resource.NetworkInterfaceId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2NetworkInterfaceAttachment(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeNetworkInterfacesRequest(&ec2.DescribeNetworkInterfacesInput{})
	p := ec2.NewDescribeNetworkInterfacesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.NetworkInterfaces {
			r.resources = append(r.resources, *resource.Attachment.AttachmentId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2NetworkInterfacePermission(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeNetworkInterfacePermissionsRequest(&ec2.DescribeNetworkInterfacePermissionsInput{})
	p := ec2.NewDescribeNetworkInterfacePermissionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.NetworkInterfacePermissions {
			r.resources = append(r.resources, *resource.NetworkInterfacePermissionId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2PlacementGroup(client *ec2.Client) (r resourceSliceError) {
	page, err := client.DescribePlacementGroupsRequest(&ec2.DescribePlacementGroupsInput{}).Send(context.Background())
	for _, resource := range page.PlacementGroups {
		r.resources = append(r.resources, *resource.GroupId)
	}
	r.err = err
	return
}

func getEc2RouteTable(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeRouteTablesRequest(&ec2.DescribeRouteTablesInput{})
	p := ec2.NewDescribeRouteTablesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.RouteTables {
			r.resources = append(r.resources, *resource.RouteTableId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2RouteTableSubnetAssociation(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeRouteTablesRequest(&ec2.DescribeRouteTablesInput{})
	p := ec2.NewDescribeRouteTablesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.RouteTables {
			for _, resource := range resource.Associations {
				r.resources = append(r.resources, *resource.RouteTableAssociationId)
			}
		}
	}
	r.err = p.Err()
	return
}

func getEc2SecurityGroup(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeSecurityGroupsRequest(&ec2.DescribeSecurityGroupsInput{})
	p := ec2.NewDescribeSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.SecurityGroups {
			r.resources = append(r.resources, *resource.GroupId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2SpotFleet(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeSpotFleetRequestsRequest(&ec2.DescribeSpotFleetRequestsInput{})
	p := ec2.NewDescribeSpotFleetRequestsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.SpotFleetRequestConfigs {
			r.resources = append(r.resources, *resource.SpotFleetRequestId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2Subnet(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeSubnetsRequest(&ec2.DescribeSubnetsInput{})
	p := ec2.NewDescribeSubnetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Subnets {
			r.resources = append(r.resources, *resource.SubnetId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2TrafficMirrorFilter(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeTrafficMirrorFiltersRequest(&ec2.DescribeTrafficMirrorFiltersInput{})
	p := ec2.NewDescribeTrafficMirrorFiltersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TrafficMirrorFilters {
			r.resources = append(r.resources, *resource.TrafficMirrorFilterId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2TrafficMirrorFilterRule(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeTrafficMirrorFiltersRequest(&ec2.DescribeTrafficMirrorFiltersInput{})
	p := ec2.NewDescribeTrafficMirrorFiltersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TrafficMirrorFilters {
			for _, resource := range resource.EgressFilterRules {
				r.resources = append(r.resources, *resource.TrafficMirrorFilterRuleId)
			}
			for _, resource := range resource.IngressFilterRules {
				r.resources = append(r.resources, *resource.TrafficMirrorFilterRuleId)
			}
		}
	}
	r.err = p.Err()
	return
}

func getEc2TrafficMirrorSession(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeTrafficMirrorSessionsRequest(&ec2.DescribeTrafficMirrorSessionsInput{})
	p := ec2.NewDescribeTrafficMirrorSessionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TrafficMirrorSessions {
			r.resources = append(r.resources, *resource.TrafficMirrorSessionId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2TrafficMirrorTarget(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeTrafficMirrorTargetsRequest(&ec2.DescribeTrafficMirrorTargetsInput{})
	p := ec2.NewDescribeTrafficMirrorTargetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TrafficMirrorTargets {
			r.resources = append(r.resources, *resource.TrafficMirrorTargetId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2TransitGateway(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeTransitGatewaysRequest(&ec2.DescribeTransitGatewaysInput{})
	p := ec2.NewDescribeTransitGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TransitGateways {
			r.resources = append(r.resources, *resource.TransitGatewayId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2TransitGatewayAttachment(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeTransitGatewayAttachmentsRequest(&ec2.DescribeTransitGatewayAttachmentsInput{})
	p := ec2.NewDescribeTransitGatewayAttachmentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TransitGatewayAttachments {
			r.resources = append(r.resources, *resource.TransitGatewayAttachmentId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2TransitGatewayRouteTable(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeTransitGatewayRouteTablesRequest(&ec2.DescribeTransitGatewayRouteTablesInput{})
	p := ec2.NewDescribeTransitGatewayRouteTablesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TransitGatewayRouteTables {
			r.resources = append(r.resources, *resource.TransitGatewayRouteTableId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2Volume(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
	p := ec2.NewDescribeVolumesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Volumes {
			r.resources = append(r.resources, *resource.VolumeId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2VPC(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})
	p := ec2.NewDescribeVpcsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Vpcs {
			r.resources = append(r.resources, *resource.VpcId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2VPCCidrBlock(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})
	p := ec2.NewDescribeVpcsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Vpcs {
			for _, resource := range resource.CidrBlockAssociationSet {
				r.resources = append(r.resources, *resource.AssociationId)
			}
		}
	}
	r.err = p.Err()
	return
}

func getEc2VPCEndpoint(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeVpcEndpointsRequest(&ec2.DescribeVpcEndpointsInput{})
	p := ec2.NewDescribeVpcEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.VpcEndpoints {
			r.resources = append(r.resources, *resource.VpcEndpointId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2VPCEndpointConnectionNotification(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeVpcEndpointConnectionNotificationsRequest(&ec2.DescribeVpcEndpointConnectionNotificationsInput{})
	p := ec2.NewDescribeVpcEndpointConnectionNotificationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ConnectionNotificationSet {
			r.resources = append(r.resources, *resource.ConnectionNotificationId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2VPCEndpointService(client *ec2.Client) (r resourceSliceError) {
	page, err := client.DescribeVpcEndpointServicesRequest(&ec2.DescribeVpcEndpointServicesInput{}).Send(context.Background())
	for _, resource := range page.ServiceDetails {
		r.resources = append(r.resources, *resource.ServiceId)
	}
	r.err = err
	return
}

func getEc2VPCPeeringConnection(client *ec2.Client) (r resourceSliceError) {
	req := client.DescribeVpcPeeringConnectionsRequest(&ec2.DescribeVpcPeeringConnectionsInput{})
	p := ec2.NewDescribeVpcPeeringConnectionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.VpcPeeringConnections {
			r.resources = append(r.resources, *resource.VpcPeeringConnectionId)
		}
	}
	r.err = p.Err()
	return
}

func getEc2VPNConnection(client *ec2.Client) (r resourceSliceError) {
	page, err := client.DescribeVpnConnectionsRequest(&ec2.DescribeVpnConnectionsInput{}).Send(context.Background())
	for _, resource := range page.VpnConnections {
		if resource.ConnectionId == nil {
			continue
		}
		r.resources = append(r.resources, *resource.ConnectionId)
	}
	r.err = err
	return
}

func getEc2VPNGateway(client *ec2.Client) (r resourceSliceError) {
	page, err := client.DescribeVpnGatewaysRequest(&ec2.DescribeVpnGatewaysInput{}).Send(context.Background())
	for _, resource := range page.VpnGateways {
		r.resources = append(r.resources, *resource.VpnGatewayId)
	}
	r.err = err
	return
}

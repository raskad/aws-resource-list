package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getEc2(session *session.Session) (resources resourceMap) {
	client := ec2.New(session)
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

func getEc2CapacityReservation(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeCapacityReservationsPages(&ec2.DescribeCapacityReservationsInput{}, func(page *ec2.DescribeCapacityReservationsOutput, lastPage bool) bool {
		for _, resource := range page.CapacityReservations {
			r.resources = append(r.resources, *resource.CapacityReservationId)
		}
		return true
	})
	return
}

func getEc2ClientVpnEndpoint(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeClientVpnEndpointsPages(&ec2.DescribeClientVpnEndpointsInput{}, func(page *ec2.DescribeClientVpnEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.ClientVpnEndpoints {
			r.resources = append(r.resources, *resource.ClientVpnEndpointId)
		}
		return true
	})
	return
}

func getEc2CustomerGateway(client *ec2.EC2) (r resourceSliceError) {
	output, err := client.DescribeCustomerGateways(&ec2.DescribeCustomerGatewaysInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.CustomerGateways {
		r.resources = append(r.resources, *resource.CustomerGatewayId)
	}
	return
}

func getEc2DHCPOptions(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeDhcpOptionsPages(&ec2.DescribeDhcpOptionsInput{}, func(page *ec2.DescribeDhcpOptionsOutput, lastPage bool) bool {
		for _, resource := range page.DhcpOptions {
			r.resources = append(r.resources, *resource.DhcpOptionsId)
		}
		return true
	})
	return
}

func getEc2Fleet(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeFleetsPages(&ec2.DescribeFleetsInput{}, func(page *ec2.DescribeFleetsOutput, lastPage bool) bool {
		for _, resource := range page.Fleets {
			r.resources = append(r.resources, *resource.FleetId)
		}
		return true
	})
	return
}

func getEc2EgressOnlyInternetGateway(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeEgressOnlyInternetGatewaysPages(&ec2.DescribeEgressOnlyInternetGatewaysInput{}, func(page *ec2.DescribeEgressOnlyInternetGatewaysOutput, lastPage bool) bool {
		for _, resource := range page.EgressOnlyInternetGateways {
			r.resources = append(r.resources, *resource.EgressOnlyInternetGatewayId)
		}
		return true
	})
	return
}

func getEc2Eip(client *ec2.EC2) (r resourceSliceError) {
	output, err := client.DescribeAddresses(&ec2.DescribeAddressesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Addresses {
		r.resources = append(r.resources, *resource.AllocationId)
	}
	return
}

func getEc2EipAssociation(client *ec2.EC2) (r resourceSliceError) {
	output, err := client.DescribeAddresses(&ec2.DescribeAddressesInput{})
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

func getEc2FlowLog(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeFlowLogsPages(&ec2.DescribeFlowLogsInput{}, func(page *ec2.DescribeFlowLogsOutput, lastPage bool) bool {
		for _, resource := range page.FlowLogs {
			r.resources = append(r.resources, *resource.FlowLogId)
		}
		return true
	})
	return
}

func getEc2Host(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeHostsPages(&ec2.DescribeHostsInput{}, func(page *ec2.DescribeHostsOutput, lastPage bool) bool {
		for _, resource := range page.Hosts {
			r.resources = append(r.resources, *resource.HostId)
		}
		return true
	})
	return
}

func getEc2Instace(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeInstancesPages(&ec2.DescribeInstancesInput{}, func(page *ec2.DescribeInstancesOutput, lastPage bool) bool {
		for _, reservation := range page.Reservations {
			for _, resource := range reservation.Instances {
				r.resources = append(r.resources, *resource.InstanceId)
			}
		}
		return true
	})
	return
}

func getEc2InternetGateway(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeInternetGatewaysPages(&ec2.DescribeInternetGatewaysInput{}, func(page *ec2.DescribeInternetGatewaysOutput, lastPage bool) bool {
		for _, resource := range page.InternetGateways {
			r.resources = append(r.resources, *resource.InternetGatewayId)
		}
		return true
	})
	return
}

func getEc2LaunchTemplate(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeLaunchTemplatesPages(&ec2.DescribeLaunchTemplatesInput{}, func(page *ec2.DescribeLaunchTemplatesOutput, lastPage bool) bool {
		for _, resource := range page.LaunchTemplates {
			r.resources = append(r.resources, *resource.LaunchTemplateId)
		}
		return true
	})
	return
}

func getEc2NatGateway(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeNatGatewaysPages(&ec2.DescribeNatGatewaysInput{}, func(page *ec2.DescribeNatGatewaysOutput, lastPage bool) bool {
		for _, resource := range page.NatGateways {
			r.resources = append(r.resources, *resource.NatGatewayId)
		}
		return true
	})
	return
}

func getEc2NetworkACL(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeNetworkAclsPages(&ec2.DescribeNetworkAclsInput{}, func(page *ec2.DescribeNetworkAclsOutput, lastPage bool) bool {
		for _, resource := range page.NetworkAcls {
			r.resources = append(r.resources, *resource.NetworkAclId)
		}
		return true
	})
	return
}

func getEc2NetworkACLSubnetAssociation(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeNetworkAclsPages(&ec2.DescribeNetworkAclsInput{}, func(page *ec2.DescribeNetworkAclsOutput, lastPage bool) bool {
		for _, resource := range page.NetworkAcls {
			for _, resource := range resource.Associations {
				r.resources = append(r.resources, *resource.NetworkAclAssociationId)
			}
		}
		return true
	})
	return
}

func getEc2NetworkInterface(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeNetworkInterfacesPages(&ec2.DescribeNetworkInterfacesInput{}, func(page *ec2.DescribeNetworkInterfacesOutput, lastPage bool) bool {
		for _, resource := range page.NetworkInterfaces {
			r.resources = append(r.resources, *resource.NetworkInterfaceId)
		}
		return true
	})
	return
}

func getEc2NetworkInterfaceAttachment(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeNetworkInterfacesPages(&ec2.DescribeNetworkInterfacesInput{}, func(page *ec2.DescribeNetworkInterfacesOutput, lastPage bool) bool {
		for _, resource := range page.NetworkInterfaces {
			r.resources = append(r.resources, *resource.Attachment.AttachmentId)
		}
		return true
	})
	return
}

func getEc2NetworkInterfacePermission(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeNetworkInterfacePermissionsPages(&ec2.DescribeNetworkInterfacePermissionsInput{}, func(page *ec2.DescribeNetworkInterfacePermissionsOutput, lastPage bool) bool {
		for _, resource := range page.NetworkInterfacePermissions {
			r.resources = append(r.resources, *resource.NetworkInterfacePermissionId)
		}
		return true
	})
	return
}

func getEc2PlacementGroup(client *ec2.EC2) (r resourceSliceError) {
	page, err := client.DescribePlacementGroups(&ec2.DescribePlacementGroupsInput{})
	for _, resource := range page.PlacementGroups {
		r.resources = append(r.resources, *resource.GroupId)
	}
	r.err = err
	return
}

func getEc2RouteTable(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeRouteTablesPages(&ec2.DescribeRouteTablesInput{}, func(page *ec2.DescribeRouteTablesOutput, lastPage bool) bool {
		for _, resource := range page.RouteTables {
			r.resources = append(r.resources, *resource.RouteTableId)
		}
		return true
	})
	return
}

func getEc2RouteTableSubnetAssociation(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeRouteTablesPages(&ec2.DescribeRouteTablesInput{}, func(page *ec2.DescribeRouteTablesOutput, lastPage bool) bool {
		for _, resource := range page.RouteTables {
			for _, resource := range resource.Associations {
				r.resources = append(r.resources, *resource.RouteTableAssociationId)
			}
		}
		return true
	})
	return
}

func getEc2SecurityGroup(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeSecurityGroupsPages(&ec2.DescribeSecurityGroupsInput{}, func(page *ec2.DescribeSecurityGroupsOutput, lastPage bool) bool {
		for _, resource := range page.SecurityGroups {
			r.resources = append(r.resources, *resource.GroupId)
		}
		return true
	})
	return
}

func getEc2SpotFleet(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeSpotFleetRequestsPages(&ec2.DescribeSpotFleetRequestsInput{}, func(page *ec2.DescribeSpotFleetRequestsOutput, lastPage bool) bool {
		for _, resource := range page.SpotFleetRequestConfigs {
			r.resources = append(r.resources, *resource.SpotFleetRequestId)
		}
		return true
	})
	return
}

func getEc2Subnet(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeSubnetsPages(&ec2.DescribeSubnetsInput{}, func(page *ec2.DescribeSubnetsOutput, lastPage bool) bool {
		for _, resource := range page.Subnets {
			r.resources = append(r.resources, *resource.SubnetId)
		}
		return true
	})
	return
}

func getEc2TrafficMirrorFilter(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeTrafficMirrorFiltersPages(&ec2.DescribeTrafficMirrorFiltersInput{}, func(page *ec2.DescribeTrafficMirrorFiltersOutput, lastPage bool) bool {
		for _, resource := range page.TrafficMirrorFilters {
			r.resources = append(r.resources, *resource.TrafficMirrorFilterId)
		}
		return true
	})
	return
}

func getEc2TrafficMirrorFilterRule(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeTrafficMirrorFiltersPages(&ec2.DescribeTrafficMirrorFiltersInput{}, func(page *ec2.DescribeTrafficMirrorFiltersOutput, lastPage bool) bool {
		for _, resource := range page.TrafficMirrorFilters {
			for _, resource := range resource.EgressFilterRules {
				r.resources = append(r.resources, *resource.TrafficMirrorFilterRuleId)
			}
			for _, resource := range resource.IngressFilterRules {
				r.resources = append(r.resources, *resource.TrafficMirrorFilterRuleId)
			}
		}
		return true
	})
	return
}

func getEc2TrafficMirrorSession(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeTrafficMirrorSessionsPages(&ec2.DescribeTrafficMirrorSessionsInput{}, func(page *ec2.DescribeTrafficMirrorSessionsOutput, lastPage bool) bool {
		for _, resource := range page.TrafficMirrorSessions {
			r.resources = append(r.resources, *resource.TrafficMirrorSessionId)
		}
		return true
	})
	return
}

func getEc2TrafficMirrorTarget(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeTrafficMirrorTargetsPages(&ec2.DescribeTrafficMirrorTargetsInput{}, func(page *ec2.DescribeTrafficMirrorTargetsOutput, lastPage bool) bool {
		for _, resource := range page.TrafficMirrorTargets {
			r.resources = append(r.resources, *resource.TrafficMirrorTargetId)
		}
		return true
	})
	return
}

func getEc2TransitGateway(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeTransitGatewaysPages(&ec2.DescribeTransitGatewaysInput{}, func(page *ec2.DescribeTransitGatewaysOutput, lastPage bool) bool {
		for _, resource := range page.TransitGateways {
			r.resources = append(r.resources, *resource.TransitGatewayId)
		}
		return true
	})
	return
}

func getEc2TransitGatewayAttachment(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeTransitGatewayAttachmentsPages(&ec2.DescribeTransitGatewayAttachmentsInput{}, func(page *ec2.DescribeTransitGatewayAttachmentsOutput, lastPage bool) bool {
		for _, resource := range page.TransitGatewayAttachments {
			r.resources = append(r.resources, *resource.TransitGatewayAttachmentId)
		}
		return true
	})
	return
}

func getEc2TransitGatewayRouteTable(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeTransitGatewayRouteTablesPages(&ec2.DescribeTransitGatewayRouteTablesInput{}, func(page *ec2.DescribeTransitGatewayRouteTablesOutput, lastPage bool) bool {
		for _, resource := range page.TransitGatewayRouteTables {
			r.resources = append(r.resources, *resource.TransitGatewayRouteTableId)
		}
		return true
	})
	return
}

func getEc2Volume(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeVolumesPages(&ec2.DescribeVolumesInput{}, func(page *ec2.DescribeVolumesOutput, lastPage bool) bool {
		for _, resource := range page.Volumes {
			r.resources = append(r.resources, *resource.VolumeId)
		}
		return true
	})
	return
}

func getEc2VPC(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeVpcsPages(&ec2.DescribeVpcsInput{}, func(page *ec2.DescribeVpcsOutput, lastPage bool) bool {
		for _, resource := range page.Vpcs {
			r.resources = append(r.resources, *resource.VpcId)
		}
		return true
	})
	return
}

func getEc2VPCCidrBlock(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeVpcsPages(&ec2.DescribeVpcsInput{}, func(page *ec2.DescribeVpcsOutput, lastPage bool) bool {
		for _, resource := range page.Vpcs {
			for _, resource := range resource.CidrBlockAssociationSet {
				r.resources = append(r.resources, *resource.AssociationId)
			}
		}
		return true
	})
	return
}

func getEc2VPCEndpoint(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeVpcEndpointsPages(&ec2.DescribeVpcEndpointsInput{}, func(page *ec2.DescribeVpcEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.VpcEndpoints {
			r.resources = append(r.resources, *resource.VpcEndpointId)
		}
		return true
	})
	return
}

func getEc2VPCEndpointConnectionNotification(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeVpcEndpointConnectionNotificationsPages(&ec2.DescribeVpcEndpointConnectionNotificationsInput{}, func(page *ec2.DescribeVpcEndpointConnectionNotificationsOutput, lastPage bool) bool {
		for _, resource := range page.ConnectionNotificationSet {
			r.resources = append(r.resources, *resource.ConnectionNotificationId)
		}
		return true
	})
	return
}

func getEc2VPCEndpointService(client *ec2.EC2) (r resourceSliceError) {
	page, err := client.DescribeVpcEndpointServices(&ec2.DescribeVpcEndpointServicesInput{})
	for _, resource := range page.ServiceDetails {
		r.resources = append(r.resources, *resource.ServiceId)
	}
	r.err = err
	return
}

func getEc2VPCPeeringConnection(client *ec2.EC2) (r resourceSliceError) {
	r.err = client.DescribeVpcPeeringConnectionsPages(&ec2.DescribeVpcPeeringConnectionsInput{}, func(page *ec2.DescribeVpcPeeringConnectionsOutput, lastPage bool) bool {
		for _, resource := range page.VpcPeeringConnections {
			r.resources = append(r.resources, *resource.VpcPeeringConnectionId)
		}
		return true
	})
	return
}

func getEc2VPNConnection(client *ec2.EC2) (r resourceSliceError) {
	page, err := client.DescribeVpnConnections(&ec2.DescribeVpnConnectionsInput{})
	for _, resource := range page.VpnConnections {
		r.resources = append(r.resources, *resource.VpnConnectionId)
	}
	r.err = err
	return
}

func getEc2VPNGateway(client *ec2.EC2) (r resourceSliceError) {
	page, err := client.DescribeVpnGateways(&ec2.DescribeVpnGatewaysInput{})
	for _, resource := range page.VpnGateways {
		r.resources = append(r.resources, *resource.VpnGatewayId)
	}
	r.err = err
	return
}

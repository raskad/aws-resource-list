package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getEc2(session *session.Session) (resources resourceMap) {
	client := ec2.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		ec2CapacityReservation:               getEc2CapacityReservation(client),
		ec2ClientVpnEndpoint:                 getEc2ClientVpnEndpoint(client),
		ec2CustomerGateway:                   getEc2CustomerGateway(client),
		ec2DHCPOptions:                       getEc2DHCPOptions(client),
		ec2EC2Fleet:                          getEc2Fleet(client),
		ec2EgressOnlyInternetGateway:         getEc2EgressOnlyInternetGateway(client),
		ec2EIP:                               getEc2Eip(client),
		ec2EIPAssociation:                    getEc2EipAssociation(client),
		ec2FlowLog:                           getEc2FlowLog(client),
		ec2Host:                              getEc2Host(client),
		ec2Instance:                          getEc2Instace(client),
		ec2InternetGateway:                   getEc2InternetGateway(client),
		ec2LaunchTemplate:                    getEc2LaunchTemplate(client),
		ec2NatGateway:                        getEc2NatGateway(client),
		ec2NetworkACL:                        getEc2NetworkACL(client),
		ec2NetworkACLSubnetAssociation:       getEc2NetworkACLSubnetAssociation(client),
		ec2NetworkInterface:                  getEc2NetworkInterface(client),
		ec2NetworkInterfaceAttachment:        getEc2NetworkInterfaceAttachment(client),
		ec2NetworkInterfacePermission:        getEc2NetworkInterfacePermission(client),
		ec2PlacementGroup:                    getEc2PlacementGroup(client),
		ec2RouteTable:                        getEc2RouteTable(client),
		ec2RouteTableSubnetAssociation:       getEc2RouteTableSubnetAssociation(client),
		ec2SecurityGroup:                     getEc2SecurityGroup(client),
		ec2SpotFleet:                         getEc2SpotFleet(client),
		ec2Subnet:                            getEc2Subnet(client),
		ec2TrafficMirrorFilter:               getEc2TrafficMirrorFilter(client),
		ec2TrafficMirrorFilterRule:           getEc2TrafficMirrorFilterRule(client),
		ec2TrafficMirrorSession:              getEc2TrafficMirrorSession(client),
		ec2TrafficMirrorTarget:               getEc2TrafficMirrorTarget(client),
		ec2TransitGateway:                    getEc2TransitGateway(client),
		ec2TransitGatewayAttachment:          getEc2TransitGatewayAttachment(client),
		ec2TransitGatewayRouteTable:          getEc2TransitGatewayRouteTable(client),
		ec2Volume:                            getEc2Volume(client),
		ec2VPC:                               getEc2VPC(client),
		ec2VPCCidrBlock:                      getEc2VPCCidrBlock(client),
		ec2VPCEndpoint:                       getEc2VPCEndpoint(client),
		ec2VPCEndpointConnectionNotification: getEc2VPCEndpointConnectionNotification(client),
		ec2VPCEndpointService:                getEc2VPCEndpointService(client),
		ec2VPCPeeringConnection:              getEc2VPCPeeringConnection(client),
		ec2VPNConnection:                     getEc2VPNConnection(client),
		ec2VPNGateway:                        getEc2VPNGateway(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getEc2CapacityReservation(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2CapacityReservation resources")
	r.err = client.DescribeCapacityReservationsPages(&ec2.DescribeCapacityReservationsInput{}, func(page *ec2.DescribeCapacityReservationsOutput, lastPage bool) bool {
		for _, resource := range page.CapacityReservations {
			logDebug("Got Ec2CapacityReservation resource with PhysicalResourceId", *resource.CapacityReservationId)
			r.resources = append(r.resources, *resource.CapacityReservationId)
		}
		return true
	})
	return
}

func getEc2ClientVpnEndpoint(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2ClientVpnEndpoint resources")
	r.err = client.DescribeClientVpnEndpointsPages(&ec2.DescribeClientVpnEndpointsInput{}, func(page *ec2.DescribeClientVpnEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.ClientVpnEndpoints {
			logDebug("Got Ec2ClientVpnEndpoint resource with PhysicalResourceId", *resource.ClientVpnEndpointId)
			r.resources = append(r.resources, *resource.ClientVpnEndpointId)
		}
		return true
	})
	return
}

func getEc2CustomerGateway(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2CustomerGateway resources")
	output, err := client.DescribeCustomerGateways(&ec2.DescribeCustomerGatewaysInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.CustomerGateways {
		logDebug("Got Ec2CustomerGateway resource with PhysicalResourceId", *resource.CustomerGatewayId)
		r.resources = append(r.resources, *resource.CustomerGatewayId)
	}
	return
}

func getEc2DHCPOptions(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2DHCPOptions resources")
	r.err = client.DescribeDhcpOptionsPages(&ec2.DescribeDhcpOptionsInput{}, func(page *ec2.DescribeDhcpOptionsOutput, lastPage bool) bool {
		for _, resource := range page.DhcpOptions {
			logDebug("Got Ec2DHCPOptions resource with PhysicalResourceId", *resource.DhcpOptionsId)
			r.resources = append(r.resources, *resource.DhcpOptionsId)
		}
		return true
	})
	return
}

func getEc2Fleet(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2Fleet resources")
	r.err = client.DescribeFleetsPages(&ec2.DescribeFleetsInput{}, func(page *ec2.DescribeFleetsOutput, lastPage bool) bool {
		for _, resource := range page.Fleets {
			logDebug("Got Ec2Fleet resource with PhysicalResourceId", *resource.FleetId)
			r.resources = append(r.resources, *resource.FleetId)
		}
		return true
	})
	return
}

func getEc2EgressOnlyInternetGateway(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2EgressOnlyInternetGateway resources")
	r.err = client.DescribeEgressOnlyInternetGatewaysPages(&ec2.DescribeEgressOnlyInternetGatewaysInput{}, func(page *ec2.DescribeEgressOnlyInternetGatewaysOutput, lastPage bool) bool {
		for _, resource := range page.EgressOnlyInternetGateways {
			logDebug("Got Ec2EgressOnlyInternetGateway resource with PhysicalResourceId", *resource.EgressOnlyInternetGatewayId)
			r.resources = append(r.resources, *resource.EgressOnlyInternetGatewayId)
		}
		return true
	})
	return
}

func getEc2Eip(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2Eip resources")
	output, err := client.DescribeAddresses(&ec2.DescribeAddressesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Addresses {
		logDebug("Got Ec2Eip resource with PhysicalResourceId", *resource.AllocationId)
		r.resources = append(r.resources, *resource.AllocationId)
	}
	return
}

func getEc2EipAssociation(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2EipAssociation resources")
	output, err := client.DescribeAddresses(&ec2.DescribeAddressesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Addresses {
		if resource.AssociationId != nil {
			logDebug("Got Ec2EipAssociation resource with PhysicalResourceId", *resource.AssociationId)
			r.resources = append(r.resources, *resource.AssociationId)
		}
	}
	return
}

func getEc2FlowLog(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2FlowLog resources")
	r.err = client.DescribeFlowLogsPages(&ec2.DescribeFlowLogsInput{}, func(page *ec2.DescribeFlowLogsOutput, lastPage bool) bool {
		for _, resource := range page.FlowLogs {
			logDebug("Got Ec2FlowLog resource with PhysicalResourceId", *resource.FlowLogId)
			r.resources = append(r.resources, *resource.FlowLogId)
		}
		return true
	})
	return
}

func getEc2Host(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2Host resources")
	r.err = client.DescribeHostsPages(&ec2.DescribeHostsInput{}, func(page *ec2.DescribeHostsOutput, lastPage bool) bool {
		for _, resource := range page.Hosts {
			logDebug("Got Ec2Host resource with PhysicalResourceId", *resource.HostId)
			r.resources = append(r.resources, *resource.HostId)
		}
		return true
	})
	return
}

func getEc2Instace(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2Instace resources")
	r.err = client.DescribeInstancesPages(&ec2.DescribeInstancesInput{}, func(page *ec2.DescribeInstancesOutput, lastPage bool) bool {
		for _, reservation := range page.Reservations {
			for _, resource := range reservation.Instances {
				logDebug("Got Ec2Instace resource with PhysicalResourceId", *resource.InstanceId)
				r.resources = append(r.resources, *resource.InstanceId)
			}
		}
		return true
	})
	return
}

func getEc2InternetGateway(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2InternetGateway resources")
	r.err = client.DescribeInternetGatewaysPages(&ec2.DescribeInternetGatewaysInput{}, func(page *ec2.DescribeInternetGatewaysOutput, lastPage bool) bool {
		for _, resource := range page.InternetGateways {
			logDebug("Got Ec2InternetGateway resource with PhysicalResourceId", *resource.InternetGatewayId)
			r.resources = append(r.resources, *resource.InternetGatewayId)
		}
		return true
	})
	return
}

func getEc2LaunchTemplate(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2LaunchTemplate resources")
	r.err = client.DescribeLaunchTemplatesPages(&ec2.DescribeLaunchTemplatesInput{}, func(page *ec2.DescribeLaunchTemplatesOutput, lastPage bool) bool {
		for _, resource := range page.LaunchTemplates {
			logDebug("Got Ec2LaunchTemplate resource with PhysicalResourceId", *resource.LaunchTemplateId)
			r.resources = append(r.resources, *resource.LaunchTemplateId)
		}
		return true
	})
	return
}

func getEc2NatGateway(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2NatGateway resources")
	r.err = client.DescribeNatGatewaysPages(&ec2.DescribeNatGatewaysInput{}, func(page *ec2.DescribeNatGatewaysOutput, lastPage bool) bool {
		for _, resource := range page.NatGateways {
			logDebug("Got Ec2NatGateway resource with PhysicalResourceId", *resource.NatGatewayId)
			r.resources = append(r.resources, *resource.NatGatewayId)
		}
		return true
	})
	return
}

func getEc2NetworkACL(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2NetworkACL resources")
	r.err = client.DescribeNetworkAclsPages(&ec2.DescribeNetworkAclsInput{}, func(page *ec2.DescribeNetworkAclsOutput, lastPage bool) bool {
		for _, resource := range page.NetworkAcls {
			logDebug("Got Ec2NetworkACL resource with PhysicalResourceId", *resource.NetworkAclId)
			r.resources = append(r.resources, *resource.NetworkAclId)
		}
		return true
	})
	return
}

func getEc2NetworkACLSubnetAssociation(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2NetworkACLSubnetAssociation resources")
	r.err = client.DescribeNetworkAclsPages(&ec2.DescribeNetworkAclsInput{}, func(page *ec2.DescribeNetworkAclsOutput, lastPage bool) bool {
		for _, resource := range page.NetworkAcls {
			for _, resource := range resource.Associations {
				logDebug("Got Ec2NetworkACLSubnetAssociation resource with PhysicalResourceId", *resource.NetworkAclAssociationId)
				r.resources = append(r.resources, *resource.NetworkAclAssociationId)
			}
		}
		return true
	})
	return
}

func getEc2NetworkInterface(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2NetworkInterface resources")
	r.err = client.DescribeNetworkInterfacesPages(&ec2.DescribeNetworkInterfacesInput{}, func(page *ec2.DescribeNetworkInterfacesOutput, lastPage bool) bool {
		for _, resource := range page.NetworkInterfaces {
			logDebug("Got Ec2NetworkInterface resource with PhysicalResourceId", *resource.NetworkInterfaceId)
			r.resources = append(r.resources, *resource.NetworkInterfaceId)
		}
		return true
	})
	return
}

func getEc2NetworkInterfaceAttachment(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2NetworkInterfaceAttachment resources")
	r.err = client.DescribeNetworkInterfacesPages(&ec2.DescribeNetworkInterfacesInput{}, func(page *ec2.DescribeNetworkInterfacesOutput, lastPage bool) bool {
		for _, resource := range page.NetworkInterfaces {
			logDebug("Got Ec2NetworkInterfaceAttachment resource with PhysicalResourceId", *resource.Attachment.AttachmentId)
			r.resources = append(r.resources, *resource.Attachment.AttachmentId)
		}
		return true
	})
	return
}

func getEc2NetworkInterfacePermission(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2NetworkInterfacePermission resources")
	r.err = client.DescribeNetworkInterfacePermissionsPages(&ec2.DescribeNetworkInterfacePermissionsInput{}, func(page *ec2.DescribeNetworkInterfacePermissionsOutput, lastPage bool) bool {
		for _, resource := range page.NetworkInterfacePermissions {
			logDebug("Got Ec2NetworkInterfacePermission resource with PhysicalResourceId", *resource.NetworkInterfacePermissionId)
			r.resources = append(r.resources, *resource.NetworkInterfacePermissionId)
		}
		return true
	})
	return
}

func getEc2PlacementGroup(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2PlacementGroup resources")
	page, err := client.DescribePlacementGroups(&ec2.DescribePlacementGroupsInput{})
	for _, resource := range page.PlacementGroups {
		logDebug("Got Ec2PlacementGroup resource with PhysicalResourceId", *resource.GroupId)
		r.resources = append(r.resources, *resource.GroupId)
	}
	r.err = err
	return
}

func getEc2RouteTable(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2RouteTable resources")
	r.err = client.DescribeRouteTablesPages(&ec2.DescribeRouteTablesInput{}, func(page *ec2.DescribeRouteTablesOutput, lastPage bool) bool {
		for _, resource := range page.RouteTables {
			logDebug("Got Ec2RouteTable resource with PhysicalResourceId", *resource.RouteTableId)
			r.resources = append(r.resources, *resource.RouteTableId)
		}
		return true
	})
	return
}

func getEc2RouteTableSubnetAssociation(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2RouteTableSubnetAssociation resources")
	r.err = client.DescribeRouteTablesPages(&ec2.DescribeRouteTablesInput{}, func(page *ec2.DescribeRouteTablesOutput, lastPage bool) bool {
		for _, resource := range page.RouteTables {
			for _, resource := range resource.Associations {
				logDebug("Got Ec2RouteTableSubnetAssociation resource with PhysicalResourceId", *resource.RouteTableAssociationId)
				r.resources = append(r.resources, *resource.RouteTableAssociationId)
			}
		}
		return true
	})
	return
}

func getEc2SecurityGroup(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2SecurityGroup resources")
	r.err = client.DescribeSecurityGroupsPages(&ec2.DescribeSecurityGroupsInput{}, func(page *ec2.DescribeSecurityGroupsOutput, lastPage bool) bool {
		for _, resource := range page.SecurityGroups {
			logDebug("Got Ec2SecurityGroup resource with PhysicalResourceId", *resource.GroupId)
			r.resources = append(r.resources, *resource.GroupId)
		}
		return true
	})
	return
}

func getEc2SpotFleet(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2SpotFleetRequest resources")
	r.err = client.DescribeSpotFleetRequestsPages(&ec2.DescribeSpotFleetRequestsInput{}, func(page *ec2.DescribeSpotFleetRequestsOutput, lastPage bool) bool {
		for _, resource := range page.SpotFleetRequestConfigs {
			logDebug("Got Ec2SpotFleetRequest resource with PhysicalResourceId", *resource.SpotFleetRequestId)
			r.resources = append(r.resources, *resource.SpotFleetRequestId)
		}
		return true
	})
	return
}

func getEc2Subnet(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2Subnet resources")
	r.err = client.DescribeSubnetsPages(&ec2.DescribeSubnetsInput{}, func(page *ec2.DescribeSubnetsOutput, lastPage bool) bool {
		for _, resource := range page.Subnets {
			logDebug("Got Ec2Subnet resource with PhysicalResourceId", *resource.SubnetId)
			r.resources = append(r.resources, *resource.SubnetId)
		}
		return true
	})
	return
}

func getEc2TrafficMirrorFilter(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2TrafficMirrorFilter resources")
	r.err = client.DescribeTrafficMirrorFiltersPages(&ec2.DescribeTrafficMirrorFiltersInput{}, func(page *ec2.DescribeTrafficMirrorFiltersOutput, lastPage bool) bool {
		for _, resource := range page.TrafficMirrorFilters {
			logDebug("Got Ec2TrafficMirrorFilter resource with PhysicalResourceId", *resource.TrafficMirrorFilterId)
			r.resources = append(r.resources, *resource.TrafficMirrorFilterId)
		}
		return true
	})
	return
}

func getEc2TrafficMirrorFilterRule(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2TrafficMirrorFilterRule resources")
	r.err = client.DescribeTrafficMirrorFiltersPages(&ec2.DescribeTrafficMirrorFiltersInput{}, func(page *ec2.DescribeTrafficMirrorFiltersOutput, lastPage bool) bool {
		for _, resource := range page.TrafficMirrorFilters {
			for _, resource := range resource.EgressFilterRules {
				logDebug("Got Ec2TrafficMirrorFilterRule resource with PhysicalResourceId", *resource.TrafficMirrorFilterRuleId)
				r.resources = append(r.resources, *resource.TrafficMirrorFilterRuleId)
			}
			for _, resource := range resource.IngressFilterRules {
				logDebug("Got Ec2TrafficMirrorFilterRule resource with PhysicalResourceId", *resource.TrafficMirrorFilterRuleId)
				r.resources = append(r.resources, *resource.TrafficMirrorFilterRuleId)
			}
		}
		return true
	})
	return
}

func getEc2TrafficMirrorSession(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2TrafficMirrorSession resources")
	r.err = client.DescribeTrafficMirrorSessionsPages(&ec2.DescribeTrafficMirrorSessionsInput{}, func(page *ec2.DescribeTrafficMirrorSessionsOutput, lastPage bool) bool {
		for _, resource := range page.TrafficMirrorSessions {
			logDebug("Got Ec2TrafficMirrorSession resource with PhysicalResourceId", *resource.TrafficMirrorSessionId)
			r.resources = append(r.resources, *resource.TrafficMirrorSessionId)
		}
		return true
	})
	return
}

func getEc2TrafficMirrorTarget(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2TrafficMirrorTarget resources")
	r.err = client.DescribeTrafficMirrorTargetsPages(&ec2.DescribeTrafficMirrorTargetsInput{}, func(page *ec2.DescribeTrafficMirrorTargetsOutput, lastPage bool) bool {
		for _, resource := range page.TrafficMirrorTargets {
			logDebug("Got Ec2TrafficMirrorTarget resource with PhysicalResourceId", *resource.TrafficMirrorTargetId)
			r.resources = append(r.resources, *resource.TrafficMirrorTargetId)
		}
		return true
	})
	return
}

func getEc2TransitGateway(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2TransitGateway resources")
	r.err = client.DescribeTransitGatewaysPages(&ec2.DescribeTransitGatewaysInput{}, func(page *ec2.DescribeTransitGatewaysOutput, lastPage bool) bool {
		for _, resource := range page.TransitGateways {
			logDebug("Got Ec2TransitGateway resource with PhysicalResourceId", *resource.TransitGatewayId)
			r.resources = append(r.resources, *resource.TransitGatewayId)
		}
		return true
	})
	return
}

func getEc2TransitGatewayAttachment(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2TransitGatewayAttachment resources")
	r.err = client.DescribeTransitGatewayAttachmentsPages(&ec2.DescribeTransitGatewayAttachmentsInput{}, func(page *ec2.DescribeTransitGatewayAttachmentsOutput, lastPage bool) bool {
		for _, resource := range page.TransitGatewayAttachments {
			logDebug("Got Ec2TransitGatewayAttachment resource with PhysicalResourceId", *resource.TransitGatewayAttachmentId)
			r.resources = append(r.resources, *resource.TransitGatewayAttachmentId)
		}
		return true
	})
	return
}

func getEc2TransitGatewayRouteTable(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2TransitGatewayRouteTable resources")
	r.err = client.DescribeTransitGatewayRouteTablesPages(&ec2.DescribeTransitGatewayRouteTablesInput{}, func(page *ec2.DescribeTransitGatewayRouteTablesOutput, lastPage bool) bool {
		for _, resource := range page.TransitGatewayRouteTables {
			logDebug("Got Ec2TransitGatewayRouteTable resource with PhysicalResourceId", *resource.TransitGatewayRouteTableId)
			r.resources = append(r.resources, *resource.TransitGatewayRouteTableId)
		}
		return true
	})
	return
}

func getEc2Volume(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2Volume resources")
	r.err = client.DescribeVolumesPages(&ec2.DescribeVolumesInput{}, func(page *ec2.DescribeVolumesOutput, lastPage bool) bool {
		for _, resource := range page.Volumes {
			logDebug("Got Ec2Volume resource with PhysicalResourceId", *resource.VolumeId)
			r.resources = append(r.resources, *resource.VolumeId)
		}
		return true
	})
	return
}

func getEc2VPC(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2VPC resources")
	r.err = client.DescribeVpcsPages(&ec2.DescribeVpcsInput{}, func(page *ec2.DescribeVpcsOutput, lastPage bool) bool {
		for _, resource := range page.Vpcs {
			logDebug("Got Ec2VPC resource with PhysicalResourceId", *resource.VpcId)
			r.resources = append(r.resources, *resource.VpcId)
		}
		return true
	})
	return
}

func getEc2VPCCidrBlock(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2VPCCidrBlock resources")
	r.err = client.DescribeVpcsPages(&ec2.DescribeVpcsInput{}, func(page *ec2.DescribeVpcsOutput, lastPage bool) bool {
		for _, resource := range page.Vpcs {
			for _, resource := range resource.CidrBlockAssociationSet {
				logDebug("Got Ec2VPCCidrBlock resource with PhysicalResourceId", *resource.AssociationId)
				r.resources = append(r.resources, *resource.AssociationId)
			}
		}
		return true
	})
	return
}

func getEc2VPCEndpoint(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2VPCEndpoint resources")
	r.err = client.DescribeVpcEndpointsPages(&ec2.DescribeVpcEndpointsInput{}, func(page *ec2.DescribeVpcEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.VpcEndpoints {
			logDebug("Got Ec2VPCEndpoint resource with PhysicalResourceId", *resource.VpcEndpointId)
			r.resources = append(r.resources, *resource.VpcEndpointId)
		}
		return true
	})
	return
}

func getEc2VPCEndpointConnectionNotification(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2VPCEndpointConnectionNotification resources")
	r.err = client.DescribeVpcEndpointConnectionNotificationsPages(&ec2.DescribeVpcEndpointConnectionNotificationsInput{}, func(page *ec2.DescribeVpcEndpointConnectionNotificationsOutput, lastPage bool) bool {
		for _, resource := range page.ConnectionNotificationSet {
			logDebug("Got Ec2VPCEndpointConnectionNotification resource with PhysicalResourceId", *resource.ConnectionNotificationId)
			r.resources = append(r.resources, *resource.ConnectionNotificationId)
		}
		return true
	})
	return
}

func getEc2VPCEndpointService(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2VPCEndpointService resources")
	page, err := client.DescribeVpcEndpointServices(&ec2.DescribeVpcEndpointServicesInput{})
	for _, resource := range page.ServiceDetails {
		logDebug("Got Ec2VPCEndpointService resource with PhysicalResourceId", *resource.ServiceId)
		r.resources = append(r.resources, *resource.ServiceId)
	}
	r.err = err
	return
}

func getEc2VPCPeeringConnection(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2VPCPeeringConnection resources")
	r.err = client.DescribeVpcPeeringConnectionsPages(&ec2.DescribeVpcPeeringConnectionsInput{}, func(page *ec2.DescribeVpcPeeringConnectionsOutput, lastPage bool) bool {
		for _, resource := range page.VpcPeeringConnections {
			logDebug("Got Ec2VPCPeeringConnection resource with PhysicalResourceId", *resource.VpcPeeringConnectionId)
			r.resources = append(r.resources, *resource.VpcPeeringConnectionId)
		}
		return true
	})
	return
}

func getEc2VPNConnection(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2VPNConnection resources")
	page, err := client.DescribeVpnConnections(&ec2.DescribeVpnConnectionsInput{})
	for _, resource := range page.VpnConnections {
		logDebug("Got Ec2VPNConnection resource with PhysicalResourceId", *resource.VpnConnectionId)
		r.resources = append(r.resources, *resource.VpnConnectionId)
	}
	r.err = err
	return
}

func getEc2VPNGateway(client *ec2.EC2) (r resourceSliceError) {
	logDebug("Listing Ec2VPNConnection resources")
	page, err := client.DescribeVpnGateways(&ec2.DescribeVpnGatewaysInput{})
	for _, resource := range page.VpnGateways {
		logDebug("Got Ec2VPNConnection resource with PhysicalResourceId", *resource.VpnGatewayId)
		r.resources = append(r.resources, *resource.VpnGatewayId)
	}
	r.err = err
	return
}

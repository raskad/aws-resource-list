package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func getRds(config aws.Config) (resources awsResourceMap) {
	client := rds.New(config)

	rdsDBClusterIDs := getRdsDBClusterIDs(client)
	rdsDBClusterEndpointIDs := getRdsDBClusterEndpointIDs(client)
	rdsDBClusterParameterGroupNames := getRdsDBClusterParameterGroupNames(client)
	rdsDBClusterSnapshotIDs := getRdsDBClusterSnapshotIDs(client)
	rdsDBInstanceIDs := getRdsDBInstanceIDs(client)
	rdsDBParameterGroupNames := getRdsDBParameterGroupNames(client)
	rdsDBSecurityGroupNames := getRdsDBSecurityGroupNames(client)
	rdsDBSnapshotIDs := getRdsDBSnapshotIDs(client)
	rdsDBSubnetGroupNames := getRdsDBSubnetGroupNames(client)
	rdsEventSubscriptionIDs := getRdsEventSubscriptionIDs(client)
	rdsGlobalClusterIDs := getRdsGlobalClusterIDs(client)
	rdsOptionGroupNames := getRdsOptionGroupNames(client)

	resources = awsResourceMap{
		rdsDBCluster:               rdsDBClusterIDs,
		rdsDBClusterEndpoint:       rdsDBClusterEndpointIDs,
		rdsDBClusterParameterGroup: rdsDBClusterParameterGroupNames,
		rdsDBClusterSnapshot:       rdsDBClusterSnapshotIDs,
		rdsDBInstance:              rdsDBInstanceIDs,
		rdsDBParameterGroup:        rdsDBParameterGroupNames,
		rdsDBSecurityGroup:         rdsDBSecurityGroupNames,
		rdsDBSnapshot:              rdsDBSnapshotIDs,
		rdsDBSubnetGroup:           rdsDBSubnetGroupNames,
		rdsEventSubscription:       rdsEventSubscriptionIDs,
		rdsGlobalCluster:           rdsGlobalClusterIDs,
		rdsOptionGroup:             rdsOptionGroupNames,
	}
	return
}

func getRdsDBClusterIDs(client *rds.Client) (resources []string) {
	req := client.DescribeDBClustersRequest(&rds.DescribeDBClustersInput{})
	p := rds.NewDescribeDBClustersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DBClusters {
			resources = append(resources, *resource.DBClusterIdentifier)
		}
	}
	return
}

func getRdsDBClusterEndpointIDs(client *rds.Client) (resources []string) {
	input := rds.DescribeDBClusterEndpointsInput{}
	for {
		page, err := client.DescribeDBClusterEndpointsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.DBClusterEndpoints {
			resources = append(resources, *resource.DBClusterEndpointIdentifier)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getRdsDBClusterParameterGroupNames(client *rds.Client) (resources []string) {
	input := rds.DescribeDBClusterParameterGroupsInput{}
	for {
		page, err := client.DescribeDBClusterParameterGroupsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.DBClusterParameterGroups {
			resources = append(resources, *resource.DBClusterParameterGroupName)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getRdsDBClusterSnapshotIDs(client *rds.Client) (resources []string) {
	input := rds.DescribeDBClusterSnapshotsInput{}
	for {
		page, err := client.DescribeDBClusterSnapshotsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.DBClusterSnapshots {
			resources = append(resources, *resource.DBClusterSnapshotIdentifier)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getRdsDBInstanceIDs(client *rds.Client) (resources []string) {
	req := client.DescribeDBInstancesRequest(&rds.DescribeDBInstancesInput{})
	p := rds.NewDescribeDBInstancesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DBInstances {
			resources = append(resources, *resource.DBInstanceIdentifier)
		}
	}
	return
}

func getRdsDBParameterGroupNames(client *rds.Client) (resources []string) {
	req := client.DescribeDBParameterGroupsRequest(&rds.DescribeDBParameterGroupsInput{})
	p := rds.NewDescribeDBParameterGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DBParameterGroups {
			resources = append(resources, *resource.DBParameterGroupName)
		}
	}
	return
}

func getRdsDBSecurityGroupNames(client *rds.Client) (resources []string) {
	req := client.DescribeDBSecurityGroupsRequest(&rds.DescribeDBSecurityGroupsInput{})
	p := rds.NewDescribeDBSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DBSecurityGroups {
			resources = append(resources, *resource.DBSecurityGroupName)
		}
	}
	return
}

func getRdsDBSnapshotIDs(client *rds.Client) (resources []string) {
	req := client.DescribeDBSnapshotsRequest(&rds.DescribeDBSnapshotsInput{})
	p := rds.NewDescribeDBSnapshotsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DBSnapshots {
			resources = append(resources, *resource.DBSnapshotIdentifier)
		}
	}
	return
}

func getRdsDBSubnetGroupNames(client *rds.Client) (resources []string) {
	req := client.DescribeDBSubnetGroupsRequest(&rds.DescribeDBSubnetGroupsInput{})
	p := rds.NewDescribeDBSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DBSubnetGroups {
			resources = append(resources, *resource.DBSubnetGroupName)
		}
	}
	return
}

func getRdsEventSubscriptionIDs(client *rds.Client) (resources []string) {
	req := client.DescribeEventSubscriptionsRequest(&rds.DescribeEventSubscriptionsInput{})
	p := rds.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.EventSubscriptionsList {
			resources = append(resources, *resource.CustSubscriptionId)
		}
	}
	return
}

func getRdsGlobalClusterIDs(client *rds.Client) (resources []string) {
	input := rds.DescribeGlobalClustersInput{}
	for {
		page, err := client.DescribeGlobalClustersRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.GlobalClusters {
			resources = append(resources, *resource.GlobalClusterIdentifier)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getRdsOptionGroupNames(client *rds.Client) (resources []string) {
	req := client.DescribeOptionGroupsRequest(&rds.DescribeOptionGroupsInput{})
	p := rds.NewDescribeOptionGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.OptionGroupsList {
			resources = append(resources, *resource.OptionGroupName)
		}
	}
	return
}

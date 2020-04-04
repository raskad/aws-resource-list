package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
)

func getNeptune(config aws.Config) (resources awsResourceMap) {
	client := neptune.New(config)

	getNeptuneDBClusterIDs := getNeptuneDBClusterIDs(client)
	getNeptuneDBClusterParameterGroupNames := getNeptuneDBClusterParameterGroupNames(client)
	neptuneDBClusterSnapshotIDs := getNeptuneDBClusterSnapshotIDs(client)
	neptuneDBEventSubscriptionIDs := getNeptuneDBEventSubscriptionIDs(client)
	getNeptuneDBInstanceIDs := getNeptuneDBInstanceIDs(client)
	getNeptuneDBParameterGroupNames := getNeptuneDBParameterGroupNames(client)
	getNeptuneDBSubnetGroupNames := getNeptuneDBSubnetGroupNames(client)

	resources = awsResourceMap{
		neptuneDBCluster:               getNeptuneDBClusterIDs,
		neptuneDBClusterParameterGroup: getNeptuneDBClusterParameterGroupNames,
		neptuneDBClusterSnapshot:       neptuneDBClusterSnapshotIDs,
		neptuneDBEventSubscription:     neptuneDBEventSubscriptionIDs,
		neptuneDBInstance:              getNeptuneDBInstanceIDs,
		neptuneDBParameterGroup:        getNeptuneDBParameterGroupNames,
		neptuneDBSubnetGroup:           getNeptuneDBSubnetGroupNames,
	}
	return
}

func getNeptuneDBClusterIDs(client *neptune.Client) (resources []string) {
	input := neptune.DescribeDBClustersInput{}
	for {
		page, err := client.DescribeDBClustersRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.DBClusters {
			resources = append(resources, *resource.DBClusterIdentifier)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getNeptuneDBClusterParameterGroupNames(client *neptune.Client) (resources []string) {
	input := neptune.DescribeDBClusterParameterGroupsInput{}
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

func getNeptuneDBClusterSnapshotIDs(client *neptune.Client) (resources []string) {
	input := neptune.DescribeDBClusterSnapshotsInput{}
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

func getNeptuneDBEventSubscriptionIDs(client *neptune.Client) (resources []string) {
	input := neptune.DescribeEventSubscriptionsInput{}
	for {
		page, err := client.DescribeEventSubscriptionsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.EventSubscriptionsList {
			resources = append(resources, *resource.CustSubscriptionId)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getNeptuneDBInstanceIDs(client *neptune.Client) (resources []string) {
	req := client.DescribeDBInstancesRequest(&neptune.DescribeDBInstancesInput{})
	p := neptune.NewDescribeDBInstancesPaginator(req)
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

func getNeptuneDBParameterGroupNames(client *neptune.Client) (resources []string) {
	req := client.DescribeDBParameterGroupsRequest(&neptune.DescribeDBParameterGroupsInput{})
	p := neptune.NewDescribeDBParameterGroupsPaginator(req)
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

func getNeptuneDBSubnetGroupNames(client *neptune.Client) (resources []string) {
	req := client.DescribeDBSubnetGroupsRequest(&neptune.DescribeDBSubnetGroupsInput{})
	p := neptune.NewDescribeDBSubnetGroupsPaginator(req)
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

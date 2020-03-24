package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func getRds(config aws.Config) (resources resourceMap) {
	client := rds.New(config)

	rdsDBClusterIDs := getRdsDBClusterIDs(client)
	rdsDBClusterParameterGroupNames := getRdsDBClusterParameterGroupNames(client)
	rdsDBInstanceIDs := getRdsDBInstanceIDs(client)
	rdsDBParameterGroupNames := getRdsDBParameterGroupNames(client)
	rdsDBSecurityGroupNames := getRdsDBSecurityGroupNames(client)
	rdsDBSubnetGroupNames := getRdsDBSubnetGroupNames(client)
	rdsEventSubscriptionIDs := getRdsEventSubscriptionIDs(client)
	rdsOptionGroupNames := getRdsOptionGroupNames(client)

	resources = resourceMap{
		rdsDBCluster:               rdsDBClusterIDs,
		rdsDBClusterParameterGroup: rdsDBClusterParameterGroupNames,
		rdsDBInstance:              rdsDBInstanceIDs,
		rdsDBParameterGroup:        rdsDBParameterGroupNames,
		rdsDBSecurityGroup:         rdsDBSecurityGroupNames,
		rdsDBSubnetGroup:           rdsDBSubnetGroupNames,
		rdsEventSubscription:       rdsEventSubscriptionIDs,
		rdsOptionGroup:             rdsOptionGroupNames,
	}
	return
}

func getRdsDBClusterIDs(client *rds.Client) (resources []string) {
	req := client.DescribeDBClustersRequest(&rds.DescribeDBClustersInput{})
	p := rds.NewDescribeDBClustersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.DBClusters {
			resources = append(resources, *resource.DBClusterIdentifier)
		}
	}
	return
}

func getRdsDBClusterParameterGroupNames(client *rds.Client) (resources []string) {
	input := rds.DescribeDBClusterParameterGroupsInput{}
	for {
		page, err := client.DescribeDBClusterParameterGroupsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.DBClusterParameterGroups {
			resources = append(resources, *resource.DBClusterParameterGroupName)
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
		logErr(p.Err())
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
		logErr(p.Err())
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
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.DBSecurityGroups {
			resources = append(resources, *resource.DBSecurityGroupName)
		}
	}
	return
}

func getRdsDBSubnetGroupNames(client *rds.Client) (resources []string) {
	req := client.DescribeDBSubnetGroupsRequest(&rds.DescribeDBSubnetGroupsInput{})
	p := rds.NewDescribeDBSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
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
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.EventSubscriptionsList {
			resources = append(resources, *resource.CustSubscriptionId)
		}
	}
	return
}

func getRdsOptionGroupNames(client *rds.Client) (resources []string) {
	req := client.DescribeOptionGroupsRequest(&rds.DescribeOptionGroupsInput{})
	p := rds.NewDescribeOptionGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.OptionGroupsList {
			resources = append(resources, *resource.OptionGroupName)
		}
	}
	return
}

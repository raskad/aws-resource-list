package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
)

func getDocDB(config aws.Config) (resources resourceMap) {
	client := docdb.New(config)

	docDBDBClusterIDs := getDocDBDBClusterIDs(client)
	docDBDBClusterParameterGroupNames := getDocDBDBClusterParameterGroupNames(client)
	docDBDBInstanceIDs := getDocDBDBInstanceIDs(client)
	docDBDBSubnetGroupNames := getDocDBDBSubnetGroupNames(client)

	resources = resourceMap{
		docDBDBCluster:               docDBDBClusterIDs,
		docDBDBClusterParameterGroup: docDBDBClusterParameterGroupNames,
		docDBDBInstance:              docDBDBInstanceIDs,
		docDBDBSubnetGroup:           docDBDBSubnetGroupNames,
	}
	return
}

func getDocDBDBClusterIDs(client *docdb.Client) (resources []string) {
	req := client.DescribeDBClustersRequest(&docdb.DescribeDBClustersInput{})
	p := docdb.NewDescribeDBClustersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.DBClusters {
			resources = append(resources, *resource.DBClusterIdentifier)
		}
	}
	return
}

func getDocDBDBClusterParameterGroupNames(client *docdb.Client) (resources []string) {
	input := docdb.DescribeDBClusterParameterGroupsInput{}
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

func getDocDBDBInstanceIDs(client *docdb.Client) (resources []string) {
	req := client.DescribeDBInstancesRequest(&docdb.DescribeDBInstancesInput{})
	p := docdb.NewDescribeDBInstancesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.DBInstances {
			resources = append(resources, *resource.DBInstanceIdentifier)
		}
	}
	return
}

func getDocDBDBSubnetGroupNames(client *docdb.Client) (resources []string) {
	req := client.DescribeDBSubnetGroupsRequest(&docdb.DescribeDBSubnetGroupsInput{})
	p := docdb.NewDescribeDBSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.DBSubnetGroups {
			resources = append(resources, *resource.DBSubnetGroupName)
		}
	}
	return
}

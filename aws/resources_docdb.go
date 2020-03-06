package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
)

func getDocDB(config aws.Config) (resources resourceMap) {
	client := docdb.New(config)
	resources = reduce(
		getDocDBDBCluster(client).unwrap(docDBDBCluster),
		getDocDBDBClusterParameterGroup(client).unwrap(docDBDBClusterParameterGroup),
		getDocDBDBInstance(client).unwrap(docDBDBInstance),
		getDocDBDBSubnetGroup(client).unwrap(docDBDBSubnetGroup),
	)
	return
}

func getDocDBDBCluster(client *docdb.Client) (r resourceSliceError) {
	req := client.DescribeDBClustersRequest(&docdb.DescribeDBClustersInput{})
	p := docdb.NewDescribeDBClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBClusters {
			r.resources = append(r.resources, *resource.DBClusterIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getDocDBDBClusterParameterGroup(client *docdb.Client) (r resourceSliceError) {
	input := docdb.DescribeDBClusterParameterGroupsInput{}
	for {
		page, err := client.DescribeDBClusterParameterGroupsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DBClusterParameterGroups {
			r.resources = append(r.resources, *resource.DBClusterParameterGroupName)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getDocDBDBInstance(client *docdb.Client) (r resourceSliceError) {
	req := client.DescribeDBInstancesRequest(&docdb.DescribeDBInstancesInput{})
	p := docdb.NewDescribeDBInstancesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBInstances {
			r.resources = append(r.resources, *resource.DBInstanceIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getDocDBDBSubnetGroup(client *docdb.Client) (r resourceSliceError) {
	req := client.DescribeDBSubnetGroupsRequest(&docdb.DescribeDBSubnetGroupsInput{})
	p := docdb.NewDescribeDBSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBSubnetGroups {
			r.resources = append(r.resources, *resource.DBSubnetGroupName)
		}
	}
	r.err = p.Err()
	return
}

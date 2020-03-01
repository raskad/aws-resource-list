package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/docdb"
)

func getDocDB(session *session.Session) (resources resourceMap) {
	client := docdb.New(session)
	resources = reduce(
		getDocDBDBCluster(client).unwrap(docDBDBCluster),
		getDocDBDBClusterParameterGroup(client).unwrap(docDBDBClusterParameterGroup),
		getDocDBDBInstance(client).unwrap(docDBDBInstance),
		getDocDBDBSubnetGroup(client).unwrap(docDBDBSubnetGroup),
	)
	return
}

func getDocDBDBCluster(client *docdb.DocDB) (r resourceSliceError) {
	r.err = client.DescribeDBClustersPages(&docdb.DescribeDBClustersInput{}, func(page *docdb.DescribeDBClustersOutput, lastPage bool) bool {
		for _, resource := range page.DBClusters {
			r.resources = append(r.resources, *resource.DBClusterIdentifier)
		}
		return true
	})
	return
}

func getDocDBDBClusterParameterGroup(client *docdb.DocDB) (r resourceSliceError) {
	input := docdb.DescribeDBClusterParameterGroupsInput{}
	for {
		page, err := client.DescribeDBClusterParameterGroups(&input)
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

func getDocDBDBInstance(client *docdb.DocDB) (r resourceSliceError) {
	r.err = client.DescribeDBInstancesPages(&docdb.DescribeDBInstancesInput{}, func(page *docdb.DescribeDBInstancesOutput, lastPage bool) bool {
		for _, resource := range page.DBInstances {
			r.resources = append(r.resources, *resource.DBInstanceIdentifier)
		}
		return true
	})
	return
}

func getDocDBDBSubnetGroup(client *docdb.DocDB) (r resourceSliceError) {
	r.err = client.DescribeDBSubnetGroupsPages(&docdb.DescribeDBSubnetGroupsInput{}, func(page *docdb.DescribeDBSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBSubnetGroups {
			r.resources = append(r.resources, *resource.DBSubnetGroupName)
		}
		return true
	})
	return
}

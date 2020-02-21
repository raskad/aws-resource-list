package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/docdb"
)

func getDocDB(session *session.Session) (resources resourceMap) {
	client := docdb.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		docDBDBCluster:               getDocDBDBCluster(client),
		docDBDBClusterParameterGroup: getDocDBDBClusterParameterGroup(client),
		docDBDBInstance:              getDocDBDBInstance(client),
		docDBDBSubnetGroup:           getDocDBDBSubnetGroup(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getDocDBDBCluster(client *docdb.DocDB) (r resourceSliceError) {
	r.err = client.DescribeDBClustersPages(&docdb.DescribeDBClustersInput{}, func(page *docdb.DescribeDBClustersOutput, lastPage bool) bool {
		logDebug("List DocDBDBCluster resources page. Remaining pages", page.Marker)
		for _, resource := range page.DBClusters {
			logDebug("Got DocDBDBCluster resource with PhysicalResourceId", *resource.DBClusterIdentifier)
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
		logDebug("Listing DocDBDBClusterParameterGroup resources page. Remaining pages", page.Marker)
		for _, resource := range page.DBClusterParameterGroups {
			logDebug("Got DocDBDBClusterParameterGroup resource with PhysicalResourceId", *resource.DBClusterParameterGroupName)
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
		logDebug("List DocDBDBInstance resources page. Remaining pages", page.Marker)
		for _, resource := range page.DBInstances {
			logDebug("Got DocDBDBInstance resource with PhysicalResourceId", *resource.DBInstanceIdentifier)
			r.resources = append(r.resources, *resource.DBInstanceIdentifier)
		}
		return true
	})
	return
}

func getDocDBDBSubnetGroup(client *docdb.DocDB) (r resourceSliceError) {
	r.err = client.DescribeDBSubnetGroupsPages(&docdb.DescribeDBSubnetGroupsInput{}, func(page *docdb.DescribeDBSubnetGroupsOutput, lastPage bool) bool {
		logDebug("List DocDBDBSubnetGroup resources page. Remaining pages", page.Marker)
		for _, resource := range page.DBSubnetGroups {
			logDebug("Got DocDBDBSubnetGroup resource with PhysicalResourceId", *resource.DBSubnetGroupName)
			r.resources = append(r.resources, *resource.DBSubnetGroupName)
		}
		return true
	})
	return
}

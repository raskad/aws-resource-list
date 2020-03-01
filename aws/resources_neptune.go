package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/neptune"
)

func getNeptune(session *session.Session) (resources resourceMap) {
	client := neptune.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		neptuneDBCluster:               getNeptuneDBCluster(client),
		neptuneDBClusterParameterGroup: getNeptuneDBClusterParameterGroup(client),
		neptuneDBInstance:              getNeptuneDBInstance(client),
		neptuneDBParameterGroup:        getNeptuneDBParameterGroup(client),
		neptuneDBSubnetGroup:           getNeptuneDBSubnetGroup(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getNeptuneDBCluster(client *neptune.Neptune) (r resourceSliceError) {
	logDebug("Listing NeptuneDBCluster resources")
	input := neptune.DescribeDBClustersInput{}
	for {
		page, err := client.DescribeDBClusters(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DBClusters {
			logDebug("Got NeptuneDBCluster resource with PhysicalResourceId", *resource.DBClusterIdentifier)
			r.resources = append(r.resources, *resource.DBClusterIdentifier)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getNeptuneDBClusterParameterGroup(client *neptune.Neptune) (r resourceSliceError) {
	logDebug("Listing NeptuneDBClusterParameterGroup resources")
	input := neptune.DescribeDBClusterParameterGroupsInput{}
	for {
		page, err := client.DescribeDBClusterParameterGroups(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DBClusterParameterGroups {
			logDebug("Got NeptuneDBClusterParameterGroup resource with PhysicalResourceId", *resource.DBClusterParameterGroupName)
			r.resources = append(r.resources, *resource.DBClusterParameterGroupName)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getNeptuneDBInstance(client *neptune.Neptune) (r resourceSliceError) {
	logDebug("Listing NeptuneDBInstance resources")
	r.err = client.DescribeDBInstancesPages(&neptune.DescribeDBInstancesInput{}, func(page *neptune.DescribeDBInstancesOutput, lastPage bool) bool {
		for _, resource := range page.DBInstances {
			logDebug("Got NeptuneDBInstance resource with PhysicalResourceId", *resource.DBInstanceIdentifier)
			r.resources = append(r.resources, *resource.DBInstanceIdentifier)
		}
		return true
	})
	return
}

func getNeptuneDBParameterGroup(client *neptune.Neptune) (r resourceSliceError) {
	logDebug("Listing NeptuneDBParameterGroup resources")
	r.err = client.DescribeDBParameterGroupsPages(&neptune.DescribeDBParameterGroupsInput{}, func(page *neptune.DescribeDBParameterGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBParameterGroups {
			logDebug("Got NeptuneDBParameterGroup resource with PhysicalResourceId", *resource.DBParameterGroupName)
			r.resources = append(r.resources, *resource.DBParameterGroupName)
		}
		return true
	})
	return
}

func getNeptuneDBSubnetGroup(client *neptune.Neptune) (r resourceSliceError) {
	logDebug("Listing NeptuneDBSubnetGroup resources")
	r.err = client.DescribeDBSubnetGroupsPages(&neptune.DescribeDBSubnetGroupsInput{}, func(page *neptune.DescribeDBSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBSubnetGroups {
			logDebug("Got NeptuneDBSubnetGroup resource with PhysicalResourceId", *resource.DBSubnetGroupName)
			r.resources = append(r.resources, *resource.DBSubnetGroupName)
		}
		return true
	})
	return
}

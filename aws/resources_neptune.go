package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/neptune"
)

func getNeptune(session *session.Session) (resources resourceMap) {
	client := neptune.New(session)
	resources = reduce(
		getNeptuneDBCluster(client).unwrap(neptuneDBCluster),
		getNeptuneDBClusterParameterGroup(client).unwrap(neptuneDBClusterParameterGroup),
		getNeptuneDBInstance(client).unwrap(neptuneDBInstance),
		getNeptuneDBParameterGroup(client).unwrap(neptuneDBParameterGroup),
		getNeptuneDBSubnetGroup(client).unwrap(neptuneDBSubnetGroup),
	)
	return
}

func getNeptuneDBCluster(client *neptune.Neptune) (r resourceSliceError) {
	input := neptune.DescribeDBClustersInput{}
	for {
		page, err := client.DescribeDBClusters(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DBClusters {
			r.resources = append(r.resources, *resource.DBClusterIdentifier)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getNeptuneDBClusterParameterGroup(client *neptune.Neptune) (r resourceSliceError) {
	input := neptune.DescribeDBClusterParameterGroupsInput{}
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

func getNeptuneDBInstance(client *neptune.Neptune) (r resourceSliceError) {
	r.err = client.DescribeDBInstancesPages(&neptune.DescribeDBInstancesInput{}, func(page *neptune.DescribeDBInstancesOutput, lastPage bool) bool {
		for _, resource := range page.DBInstances {
			r.resources = append(r.resources, *resource.DBInstanceIdentifier)
		}
		return true
	})
	return
}

func getNeptuneDBParameterGroup(client *neptune.Neptune) (r resourceSliceError) {
	r.err = client.DescribeDBParameterGroupsPages(&neptune.DescribeDBParameterGroupsInput{}, func(page *neptune.DescribeDBParameterGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBParameterGroups {
			r.resources = append(r.resources, *resource.DBParameterGroupName)
		}
		return true
	})
	return
}

func getNeptuneDBSubnetGroup(client *neptune.Neptune) (r resourceSliceError) {
	r.err = client.DescribeDBSubnetGroupsPages(&neptune.DescribeDBSubnetGroupsInput{}, func(page *neptune.DescribeDBSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBSubnetGroups {
			r.resources = append(r.resources, *resource.DBSubnetGroupName)
		}
		return true
	})
	return
}

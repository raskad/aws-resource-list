package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
)

func getNeptune(config aws.Config) (resources resourceMap) {
	client := neptune.New(config)
	resources = reduce(
		getNeptuneDBCluster(client).unwrap(neptuneDBCluster),
		getNeptuneDBClusterParameterGroup(client).unwrap(neptuneDBClusterParameterGroup),
		getNeptuneDBInstance(client).unwrap(neptuneDBInstance),
		getNeptuneDBParameterGroup(client).unwrap(neptuneDBParameterGroup),
		getNeptuneDBSubnetGroup(client).unwrap(neptuneDBSubnetGroup),
	)
	return
}

func getNeptuneDBCluster(client *neptune.Client) (r resourceSliceError) {
	input := neptune.DescribeDBClustersInput{}
	for {
		page, err := client.DescribeDBClustersRequest(&input).Send(context.Background())
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

func getNeptuneDBClusterParameterGroup(client *neptune.Client) (r resourceSliceError) {
	input := neptune.DescribeDBClusterParameterGroupsInput{}
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

func getNeptuneDBInstance(client *neptune.Client) (r resourceSliceError) {
	req := client.DescribeDBInstancesRequest(&neptune.DescribeDBInstancesInput{})
	p := neptune.NewDescribeDBInstancesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBInstances {
			r.resources = append(r.resources, *resource.DBInstanceIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getNeptuneDBParameterGroup(client *neptune.Client) (r resourceSliceError) {
	req := client.DescribeDBParameterGroupsRequest(&neptune.DescribeDBParameterGroupsInput{})
	p := neptune.NewDescribeDBParameterGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBParameterGroups {
			r.resources = append(r.resources, *resource.DBParameterGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getNeptuneDBSubnetGroup(client *neptune.Client) (r resourceSliceError) {
	req := client.DescribeDBSubnetGroupsRequest(&neptune.DescribeDBSubnetGroupsInput{})
	p := neptune.NewDescribeDBSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBSubnetGroups {
			r.resources = append(r.resources, *resource.DBSubnetGroupName)
		}
	}
	r.err = p.Err()
	return
}

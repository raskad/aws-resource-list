package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func getRds(config aws.Config) (resources resourceMap) {
	client := rds.New(config)
	resources = reduce(
		getRdsDBCluster(client).unwrap(rdsDBCluster),
		getRdsDBClusterParameterGroup(client).unwrap(rdsDBClusterParameterGroup),
		getRdsDBInstance(client).unwrap(rdsDBInstance),
		getRdsDBParameterGroup(client).unwrap(rdsDBParameterGroup),
		getRdsDBSecurityGroup(client).unwrap(rdsDBSecurityGroup),
		getRdsDBSubnetGroup(client).unwrap(rdsDBSubnetGroup),
		getRdsEventSubscription(client).unwrap(rdsEventSubscription),
		getRdsOptionGroup(client).unwrap(rdsOptionGroup),
	)
	return
}

func getRdsDBCluster(client *rds.Client) (r resourceSliceError) {
	req := client.DescribeDBClustersRequest(&rds.DescribeDBClustersInput{})
	p := rds.NewDescribeDBClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBClusters {
			r.resources = append(r.resources, *resource.DBClusterIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getRdsDBClusterParameterGroup(client *rds.Client) (r resourceSliceError) {
	input := rds.DescribeDBClusterParameterGroupsInput{}
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

func getRdsDBInstance(client *rds.Client) (r resourceSliceError) {
	req := client.DescribeDBInstancesRequest(&rds.DescribeDBInstancesInput{})
	p := rds.NewDescribeDBInstancesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBInstances {
			r.resources = append(r.resources, *resource.DBInstanceIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getRdsDBParameterGroup(client *rds.Client) (r resourceSliceError) {
	req := client.DescribeDBParameterGroupsRequest(&rds.DescribeDBParameterGroupsInput{})
	p := rds.NewDescribeDBParameterGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBParameterGroups {
			r.resources = append(r.resources, *resource.DBParameterGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getRdsDBSecurityGroup(client *rds.Client) (r resourceSliceError) {
	req := client.DescribeDBSecurityGroupsRequest(&rds.DescribeDBSecurityGroupsInput{})
	p := rds.NewDescribeDBSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBSecurityGroups {
			r.resources = append(r.resources, *resource.DBSecurityGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getRdsDBSubnetGroup(client *rds.Client) (r resourceSliceError) {
	req := client.DescribeDBSubnetGroupsRequest(&rds.DescribeDBSubnetGroupsInput{})
	p := rds.NewDescribeDBSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DBSubnetGroups {
			r.resources = append(r.resources, *resource.DBSubnetGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getRdsEventSubscription(client *rds.Client) (r resourceSliceError) {
	req := client.DescribeEventSubscriptionsRequest(&rds.DescribeEventSubscriptionsInput{})
	p := rds.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.EventSubscriptionsList {
			r.resources = append(r.resources, *resource.CustSubscriptionId)
		}
	}
	r.err = p.Err()
	return
}

func getRdsOptionGroup(client *rds.Client) (r resourceSliceError) {
	req := client.DescribeOptionGroupsRequest(&rds.DescribeOptionGroupsInput{})
	p := rds.NewDescribeOptionGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.OptionGroupsList {
			r.resources = append(r.resources, *resource.OptionGroupName)
		}
	}
	r.err = p.Err()
	return
}

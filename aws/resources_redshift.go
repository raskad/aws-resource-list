package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func getRedshift(config aws.Config) (resources resourceMap) {
	client := redshift.New(config)
	resources = reduce(
		getRedshiftCluster(client).unwrap(redshiftCluster),
		getRedshiftClusterParameterGroup(client).unwrap(redshiftClusterParameterGroup),
		getRedshiftClusterSecurityGroup(client).unwrap(redshiftClusterSecurityGroup),
		getRedshiftClusterSubnetGroup(client).unwrap(redshiftClusterSubnetGroup),
	)
	return
}

func getRedshiftCluster(client *redshift.Client) (r resourceSliceError) {
	req := client.DescribeClustersRequest(&redshift.DescribeClustersInput{})
	p := redshift.NewDescribeClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Clusters {
			r.resources = append(r.resources, *resource.ClusterIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getRedshiftClusterParameterGroup(client *redshift.Client) (r resourceSliceError) {
	req := client.DescribeClusterParameterGroupsRequest(&redshift.DescribeClusterParameterGroupsInput{})
	p := redshift.NewDescribeClusterParameterGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ParameterGroups {
			r.resources = append(r.resources, *resource.ParameterGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getRedshiftClusterSecurityGroup(client *redshift.Client) (r resourceSliceError) {
	req := client.DescribeClusterSecurityGroupsRequest(&redshift.DescribeClusterSecurityGroupsInput{})
	p := redshift.NewDescribeClusterSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ClusterSecurityGroups {
			r.resources = append(r.resources, *resource.ClusterSecurityGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getRedshiftClusterSubnetGroup(client *redshift.Client) (r resourceSliceError) {
	req := client.DescribeClusterSubnetGroupsRequest(&redshift.DescribeClusterSubnetGroupsInput{})
	p := redshift.NewDescribeClusterSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ClusterSubnetGroups {
			r.resources = append(r.resources, *resource.ClusterSubnetGroupName)
		}
	}
	r.err = p.Err()
	return
}

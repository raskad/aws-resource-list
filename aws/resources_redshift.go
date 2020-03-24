package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func getRedshift(config aws.Config) (resources resourceMap) {
	client := redshift.New(config)

	redshiftClusterIDs := getRedshiftClusterIDs(client)
	redshiftClusterParameterGroupNames := getRedshiftClusterParameterGroupNames(client)
	redshiftClusterSecurityGroupNames := getRedshiftClusterSecurityGroupNames(client)
	redshiftClusterSubnetGroupNames := getRedshiftClusterSubnetGroupNames(client)

	resources = resourceMap{
		redshiftCluster:               redshiftClusterIDs,
		redshiftClusterParameterGroup: redshiftClusterParameterGroupNames,
		redshiftClusterSecurityGroup:  redshiftClusterSecurityGroupNames,
		redshiftClusterSubnetGroup:    redshiftClusterSubnetGroupNames,
	}
	return
}

func getRedshiftClusterIDs(client *redshift.Client) (resources []string) {
	req := client.DescribeClustersRequest(&redshift.DescribeClustersInput{})
	p := redshift.NewDescribeClustersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Clusters {
			resources = append(resources, *resource.ClusterIdentifier)
		}
	}
	return
}

func getRedshiftClusterParameterGroupNames(client *redshift.Client) (resources []string) {
	req := client.DescribeClusterParameterGroupsRequest(&redshift.DescribeClusterParameterGroupsInput{})
	p := redshift.NewDescribeClusterParameterGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ParameterGroups {
			resources = append(resources, *resource.ParameterGroupName)
		}
	}
	return
}

func getRedshiftClusterSecurityGroupNames(client *redshift.Client) (resources []string) {
	req := client.DescribeClusterSecurityGroupsRequest(&redshift.DescribeClusterSecurityGroupsInput{})
	p := redshift.NewDescribeClusterSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ClusterSecurityGroups {
			resources = append(resources, *resource.ClusterSecurityGroupName)
		}
	}
	return
}

func getRedshiftClusterSubnetGroupNames(client *redshift.Client) (resources []string) {
	req := client.DescribeClusterSubnetGroupsRequest(&redshift.DescribeClusterSubnetGroupsInput{})
	p := redshift.NewDescribeClusterSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ClusterSubnetGroups {
			resources = append(resources, *resource.ClusterSubnetGroupName)
		}
	}
	return
}

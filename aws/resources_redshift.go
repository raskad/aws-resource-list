package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshift"
)

func getRedshift(session *session.Session) (resources resourceMap) {
	client := redshift.New(session)
	resources = reduce(
		getRedshiftCluster(client).unwrap(redshiftCluster),
		getRedshiftClusterParameterGroup(client).unwrap(redshiftClusterParameterGroup),
		getRedshiftClusterSecurityGroup(client).unwrap(redshiftClusterSecurityGroup),
		getRedshiftClusterSubnetGroup(client).unwrap(redshiftClusterSubnetGroup),
	)
	return
}

func getRedshiftCluster(client *redshift.Redshift) (r resourceSliceError) {
	r.err = client.DescribeClustersPages(&redshift.DescribeClustersInput{}, func(page *redshift.DescribeClustersOutput, lastPage bool) bool {
		for _, resource := range page.Clusters {
			r.resources = append(r.resources, *resource.ClusterIdentifier)
		}
		return true
	})
	return
}

func getRedshiftClusterParameterGroup(client *redshift.Redshift) (r resourceSliceError) {
	r.err = client.DescribeClusterParameterGroupsPages(&redshift.DescribeClusterParameterGroupsInput{}, func(page *redshift.DescribeClusterParameterGroupsOutput, lastPage bool) bool {
		for _, resource := range page.ParameterGroups {
			r.resources = append(r.resources, *resource.ParameterGroupName)
		}
		return true
	})
	return
}

func getRedshiftClusterSecurityGroup(client *redshift.Redshift) (r resourceSliceError) {
	r.err = client.DescribeClusterSecurityGroupsPages(&redshift.DescribeClusterSecurityGroupsInput{}, func(page *redshift.DescribeClusterSecurityGroupsOutput, lastPage bool) bool {
		for _, resource := range page.ClusterSecurityGroups {
			r.resources = append(r.resources, *resource.ClusterSecurityGroupName)
		}
		return true
	})
	return
}

func getRedshiftClusterSubnetGroup(client *redshift.Redshift) (r resourceSliceError) {
	r.err = client.DescribeClusterSubnetGroupsPages(&redshift.DescribeClusterSubnetGroupsInput{}, func(page *redshift.DescribeClusterSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.ClusterSubnetGroups {
			r.resources = append(r.resources, *resource.ClusterSubnetGroupName)
		}
		return true
	})
	return
}

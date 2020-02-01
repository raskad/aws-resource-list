package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshift"
)

func getRedshift(session *session.Session) (resources resourceMap) {
	client := redshift.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		redshiftCluster:               getRedshiftCluster(client),
		redshiftClusterParameterGroup: getRedshiftClusterParameterGroup(client),
		redshiftClusterSecurityGroup:  getRedshiftClusterSecurityGroup(client),
		redshiftClusterSubnetGroup:    getRedshiftClusterSubnetGroup(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getRedshiftCluster(client *redshift.Redshift) (r resourceSliceError) {
	r.err = client.DescribeClustersPages(&redshift.DescribeClustersInput{}, func(page *redshift.DescribeClustersOutput, lastPage bool) bool {
		logDebug("Listing RedshiftCluster resources page")
		for _, resource := range page.Clusters {
			logDebug("Got RedshiftCluster resource with PhysicalResourceId", *resource.ClusterIdentifier)
			r.resources = append(r.resources, *resource.ClusterIdentifier)
		}
		return true
	})
	return
}

func getRedshiftClusterParameterGroup(client *redshift.Redshift) (r resourceSliceError) {
	r.err = client.DescribeClusterParameterGroupsPages(&redshift.DescribeClusterParameterGroupsInput{}, func(page *redshift.DescribeClusterParameterGroupsOutput, lastPage bool) bool {
		logDebug("Listing RedshiftClusterParameterGroup resources page")
		for _, resource := range page.ParameterGroups {
			logDebug("Got RedshiftClusterParameterGroup resource with PhysicalResourceId", *resource.ParameterGroupName)
			r.resources = append(r.resources, *resource.ParameterGroupName)
		}
		return true
	})
	return
}

func getRedshiftClusterSecurityGroup(client *redshift.Redshift) (r resourceSliceError) {
	r.err = client.DescribeClusterSecurityGroupsPages(&redshift.DescribeClusterSecurityGroupsInput{}, func(page *redshift.DescribeClusterSecurityGroupsOutput, lastPage bool) bool {
		logDebug("Listing RedshiftClusterSecurityGroup resources page")
		for _, resource := range page.ClusterSecurityGroups {
			logDebug("Got RedshiftClusterSecurityGroup resource with PhysicalResourceId", *resource.ClusterSecurityGroupName)
			r.resources = append(r.resources, *resource.ClusterSecurityGroupName)
		}
		return true
	})
	return
}

func getRedshiftClusterSubnetGroup(client *redshift.Redshift) (r resourceSliceError) {
	r.err = client.DescribeClusterSubnetGroupsPages(&redshift.DescribeClusterSubnetGroupsInput{}, func(page *redshift.DescribeClusterSubnetGroupsOutput, lastPage bool) bool {
		logDebug("Listing RedshiftClusterSubnetGroup resources page")
		for _, resource := range page.ClusterSubnetGroups {
			logDebug("Got RedshiftClusterSubnetGroup resource with PhysicalResourceId", *resource.ClusterSubnetGroupName)
			r.resources = append(r.resources, *resource.ClusterSubnetGroupName)
		}
		return true
	})
	return
}

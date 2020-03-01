package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func getRds(session *session.Session) (resources resourceMap) {
	client := rds.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		rdsDBCluster:               getRdsDBCluster(client),
		rdsDBClusterParameterGroup: getRdsDBClusterParameterGroup(client),
		rdsDBInstance:              getRdsDBInstance(client),
		rdsDBParameterGroup:        getRdsDBParameterGroup(client),
		rdsDBSecurityGroup:         getRdsDBSecurityGroup(client),
		rdsDBSubnetGroup:           getRdsDBSubnetGroup(client),
		rdsEventSubscription:       getRdsEventSubscription(client),
		rdsOptionGroup:             getRdsOptionGroup(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getRdsDBCluster(client *rds.RDS) (r resourceSliceError) {
	logDebug("Listing RdsDBCluster resources")
	r.err = client.DescribeDBClustersPages(&rds.DescribeDBClustersInput{}, func(page *rds.DescribeDBClustersOutput, lastPage bool) bool {
		for _, resource := range page.DBClusters {
			logDebug("Got RdsDBCluster resource with PhysicalResourceId", *resource.DBClusterIdentifier)
			r.resources = append(r.resources, *resource.DBClusterIdentifier)
		}
		return true
	})
	return
}

func getRdsDBClusterParameterGroup(client *rds.RDS) (r resourceSliceError) {
	logDebug("Listing RdsDBClusterParameterGroup resources")
	input := rds.DescribeDBClusterParameterGroupsInput{}
	for {
		page, err := client.DescribeDBClusterParameterGroups(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DBClusterParameterGroups {
			logDebug("Got RdsDBClusterParameterGroup resource with PhysicalResourceId", *resource.DBClusterParameterGroupName)
			r.resources = append(r.resources, *resource.DBClusterParameterGroupName)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getRdsDBInstance(client *rds.RDS) (r resourceSliceError) {
	logDebug("Listing RdsDBInstance resources")
	r.err = client.DescribeDBInstancesPages(&rds.DescribeDBInstancesInput{}, func(page *rds.DescribeDBInstancesOutput, lastPage bool) bool {
		for _, resource := range page.DBInstances {
			logDebug("Got RdsDBInstance resource with PhysicalResourceId", *resource.DBInstanceIdentifier)
			r.resources = append(r.resources, *resource.DBInstanceIdentifier)
		}
		return true
	})
	return
}

func getRdsDBParameterGroup(client *rds.RDS) (r resourceSliceError) {
	logDebug("Listing RdsDBParameterGroup resources")
	r.err = client.DescribeDBParameterGroupsPages(&rds.DescribeDBParameterGroupsInput{}, func(page *rds.DescribeDBParameterGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBParameterGroups {
			logDebug("Got RdsDBParameterGroup resource with PhysicalResourceId", *resource.DBParameterGroupName)
			r.resources = append(r.resources, *resource.DBParameterGroupName)
		}
		return true
	})
	return
}

func getRdsDBSecurityGroup(client *rds.RDS) (r resourceSliceError) {
	logDebug("Listing RdsDBSecurityGroup resources")
	r.err = client.DescribeDBSecurityGroupsPages(&rds.DescribeDBSecurityGroupsInput{}, func(page *rds.DescribeDBSecurityGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBSecurityGroups {
			logDebug("Got RdsDBSecurityGroup resource with PhysicalResourceId", *resource.DBSecurityGroupName)
			r.resources = append(r.resources, *resource.DBSecurityGroupName)
		}
		return true
	})
	return
}

func getRdsDBSubnetGroup(client *rds.RDS) (r resourceSliceError) {
	logDebug("Listing RdsDBSubnetGroup resources")
	r.err = client.DescribeDBSubnetGroupsPages(&rds.DescribeDBSubnetGroupsInput{}, func(page *rds.DescribeDBSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBSubnetGroups {
			logDebug("Got RdsDBSubnetGroup resource with PhysicalResourceId", *resource.DBSubnetGroupName)
			r.resources = append(r.resources, *resource.DBSubnetGroupName)
		}
		return true
	})
	return
}

func getRdsEventSubscription(client *rds.RDS) (r resourceSliceError) {
	logDebug("Listing RdsEventSubscription resources")
	r.err = client.DescribeEventSubscriptionsPages(&rds.DescribeEventSubscriptionsInput{}, func(page *rds.DescribeEventSubscriptionsOutput, lastPage bool) bool {
		for _, resource := range page.EventSubscriptionsList {
			logDebug("Got RdsEventSubscription resource with PhysicalResourceId", *resource.CustSubscriptionId)
			r.resources = append(r.resources, *resource.CustSubscriptionId)
		}
		return true
	})
	return
}

func getRdsOptionGroup(client *rds.RDS) (r resourceSliceError) {
	logDebug("Listing RdsOptionGroup resources")
	r.err = client.DescribeOptionGroupsPages(&rds.DescribeOptionGroupsInput{}, func(page *rds.DescribeOptionGroupsOutput, lastPage bool) bool {
		for _, resource := range page.OptionGroupsList {
			logDebug("Got RdsOptionGroup resource with PhysicalResourceId", *resource.OptionGroupName)
			r.resources = append(r.resources, *resource.OptionGroupName)
		}
		return true
	})
	return
}

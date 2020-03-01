package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func getRds(session *session.Session) (resources resourceMap) {
	client := rds.New(session)
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

func getRdsDBCluster(client *rds.RDS) (r resourceSliceError) {
	r.err = client.DescribeDBClustersPages(&rds.DescribeDBClustersInput{}, func(page *rds.DescribeDBClustersOutput, lastPage bool) bool {
		for _, resource := range page.DBClusters {
			r.resources = append(r.resources, *resource.DBClusterIdentifier)
		}
		return true
	})
	return
}

func getRdsDBClusterParameterGroup(client *rds.RDS) (r resourceSliceError) {
	input := rds.DescribeDBClusterParameterGroupsInput{}
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

func getRdsDBInstance(client *rds.RDS) (r resourceSliceError) {
	r.err = client.DescribeDBInstancesPages(&rds.DescribeDBInstancesInput{}, func(page *rds.DescribeDBInstancesOutput, lastPage bool) bool {
		for _, resource := range page.DBInstances {
			r.resources = append(r.resources, *resource.DBInstanceIdentifier)
		}
		return true
	})
	return
}

func getRdsDBParameterGroup(client *rds.RDS) (r resourceSliceError) {
	r.err = client.DescribeDBParameterGroupsPages(&rds.DescribeDBParameterGroupsInput{}, func(page *rds.DescribeDBParameterGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBParameterGroups {
			r.resources = append(r.resources, *resource.DBParameterGroupName)
		}
		return true
	})
	return
}

func getRdsDBSecurityGroup(client *rds.RDS) (r resourceSliceError) {
	r.err = client.DescribeDBSecurityGroupsPages(&rds.DescribeDBSecurityGroupsInput{}, func(page *rds.DescribeDBSecurityGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBSecurityGroups {
			r.resources = append(r.resources, *resource.DBSecurityGroupName)
		}
		return true
	})
	return
}

func getRdsDBSubnetGroup(client *rds.RDS) (r resourceSliceError) {
	r.err = client.DescribeDBSubnetGroupsPages(&rds.DescribeDBSubnetGroupsInput{}, func(page *rds.DescribeDBSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DBSubnetGroups {
			r.resources = append(r.resources, *resource.DBSubnetGroupName)
		}
		return true
	})
	return
}

func getRdsEventSubscription(client *rds.RDS) (r resourceSliceError) {
	r.err = client.DescribeEventSubscriptionsPages(&rds.DescribeEventSubscriptionsInput{}, func(page *rds.DescribeEventSubscriptionsOutput, lastPage bool) bool {
		for _, resource := range page.EventSubscriptionsList {
			r.resources = append(r.resources, *resource.CustSubscriptionId)
		}
		return true
	})
	return
}

func getRdsOptionGroup(client *rds.RDS) (r resourceSliceError) {
	r.err = client.DescribeOptionGroupsPages(&rds.DescribeOptionGroupsInput{}, func(page *rds.DescribeOptionGroupsOutput, lastPage bool) bool {
		for _, resource := range page.OptionGroupsList {
			r.resources = append(r.resources, *resource.OptionGroupName)
		}
		return true
	})
	return
}

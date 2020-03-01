package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
)

func getDms(session *session.Session) (resources resourceMap) {
	client := databasemigrationservice.New(session)
	resources = reduce(
		getDmsCertificate(client).unwrap(dmsCertificate),
		getDmsEndpoint(client).unwrap(dmsEndpoint),
		getDmsEventSubscription(client).unwrap(dmsEventSubscription),
		getDmsReplicationInstance(client).unwrap(dmsReplicationInstance),
		getDmsReplicationSubnetGroup(client).unwrap(dmsReplicationSubnetGroup),
		getDmsReplicationTask(client).unwrap(dmsReplicationTask),
	)
	return
}

func getDmsCertificate(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	r.err = client.DescribeCertificatesPages(&databasemigrationservice.DescribeCertificatesInput{}, func(page *databasemigrationservice.DescribeCertificatesOutput, lastPage bool) bool {
		for _, resource := range page.Certificates {
			r.resources = append(r.resources, *resource.CertificateIdentifier)
		}
		return true
	})
	return
}

func getDmsEndpoint(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	r.err = client.DescribeEndpointsPages(&databasemigrationservice.DescribeEndpointsInput{}, func(page *databasemigrationservice.DescribeEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.Endpoints {
			r.resources = append(r.resources, *resource.EndpointIdentifier)
		}
		return true
	})
	return
}

func getDmsEventSubscription(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	r.err = client.DescribeEventSubscriptionsPages(&databasemigrationservice.DescribeEventSubscriptionsInput{}, func(page *databasemigrationservice.DescribeEventSubscriptionsOutput, lastPage bool) bool {
		for _, resource := range page.EventSubscriptionsList {
			r.resources = append(r.resources, *resource.CustSubscriptionId)
		}
		return true
	})
	return
}

func getDmsReplicationInstance(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	r.err = client.DescribeReplicationInstancesPages(&databasemigrationservice.DescribeReplicationInstancesInput{}, func(page *databasemigrationservice.DescribeReplicationInstancesOutput, lastPage bool) bool {
		for _, resource := range page.ReplicationInstances {
			r.resources = append(r.resources, *resource.ReplicationInstanceIdentifier)
		}
		return true
	})
	return
}

func getDmsReplicationSubnetGroup(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	r.err = client.DescribeReplicationSubnetGroupsPages(&databasemigrationservice.DescribeReplicationSubnetGroupsInput{}, func(page *databasemigrationservice.DescribeReplicationSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.ReplicationSubnetGroups {
			r.resources = append(r.resources, *resource.ReplicationSubnetGroupIdentifier)
		}
		return true
	})
	return
}

func getDmsReplicationTask(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	r.err = client.DescribeReplicationTasksPages(&databasemigrationservice.DescribeReplicationTasksInput{}, func(page *databasemigrationservice.DescribeReplicationTasksOutput, lastPage bool) bool {
		for _, resource := range page.ReplicationTasks {
			r.resources = append(r.resources, *resource.ReplicationTaskIdentifier)
		}
		return true
	})
	return
}

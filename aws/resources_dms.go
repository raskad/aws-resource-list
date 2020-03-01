package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
)

func getDms(session *session.Session) (resources resourceMap) {
	client := databasemigrationservice.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		dmsCertificate:            getDmsCertificate(client),
		dmsEndpoint:               getDmsEndpoint(client),
		dmsEventSubscription:      getDmsEventSubscription(client),
		dmsReplicationInstance:    getDmsReplicationInstance(client),
		dmsReplicationSubnetGroup: getDmsReplicationSubnetGroup(client),
		dmsReplicationTask:        getDmsReplicationTask(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getDmsCertificate(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	logDebug("Listing DmsCertificate resources")
	r.err = client.DescribeCertificatesPages(&databasemigrationservice.DescribeCertificatesInput{}, func(page *databasemigrationservice.DescribeCertificatesOutput, lastPage bool) bool {
		for _, resource := range page.Certificates {
			logDebug("Got DmsCertificate resource with PhysicalResourceId", *resource.CertificateIdentifier)
			r.resources = append(r.resources, *resource.CertificateIdentifier)
		}
		return true
	})
	return
}

func getDmsEndpoint(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	logDebug("Listing DmsEndpoint resources")
	r.err = client.DescribeEndpointsPages(&databasemigrationservice.DescribeEndpointsInput{}, func(page *databasemigrationservice.DescribeEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.Endpoints {
			logDebug("Got DmsEndpoint resource with PhysicalResourceId", *resource.EndpointIdentifier)
			r.resources = append(r.resources, *resource.EndpointIdentifier)
		}
		return true
	})
	return
}

func getDmsEventSubscription(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	logDebug("Listing DmsEventSubscription resources")
	r.err = client.DescribeEventSubscriptionsPages(&databasemigrationservice.DescribeEventSubscriptionsInput{}, func(page *databasemigrationservice.DescribeEventSubscriptionsOutput, lastPage bool) bool {
		for _, resource := range page.EventSubscriptionsList {
			logDebug("Got DmsEventSubscription resource with PhysicalResourceId", *resource.CustSubscriptionId)
			r.resources = append(r.resources, *resource.CustSubscriptionId)
		}
		return true
	})
	return
}

func getDmsReplicationInstance(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	logDebug("Listing DmsReplicationInstance resources")
	r.err = client.DescribeReplicationInstancesPages(&databasemigrationservice.DescribeReplicationInstancesInput{}, func(page *databasemigrationservice.DescribeReplicationInstancesOutput, lastPage bool) bool {
		for _, resource := range page.ReplicationInstances {
			logDebug("Got DmsReplicationInstance resource with PhysicalResourceId", *resource.ReplicationInstanceIdentifier)
			r.resources = append(r.resources, *resource.ReplicationInstanceIdentifier)
		}
		return true
	})
	return
}

func getDmsReplicationSubnetGroup(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	logDebug("Listing DmsReplicationSubnetGroup resources")
	r.err = client.DescribeReplicationSubnetGroupsPages(&databasemigrationservice.DescribeReplicationSubnetGroupsInput{}, func(page *databasemigrationservice.DescribeReplicationSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.ReplicationSubnetGroups {
			logDebug("Got DmsReplicationSubnetGroup resource with PhysicalResourceId", *resource.ReplicationSubnetGroupIdentifier)
			r.resources = append(r.resources, *resource.ReplicationSubnetGroupIdentifier)
		}
		return true
	})
	return
}

func getDmsReplicationTask(client *databasemigrationservice.DatabaseMigrationService) (r resourceSliceError) {
	logDebug("Listing DmsReplicationTask resources")
	r.err = client.DescribeReplicationTasksPages(&databasemigrationservice.DescribeReplicationTasksInput{}, func(page *databasemigrationservice.DescribeReplicationTasksOutput, lastPage bool) bool {
		for _, resource := range page.ReplicationTasks {
			logDebug("Got DmsReplicationTask resource with PhysicalResourceId", *resource.ReplicationTaskIdentifier)
			r.resources = append(r.resources, *resource.ReplicationTaskIdentifier)
		}
		return true
	})
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func getDatabaseMigrationService(config aws.Config) (resources awsResourceMap) {
	client := databasemigrationservice.New(config)

	databaseMigrationServiceCertificateIDs := getDatabaseMigrationServiceCertificateIDs(client)
	databaseMigrationServiceEndpointIDs := getDatabaseMigrationServiceEndpointIDs(client)
	databaseMigrationServiceEventSubscriptionIDs := getDatabaseMigrationServiceEventSubscriptionIDs(client)
	databaseMigrationServiceReplicationInstanceIDs := getDatabaseMigrationServiceReplicationInstanceIDs(client)
	databaseMigrationServiceReplicationSubnetGroupIDs := getDatabaseMigrationServiceReplicationSubnetGroupIDs(client)
	databaseMigrationServiceReplicationTaskIDs := getDatabaseMigrationServiceReplicationTaskIDs(client)

	resources = awsResourceMap{
		databaseMigrationServiceCertificate:            databaseMigrationServiceCertificateIDs,
		databaseMigrationServiceEndpoint:               databaseMigrationServiceEndpointIDs,
		databaseMigrationServiceEventSubscription:      databaseMigrationServiceEventSubscriptionIDs,
		databaseMigrationServiceReplicationInstance:    databaseMigrationServiceReplicationInstanceIDs,
		databaseMigrationServiceReplicationSubnetGroup: databaseMigrationServiceReplicationSubnetGroupIDs,
		databaseMigrationServiceReplicationTask:        databaseMigrationServiceReplicationTaskIDs,
	}
	return
}

func getDatabaseMigrationServiceCertificateIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeCertificatesRequest(&databasemigrationservice.DescribeCertificatesInput{})
	p := databasemigrationservice.NewDescribeCertificatesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Certificates {
			resources = append(resources, *resource.CertificateIdentifier)
		}
	}
	return
}

func getDatabaseMigrationServiceEndpointIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeEndpointsRequest(&databasemigrationservice.DescribeEndpointsInput{})
	p := databasemigrationservice.NewDescribeEndpointsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Endpoints {
			resources = append(resources, *resource.EndpointIdentifier)
		}
	}
	return
}

func getDatabaseMigrationServiceEventSubscriptionIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeEventSubscriptionsRequest(&databasemigrationservice.DescribeEventSubscriptionsInput{})
	p := databasemigrationservice.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.EventSubscriptionsList {
			resources = append(resources, *resource.CustSubscriptionId)
		}
	}
	return
}

func getDatabaseMigrationServiceReplicationInstanceIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeReplicationInstancesRequest(&databasemigrationservice.DescribeReplicationInstancesInput{})
	p := databasemigrationservice.NewDescribeReplicationInstancesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ReplicationInstances {
			resources = append(resources, *resource.ReplicationInstanceIdentifier)
		}
	}
	return
}

func getDatabaseMigrationServiceReplicationSubnetGroupIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeReplicationSubnetGroupsRequest(&databasemigrationservice.DescribeReplicationSubnetGroupsInput{})
	p := databasemigrationservice.NewDescribeReplicationSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ReplicationSubnetGroups {
			resources = append(resources, *resource.ReplicationSubnetGroupIdentifier)
		}
	}
	return
}

func getDatabaseMigrationServiceReplicationTaskIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeReplicationTasksRequest(&databasemigrationservice.DescribeReplicationTasksInput{})
	p := databasemigrationservice.NewDescribeReplicationTasksPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ReplicationTasks {
			resources = append(resources, *resource.ReplicationTaskIdentifier)
		}
	}
	return
}

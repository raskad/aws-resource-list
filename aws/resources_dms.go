package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func getDms(config aws.Config) (resources resourceMap) {
	client := databasemigrationservice.New(config)

	dmsCertificateIDs := getDmsCertificateIDs(client)
	dmsEndpointIDs := getDmsEndpointIDs(client)
	dmsEventSubscriptionIDs := getDmsEventSubscriptionIDs(client)
	dmsReplicationInstanceIDs := getDmsReplicationInstanceIDs(client)
	dmsReplicationSubnetGroupIDs := getDmsReplicationSubnetGroupIDs(client)
	dmsReplicationTaskIDs := getDmsReplicationTaskIDs(client)

	resources = resourceMap{
		dmsCertificate:            dmsCertificateIDs,
		dmsEndpoint:               dmsEndpointIDs,
		dmsEventSubscription:      dmsEventSubscriptionIDs,
		dmsReplicationInstance:    dmsReplicationInstanceIDs,
		dmsReplicationSubnetGroup: dmsReplicationSubnetGroupIDs,
		dmsReplicationTask:        dmsReplicationTaskIDs,
	}
	return
}

func getDmsCertificateIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeCertificatesRequest(&databasemigrationservice.DescribeCertificatesInput{})
	p := databasemigrationservice.NewDescribeCertificatesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Certificates {
			resources = append(resources, *resource.CertificateIdentifier)
		}
	}
	return
}

func getDmsEndpointIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeEndpointsRequest(&databasemigrationservice.DescribeEndpointsInput{})
	p := databasemigrationservice.NewDescribeEndpointsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Endpoints {
			resources = append(resources, *resource.EndpointIdentifier)
		}
	}
	return
}

func getDmsEventSubscriptionIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeEventSubscriptionsRequest(&databasemigrationservice.DescribeEventSubscriptionsInput{})
	p := databasemigrationservice.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.EventSubscriptionsList {
			resources = append(resources, *resource.CustSubscriptionId)
		}
	}
	return
}

func getDmsReplicationInstanceIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeReplicationInstancesRequest(&databasemigrationservice.DescribeReplicationInstancesInput{})
	p := databasemigrationservice.NewDescribeReplicationInstancesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ReplicationInstances {
			resources = append(resources, *resource.ReplicationInstanceIdentifier)
		}
	}
	return
}

func getDmsReplicationSubnetGroupIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeReplicationSubnetGroupsRequest(&databasemigrationservice.DescribeReplicationSubnetGroupsInput{})
	p := databasemigrationservice.NewDescribeReplicationSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ReplicationSubnetGroups {
			resources = append(resources, *resource.ReplicationSubnetGroupIdentifier)
		}
	}
	return
}

func getDmsReplicationTaskIDs(client *databasemigrationservice.Client) (resources []string) {
	req := client.DescribeReplicationTasksRequest(&databasemigrationservice.DescribeReplicationTasksInput{})
	p := databasemigrationservice.NewDescribeReplicationTasksPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.ReplicationTasks {
			resources = append(resources, *resource.ReplicationTaskIdentifier)
		}
	}
	return
}

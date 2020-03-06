package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func getDms(config aws.Config) (resources resourceMap) {
	client := databasemigrationservice.New(config)
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

func getDmsCertificate(client *databasemigrationservice.Client) (r resourceSliceError) {
	req := client.DescribeCertificatesRequest(&databasemigrationservice.DescribeCertificatesInput{})
	p := databasemigrationservice.NewDescribeCertificatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Certificates {
			r.resources = append(r.resources, *resource.CertificateIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getDmsEndpoint(client *databasemigrationservice.Client) (r resourceSliceError) {
	req := client.DescribeEndpointsRequest(&databasemigrationservice.DescribeEndpointsInput{})
	p := databasemigrationservice.NewDescribeEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Endpoints {
			r.resources = append(r.resources, *resource.EndpointIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getDmsEventSubscription(client *databasemigrationservice.Client) (r resourceSliceError) {
	req := client.DescribeEventSubscriptionsRequest(&databasemigrationservice.DescribeEventSubscriptionsInput{})
	p := databasemigrationservice.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.EventSubscriptionsList {
			r.resources = append(r.resources, *resource.CustSubscriptionId)
		}
	}
	r.err = p.Err()
	return
}

func getDmsReplicationInstance(client *databasemigrationservice.Client) (r resourceSliceError) {
	req := client.DescribeReplicationInstancesRequest(&databasemigrationservice.DescribeReplicationInstancesInput{})
	p := databasemigrationservice.NewDescribeReplicationInstancesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ReplicationInstances {
			r.resources = append(r.resources, *resource.ReplicationInstanceIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getDmsReplicationSubnetGroup(client *databasemigrationservice.Client) (r resourceSliceError) {
	req := client.DescribeReplicationSubnetGroupsRequest(&databasemigrationservice.DescribeReplicationSubnetGroupsInput{})
	p := databasemigrationservice.NewDescribeReplicationSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ReplicationSubnetGroups {
			r.resources = append(r.resources, *resource.ReplicationSubnetGroupIdentifier)
		}
	}
	r.err = p.Err()
	return
}

func getDmsReplicationTask(client *databasemigrationservice.Client) (r resourceSliceError) {
	req := client.DescribeReplicationTasksRequest(&databasemigrationservice.DescribeReplicationTasksInput{})
	p := databasemigrationservice.NewDescribeReplicationTasksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ReplicationTasks {
			r.resources = append(r.resources, *resource.ReplicationTaskIdentifier)
		}
	}
	r.err = p.Err()
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func getRedshift(config aws.Config) (resources awsResourceMap) {
	client := redshift.New(config)

	redshiftClusterIDs := getRedshiftClusterIDs(client)
	redshiftClusterParameterGroupNames := getRedshiftClusterParameterGroupNames(client)
	redshiftClusterSecurityGroupNames := getRedshiftClusterSecurityGroupNames(client)
	redshiftClusterSubnetGroupNames := getRedshiftClusterSubnetGroupNames(client)
	redshiftEventSubscriptionNames := getRedshiftEventSubscriptionNames(client)
	redshiftSnapshotCopyGrantNames := getRedshiftSnapshotCopyGrantNames(client)
	redshiftSnapshotScheduleNames := getRedshiftSnapshotScheduleNames(client)

	resources = awsResourceMap{
		redshiftCluster:               redshiftClusterIDs,
		redshiftClusterParameterGroup: redshiftClusterParameterGroupNames,
		redshiftClusterSecurityGroup:  redshiftClusterSecurityGroupNames,
		redshiftClusterSubnetGroup:    redshiftClusterSubnetGroupNames,
		redshiftEventSubscription:     redshiftEventSubscriptionNames,
		redshiftSnapshotCopyGrant:     redshiftSnapshotCopyGrantNames,
		redshiftSnapshotSchedule:      redshiftSnapshotScheduleNames,
	}
	return
}

func getRedshiftClusterIDs(client *redshift.Client) (resources []string) {
	req := client.DescribeClustersRequest(&redshift.DescribeClustersInput{})
	p := redshift.NewDescribeClustersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ClusterSubnetGroups {
			resources = append(resources, *resource.ClusterSubnetGroupName)
		}
	}
	return
}

func getRedshiftEventSubscriptionNames(client *redshift.Client) (resources []string) {
	req := client.DescribeEventSubscriptionsRequest(&redshift.DescribeEventSubscriptionsInput{})
	p := redshift.NewDescribeEventSubscriptionsPaginator(req)
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

func getRedshiftSnapshotCopyGrantNames(client *redshift.Client) (resources []string) {
	input := redshift.DescribeSnapshotCopyGrantsInput{}
	for {
		page, err := client.DescribeSnapshotCopyGrantsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.SnapshotCopyGrants {
			resources = append(resources, *resource.SnapshotCopyGrantName)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

func getRedshiftSnapshotScheduleNames(client *redshift.Client) (resources []string) {
	input := redshift.DescribeSnapshotSchedulesInput{}
	for {
		page, err := client.DescribeSnapshotSchedulesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.SnapshotSchedules {
			resources = append(resources, *resource.ScheduleIdentifier)
		}
		if page.Marker == nil {
			return
		}
		input.Marker = page.Marker
	}
}

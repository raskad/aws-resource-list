package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
)

func getDAX(config aws.Config) (resources resourceMap) {
	client := dax.New(config)

	daxClusterNames := getDaxClusterNames(client)
	daxParameterGroupNames := getDaxParameterGroupNames(client)
	daxSubnetGroupNames := getDaxSubnetGroupNames(client)

	resources = resourceMap{
		daxCluster:        daxClusterNames,
		daxParameterGroup: daxParameterGroupNames,
		daxSubnetGroup:    daxSubnetGroupNames,
	}
	return
}

func getDaxClusterNames(client *dax.Client) (resources []string) {
	input := dax.DescribeClustersInput{}
	for {
		page, err := client.DescribeClustersRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Clusters {
			resources = append(resources, *resource.ClusterName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getDaxParameterGroupNames(client *dax.Client) (resources []string) {
	input := dax.DescribeParameterGroupsInput{}
	for {
		page, err := client.DescribeParameterGroupsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.ParameterGroups {
			resources = append(resources, *resource.ParameterGroupName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getDaxSubnetGroupNames(client *dax.Client) (resources []string) {
	input := dax.DescribeSubnetGroupsInput{}
	for {
		page, err := client.DescribeSubnetGroupsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.SubnetGroups {
			resources = append(resources, *resource.SubnetGroupName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

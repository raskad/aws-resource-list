package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
)

func getDAX(config aws.Config) (resources resourceMap) {
	client := dax.New(config)
	resources = reduce(
		getDaxCluster(client).unwrap(daxCluster),
		getDaxParameterGroup(client).unwrap(daxParameterGroup),
		getDaxSubnetGroup(client).unwrap(daxSubnetGroup),
	)
	return
}

func getDaxCluster(client *dax.Client) (r resourceSliceError) {
	input := dax.DescribeClustersInput{}
	for {
		page, err := client.DescribeClustersRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Clusters {
			r.resources = append(r.resources, *resource.ClusterName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getDaxParameterGroup(client *dax.Client) (r resourceSliceError) {
	input := dax.DescribeParameterGroupsInput{}
	for {
		page, err := client.DescribeParameterGroupsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ParameterGroups {
			r.resources = append(r.resources, *resource.ParameterGroupName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getDaxSubnetGroup(client *dax.Client) (r resourceSliceError) {
	input := dax.DescribeSubnetGroupsInput{}
	for {
		page, err := client.DescribeSubnetGroupsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.SubnetGroups {
			r.resources = append(r.resources, *resource.SubnetGroupName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

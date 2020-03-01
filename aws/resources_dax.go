package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dax"
)

func getDAX(session *session.Session) (resources resourceMap) {
	client := dax.New(session)
	resources = reduce(
		getDaxCluster(client).unwrap(daxCluster),
		getDaxParameterGroup(client).unwrap(daxParameterGroup),
		getDaxSubnetGroup(client).unwrap(daxSubnetGroup),
	)
	return
}

func getDaxCluster(client *dax.DAX) (r resourceSliceError) {
	input := dax.DescribeClustersInput{}
	for {
		page, err := client.DescribeClusters(&input)
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

func getDaxParameterGroup(client *dax.DAX) (r resourceSliceError) {
	input := dax.DescribeParameterGroupsInput{}
	for {
		page, err := client.DescribeParameterGroups(&input)
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

func getDaxSubnetGroup(client *dax.DAX) (r resourceSliceError) {
	input := dax.DescribeSubnetGroupsInput{}
	for {
		page, err := client.DescribeSubnetGroups(&input)
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

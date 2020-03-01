package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dax"
)

func getDAX(session *session.Session) (resources resourceMap) {
	client := dax.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		daxCluster:        getDaxCluster(client),
		daxParameterGroup: getDaxParameterGroup(client),
		daxSubnetGroup:    getDaxSubnetGroup(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getDaxCluster(client *dax.DAX) (r resourceSliceError) {
	logDebug("Listing DaxCluster resources")
	input := dax.DescribeClustersInput{}
	for {
		page, err := client.DescribeClusters(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Clusters {
			logDebug("Got DaxCluster resource with PhysicalResourceId", *resource.ClusterName)
			r.resources = append(r.resources, *resource.ClusterName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getDaxParameterGroup(client *dax.DAX) (r resourceSliceError) {
	logDebug("Listing DaxParameterGroup resources")
	input := dax.DescribeParameterGroupsInput{}
	for {
		page, err := client.DescribeParameterGroups(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ParameterGroups {
			logDebug("Got DaxParameterGroup resource with PhysicalResourceId", *resource.ParameterGroupName)
			r.resources = append(r.resources, *resource.ParameterGroupName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getDaxSubnetGroup(client *dax.DAX) (r resourceSliceError) {
	logDebug("Listing DaxSubnetGroup resources")
	input := dax.DescribeSubnetGroupsInput{}
	for {
		page, err := client.DescribeSubnetGroups(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.SubnetGroups {
			logDebug("Got DaxSubnetGroup resource with PhysicalResourceId", *resource.SubnetGroupName)
			r.resources = append(r.resources, *resource.SubnetGroupName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

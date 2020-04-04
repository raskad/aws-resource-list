package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
)

func getResourceGroups(config aws.Config) (resources awsResourceMap) {
	client := resourcegroups.New(config)

	resourceGroupsGroupNames := getResourceGroupsGroupNames(client)

	resources = awsResourceMap{
		resourceGroupsGroup: resourceGroupsGroupNames,
	}
	return
}

func getResourceGroupsGroupNames(client *resourcegroups.Client) (resources []string) {
	input := resourcegroups.ListGroupsInput{}
	for {
		page, err := client.ListGroupsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Groups {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

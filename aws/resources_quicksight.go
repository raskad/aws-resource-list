package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
)

func getQuickSight(config aws.Config) (resources awsResourceMap) {
	client := quicksight.New(config)

	quickSightGroupNames := getQuickSightGroupNames(client)
	quickSightUserNames := getQuickSightUserNames(client)

	resources = awsResourceMap{
		quickSightGroup: quickSightGroupNames,
		quickSightUser:  quickSightUserNames,
	}
	return
}

func getQuickSightGroupNames(client *quicksight.Client) (resources []string) {
	input := quicksight.ListGroupsInput{
		AwsAccountId: &accountID,
		Namespace:    aws.String("default"),
	}
	for {
		page, err := client.ListGroupsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.GroupList {
			resources = append(resources, *resource.GroupName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getQuickSightUserNames(client *quicksight.Client) (resources []string) {
	input := quicksight.ListUsersInput{
		AwsAccountId: &accountID,
		Namespace:    aws.String("default"),
	}
	for {
		page, err := client.ListUsersRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.UserList {
			resources = append(resources, *resource.UserName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

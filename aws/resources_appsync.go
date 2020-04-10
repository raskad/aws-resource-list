package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

func getAppSync(config aws.Config) (resources awsResourceMap) {
	client := appsync.New(config)

	appSyncGraphQLApiIDs := getAppSyncGraphQLApiIDs(client)
	appSyncFunctionsIDs := getAppSyncFunctionsIDs(client, appSyncGraphQLApiIDs)

	resources = awsResourceMap{
		appSyncFunctions:  appSyncFunctionsIDs,
		appSyncGraphQLApi: appSyncGraphQLApiIDs,
	}
	return
}

func getAppSyncFunctionsIDs(client *appsync.Client, appSyncGraphQLApiIDs []string) (resources []string) {
	for _, appSyncGraphQLApiID := range appSyncGraphQLApiIDs {
		input := appsync.ListFunctionsInput{
			ApiId: &appSyncGraphQLApiID,
		}
		for {
			page, err := client.ListFunctionsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.Functions {
				resources = append(resources, *resource.FunctionId)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getAppSyncGraphQLApiIDs(client *appsync.Client) (resources []string) {
	input := appsync.ListGraphqlApisInput{}
	for {
		page, err := client.ListGraphqlApisRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.GraphqlApis {
			resources = append(resources, *resource.ApiId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

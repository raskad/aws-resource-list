package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

func getAppSync(config aws.Config) (resources resourceMap) {
	client := appsync.New(config)

	appSyncGraphQLApiIDs := getAppSyncGraphQLApiIDs(client)

	resources = resourceMap{
		appSyncGraphQLApi: appSyncGraphQLApiIDs,
	}
	return
}

func getAppSyncGraphQLApiIDs(client *appsync.Client) (resources []string) {
	input := appsync.ListGraphqlApisInput{}
	for {
		page, err := client.ListGraphqlApisRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.GraphqlApis {
			resources = append(resources, *resource.ApiId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

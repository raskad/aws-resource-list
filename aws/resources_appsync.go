package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

func getAppSync(config aws.Config) (resources resourceMap) {
	client := appsync.New(config)
	resources = reduce(
		getAppSyncGraphQLApi(client).unwrap(appSyncGraphQLApi),
	)
	return
}

func getAppSyncGraphQLApi(client *appsync.Client) (r resourceSliceError) {
	input := appsync.ListGraphqlApisInput{}
	for {
		page, err := client.ListGraphqlApisRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.GraphqlApis {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

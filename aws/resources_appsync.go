package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appsync"
)

func getAppSync(session *session.Session) (resources resourceMap) {
	client := appsync.New(session)
	resources = reduce(
		getAppSyncGraphQLApi(client).unwrap(appSyncGraphQLApi),
	)
	return
}

func getAppSyncGraphQLApi(client *appsync.AppSync) (r resourceSliceError) {
	input := appsync.ListGraphqlApisInput{}
	for {
		page, err := client.ListGraphqlApis(&input)
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

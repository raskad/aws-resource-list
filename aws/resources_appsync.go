package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appsync"
)

func getAppSync(session *session.Session) (resources resourceMap) {
	client := appsync.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		appSyncGraphQLApi: getAppSyncGraphQLApi(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAppSyncGraphQLApi(client *appsync.AppSync) (r resourceSliceError) {
	logDebug("Listing AppSyncGraphQLApi resources")
	input := appsync.ListGraphqlApisInput{}
	for {
		page, err := client.ListGraphqlApis(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.GraphqlApis {
			logDebug("Got AppSyncGraphQLApi resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

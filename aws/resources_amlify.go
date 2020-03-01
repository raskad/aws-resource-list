package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/amplify"
)

func getAmplify(session *session.Session) (resources resourceMap) {
	client := amplify.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		amplifyApp: getAmplifyApp(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAmplifyApp(client *amplify.Amplify) (r resourceSliceError) {
	logDebug("Listing AmplifyApp resources")
	input := amplify.ListAppsInput{}
	for {
		page, err := client.ListApps(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Apps {
			logDebug("Got AmplifyApp resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

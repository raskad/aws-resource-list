package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
)

func getCognitoIdentity(session *session.Session) (resources resourceMap) {
	client := cognitoidentity.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		cognitoIdentityPool: getCognitoIdentityPool(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCognitoIdentityPool(client *cognitoidentity.CognitoIdentity) (r resourceSliceError) {
	input := cognitoidentity.ListIdentityPoolsInput{}
	for {
		page, err := client.ListIdentityPools(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing CognitoIdentityPool resources page. Remaining pages", page.NextToken)
		for _, resource := range page.IdentityPools {
			logDebug("Got CognitoIdentityPool resource with PhysicalResourceId", *resource.IdentityPoolName)
			r.resources = append(r.resources, *resource.IdentityPoolName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

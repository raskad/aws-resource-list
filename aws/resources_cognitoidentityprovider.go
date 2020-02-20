package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func getCognitoIdentityProvider(session *session.Session) (resources resourceMap) {
	client := cognitoidentityprovider.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		cognitoUserPool:                 getCognitoUserPool(client),
		cognitoUserPoolClient:           getCognitoUserPoolClient(client),
		cognitoUserPoolGroup:            getCognitoUserPoolGroup(client),
		cognitoUserPoolIdentityProvider: getCognitoUserPoolIdentityProvider(client),
		cognitoUserPoolResourceServer:   getCognitoUserPoolResourceServer(client),
		cognitoUserPoolUser:             getCognitoUserPoolUser(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCognitoUserPool(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListUserPoolsPages(&cognitoidentityprovider.ListUserPoolsInput{}, func(page *cognitoidentityprovider.ListUserPoolsOutput, lastPage bool) bool {
		logDebug("List CognitoUserPool resources page")
		for _, resource := range page.UserPools {
			logDebug("Got CognitoUserPool resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getCognitoUserPoolClient(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListUserPoolClientsPages(&cognitoidentityprovider.ListUserPoolClientsInput{}, func(page *cognitoidentityprovider.ListUserPoolClientsOutput, lastPage bool) bool {
		logDebug("List CognitoUserPoolClient resources page")
		for _, resource := range page.UserPoolClients {
			logDebug("Got CognitoUserPoolClient resource with PhysicalResourceId", *resource.ClientName)
			r.resources = append(r.resources, *resource.ClientName)
		}
		return true
	})
	return
}

func getCognitoUserPoolGroup(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListGroupsPages(&cognitoidentityprovider.ListGroupsInput{}, func(page *cognitoidentityprovider.ListGroupsOutput, lastPage bool) bool {
		logDebug("List CognitoUserPoolGroup resources page")
		for _, resource := range page.Groups {
			logDebug("Got CognitoUserPoolGroup resource with PhysicalResourceId", *resource.GroupName)
			r.resources = append(r.resources, *resource.GroupName)
		}
		return true
	})
	return
}

func getCognitoUserPoolIdentityProvider(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListIdentityProvidersPages(&cognitoidentityprovider.ListIdentityProvidersInput{}, func(page *cognitoidentityprovider.ListIdentityProvidersOutput, lastPage bool) bool {
		logDebug("List CognitoUserPoolIdentityProvider resources page")
		for _, resource := range page.Providers {
			logDebug("Got CognitoUserPoolIdentityProvider resource with PhysicalResourceId", *resource.ProviderName)
			r.resources = append(r.resources, *resource.ProviderName)
		}
		return true
	})
	return
}

func getCognitoUserPoolResourceServer(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListResourceServersPages(&cognitoidentityprovider.ListResourceServersInput{}, func(page *cognitoidentityprovider.ListResourceServersOutput, lastPage bool) bool {
		logDebug("List CognitoUserPoolResourceServer resources page")
		for _, resource := range page.ResourceServers {
			logDebug("Got CognitoUserPoolResourceServer resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getCognitoUserPoolUser(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListUsersPages(&cognitoidentityprovider.ListUsersInput{}, func(page *cognitoidentityprovider.ListUsersOutput, lastPage bool) bool {
		logDebug("List CognitoUserPoolUser resources page")
		for _, resource := range page.Users {
			logDebug("Got CognitoUserPoolUser resource with PhysicalResourceId", *resource.Username)
			r.resources = append(r.resources, *resource.Username)
		}
		return true
	})
	return
}

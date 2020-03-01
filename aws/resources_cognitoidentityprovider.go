package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func getCognitoIdentityProvider(session *session.Session) (resources resourceMap) {
	client := cognitoidentityprovider.New(session)
	resources = reduce(
		getCognitoUserPool(client).unwrap(cognitoUserPool),
		getCognitoUserPoolClient(client).unwrap(cognitoUserPoolClient),
		getCognitoUserPoolGroup(client).unwrap(cognitoUserPoolGroup),
		getCognitoUserPoolIdentityProvider(client).unwrap(cognitoUserPoolIdentityProvider),
		getCognitoUserPoolResourceServer(client).unwrap(cognitoUserPoolResourceServer),
		getCognitoUserPoolUser(client).unwrap(cognitoUserPoolUser),
	)
	return
}

func getCognitoUserPool(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListUserPoolsPages(&cognitoidentityprovider.ListUserPoolsInput{}, func(page *cognitoidentityprovider.ListUserPoolsOutput, lastPage bool) bool {
		for _, resource := range page.UserPools {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getCognitoUserPoolClient(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListUserPoolClientsPages(&cognitoidentityprovider.ListUserPoolClientsInput{}, func(page *cognitoidentityprovider.ListUserPoolClientsOutput, lastPage bool) bool {
		for _, resource := range page.UserPoolClients {
			r.resources = append(r.resources, *resource.ClientName)
		}
		return true
	})
	return
}

func getCognitoUserPoolGroup(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListGroupsPages(&cognitoidentityprovider.ListGroupsInput{}, func(page *cognitoidentityprovider.ListGroupsOutput, lastPage bool) bool {
		for _, resource := range page.Groups {
			r.resources = append(r.resources, *resource.GroupName)
		}
		return true
	})
	return
}

func getCognitoUserPoolIdentityProvider(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListIdentityProvidersPages(&cognitoidentityprovider.ListIdentityProvidersInput{}, func(page *cognitoidentityprovider.ListIdentityProvidersOutput, lastPage bool) bool {
		for _, resource := range page.Providers {
			r.resources = append(r.resources, *resource.ProviderName)
		}
		return true
	})
	return
}

func getCognitoUserPoolResourceServer(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListResourceServersPages(&cognitoidentityprovider.ListResourceServersInput{}, func(page *cognitoidentityprovider.ListResourceServersOutput, lastPage bool) bool {
		for _, resource := range page.ResourceServers {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getCognitoUserPoolUser(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListUsersPages(&cognitoidentityprovider.ListUsersInput{}, func(page *cognitoidentityprovider.ListUsersOutput, lastPage bool) bool {
		for _, resource := range page.Users {
			r.resources = append(r.resources, *resource.Username)
		}
		return true
	})
	return
}

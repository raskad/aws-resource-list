package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func getCognitoIdentityProvider(session *session.Session) (resources resourceMap) {
	client := cognitoidentityprovider.New(session)

	userPoolResourceMap := getCognitoUserPool(client).unwrap(cognitoUserPool)
	userPoolIDs := userPoolResourceMap[cognitoUserPool]

	resources = reduce(
		userPoolResourceMap,
		getCognitoUserPoolClient(client, userPoolIDs).unwrap(cognitoUserPoolClient),
		getCognitoUserPoolGroup(client, userPoolIDs).unwrap(cognitoUserPoolGroup),
		getCognitoUserPoolIdentityProvider(client, userPoolIDs).unwrap(cognitoUserPoolIdentityProvider),
		getCognitoUserPoolResourceServer(client, userPoolIDs).unwrap(cognitoUserPoolResourceServer),
		getCognitoUserPoolUser(client, userPoolIDs).unwrap(cognitoUserPoolUser),
	)
	return
}

func getCognitoUserPool(client *cognitoidentityprovider.CognitoIdentityProvider) (r resourceSliceError) {
	r.err = client.ListUserPoolsPages(&cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: aws.Int64(16),
	}, func(page *cognitoidentityprovider.ListUserPoolsOutput, lastPage bool) bool {
		for _, resource := range page.UserPools {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getCognitoUserPoolClient(client *cognitoidentityprovider.CognitoIdentityProvider, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		r.err = client.ListUserPoolClientsPages(&cognitoidentityprovider.ListUserPoolClientsInput{
			UserPoolId: aws.String(useruserPoolID),
		}, func(page *cognitoidentityprovider.ListUserPoolClientsOutput, lastPage bool) bool {
			for _, resource := range page.UserPoolClients {
				r.resources = append(r.resources, *resource.ClientName)
			}
			return true
		})
	}
	return
}

func getCognitoUserPoolGroup(client *cognitoidentityprovider.CognitoIdentityProvider, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		r.err = client.ListGroupsPages(&cognitoidentityprovider.ListGroupsInput{
			UserPoolId: aws.String(useruserPoolID),
		}, func(page *cognitoidentityprovider.ListGroupsOutput, lastPage bool) bool {
			for _, resource := range page.Groups {
				r.resources = append(r.resources, *resource.GroupName)
			}
			return true
		})
	}
	return
}

func getCognitoUserPoolIdentityProvider(client *cognitoidentityprovider.CognitoIdentityProvider, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		r.err = client.ListIdentityProvidersPages(&cognitoidentityprovider.ListIdentityProvidersInput{
			UserPoolId: aws.String(useruserPoolID),
		}, func(page *cognitoidentityprovider.ListIdentityProvidersOutput, lastPage bool) bool {
			for _, resource := range page.Providers {
				r.resources = append(r.resources, *resource.ProviderName)
			}
			return true
		})
	}
	return
}

func getCognitoUserPoolResourceServer(client *cognitoidentityprovider.CognitoIdentityProvider, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		r.err = client.ListResourceServersPages(&cognitoidentityprovider.ListResourceServersInput{
			UserPoolId: aws.String(useruserPoolID),
			MaxResults: aws.Int64(16),
		}, func(page *cognitoidentityprovider.ListResourceServersOutput, lastPage bool) bool {
			for _, resource := range page.ResourceServers {
				r.resources = append(r.resources, *resource.Name)
			}
			return true
		})
	}
	return
}

func getCognitoUserPoolUser(client *cognitoidentityprovider.CognitoIdentityProvider, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		r.err = client.ListUsersPages(&cognitoidentityprovider.ListUsersInput{
			UserPoolId: aws.String(useruserPoolID),
		}, func(page *cognitoidentityprovider.ListUsersOutput, lastPage bool) bool {
			for _, resource := range page.Users {
				r.resources = append(r.resources, *resource.Username)
			}
			return true
		})
	}
	return
}

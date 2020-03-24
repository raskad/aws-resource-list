package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func getCognitoIdentityProvider(config aws.Config) (resources resourceMap) {
	client := cognitoidentityprovider.New(config)

	cognitoUserPoolIDs := getCognitoUserPoolIDs(client)
	cognitoUserPoolClientNames := getCognitoUserPoolClientNames(client, cognitoUserPoolIDs)
	cognitoUserPoolGroupNames := getCognitoUserPoolGroupNames(client, cognitoUserPoolIDs)
	cognitoUserPoolIdentityProviderNames := getCognitoUserPoolIdentityProviderNames(client, cognitoUserPoolIDs)
	cognitoUserPoolResourceServerNames := getCognitoUserPoolResourceServerNames(client, cognitoUserPoolIDs)
	cognitoUserPoolUserNames := getCognitoUserPoolUserNames(client, cognitoUserPoolIDs)

	resources = resourceMap{
		cognitoUserPool:                 cognitoUserPoolIDs,
		cognitoUserPoolClient:           cognitoUserPoolClientNames,
		cognitoUserPoolGroup:            cognitoUserPoolGroupNames,
		cognitoUserPoolIdentityProvider: cognitoUserPoolIdentityProviderNames,
		cognitoUserPoolResourceServer:   cognitoUserPoolResourceServerNames,
		cognitoUserPoolUser:             cognitoUserPoolUserNames,
	}
	return
}

func getCognitoUserPoolIDs(client *cognitoidentityprovider.Client) (resources []string) {
	req := client.ListUserPoolsRequest(&cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: aws.Int64(16),
	})
	p := cognitoidentityprovider.NewListUserPoolsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.UserPools {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getCognitoUserPoolClientNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListUserPoolClientsRequest(&cognitoidentityprovider.ListUserPoolClientsInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListUserPoolClientsPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.UserPoolClients {
				resources = append(resources, *resource.ClientName)
			}
		}
	}
	return
}

func getCognitoUserPoolGroupNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListGroupsRequest(&cognitoidentityprovider.ListGroupsInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListGroupsPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.Groups {
				resources = append(resources, *resource.GroupName)
			}
		}
	}
	return
}

func getCognitoUserPoolIdentityProviderNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListIdentityProvidersRequest(&cognitoidentityprovider.ListIdentityProvidersInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListIdentityProvidersPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.Providers {
				resources = append(resources, *resource.ProviderName)
			}
		}
	}
	return
}

func getCognitoUserPoolResourceServerNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListResourceServersRequest(&cognitoidentityprovider.ListResourceServersInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListResourceServersPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.ResourceServers {
				resources = append(resources, *resource.Name)
			}
		}
	}
	return
}

func getCognitoUserPoolUserNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListUsersRequest(&cognitoidentityprovider.ListUsersInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListUsersPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.Users {
				resources = append(resources, *resource.Username)
			}
		}
	}
	return
}

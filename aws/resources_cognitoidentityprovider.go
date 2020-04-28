package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func getCognitoIdentityProvider(config aws.Config) (resources awsResourceMap) {
	client := cognitoidentityprovider.New(config)

	cognitoIdentityProviderUserPoolIDs := getCognitoIdentityProviderUserPoolIDs(client)
	cognitoIdentityProviderUserPoolClientNames := getCognitoIdentityProviderUserPoolClientNames(client, cognitoIdentityProviderUserPoolIDs)
	cognitoIdentityProviderUserPoolGroupNames := getCognitoIdentityProviderUserPoolGroupNames(client, cognitoIdentityProviderUserPoolIDs)
	cognitoIdentityProviderUserPoolIdentityProviderNames := getCognitoIdentityProviderUserPoolIdentityProviderNames(client, cognitoIdentityProviderUserPoolIDs)
	cognitoIdentityProviderUserPoolResourceServerNames := getCognitoIdentityProviderUserPoolResourceServerNames(client, cognitoIdentityProviderUserPoolIDs)
	cognitoIdentityProviderUserPoolUserNames := getCognitoIdentityProviderUserPoolUserNames(client, cognitoIdentityProviderUserPoolIDs)

	resources = awsResourceMap{
		cognitoIdentityProviderUserPool:                 cognitoIdentityProviderUserPoolIDs,
		cognitoIdentityProviderUserPoolClient:           cognitoIdentityProviderUserPoolClientNames,
		cognitoIdentityProviderUserPoolGroup:            cognitoIdentityProviderUserPoolGroupNames,
		cognitoIdentityProviderUserPoolIdentityProvider: cognitoIdentityProviderUserPoolIdentityProviderNames,
		cognitoIdentityProviderUserPoolResourceServer:   cognitoIdentityProviderUserPoolResourceServerNames,
		cognitoIdentityProviderUserPoolUser:             cognitoIdentityProviderUserPoolUserNames,
	}
	return
}

func getCognitoIdentityProviderUserPoolIDs(client *cognitoidentityprovider.Client) (resources []string) {
	req := client.ListUserPoolsRequest(&cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: aws.Int64(16),
	})
	p := cognitoidentityprovider.NewListUserPoolsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.UserPools {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getCognitoIdentityProviderUserPoolClientNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListUserPoolClientsRequest(&cognitoidentityprovider.ListUserPoolClientsInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListUserPoolClientsPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.UserPoolClients {
				resources = append(resources, *resource.ClientName)
			}
		}
	}
	return
}

func getCognitoIdentityProviderUserPoolGroupNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListGroupsRequest(&cognitoidentityprovider.ListGroupsInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListGroupsPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.Groups {
				resources = append(resources, *resource.GroupName)
			}
		}
	}
	return
}

func getCognitoIdentityProviderUserPoolIdentityProviderNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListIdentityProvidersRequest(&cognitoidentityprovider.ListIdentityProvidersInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListIdentityProvidersPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.Providers {
				resources = append(resources, *resource.ProviderName)
			}
		}
	}
	return
}

func getCognitoIdentityProviderUserPoolResourceServerNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListResourceServersRequest(&cognitoidentityprovider.ListResourceServersInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListResourceServersPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.ResourceServers {
				resources = append(resources, *resource.Name)
			}
		}
	}
	return
}

func getCognitoIdentityProviderUserPoolUserNames(client *cognitoidentityprovider.Client, userPoolIDs []string) (resources []string) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListUsersRequest(&cognitoidentityprovider.ListUsersInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListUsersPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.Users {
				resources = append(resources, *resource.Username)
			}
		}
	}
	return
}

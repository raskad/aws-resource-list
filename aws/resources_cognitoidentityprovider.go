package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func getCognitoIdentityProvider(config aws.Config) (resources resourceMap) {
	client := cognitoidentityprovider.New(config)

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

func getCognitoUserPool(client *cognitoidentityprovider.Client) (r resourceSliceError) {
	req := client.ListUserPoolsRequest(&cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: aws.Int64(16),
	})
	p := cognitoidentityprovider.NewListUserPoolsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.UserPools {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getCognitoUserPoolClient(client *cognitoidentityprovider.Client, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListUserPoolClientsRequest(&cognitoidentityprovider.ListUserPoolClientsInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListUserPoolClientsPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.UserPoolClients {
				r.resources = append(r.resources, *resource.ClientName)
			}
		}
		r.err = p.Err()
	}
	return
}

func getCognitoUserPoolGroup(client *cognitoidentityprovider.Client, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListGroupsRequest(&cognitoidentityprovider.ListGroupsInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListGroupsPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.Groups {
				r.resources = append(r.resources, *resource.GroupName)
			}
		}
		r.err = p.Err()
	}
	return
}

func getCognitoUserPoolIdentityProvider(client *cognitoidentityprovider.Client, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListIdentityProvidersRequest(&cognitoidentityprovider.ListIdentityProvidersInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListIdentityProvidersPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.Providers {
				r.resources = append(r.resources, *resource.ProviderName)
			}
		}
		r.err = p.Err()
	}
	return
}

func getCognitoUserPoolResourceServer(client *cognitoidentityprovider.Client, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListResourceServersRequest(&cognitoidentityprovider.ListResourceServersInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListResourceServersPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.ResourceServers {
				r.resources = append(r.resources, *resource.Name)
			}
		}
		r.err = p.Err()
	}
	return
}

func getCognitoUserPoolUser(client *cognitoidentityprovider.Client, userPoolIDs []string) (r resourceSliceError) {
	for _, useruserPoolID := range userPoolIDs {
		req := client.ListUsersRequest(&cognitoidentityprovider.ListUsersInput{
			UserPoolId: aws.String(useruserPoolID),
		})
		p := cognitoidentityprovider.NewListUsersPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.Users {
				r.resources = append(r.resources, *resource.Username)
			}
		}
		r.err = p.Err()
	}
	return
}

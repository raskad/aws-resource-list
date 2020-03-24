package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
)

func getCognitoIdentity(config aws.Config) (resources resourceMap) {
	client := cognitoidentity.New(config)

	cognitoIdentityPoolNames := getCognitoIdentityPoolNames(client)

	resources = resourceMap{
		cognitoIdentityPool: cognitoIdentityPoolNames,
	}
	return
}

func getCognitoIdentityPoolNames(client *cognitoidentity.Client) (resources []string) {
	input := cognitoidentity.ListIdentityPoolsInput{
		MaxResults: aws.Int64(16),
	}
	for {
		page, err := client.ListIdentityPoolsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.IdentityPools {
			resources = append(resources, *resource.IdentityPoolName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

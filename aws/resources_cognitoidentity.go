package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
)

func getCognitoIdentity(config aws.Config) (resources awsResourceMap) {
	client := cognitoidentity.New(config)

	cognitoIdentityPoolNames := getCognitoIdentityPoolNames(client)

	resources = awsResourceMap{
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
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.IdentityPools {
			resources = append(resources, *resource.IdentityPoolName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

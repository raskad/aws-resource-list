package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
)

func getCognitoIdentity(config aws.Config) (resources resourceMap) {
	client := cognitoidentity.New(config)
	resources = reduce(
		getCognitoIdentityPool(client).unwrap(cognitoIdentityPool),
	)
	return
}

func getCognitoIdentityPool(client *cognitoidentity.Client) (r resourceSliceError) {
	input := cognitoidentity.ListIdentityPoolsInput{
		MaxResults: aws.Int64(16),
	}
	for {
		page, err := client.ListIdentityPoolsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.IdentityPools {
			r.resources = append(r.resources, *resource.IdentityPoolName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

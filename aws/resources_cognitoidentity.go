package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
)

func getCognitoIdentity(session *session.Session) (resources resourceMap) {
	client := cognitoidentity.New(session)
	resources = reduce(
		getCognitoIdentityPool(client).unwrap(cognitoIdentityPool),
	)
	return
}

func getCognitoIdentityPool(client *cognitoidentity.CognitoIdentity) (r resourceSliceError) {
	input := cognitoidentity.ListIdentityPoolsInput{
		MaxResults: aws.Int64(16),
	}
	for {
		page, err := client.ListIdentityPools(&input)
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

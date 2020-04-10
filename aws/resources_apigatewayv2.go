package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

func getAPIGatewayV2(config aws.Config) (resources awsResourceMap) {
	client := apigatewayv2.New(config)

	apiGatewayV2APIIDs := getAPIGatewayV2APIIDs(client)
	apiGatewayV2DomainNames := getAPIGatewayV2DomainNames(client)

	resources = awsResourceMap{
		apiGatewayV2Api:        apiGatewayV2APIIDs,
		apiGatewayV2DomainName: apiGatewayV2DomainNames,
	}
	return
}

func getAPIGatewayV2APIIDs(client *apigatewayv2.Client) (resources []string) {
	input := apigatewayv2.GetApisInput{}
	for {
		page, err := client.GetApisRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Items {
			resources = append(resources, *resource.ApiId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAPIGatewayV2DomainNames(client *apigatewayv2.Client) (resources []string) {
	input := apigatewayv2.GetDomainNamesInput{}
	for {
		page, err := client.GetDomainNamesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Items {
			resources = append(resources, *resource.DomainName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

func getAPIGatewayV2(config aws.Config) (resources resourceMap) {
	client := apigatewayv2.New(config)
	resources = reduce(
		getAPIGatewayV2API(client).unwrap(apiGatewayV2Api),
		getAPIGatewayV2DomainName(client).unwrap(apiGatewayV2DomainName),
	)
	return
}

func getAPIGatewayV2API(client *apigatewayv2.Client) (r resourceSliceError) {
	input := apigatewayv2.GetApisInput{}
	for {
		page, err := client.GetApisRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.ApiId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAPIGatewayV2DomainName(client *apigatewayv2.Client) (r resourceSliceError) {
	input := apigatewayv2.GetDomainNamesInput{}
	for {
		page, err := client.GetDomainNamesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.DomainName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

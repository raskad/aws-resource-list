package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
)

func getAPIGatewayV2(session *session.Session) (resources resourceMap) {
	client := apigatewayv2.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		apiGatewayV2Api:        getAPIGatewayV2API(client),
		apiGatewayV2DomainName: getAPIGatewayV2DomainName(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAPIGatewayV2API(client *apigatewayv2.ApiGatewayV2) (r resourceSliceError) {
	logDebug("Listing ApiGatewayV2Api resources")
	input := apigatewayv2.GetApisInput{}
	for {
		page, err := client.GetApis(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Items {
			logDebug("Got ApiGatewayV2Api resource with PhysicalResourceId", *resource.ApiId)
			r.resources = append(r.resources, *resource.ApiId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAPIGatewayV2DomainName(client *apigatewayv2.ApiGatewayV2) (r resourceSliceError) {
	logDebug("Listing APIGatewayV2DomainName resources")
	input := apigatewayv2.GetDomainNamesInput{}
	for {
		page, err := client.GetDomainNames(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Items {
			logDebug("Got APIGatewayV2DomainName resource with PhysicalResourceId", *resource.DomainName)
			r.resources = append(r.resources, *resource.DomainName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

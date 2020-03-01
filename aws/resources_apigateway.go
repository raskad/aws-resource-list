package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func getAPIGateway(session *session.Session) (resources resourceMap) {
	client := apigateway.New(session)
	resources = reduce(
		getAPIGatewayAPIKey(client).unwrap(apiGatewayAPIKey),
		getAPIGatewayClientCertificate(client).unwrap(apiGatewayClientCertificate),
		getAPIGatewayDomainName(client).unwrap(apiGatewayDomainName),
		getAPIGatewayRestAPI(client).unwrap(apiGatewayRestAPI),
		getAPIGatewayUsagePlan(client).unwrap(apiGatewayUsagePlan),
		getAPIGatewayVpcLink(client).unwrap(apiGatewayVpcLink),
	)
	return
}

func getAPIGatewayAPIKey(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetApiKeysPages(&apigateway.GetApiKeysInput{}, func(page *apigateway.GetApiKeysOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getAPIGatewayClientCertificate(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetClientCertificatesPages(&apigateway.GetClientCertificatesInput{}, func(page *apigateway.GetClientCertificatesOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.ClientCertificateId)
		}
		return true
	})
	return
}

func getAPIGatewayDomainName(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetDomainNamesPages(&apigateway.GetDomainNamesInput{}, func(page *apigateway.GetDomainNamesOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.DomainName)
		}
		return true
	})
	return
}

func getAPIGatewayRestAPI(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetRestApisPages(&apigateway.GetRestApisInput{}, func(page *apigateway.GetRestApisOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getAPIGatewayUsagePlan(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetUsagePlansPages(&apigateway.GetUsagePlansInput{}, func(page *apigateway.GetUsagePlansOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getAPIGatewayVpcLink(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetVpcLinksPages(&apigateway.GetVpcLinksInput{}, func(page *apigateway.GetVpcLinksOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

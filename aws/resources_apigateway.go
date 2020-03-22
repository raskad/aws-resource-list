package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func getAPIGateway(config aws.Config) (resources resourceMap) {
	client := apigateway.New(config)
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

func getAPIGatewayAPIKey(client *apigateway.Client) (r resourceSliceError) {
	req := client.GetApiKeysRequest(&apigateway.GetApiKeysInput{})
	p := apigateway.NewGetApiKeysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getAPIGatewayClientCertificate(client *apigateway.Client) (r resourceSliceError) {
	req := client.GetClientCertificatesRequest(&apigateway.GetClientCertificatesInput{})
	p := apigateway.NewGetClientCertificatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.ClientCertificateId)
		}
	}
	r.err = p.Err()
	return
}

func getAPIGatewayDomainName(client *apigateway.Client) (r resourceSliceError) {
	req := client.GetDomainNamesRequest(&apigateway.GetDomainNamesInput{})
	p := apigateway.NewGetDomainNamesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.DomainName)
		}
	}
	r.err = p.Err()
	return
}

func getAPIGatewayRestAPI(client *apigateway.Client) (r resourceSliceError) {
	req := client.GetRestApisRequest(&apigateway.GetRestApisInput{})
	p := apigateway.NewGetRestApisPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getAPIGatewayUsagePlan(client *apigateway.Client) (r resourceSliceError) {
	req := client.GetUsagePlanKeysRequest(&apigateway.GetUsagePlanKeysInput{})
	p := apigateway.NewGetUsagePlanKeysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getAPIGatewayVpcLink(client *apigateway.Client) (r resourceSliceError) {
	req := client.GetVpcLinksRequest(&apigateway.GetVpcLinksInput{})
	p := apigateway.NewGetVpcLinksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

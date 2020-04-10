package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func getAPIGateway(config aws.Config) (resources awsResourceMap) {
	client := apigateway.New(config)

	apiGatewayAPIKeyIDs := getAPIGatewayAPIKeyIDs(client)
	apiGatewayClientCertificateIDs := getAPIGatewayClientCertificateIDs(client)
	apiGatewayDomainNames := getAPIGatewayDomainNames(client)
	apiGatewayRestAPIIDs := getAPIGatewayRestAPIIDs(client)
	apiGatewayUsagePlanIDs := getAPIGatewayUsagePlanIDs(client)
	apiGatewayVpcLinkIDs := getAPIGatewayVpcLinkIDs(client)

	resources = awsResourceMap{
		apiGatewayAPIKey:            apiGatewayAPIKeyIDs,
		apiGatewayClientCertificate: apiGatewayClientCertificateIDs,
		apiGatewayDomainName:        apiGatewayDomainNames,
		apiGatewayRestAPI:           apiGatewayRestAPIIDs,
		apiGatewayUsagePlan:         apiGatewayUsagePlanIDs,
		apiGatewayVpcLink:           apiGatewayVpcLinkIDs,
	}
	return
}

func getAPIGatewayAPIKeyIDs(client *apigateway.Client) (resources []string) {
	req := client.GetApiKeysRequest(&apigateway.GetApiKeysInput{})
	p := apigateway.NewGetApiKeysPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Items {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getAPIGatewayClientCertificateIDs(client *apigateway.Client) (resources []string) {
	req := client.GetClientCertificatesRequest(&apigateway.GetClientCertificatesInput{})
	p := apigateway.NewGetClientCertificatesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Items {
			resources = append(resources, *resource.ClientCertificateId)
		}
	}
	return
}

func getAPIGatewayDomainNames(client *apigateway.Client) (resources []string) {
	req := client.GetDomainNamesRequest(&apigateway.GetDomainNamesInput{})
	p := apigateway.NewGetDomainNamesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Items {
			resources = append(resources, *resource.DomainName)
		}
	}
	return
}

func getAPIGatewayRestAPIIDs(client *apigateway.Client) (resources []string) {
	req := client.GetRestApisRequest(&apigateway.GetRestApisInput{})
	p := apigateway.NewGetRestApisPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Items {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getAPIGatewayUsagePlanIDs(client *apigateway.Client) (resources []string) {
	req := client.GetUsagePlanKeysRequest(&apigateway.GetUsagePlanKeysInput{})
	p := apigateway.NewGetUsagePlanKeysPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Items {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getAPIGatewayVpcLinkIDs(client *apigateway.Client) (resources []string) {
	req := client.GetVpcLinksRequest(&apigateway.GetVpcLinksInput{})
	p := apigateway.NewGetVpcLinksPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Items {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

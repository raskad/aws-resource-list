package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func getAPIGateway(session *session.Session) (resources resourceMap) {
	client := apigateway.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		apiGatewayAPIKey:            getAPIGatewayAPIKey(client),
		apiGatewayClientCertificate: getAPIGatewayClientCertificate(client),
		apiGatewayDomainName:        getAPIGatewayDomainName(client),
		apiGatewayRestAPI:           getAPIGatewayRestAPI(client),
		apiGatewayUsagePlan:         getAPIGatewayUsagePlan(client),
		apiGatewayVpcLink:           getAPIGatewayVpcLink(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAPIGatewayAPIKey(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetApiKeysPages(&apigateway.GetApiKeysInput{}, func(page *apigateway.GetApiKeysOutput, lastPage bool) bool {
		logDebug("List APIGatewayAPIKey resources page. Remaining pages", page.Position)
		for _, resource := range page.Items {
			logDebug("Got APIGatewayAPIKey resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getAPIGatewayClientCertificate(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetClientCertificatesPages(&apigateway.GetClientCertificatesInput{}, func(page *apigateway.GetClientCertificatesOutput, lastPage bool) bool {
		logDebug("List APIGatewayClientCertificate resources page. Remaining pages", page.Position)
		for _, resource := range page.Items {
			logDebug("Got APIGatewayClientCertificate resource with PhysicalResourceId", *resource.ClientCertificateId)
			r.resources = append(r.resources, *resource.ClientCertificateId)
		}
		return true
	})
	return
}

func getAPIGatewayDomainName(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetDomainNamesPages(&apigateway.GetDomainNamesInput{}, func(page *apigateway.GetDomainNamesOutput, lastPage bool) bool {
		logDebug("List APIGatewayDomainName resources page. Remaining pages", page.Position)
		for _, resource := range page.Items {
			logDebug("Got APIGatewayDomainName resource with PhysicalResourceId", *resource.DomainName)
			r.resources = append(r.resources, *resource.DomainName)
		}
		return true
	})
	return
}

func getAPIGatewayRestAPI(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetRestApisPages(&apigateway.GetRestApisInput{}, func(page *apigateway.GetRestApisOutput, lastPage bool) bool {
		logDebug("List APIGatewayRestAPI resources page. Remaining pages", page.Position)
		for _, resource := range page.Items {
			logDebug("Got APIGatewayRestAPI resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getAPIGatewayUsagePlan(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetUsagePlansPages(&apigateway.GetUsagePlansInput{}, func(page *apigateway.GetUsagePlansOutput, lastPage bool) bool {
		logDebug("List APIGatewayUsagePlan resources page. Remaining pages", page.Position)
		for _, resource := range page.Items {
			logDebug("Got APIGatewayUsagePlan resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getAPIGatewayVpcLink(client *apigateway.APIGateway) (r resourceSliceError) {
	r.err = client.GetVpcLinksPages(&apigateway.GetVpcLinksInput{}, func(page *apigateway.GetVpcLinksOutput, lastPage bool) bool {
		logDebug("List APIGatewayVpcLink resources page. Remaining pages", page.Position)
		for _, resource := range page.Items {
			logDebug("Got APIGatewayVpcLink resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

func getServiceDiscovery(config aws.Config) (resources awsResourceMap) {
	client := servicediscovery.New(config)

	serviceDiscoveryHTTPNamespaceNames := getServiceDiscoveryHTTPNamespaceNames(client)
	serviceDiscoveryPrivateDNSNamespaceNames := getServiceDiscoveryPrivateDNSNamespaceNames(client)
	serviceDiscoveryPublicDNSNamespaceNames := getServiceDiscoveryPublicDNSNamespaceNames(client)
	serviceDiscoveryServiceIDs := getServiceDiscoveryServiceIDs(client)

	resources = awsResourceMap{
		serviceDiscoveryHTTPNamespace:       serviceDiscoveryHTTPNamespaceNames,
		serviceDiscoveryPrivateDNSNamespace: serviceDiscoveryPrivateDNSNamespaceNames,
		serviceDiscoveryPublicDNSNamespace:  serviceDiscoveryPublicDNSNamespaceNames,
		serviceDiscoveryService:             serviceDiscoveryServiceIDs,
	}
	return
}

func getServiceDiscoveryHTTPNamespaceNames(client *servicediscovery.Client) (resources []string) {
	req := client.ListNamespacesRequest(&servicediscovery.ListNamespacesInput{})
	p := servicediscovery.NewListNamespacesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Namespaces {
			if resource.Type == servicediscovery.NamespaceTypeHttp {
				resources = append(resources, *resource.Name)
			}
		}
	}
	return
}

func getServiceDiscoveryPrivateDNSNamespaceNames(client *servicediscovery.Client) (resources []string) {
	req := client.ListNamespacesRequest(&servicediscovery.ListNamespacesInput{})
	p := servicediscovery.NewListNamespacesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Namespaces {
			if resource.Type == servicediscovery.NamespaceTypeDnsPrivate {
				resources = append(resources, *resource.Name)
			}
		}
	}
	return
}

func getServiceDiscoveryPublicDNSNamespaceNames(client *servicediscovery.Client) (resources []string) {
	req := client.ListNamespacesRequest(&servicediscovery.ListNamespacesInput{})
	p := servicediscovery.NewListNamespacesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Namespaces {
			if resource.Type == servicediscovery.NamespaceTypeDnsPublic {
				resources = append(resources, *resource.Name)
			}
		}
	}
	return
}

func getServiceDiscoveryServiceIDs(client *servicediscovery.Client) (resources []string) {
	req := client.ListServicesRequest(&servicediscovery.ListServicesInput{})
	p := servicediscovery.NewListServicesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Services {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

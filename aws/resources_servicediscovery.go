package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

func getServiceDiscovery(config aws.Config) (resources resourceMap) {
	client := servicediscovery.New(config)

	resources = reduce(
		getServiceDiscoveryHTTPNamespace(client).unwrap(serviceDiscoveryHTTPNamespace),
		getServiceDiscoveryPrivateDNSNamespace(client).unwrap(serviceDiscoveryPrivateDNSNamespace),
		getServiceDiscoveryPublicDNSNamespace(client).unwrap(serviceDiscoveryPublicDNSNamespace),
		getServiceDiscoveryService(client).unwrap(serviceDiscoveryService),
	)
	return
}

func getServiceDiscoveryHTTPNamespace(client *servicediscovery.Client) (r resourceSliceError) {
	req := client.ListNamespacesRequest(&servicediscovery.ListNamespacesInput{})
	p := servicediscovery.NewListNamespacesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Namespaces {
			if resource.Type == servicediscovery.NamespaceTypeHttp {
				r.resources = append(r.resources, *resource.Name)
			}
		}
	}
	r.err = p.Err()
	return
}

func getServiceDiscoveryPrivateDNSNamespace(client *servicediscovery.Client) (r resourceSliceError) {
	req := client.ListNamespacesRequest(&servicediscovery.ListNamespacesInput{})
	p := servicediscovery.NewListNamespacesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Namespaces {
			if resource.Type == servicediscovery.NamespaceTypeDnsPrivate {
				r.resources = append(r.resources, *resource.Name)
			}
		}
	}
	r.err = p.Err()
	return
}

func getServiceDiscoveryPublicDNSNamespace(client *servicediscovery.Client) (r resourceSliceError) {
	req := client.ListNamespacesRequest(&servicediscovery.ListNamespacesInput{})
	p := servicediscovery.NewListNamespacesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Namespaces {
			if resource.Type == servicediscovery.NamespaceTypeDnsPublic {
				r.resources = append(r.resources, *resource.Name)
			}
		}
	}
	r.err = p.Err()
	return
}

func getServiceDiscoveryService(client *servicediscovery.Client) (r resourceSliceError) {
	req := client.ListServicesRequest(&servicediscovery.ListServicesInput{})
	p := servicediscovery.NewListServicesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Services {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

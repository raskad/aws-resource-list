package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
)

func getServiceDiscovery(session *session.Session) (resources resourceMap) {
	client := servicediscovery.New(session)

	serviceDiscoveryServiceResourceMap := getServiceDiscoveryService(client).unwrap(serviceDiscoveryService)
	serviceDiscoveryServiceIDs := serviceDiscoveryServiceResourceMap[serviceDiscoveryService]

	resources = reduce(
		getServiceDiscoveryHTTPNamespace(client).unwrap(serviceDiscoveryHTTPNamespace),
		getServiceDiscoveryInstance(client, serviceDiscoveryServiceIDs).unwrap(serviceDiscoveryInstance),
		getServiceDiscoveryPrivateDNSNamespace(client).unwrap(serviceDiscoveryPrivateDNSNamespace),
		getServiceDiscoveryPublicDNSNamespace(client).unwrap(serviceDiscoveryPublicDNSNamespace),
		serviceDiscoveryServiceResourceMap,
	)
	return
}

func getServiceDiscoveryHTTPNamespace(client *servicediscovery.ServiceDiscovery) (r resourceSliceError) {
	r.err = client.ListNamespacesPages(&servicediscovery.ListNamespacesInput{}, func(page *servicediscovery.ListNamespacesOutput, lastPage bool) bool {
		for _, resource := range page.Namespaces {
			if *resource.Type == "HTTP" {
				r.resources = append(r.resources, *resource.Name)
			}
		}
		return true
	})
	return
}

func getServiceDiscoveryInstance(client *servicediscovery.ServiceDiscovery, serviceIDs []string) (r resourceSliceError) {
	for _, serviceID := range serviceIDs {
		r.err = client.ListInstancesPages(&servicediscovery.ListInstancesInput{
			ServiceId: aws.String(serviceID),
		}, func(page *servicediscovery.ListInstancesOutput, lastPage bool) bool {
			for _, resource := range page.Instances {
				r.resources = append(r.resources, *resource.Id)
			}
			return true
		})
	}
	return
}

func getServiceDiscoveryPrivateDNSNamespace(client *servicediscovery.ServiceDiscovery) (r resourceSliceError) {
	r.err = client.ListNamespacesPages(&servicediscovery.ListNamespacesInput{}, func(page *servicediscovery.ListNamespacesOutput, lastPage bool) bool {
		for _, resource := range page.Namespaces {
			if *resource.Type == "DNS_PRIVATE" {
				r.resources = append(r.resources, *resource.Name)
			}
		}
		return true
	})
	return
}

func getServiceDiscoveryPublicDNSNamespace(client *servicediscovery.ServiceDiscovery) (r resourceSliceError) {
	r.err = client.ListNamespacesPages(&servicediscovery.ListNamespacesInput{}, func(page *servicediscovery.ListNamespacesOutput, lastPage bool) bool {
		for _, resource := range page.Namespaces {
			if *resource.Type == "DNS_PUBLIC" {
				r.resources = append(r.resources, *resource.Name)
			}
		}
		return true
	})
	return
}

func getServiceDiscoveryService(client *servicediscovery.ServiceDiscovery) (r resourceSliceError) {
	r.err = client.ListServicesPages(&servicediscovery.ListServicesInput{}, func(page *servicediscovery.ListServicesOutput, lastPage bool) bool {
		for _, resource := range page.Services {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

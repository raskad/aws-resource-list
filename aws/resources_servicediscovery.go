package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
)

func getServiceDiscovery(session *session.Session) (resources resourceMap) {
	client := servicediscovery.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		serviceDiscoveryHTTPNamespace:       getServiceDiscoveryHTTPNamespace(client),
		serviceDiscoveryInstance:            getServiceDiscoveryInstance(client),
		serviceDiscoveryPrivateDNSNamespace: getServiceDiscoveryPrivateDNSNamespace(client),
		serviceDiscoveryPublicDNSNamespace:  getServiceDiscoveryPublicDNSNamespace(client),
		serviceDiscoveryService:             getServiceDiscoveryService(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getServiceDiscoveryHTTPNamespace(client *servicediscovery.ServiceDiscovery) (r resourceSliceError) {
	logDebug("Listing ServiceDiscoveryHTTPNamespace resources")
	r.err = client.ListNamespacesPages(&servicediscovery.ListNamespacesInput{}, func(page *servicediscovery.ListNamespacesOutput, lastPage bool) bool {
		for _, resource := range page.Namespaces {
			if *resource.Type == "HTTP" {
				logDebug("Got ServiceDiscoveryHTTPNamespace resource with PhysicalResourceId", *resource.Name)
				r.resources = append(r.resources, *resource.Name)
			}
		}
		return true
	})
	return
}

func getServiceDiscoveryInstance(client *servicediscovery.ServiceDiscovery) (r resourceSliceError) {
	logDebug("Listing ServiceDiscoveryInstance resources")
	r.err = client.ListInstancesPages(&servicediscovery.ListInstancesInput{}, func(page *servicediscovery.ListInstancesOutput, lastPage bool) bool {
		for _, resource := range page.Instances {
			logDebug("Got ServiceDiscoveryInstance resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getServiceDiscoveryPrivateDNSNamespace(client *servicediscovery.ServiceDiscovery) (r resourceSliceError) {
	logDebug("Listing ServiceDiscoveryPrivateDNSNamespace resources")
	r.err = client.ListNamespacesPages(&servicediscovery.ListNamespacesInput{}, func(page *servicediscovery.ListNamespacesOutput, lastPage bool) bool {
		for _, resource := range page.Namespaces {
			if *resource.Type == "DNS_PRIVATE" {
				logDebug("Got ServiceDiscoveryPrivateDNSNamespace resource with PhysicalResourceId", *resource.Name)
				r.resources = append(r.resources, *resource.Name)
			}
		}
		return true
	})
	return
}

func getServiceDiscoveryPublicDNSNamespace(client *servicediscovery.ServiceDiscovery) (r resourceSliceError) {
	logDebug("Listing ServiceDiscoveryPublicDNSNamespace resources")
	r.err = client.ListNamespacesPages(&servicediscovery.ListNamespacesInput{}, func(page *servicediscovery.ListNamespacesOutput, lastPage bool) bool {
		for _, resource := range page.Namespaces {
			if *resource.Type == "DNS_PUBLIC" {
				logDebug("Got ServiceDiscoveryPublicDNSNamespace resource with PhysicalResourceId", *resource.Name)
				r.resources = append(r.resources, *resource.Name)
			}
		}
		return true
	})
	return
}

func getServiceDiscoveryService(client *servicediscovery.ServiceDiscovery) (r resourceSliceError) {
	logDebug("Listing ServiceDiscoveryService resources")
	r.err = client.ListServicesPages(&servicediscovery.ListServicesInput{}, func(page *servicediscovery.ListServicesOutput, lastPage bool) bool {
		for _, resource := range page.Services {
			logDebug("Got ServiceDiscoveryService resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

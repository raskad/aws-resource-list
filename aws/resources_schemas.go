package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/schemas"
)

func getSchemas(session *session.Session) (resources resourceMap) {
	client := schemas.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		eventSchemasDiscoverer: getEventSchemasDiscoverer(client),
		eventSchemasRegistry:   getEventSchemasRegistry(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getEventSchemasDiscoverer(client *schemas.Schemas) (r resourceSliceError) {
	r.err = client.ListDiscoverersPages(&schemas.ListDiscoverersInput{}, func(page *schemas.ListDiscoverersOutput, lastPage bool) bool {
		logDebug("Listing EventSchemasDiscoverer resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Discoverers {
			logDebug("Got EventSchemasDiscoverer resource with PhysicalResourceId", *resource.DiscovererId)
			r.resources = append(r.resources, *resource.DiscovererId)
		}
		return true
	})
	return
}

func getEventSchemasRegistry(client *schemas.Schemas) (r resourceSliceError) {
	r.err = client.ListRegistriesPages(&schemas.ListRegistriesInput{}, func(page *schemas.ListRegistriesOutput, lastPage bool) bool {
		logDebug("Listing EventSchemasRegistry resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Registries {
			logDebug("Got EventSchemasRegistry resource with PhysicalResourceId", *resource.RegistryName)
			r.resources = append(r.resources, *resource.RegistryName)
		}
		return true
	})
	return
}

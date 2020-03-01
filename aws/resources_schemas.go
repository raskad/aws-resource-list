package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/schemas"
)

func getSchemas(session *session.Session) (resources resourceMap) {
	client := schemas.New(session)
	resources = reduce(
		getEventSchemasDiscoverer(client).unwrap(eventSchemasDiscoverer),
		getEventSchemasRegistry(client).unwrap(eventSchemasRegistry),
	)
	return
}

func getEventSchemasDiscoverer(client *schemas.Schemas) (r resourceSliceError) {
	r.err = client.ListDiscoverersPages(&schemas.ListDiscoverersInput{}, func(page *schemas.ListDiscoverersOutput, lastPage bool) bool {
		for _, resource := range page.Discoverers {
			r.resources = append(r.resources, *resource.DiscovererId)
		}
		return true
	})
	return
}

func getEventSchemasRegistry(client *schemas.Schemas) (r resourceSliceError) {
	r.err = client.ListRegistriesPages(&schemas.ListRegistriesInput{}, func(page *schemas.ListRegistriesOutput, lastPage bool) bool {
		for _, resource := range page.Registries {
			r.resources = append(r.resources, *resource.RegistryName)
		}
		return true
	})
	return
}

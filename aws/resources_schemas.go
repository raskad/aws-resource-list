package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/schemas"
)

func getSchemas(config aws.Config) (resources resourceMap) {
	client := schemas.New(config)
	resources = reduce(
		getEventSchemasDiscoverer(client).unwrap(eventSchemasDiscoverer),
		getEventSchemasRegistry(client).unwrap(eventSchemasRegistry),
	)
	return
}

func getEventSchemasDiscoverer(client *schemas.Client) (r resourceSliceError) {
	req := client.ListDiscoverersRequest(&schemas.ListDiscoverersInput{})
	p := schemas.NewListDiscoverersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Discoverers {
			r.resources = append(r.resources, *resource.DiscovererId)
		}
	}
	r.err = p.Err()
	return
}

func getEventSchemasRegistry(client *schemas.Client) (r resourceSliceError) {
	req := client.ListRegistriesRequest(&schemas.ListRegistriesInput{})
	p := schemas.NewListRegistriesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Registries {
			r.resources = append(r.resources, *resource.RegistryName)
		}
	}
	r.err = p.Err()
	return
}

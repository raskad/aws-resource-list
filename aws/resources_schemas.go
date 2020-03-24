package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/schemas"
)

func getSchemas(config aws.Config) (resources resourceMap) {
	client := schemas.New(config)

	eventSchemasDiscovererIDs := getEventSchemasDiscovererIDs(client)
	eventSchemasRegistryNames := getEventSchemasRegistryNames(client)

	resources = resourceMap{
		eventSchemasDiscoverer: eventSchemasDiscovererIDs,
		eventSchemasRegistry:   eventSchemasRegistryNames,
	}
	return
}

func getEventSchemasDiscovererIDs(client *schemas.Client) (resources []string) {
	req := client.ListDiscoverersRequest(&schemas.ListDiscoverersInput{})
	p := schemas.NewListDiscoverersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Discoverers {
			resources = append(resources, *resource.DiscovererId)
		}
	}
	return
}

func getEventSchemasRegistryNames(client *schemas.Client) (resources []string) {
	req := client.ListRegistriesRequest(&schemas.ListRegistriesInput{})
	p := schemas.NewListRegistriesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Registries {
			resources = append(resources, *resource.RegistryName)
		}
	}
	return
}

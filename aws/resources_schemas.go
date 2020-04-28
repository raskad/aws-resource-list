package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/schemas"
)

func getSchemas(config aws.Config) (resources awsResourceMap) {
	client := schemas.New(config)

	schemasDiscovererIDs := getSchemasDiscovererIDs(client)
	schemasRegistryNames := getSchemasRegistryNames(client)

	resources = awsResourceMap{
		schemasDiscoverer: schemasDiscovererIDs,
		schemasRegistry:   schemasRegistryNames,
	}
	return
}

func getSchemasDiscovererIDs(client *schemas.Client) (resources []string) {
	req := client.ListDiscoverersRequest(&schemas.ListDiscoverersInput{})
	p := schemas.NewListDiscoverersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Discoverers {
			resources = append(resources, *resource.DiscovererId)
		}
	}
	return
}

func getSchemasRegistryNames(client *schemas.Client) (resources []string) {
	req := client.ListRegistriesRequest(&schemas.ListRegistriesInput{})
	p := schemas.NewListRegistriesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Registries {
			resources = append(resources, *resource.RegistryName)
		}
	}
	return
}

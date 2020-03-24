package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
)

func getMediaStore(config aws.Config) (resources resourceMap) {
	client := mediastore.New(config)

	mediaStoreContainerNames := getMediaStoreContainerNames(client)

	resources = resourceMap{
		mediaStoreContainer: mediaStoreContainerNames,
	}
	return
}

func getMediaStoreContainerNames(client *mediastore.Client) (resources []string) {
	req := client.ListContainersRequest(&mediastore.ListContainersInput{})
	p := mediastore.NewListContainersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Containers {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

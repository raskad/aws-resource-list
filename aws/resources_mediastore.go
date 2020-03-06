package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
)

func getMediaStore(config aws.Config) (resources resourceMap) {
	client := mediastore.New(config)
	resources = reduce(
		getMediaStoreContainer(client).unwrap(mediaStoreContainer),
	)
	return
}

func getMediaStoreContainer(client *mediastore.Client) (r resourceSliceError) {
	req := client.ListContainersRequest(&mediastore.ListContainersInput{})
	p := mediastore.NewListContainersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Containers {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediastore"
)

func getMediaStore(session *session.Session) (resources resourceMap) {
	client := mediastore.New(session)
	resources = reduce(
		getMediaStoreContainer(client).unwrap(mediaStoreContainer),
	)
	return
}

func getMediaStoreContainer(client *mediastore.MediaStore) (r resourceSliceError) {
	r.err = client.ListContainersPages(&mediastore.ListContainersInput{}, func(page *mediastore.ListContainersOutput, lastPage bool) bool {
		for _, resource := range page.Containers {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediastore"
)

func getMediaStore(session *session.Session) (resources resourceMap) {
	client := mediastore.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		mediaStoreContainer: getMediaStoreContainer(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getMediaStoreContainer(client *mediastore.MediaStore) (r resourceSliceError) {
	r.err = client.ListContainersPages(&mediastore.ListContainersInput{}, func(page *mediastore.ListContainersOutput, lastPage bool) bool {
		logDebug("Listing MediaStoreContainer resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Containers {
			logDebug("Got MediaStoreContainer resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

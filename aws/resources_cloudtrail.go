package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

func getCloudTrail(session *session.Session) (resources resourceMap) {
	client := cloudtrail.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		cloudTrailTrail: getCloudTrailTrail(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCloudTrailTrail(client *cloudtrail.CloudTrail) (r resourceSliceError) {
	logDebug("Listing CloudTrailTrail resources")
	r.err = client.ListTrailsPages(&cloudtrail.ListTrailsInput{}, func(page *cloudtrail.ListTrailsOutput, lastPage bool) bool {
		for _, resource := range page.Trails {
			logDebug("Got CloudTrailTrail resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

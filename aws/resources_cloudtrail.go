package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

func getCloudTrail(session *session.Session) (resources resourceMap) {
	client := cloudtrail.New(session)
	resources = reduce(
		getCloudTrailTrail(client).unwrap(cloudTrailTrail),
	)
	return
}

func getCloudTrailTrail(client *cloudtrail.CloudTrail) (r resourceSliceError) {
	r.err = client.ListTrailsPages(&cloudtrail.ListTrailsInput{}, func(page *cloudtrail.ListTrailsOutput, lastPage bool) bool {
		for _, resource := range page.Trails {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

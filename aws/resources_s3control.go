package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3control"
)

func getS3Control(session *session.Session) (resources resourceMap) {
	client := s3control.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		s3AccessPoint: getS3AccessPoint(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getS3AccessPoint(client *s3control.S3Control) (r resourceSliceError) {
	r.err = client.ListAccessPointsPages(&s3control.ListAccessPointsInput{AccountId: aws.String(accountID)}, func(page *s3control.ListAccessPointsOutput, lastPage bool) bool {
		logDebug("List S3AccessPoint resources page. Remaining pages", page.NextToken)
		for _, resource := range page.AccessPointList {
			logDebug("Got S3AccessPoint resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

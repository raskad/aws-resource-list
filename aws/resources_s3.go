package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getS3(session *session.Session) (resources resourceMap) {
	client := s3.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		s3Bucket:       getS3Bucket(client),
		s3BucketPolicy: getS3Bucket(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getS3Bucket(client *s3.S3) (r resourceSliceError) {
	logInfo("Start fetching S3Bucket resources")
	buckets, err := client.ListBuckets(&s3.ListBucketsInput{})
	for _, resource := range buckets.Buckets {
		logDebug("Got S3Bucket resource with PhysicalResourceId", *resource.Name)
		r.resources = append(r.resources, *resource.Name)
	}
	r.err = err
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getS3(session *session.Session) (resources resourceMap) {
	client := s3.New(session)
	resources = reduce(
		getS3Bucket(client).unwrap(s3Bucket),
		getS3Bucket(client).unwrap(s3BucketPolicy),
	)
	return
}

func getS3Bucket(client *s3.S3) (r resourceSliceError) {
	buckets, err := client.ListBuckets(&s3.ListBucketsInput{})
	for _, resource := range buckets.Buckets {
		r.resources = append(r.resources, *resource.Name)
	}
	r.err = err
	return
}

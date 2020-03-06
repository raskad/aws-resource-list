package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getS3(config aws.Config) (resources resourceMap) {
	client := s3.New(config)
	resources = reduce(
		getS3Bucket(client).unwrap(s3Bucket),
		getS3Bucket(client).unwrap(s3BucketPolicy),
	)
	return
}

func getS3Bucket(client *s3.Client) (r resourceSliceError) {
	buckets, err := client.ListBucketsRequest(&s3.ListBucketsInput{}).Send(context.Background())
	for _, resource := range buckets.Buckets {
		r.resources = append(r.resources, *resource.Name)
	}
	r.err = err
	return
}

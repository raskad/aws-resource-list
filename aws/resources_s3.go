package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getS3(config aws.Config) (resources resourceMap) {
	client := s3.New(config)

	s3BucketNames := getS3BucketNames(client)

	resources = resourceMap{
		s3Bucket: s3BucketNames,
	}
	return
}

func getS3BucketNames(client *s3.Client) (resources []string) {
	buckets, err := client.ListBucketsRequest(&s3.ListBucketsInput{}).Send(context.Background())
	logErr(err)
	for _, resource := range buckets.Buckets {
		resources = append(resources, *resource.Name)
	}
	return
}

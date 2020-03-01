package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

func getCloudfront(session *session.Session) (resources resourceMap) {
	client := cloudfront.New(session)
	resources = reduce(
		getCloudFrontCloudFrontOriginAccessIdentity(client).unwrap(cloudFrontCloudFrontOriginAccessIdentity),
		getCloudFrontDistribution(client).unwrap(cloudFrontDistribution),
		getCloudFrontStreamingDistribution(client).unwrap(cloudFrontStreamingDistribution),
	)
	return
}

func getCloudFrontCloudFrontOriginAccessIdentity(client *cloudfront.CloudFront) (r resourceSliceError) {
	r.err = client.ListCloudFrontOriginAccessIdentitiesPages(&cloudfront.ListCloudFrontOriginAccessIdentitiesInput{}, func(page *cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, lastPage bool) bool {
		for _, resource := range page.CloudFrontOriginAccessIdentityList.Items {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getCloudFrontDistribution(client *cloudfront.CloudFront) (r resourceSliceError) {
	r.err = client.ListDistributionsPages(&cloudfront.ListDistributionsInput{}, func(page *cloudfront.ListDistributionsOutput, lastPage bool) bool {
		for _, resource := range page.DistributionList.Items {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getCloudFrontStreamingDistribution(client *cloudfront.CloudFront) (r resourceSliceError) {
	r.err = client.ListStreamingDistributionsPages(&cloudfront.ListStreamingDistributionsInput{}, func(page *cloudfront.ListStreamingDistributionsOutput, lastPage bool) bool {
		for _, resource := range page.StreamingDistributionList.Items {
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

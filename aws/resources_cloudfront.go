package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

func getCloudfront(session *session.Session) (resources resourceMap) {
	client := cloudfront.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		cloudFrontCloudFrontOriginAccessIdentity: getCloudFrontCloudFrontOriginAccessIdentity(client),
		cloudFrontDistribution:                   getCloudFrontDistribution(client),
		cloudFrontStreamingDistribution:          getCloudFrontStreamingDistribution(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCloudFrontCloudFrontOriginAccessIdentity(client *cloudfront.CloudFront) (r resourceSliceError) {
	logDebug("Listing CloudFrontCloudFrontOriginAccessIdentity resources")
	r.err = client.ListCloudFrontOriginAccessIdentitiesPages(&cloudfront.ListCloudFrontOriginAccessIdentitiesInput{}, func(page *cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, lastPage bool) bool {
		for _, resource := range page.CloudFrontOriginAccessIdentityList.Items {
			logDebug("Got CloudFrontCloudFrontOriginAccessIdentity resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getCloudFrontDistribution(client *cloudfront.CloudFront) (r resourceSliceError) {
	logDebug("Listing CloudFrontDistribution resources")
	r.err = client.ListDistributionsPages(&cloudfront.ListDistributionsInput{}, func(page *cloudfront.ListDistributionsOutput, lastPage bool) bool {
		for _, resource := range page.DistributionList.Items {
			logDebug("Got CloudFrontDistribution resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

func getCloudFrontStreamingDistribution(client *cloudfront.CloudFront) (r resourceSliceError) {
	logDebug("Listing CloudFrontStreamingDistribution resources")
	r.err = client.ListStreamingDistributionsPages(&cloudfront.ListStreamingDistributionsInput{}, func(page *cloudfront.ListStreamingDistributionsOutput, lastPage bool) bool {
		for _, resource := range page.StreamingDistributionList.Items {
			logDebug("Got CloudFrontStreamingDistribution resource with PhysicalResourceId", *resource.Id)
			r.resources = append(r.resources, *resource.Id)
		}
		return true
	})
	return
}

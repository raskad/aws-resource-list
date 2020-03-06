package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func getCloudfront(config aws.Config) (resources resourceMap) {
	client := cloudfront.New(config)
	resources = reduce(
		getCloudFrontCloudFrontOriginAccessIdentity(client).unwrap(cloudFrontCloudFrontOriginAccessIdentity),
		getCloudFrontDistribution(client).unwrap(cloudFrontDistribution),
		getCloudFrontStreamingDistribution(client).unwrap(cloudFrontStreamingDistribution),
	)
	return
}

func getCloudFrontCloudFrontOriginAccessIdentity(client *cloudfront.Client) (r resourceSliceError) {
	req := client.ListCloudFrontOriginAccessIdentitiesRequest(&cloudfront.ListCloudFrontOriginAccessIdentitiesInput{})
	p := cloudfront.NewListCloudFrontOriginAccessIdentitiesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.CloudFrontOriginAccessIdentityList.Items {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getCloudFrontDistribution(client *cloudfront.Client) (r resourceSliceError) {
	req := client.ListDistributionsRequest(&cloudfront.ListDistributionsInput{})
	p := cloudfront.NewListDistributionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DistributionList.Items {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

func getCloudFrontStreamingDistribution(client *cloudfront.Client) (r resourceSliceError) {
	req := client.ListStreamingDistributionsRequest(&cloudfront.ListStreamingDistributionsInput{})
	p := cloudfront.NewListStreamingDistributionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.StreamingDistributionList.Items {
			r.resources = append(r.resources, *resource.Id)
		}
	}
	r.err = p.Err()
	return
}

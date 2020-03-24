package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func getCloudfront(config aws.Config) (resources resourceMap) {
	client := cloudfront.New(config)

	cloudFrontCloudFrontOriginAccessIdentityIDs := getCloudFrontCloudFrontOriginAccessIdentityIDs(client)
	cloudFrontDistributionIDs := getCloudFrontDistributionIDs(client)
	cloudFrontStreamingDistributionIDs := getCloudFrontStreamingDistributionIDs(client)

	resources = resourceMap{
		cloudFrontCloudFrontOriginAccessIdentity: cloudFrontCloudFrontOriginAccessIdentityIDs,
		cloudFrontDistribution:                   cloudFrontDistributionIDs,
		cloudFrontStreamingDistribution:          cloudFrontStreamingDistributionIDs,
	}
	return
}

func getCloudFrontCloudFrontOriginAccessIdentityIDs(client *cloudfront.Client) (resources []string) {
	req := client.ListCloudFrontOriginAccessIdentitiesRequest(&cloudfront.ListCloudFrontOriginAccessIdentitiesInput{})
	p := cloudfront.NewListCloudFrontOriginAccessIdentitiesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.CloudFrontOriginAccessIdentityList.Items {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getCloudFrontDistributionIDs(client *cloudfront.Client) (resources []string) {
	req := client.ListDistributionsRequest(&cloudfront.ListDistributionsInput{})
	p := cloudfront.NewListDistributionsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.DistributionList.Items {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getCloudFrontStreamingDistributionIDs(client *cloudfront.Client) (resources []string) {
	req := client.ListStreamingDistributionsRequest(&cloudfront.ListStreamingDistributionsInput{})
	p := cloudfront.NewListStreamingDistributionsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.StreamingDistributionList.Items {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

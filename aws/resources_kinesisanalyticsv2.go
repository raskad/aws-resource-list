package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
)

func getKinesisAnalyticsV2(config aws.Config) (resources awsResourceMap) {
	client := kinesisanalyticsv2.New(config)

	kinesisAnalyticsV2ApplicationNames := getKinesisAnalyticsV2ApplicationNames(client)

	resources = awsResourceMap{
		kinesisAnalyticsV2Application: kinesisAnalyticsV2ApplicationNames,
	}
	return
}

func getKinesisAnalyticsV2ApplicationNames(client *kinesisanalyticsv2.Client) (resources []string) {
	input := kinesisanalyticsv2.ListApplicationsInput{}
	for {
		page, err := client.ListApplicationsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ApplicationSummaries {
			resources = append(resources, *resource.ApplicationName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

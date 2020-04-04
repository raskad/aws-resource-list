package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
)

func getKinesisAnalytics(config aws.Config) (resources awsResourceMap) {
	client := kinesisanalytics.New(config)

	kinesisAnalyticsApplicationNames := getKinesisAnalyticsApplicationNames(client)

	resources = awsResourceMap{
		kinesisAnalyticsApplication: kinesisAnalyticsApplicationNames,
	}
	return
}

func getKinesisAnalyticsApplicationNames(client *kinesisanalytics.Client) (resources []string) {
	input := kinesisanalytics.ListApplicationsInput{}
	for {
		page, err := client.ListApplicationsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ApplicationSummaries {
			resources = append(resources, *resource.ApplicationName)
		}
		if !*page.HasMoreApplications {
			return
		}
		input.ExclusiveStartApplicationName = page.ApplicationSummaries[len(page.ApplicationSummaries)-1].ApplicationName
	}
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
)

func getKinesisAnalytics(config aws.Config) (resources resourceMap) {
	client := kinesisanalytics.New(config)
	resources = reduce(
		getKinesisAnalyticsApplication(client).unwrap(kinesisAnalyticsApplication),
	)
	return
}

func getKinesisAnalyticsApplication(client *kinesisanalytics.Client) (r resourceSliceError) {
	input := kinesisanalytics.ListApplicationsInput{}
	for {
		page, err := client.ListApplicationsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ApplicationSummaries {
			r.resources = append(r.resources, *resource.ApplicationName)
		}
		if !*page.HasMoreApplications {
			return
		}
		input.ExclusiveStartApplicationName = page.ApplicationSummaries[len(page.ApplicationSummaries)-1].ApplicationName
	}
}

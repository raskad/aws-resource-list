package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
)

func getKinesisAnalyticsV2(config aws.Config) (resources resourceMap) {
	client := kinesisanalyticsv2.New(config)
	resources = reduce(
		getKinesisAnalyticsV2Application(client).unwrap(kinesisAnalyticsV2Application),
	)
	return
}

func getKinesisAnalyticsV2Application(client *kinesisanalyticsv2.Client) (r resourceSliceError) {
	input := kinesisanalyticsv2.ListApplicationsInput{}
	for {
		page, err := client.ListApplicationsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ApplicationSummaries {
			r.resources = append(r.resources, *resource.ApplicationName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

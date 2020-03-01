package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
)

func getKinesisAnalyticsV2(session *session.Session) (resources resourceMap) {
	client := kinesisanalyticsv2.New(session)
	resources = reduce(
		getKinesisAnalyticsV2Application(client).unwrap(kinesisAnalyticsV2Application),
	)
	return
}

func getKinesisAnalyticsV2Application(client *kinesisanalyticsv2.KinesisAnalyticsV2) (r resourceSliceError) {
	input := kinesisanalyticsv2.ListApplicationsInput{}
	for {
		page, err := client.ListApplications(&input)
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

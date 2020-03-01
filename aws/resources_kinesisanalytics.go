package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
)

func getKinesisAnalytics(session *session.Session) (resources resourceMap) {
	client := kinesisanalytics.New(session)
	resources = reduce(
		getKinesisAnalyticsApplication(client).unwrap(kinesisAnalyticsApplication),
	)
	return
}

func getKinesisAnalyticsApplication(client *kinesisanalytics.KinesisAnalytics) (r resourceSliceError) {
	page, err := client.ListApplications(&kinesisanalytics.ListApplicationsInput{})
	for _, resource := range page.ApplicationSummaries {
		r.resources = append(r.resources, *resource.ApplicationName)
	}
	r.err = err
	return
}

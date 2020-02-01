package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
)

func getKinesisAnalytics(session *session.Session) (resources resourceMap) {
	client := kinesisanalytics.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		kinesisAnalyticsApplication: getKinesisAnalyticsApplication(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getKinesisAnalyticsApplication(client *kinesisanalytics.KinesisAnalytics) (r resourceSliceError) {
	logDebug("Listing KinesisAnalyticsApplication resources")
	page, err := client.ListApplications(&kinesisanalytics.ListApplicationsInput{})
	for _, resource := range page.ApplicationSummaries {
		logDebug("Got KinesisAnalyticsApplication resource with PhysicalResourceId", *resource.ApplicationName)
		r.resources = append(r.resources, *resource.ApplicationName)
	}
	r.err = err
	return
}

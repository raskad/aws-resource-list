package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
)

func getKinesisAnalyticsV2(session *session.Session) (resources resourceMap) {
	client := kinesisanalyticsv2.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		kinesisAnalyticsV2Application: getKinesisAnalyticsV2Application(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getKinesisAnalyticsV2Application(client *kinesisanalyticsv2.KinesisAnalyticsV2) (r resourceSliceError) {
	logDebug("Listing KinesisAnalyticsV2Application resources")
	input := kinesisanalyticsv2.ListApplicationsInput{}
	for {
		page, err := client.ListApplications(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ApplicationSummaries {
			logDebug("Got KinesisAnalyticsV2Application resource with PhysicalResourceId", *resource.ApplicationName)
			r.resources = append(r.resources, *resource.ApplicationName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

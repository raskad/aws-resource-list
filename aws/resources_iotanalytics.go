package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
)

func getIoTAnalytics(session *session.Session) (resources resourceMap) {
	client := iotanalytics.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		ioTAnalyticsChannel:   getIoTAnalyticsChannel(client),
		ioTAnalyticsDataset:   getIoTAnalyticsDataset(client),
		ioTAnalyticsDatastore: getIoTAnalyticsDatastore(client),
		ioTAnalyticsPipeline:  getIoTAnalyticsPipeline(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getIoTAnalyticsChannel(client *iotanalytics.IoTAnalytics) (r resourceSliceError) {
	logDebug("Listing IoTAnalyticsChannel resources")
	r.err = client.ListChannelsPages(&iotanalytics.ListChannelsInput{}, func(page *iotanalytics.ListChannelsOutput, lastPage bool) bool {
		for _, resource := range page.ChannelSummaries {
			logDebug("Got IoTAnalyticsChannel resource with PhysicalResourceId", *resource.ChannelName)
			r.resources = append(r.resources, *resource.ChannelName)
		}
		return true
	})
	return
}

func getIoTAnalyticsDataset(client *iotanalytics.IoTAnalytics) (r resourceSliceError) {
	logDebug("Listing IoTAnalyticsDataset resources")
	r.err = client.ListDatasetsPages(&iotanalytics.ListDatasetsInput{}, func(page *iotanalytics.ListDatasetsOutput, lastPage bool) bool {
		for _, resource := range page.DatasetSummaries {
			logDebug("Got IoTAnalyticsDataset resource with PhysicalResourceId", *resource.DatasetName)
			r.resources = append(r.resources, *resource.DatasetName)
		}
		return true
	})
	return
}

func getIoTAnalyticsDatastore(client *iotanalytics.IoTAnalytics) (r resourceSliceError) {
	logDebug("Listing IoTAnalyticsDatastore resources")
	r.err = client.ListDatastoresPages(&iotanalytics.ListDatastoresInput{}, func(page *iotanalytics.ListDatastoresOutput, lastPage bool) bool {
		for _, resource := range page.DatastoreSummaries {
			logDebug("Got IoTAnalyticsDatastore resource with PhysicalResourceId", *resource.DatastoreName)
			r.resources = append(r.resources, *resource.DatastoreName)
		}
		return true
	})
	return
}

func getIoTAnalyticsPipeline(client *iotanalytics.IoTAnalytics) (r resourceSliceError) {
	logDebug("Listing IoTAnalyticsPipeline resources")
	r.err = client.ListPipelinesPages(&iotanalytics.ListPipelinesInput{}, func(page *iotanalytics.ListPipelinesOutput, lastPage bool) bool {
		for _, resource := range page.PipelineSummaries {
			logDebug("Got IoTAnalyticsPipeline resource with PhysicalResourceId", *resource.PipelineName)
			r.resources = append(r.resources, *resource.PipelineName)
		}
		return true
	})
	return
}

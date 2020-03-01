package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
)

func getIoTAnalytics(session *session.Session) (resources resourceMap) {
	client := iotanalytics.New(session)
	resources = reduce(
		getIoTAnalyticsChannel(client).unwrap(ioTAnalyticsChannel),
		getIoTAnalyticsDataset(client).unwrap(ioTAnalyticsDataset),
		getIoTAnalyticsDatastore(client).unwrap(ioTAnalyticsDatastore),
		getIoTAnalyticsPipeline(client).unwrap(ioTAnalyticsPipeline),
	)
	return
}

func getIoTAnalyticsChannel(client *iotanalytics.IoTAnalytics) (r resourceSliceError) {
	r.err = client.ListChannelsPages(&iotanalytics.ListChannelsInput{}, func(page *iotanalytics.ListChannelsOutput, lastPage bool) bool {
		for _, resource := range page.ChannelSummaries {
			r.resources = append(r.resources, *resource.ChannelName)
		}
		return true
	})
	return
}

func getIoTAnalyticsDataset(client *iotanalytics.IoTAnalytics) (r resourceSliceError) {
	r.err = client.ListDatasetsPages(&iotanalytics.ListDatasetsInput{}, func(page *iotanalytics.ListDatasetsOutput, lastPage bool) bool {
		for _, resource := range page.DatasetSummaries {
			r.resources = append(r.resources, *resource.DatasetName)
		}
		return true
	})
	return
}

func getIoTAnalyticsDatastore(client *iotanalytics.IoTAnalytics) (r resourceSliceError) {
	r.err = client.ListDatastoresPages(&iotanalytics.ListDatastoresInput{}, func(page *iotanalytics.ListDatastoresOutput, lastPage bool) bool {
		for _, resource := range page.DatastoreSummaries {
			r.resources = append(r.resources, *resource.DatastoreName)
		}
		return true
	})
	return
}

func getIoTAnalyticsPipeline(client *iotanalytics.IoTAnalytics) (r resourceSliceError) {
	r.err = client.ListPipelinesPages(&iotanalytics.ListPipelinesInput{}, func(page *iotanalytics.ListPipelinesOutput, lastPage bool) bool {
		for _, resource := range page.PipelineSummaries {
			r.resources = append(r.resources, *resource.PipelineName)
		}
		return true
	})
	return
}

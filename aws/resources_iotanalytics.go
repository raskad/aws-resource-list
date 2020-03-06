package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics"
)

func getIoTAnalytics(config aws.Config) (resources resourceMap) {
	client := iotanalytics.New(config)
	resources = reduce(
		getIoTAnalyticsChannel(client).unwrap(ioTAnalyticsChannel),
		getIoTAnalyticsDataset(client).unwrap(ioTAnalyticsDataset),
		getIoTAnalyticsDatastore(client).unwrap(ioTAnalyticsDatastore),
		getIoTAnalyticsPipeline(client).unwrap(ioTAnalyticsPipeline),
	)
	return
}

func getIoTAnalyticsChannel(client *iotanalytics.Client) (r resourceSliceError) {
	req := client.ListChannelsRequest(&iotanalytics.ListChannelsInput{})
	p := iotanalytics.NewListChannelsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ChannelSummaries {
			r.resources = append(r.resources, *resource.ChannelName)
		}
	}
	r.err = p.Err()
	return
}

func getIoTAnalyticsDataset(client *iotanalytics.Client) (r resourceSliceError) {
	req := client.ListDatasetsRequest(&iotanalytics.ListDatasetsInput{})
	p := iotanalytics.NewListDatasetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DatasetSummaries {
			r.resources = append(r.resources, *resource.DatasetName)
		}
	}
	r.err = p.Err()
	return
}

func getIoTAnalyticsDatastore(client *iotanalytics.Client) (r resourceSliceError) {
	req := client.ListDatastoresRequest(&iotanalytics.ListDatastoresInput{})
	p := iotanalytics.NewListDatastoresPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DatastoreSummaries {
			r.resources = append(r.resources, *resource.DatastoreName)
		}
	}
	r.err = p.Err()
	return
}

func getIoTAnalyticsPipeline(client *iotanalytics.Client) (r resourceSliceError) {
	req := client.ListPipelinesRequest(&iotanalytics.ListPipelinesInput{})
	p := iotanalytics.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.PipelineSummaries {
			r.resources = append(r.resources, *resource.PipelineName)
		}
	}
	r.err = p.Err()
	return
}

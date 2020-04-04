package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics"
)

func getIoTAnalytics(config aws.Config) (resources awsResourceMap) {
	client := iotanalytics.New(config)

	ioTAnalyticsChannelNames := getIoTAnalyticsChannelNames(client)
	ioTAnalyticsDatasetNames := getIoTAnalyticsDatasetNames(client)
	ioTAnalyticsDatastoreNames := getIoTAnalyticsDatastoreNames(client)
	ioTAnalyticsPipelineNames := getIoTAnalyticsPipelineNames(client)

	resources = awsResourceMap{
		ioTAnalyticsChannel:   ioTAnalyticsChannelNames,
		ioTAnalyticsDataset:   ioTAnalyticsDatasetNames,
		ioTAnalyticsDatastore: ioTAnalyticsDatastoreNames,
		ioTAnalyticsPipeline:  ioTAnalyticsPipelineNames,
	}
	return
}

func getIoTAnalyticsChannelNames(client *iotanalytics.Client) (resources []string) {
	req := client.ListChannelsRequest(&iotanalytics.ListChannelsInput{})
	p := iotanalytics.NewListChannelsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ChannelSummaries {
			resources = append(resources, *resource.ChannelName)
		}
	}
	return
}

func getIoTAnalyticsDatasetNames(client *iotanalytics.Client) (resources []string) {
	req := client.ListDatasetsRequest(&iotanalytics.ListDatasetsInput{})
	p := iotanalytics.NewListDatasetsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DatasetSummaries {
			resources = append(resources, *resource.DatasetName)
		}
	}
	return
}

func getIoTAnalyticsDatastoreNames(client *iotanalytics.Client) (resources []string) {
	req := client.ListDatastoresRequest(&iotanalytics.ListDatastoresInput{})
	p := iotanalytics.NewListDatastoresPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DatastoreSummaries {
			resources = append(resources, *resource.DatastoreName)
		}
	}
	return
}

func getIoTAnalyticsPipelineNames(client *iotanalytics.Client) (resources []string) {
	req := client.ListPipelinesRequest(&iotanalytics.ListPipelinesInput{})
	p := iotanalytics.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.PipelineSummaries {
			resources = append(resources, *resource.PipelineName)
		}
	}
	return
}

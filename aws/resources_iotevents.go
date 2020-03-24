package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotevents"
)

func getIoTEvents(config aws.Config) (resources resourceMap) {
	client := iotevents.New(config)

	ioTEventsDetectorModelNames := getIoTEventsDetectorModelNames(client)
	ioTEventsInputNames := getIoTEventsInputNames(client)

	resources = resourceMap{
		ioTEventsDetectorModel: ioTEventsDetectorModelNames,
		ioTEventsInput:         ioTEventsInputNames,
	}
	return
}

func getIoTEventsDetectorModelNames(client *iotevents.Client) (resources []string) {
	input := iotevents.ListDetectorModelsInput{}
	for {
		page, err := client.ListDetectorModelsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.DetectorModelSummaries {
			resources = append(resources, *resource.DetectorModelName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getIoTEventsInputNames(client *iotevents.Client) (resources []string) {
	input := iotevents.ListInputsInput{}
	for {
		page, err := client.ListInputsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.InputSummaries {
			resources = append(resources, *resource.InputName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

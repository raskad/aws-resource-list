package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotevents"
)

func getIoTEvents(config aws.Config) (resources resourceMap) {
	client := iotevents.New(config)
	resources = reduce(
		getIoTEventsDetectorModel(client).unwrap(ioTEventsDetectorModel),
		getIoTEventsInput(client).unwrap(ioTEventsInput),
	)
	return
}

func getIoTEventsDetectorModel(client *iotevents.Client) (r resourceSliceError) {
	input := iotevents.ListDetectorModelsInput{}
	for {
		page, err := client.ListDetectorModelsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.DetectorModelSummaries {
			r.resources = append(r.resources, *resource.DetectorModelName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getIoTEventsInput(client *iotevents.Client) (r resourceSliceError) {
	input := iotevents.ListInputsInput{}
	for {
		page, err := client.ListInputsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.InputSummaries {
			r.resources = append(r.resources, *resource.InputName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

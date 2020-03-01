package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotevents"
)

func getIoTEvents(session *session.Session) (resources resourceMap) {
	client := iotevents.New(session)
	resources = reduce(
		getIoTEventsDetectorModel(client).unwrap(ioTEventsDetectorModel),
		getIoTEventsInput(client).unwrap(ioTEventsInput),
	)
	return
}

func getIoTEventsDetectorModel(client *iotevents.IoTEvents) (r resourceSliceError) {
	input := iotevents.ListDetectorModelsInput{}
	for {
		page, err := client.ListDetectorModels(&input)
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

func getIoTEventsInput(client *iotevents.IoTEvents) (r resourceSliceError) {
	input := iotevents.ListInputsInput{}
	for {
		page, err := client.ListInputs(&input)
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

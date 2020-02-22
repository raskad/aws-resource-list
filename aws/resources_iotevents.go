package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotevents"
)

func getIoTEvents(session *session.Session) (resources resourceMap) {
	client := iotevents.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		ioTEventsDetectorModel: getIoTEventsDetectorModel(client),
		ioTEventsInput:         getIoTEventsInput(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
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
		logDebug("Listing IoTEventsDetectorModel resources page. Remaining pages", page.NextToken)
		for _, resource := range page.DetectorModelSummaries {
			logDebug("Got IoTEventsDetectorModel resource with PhysicalResourceId", *resource.DetectorModelName)
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
		logDebug("Listing IoTEventsInput resources page. Remaining pages", page.NextToken)
		for _, resource := range page.InputSummaries {
			logDebug("Got IoTEventsInput resource with PhysicalResourceId", *resource.InputName)
			r.resources = append(r.resources, *resource.InputName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

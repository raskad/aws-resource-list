package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
)

func getIoT1ClickDevicesService(session *session.Session) (resources resourceMap) {
	client := iot1clickdevicesservice.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		ioT1ClickDevice: getIoT1ClickDevice(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getIoT1ClickDevice(client *iot1clickdevicesservice.IoT1ClickDevicesService) (r resourceSliceError) {
	input := iot1clickdevicesservice.ListDevicesInput{}
	for {
		page, err := client.ListDevices(&input)
		if err != nil {
			r.err = err
			return
		}
		logDebug("Listing IoT1ClickDevice resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Devices {
			logDebug("Got IoT1ClickDevice resource with PhysicalResourceId", *resource.DeviceId)
			r.resources = append(r.resources, *resource.DeviceId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

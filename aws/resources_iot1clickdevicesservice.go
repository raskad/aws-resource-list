package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
)

func getIoT1ClickDevicesService(session *session.Session) (resources resourceMap) {
	client := iot1clickdevicesservice.New(session)
	resources = reduce(
		getIoT1ClickDevice(client).unwrap(ioT1ClickDevice),
	)
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
		for _, resource := range page.Devices {
			r.resources = append(r.resources, *resource.DeviceId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

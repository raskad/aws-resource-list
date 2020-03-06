package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
)

func getIoT1ClickDevicesService(config aws.Config) (resources resourceMap) {
	client := iot1clickdevicesservice.New(config)
	resources = reduce(
		getIoT1ClickDevice(client).unwrap(ioT1ClickDevice),
	)
	return
}

func getIoT1ClickDevice(client *iot1clickdevicesservice.Client) (r resourceSliceError) {
	input := iot1clickdevicesservice.ListDevicesInput{}
	for {
		page, err := client.ListDevicesRequest(&input).Send(context.Background())
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

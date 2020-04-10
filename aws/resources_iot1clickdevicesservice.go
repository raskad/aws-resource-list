package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
)

func getIoT1ClickDevicesService(config aws.Config) (resources awsResourceMap) {
	client := iot1clickdevicesservice.New(config)

	ioT1ClickDeviceIDs := getIoT1ClickDeviceIDs(client)

	resources = awsResourceMap{
		ioT1ClickDevice: ioT1ClickDeviceIDs,
	}
	return
}

func getIoT1ClickDeviceIDs(client *iot1clickdevicesservice.Client) (resources []string) {
	input := iot1clickdevicesservice.ListDevicesInput{}
	for {
		page, err := client.ListDevicesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Devices {
			resources = append(resources, *resource.DeviceId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
)

func getGlobalAccelerator(config aws.Config) (resources awsResourceMap) {
	client := globalaccelerator.New(config)

	globalAcceleratorAcceleratorARNs := getGlobalAcceleratorAcceleratorARNs(client)
	globalAcceleratorListenerARNs := getGlobalAcceleratorListenerARNs(client, globalAcceleratorAcceleratorARNs)
	globalAcceleratorEndpointGroupARNs := getGlobalAcceleratorEndpointGroupARNs(client, globalAcceleratorListenerARNs)

	resources = awsResourceMap{
		globalAcceleratorAccelerator:   globalAcceleratorAcceleratorARNs,
		globalAcceleratorEndpointGroup: globalAcceleratorEndpointGroupARNs,
		globalAcceleratorListener:      globalAcceleratorListenerARNs,
	}
	return
}

func getGlobalAcceleratorAcceleratorARNs(client *globalaccelerator.Client) (resources []string) {
	input := globalaccelerator.ListAcceleratorsInput{}
	for {
		page, err := client.ListAcceleratorsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Accelerators {
			resources = append(resources, *resource.AcceleratorArn)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGlobalAcceleratorListenerARNs(client *globalaccelerator.Client, globalAcceleratorAcceleratorARNs []string) (resources []string) {
	for _, globalAcceleratorAcceleratorARN := range globalAcceleratorAcceleratorARNs {
		input := globalaccelerator.ListListenersInput{
			AcceleratorArn: aws.String(globalAcceleratorAcceleratorARN),
		}
		for {
			page, err := client.ListListenersRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.Listeners {
				resources = append(resources, *resource.ListenerArn)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getGlobalAcceleratorEndpointGroupARNs(client *globalaccelerator.Client, globalAcceleratorListenerARNs []string) (resources []string) {
	for _, globalAcceleratorListenerARN := range globalAcceleratorListenerARNs {
		input := globalaccelerator.ListEndpointGroupsInput{
			ListenerArn: aws.String(globalAcceleratorListenerARN),
		}
		for {
			page, err := client.ListEndpointGroupsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.EndpointGroups {
				resources = append(resources, *resource.EndpointGroupArn)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

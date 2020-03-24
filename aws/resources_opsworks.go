package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func getOpsWorks(config aws.Config) (resources resourceMap) {
	client := opsworks.New(config)

	opsWorksStackIDs := getOpsWorksStackIDs(client)
	opsWorksAppIDs := getOpsWorksAppIDs(client, opsWorksStackIDs)
	opsWorksInstanceIDs := getOpsWorksInstanceIDs(client, opsWorksStackIDs)
	opsWorksLayerIDs := getOpsWorksLayerIDs(client, opsWorksStackIDs)
	opsWorksUserProfileARNs := getOpsWorksUserProfileARNs(client)
	opsWorksVolumeIDs := getOpsWorksVolumeIDs(client, opsWorksStackIDs)

	resources = resourceMap{
		opsWorksApp:         opsWorksAppIDs,
		opsWorksStack:       opsWorksStackIDs,
		opsWorksInstance:    opsWorksInstanceIDs,
		opsWorksLayer:       opsWorksLayerIDs,
		opsWorksUserProfile: opsWorksUserProfileARNs,
		opsWorksVolume:      opsWorksVolumeIDs,
	}
	return
}

func getOpsWorksAppIDs(client *opsworks.Client, stackIDs []string) (resources []string) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeAppsRequest(&opsworks.DescribeAppsInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		logErr(err)
		for _, resource := range output.Apps {
			resources = append(resources, *resource.AppId)
		}
	}
	return
}

func getOpsWorksInstanceIDs(client *opsworks.Client, stackIDs []string) (resources []string) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeInstancesRequest(&opsworks.DescribeInstancesInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		logErr(err)
		for _, resource := range output.Instances {
			resources = append(resources, *resource.InstanceId)
		}
	}
	return
}

func getOpsWorksLayerIDs(client *opsworks.Client, stackIDs []string) (resources []string) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeLayersRequest(&opsworks.DescribeLayersInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		logErr(err)
		for _, resource := range output.Layers {
			resources = append(resources, *resource.LayerId)
		}
	}
	return
}

func getOpsWorksStackIDs(client *opsworks.Client) (resources []string) {
	output, err := client.DescribeStacksRequest(&opsworks.DescribeStacksInput{}).Send(context.Background())
	logErr(err)
	for _, resource := range output.Stacks {
		resources = append(resources, *resource.StackId)
	}
	return
}

func getOpsWorksUserProfileARNs(client *opsworks.Client) (resources []string) {
	output, err := client.DescribeUserProfilesRequest(&opsworks.DescribeUserProfilesInput{}).Send(context.Background())
	logErr(err)
	for _, resource := range output.UserProfiles {
		resources = append(resources, *resource.IamUserArn)
	}
	return
}

func getOpsWorksVolumeIDs(client *opsworks.Client, stackIDs []string) (resources []string) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeVolumesRequest(&opsworks.DescribeVolumesInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		logErr(err)
		for _, resource := range output.Volumes {
			resources = append(resources, *resource.VolumeId)
		}
	}
	return
}

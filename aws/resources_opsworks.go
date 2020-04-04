package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func getOpsWorks(config aws.Config) (resources awsResourceMap) {
	client := opsworks.New(config)

	opsWorksStackIDs := getOpsWorksStackIDs(client)
	opsWorksAppIDs := getOpsWorksAppIDs(client, opsWorksStackIDs)
	opsWorksInstanceIDs := getOpsWorksInstanceIDs(client, opsWorksStackIDs)
	opsWorksLayerIDs := getOpsWorksLayerIDs(client, opsWorksStackIDs)
	opsWorkRdsDbInstanceARNs := getOpsWorkRdsDbInstanceARNs(client, opsWorksStackIDs)
	opsWorksUserProfileARNs := getOpsWorksUserProfileARNs(client)
	opsWorksVolumeIDs := getOpsWorksVolumeIDs(client, opsWorksStackIDs)

	resources = awsResourceMap{
		opsWorksApp:          opsWorksAppIDs,
		opsWorksInstance:     opsWorksInstanceIDs,
		opsWorksLayer:        opsWorksLayerIDs,
		opsWorkRdsDbInstance: opsWorkRdsDbInstanceARNs,
		opsWorksStack:        opsWorksStackIDs,
		opsWorksUserProfile:  opsWorksUserProfileARNs,
		opsWorksVolume:       opsWorksVolumeIDs,
	}
	return
}

func getOpsWorksAppIDs(client *opsworks.Client, stackIDs []string) (resources []string) {
	for _, stackID := range stackIDs {
		page, err := client.DescribeAppsRequest(&opsworks.DescribeAppsInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Apps {
			resources = append(resources, *resource.AppId)
		}
	}
	return
}

func getOpsWorksInstanceIDs(client *opsworks.Client, stackIDs []string) (resources []string) {
	for _, stackID := range stackIDs {
		page, err := client.DescribeInstancesRequest(&opsworks.DescribeInstancesInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Instances {
			resources = append(resources, *resource.InstanceId)
		}
	}
	return
}

func getOpsWorksLayerIDs(client *opsworks.Client, stackIDs []string) (resources []string) {
	for _, stackID := range stackIDs {
		page, err := client.DescribeLayersRequest(&opsworks.DescribeLayersInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Layers {
			resources = append(resources, *resource.LayerId)
		}
	}
	return
}

func getOpsWorkRdsDbInstanceARNs(client *opsworks.Client, stackIDs []string) (resources []string) {
	for _, stackID := range stackIDs {
		page, err := client.DescribeRdsDbInstancesRequest(&opsworks.DescribeRdsDbInstancesInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.RdsDbInstances {
			resources = append(resources, *resource.RdsDbInstanceArn)
		}
	}
	return
}

func getOpsWorksStackIDs(client *opsworks.Client) (resources []string) {
	page, err := client.DescribeStacksRequest(&opsworks.DescribeStacksInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.Stacks {
		resources = append(resources, *resource.StackId)
	}
	return
}

func getOpsWorksUserProfileARNs(client *opsworks.Client) (resources []string) {
	page, err := client.DescribeUserProfilesRequest(&opsworks.DescribeUserProfilesInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.UserProfiles {
		resources = append(resources, *resource.IamUserArn)
	}
	return
}

func getOpsWorksVolumeIDs(client *opsworks.Client, stackIDs []string) (resources []string) {
	for _, stackID := range stackIDs {
		page, err := client.DescribeVolumesRequest(&opsworks.DescribeVolumesInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Volumes {
			resources = append(resources, *resource.VolumeId)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func getOpsWorks(config aws.Config) (resources resourceMap) {
	client := opsworks.New(config)

	opsWorksStackResourceMap := getOpsWorksStack(client).unwrap(opsWorksStack)
	opsWorksStackIDs := opsWorksStackResourceMap[opsWorksStack]

	resources = reduce(
		getOpsWorksApp(client, opsWorksStackIDs).unwrap(opsWorksApp),
		getOpsWorksInstance(client, opsWorksStackIDs).unwrap(opsWorksInstance),
		getOpsWorksLayer(client, opsWorksStackIDs).unwrap(opsWorksLayer),
		opsWorksStackResourceMap,
		getOpsWorksUserProfile(client).unwrap(opsWorksUserProfile),
		getOpsWorksVolume(client, opsWorksStackIDs).unwrap(opsWorksVolume),
	)
	return
}

func getOpsWorksApp(client *opsworks.Client, stackIDs []string) (r resourceSliceError) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeAppsRequest(&opsworks.DescribeAppsInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range output.Apps {
			r.resources = append(r.resources, *resource.AppId)
		}
	}
	return
}

func getOpsWorksInstance(client *opsworks.Client, stackIDs []string) (r resourceSliceError) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeInstancesRequest(&opsworks.DescribeInstancesInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range output.Instances {
			r.resources = append(r.resources, *resource.InstanceId)
		}
	}
	return
}

func getOpsWorksLayer(client *opsworks.Client, stackIDs []string) (r resourceSliceError) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeLayersRequest(&opsworks.DescribeLayersInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range output.Layers {
			r.resources = append(r.resources, *resource.LayerId)
		}
	}
	return
}

func getOpsWorksStack(client *opsworks.Client) (r resourceSliceError) {
	output, err := client.DescribeStacksRequest(&opsworks.DescribeStacksInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Stacks {
		r.resources = append(r.resources, *resource.StackId)
	}
	return
}

func getOpsWorksUserProfile(client *opsworks.Client) (r resourceSliceError) {
	output, err := client.DescribeUserProfilesRequest(&opsworks.DescribeUserProfilesInput{}).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.UserProfiles {
		r.resources = append(r.resources, *resource.IamUserArn)
	}
	return
}

func getOpsWorksVolume(client *opsworks.Client, stackIDs []string) (r resourceSliceError) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeVolumesRequest(&opsworks.DescribeVolumesInput{
			StackId: aws.String(stackID),
		}).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range output.Volumes {
			r.resources = append(r.resources, *resource.VolumeId)
		}
	}
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworks"
)

func getOpsWorks(session *session.Session) (resources resourceMap) {
	client := opsworks.New(session)

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

func getOpsWorksApp(client *opsworks.OpsWorks, stackIDs []string) (r resourceSliceError) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeApps(&opsworks.DescribeAppsInput{
			StackId: aws.String(stackID),
		})
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range output.Apps {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	return
}

func getOpsWorksInstance(client *opsworks.OpsWorks, stackIDs []string) (r resourceSliceError) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeInstances(&opsworks.DescribeInstancesInput{
			StackId: aws.String(stackID),
		})
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

func getOpsWorksLayer(client *opsworks.OpsWorks, stackIDs []string) (r resourceSliceError) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeLayers(&opsworks.DescribeLayersInput{
			StackId: aws.String(stackID),
		})
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

func getOpsWorksStack(client *opsworks.OpsWorks) (r resourceSliceError) {
	output, err := client.DescribeStacks(&opsworks.DescribeStacksInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Stacks {
		r.resources = append(r.resources, *resource.StackId)
	}
	return
}

func getOpsWorksUserProfile(client *opsworks.OpsWorks) (r resourceSliceError) {
	output, err := client.DescribeUserProfiles(&opsworks.DescribeUserProfilesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.UserProfiles {
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getOpsWorksVolume(client *opsworks.OpsWorks, stackIDs []string) (r resourceSliceError) {
	for _, stackID := range stackIDs {
		output, err := client.DescribeVolumes(&opsworks.DescribeVolumesInput{
			StackId: aws.String(stackID),
		})
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

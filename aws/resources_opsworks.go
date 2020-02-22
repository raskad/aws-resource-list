package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworks"
)

func getOpsWorks(session *session.Session) (resources resourceMap) {
	client := opsworks.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		opsWorksApp:         getOpsWorksApp(client),
		opsWorksInstance:    getOpsWorksInstance(client),
		opsWorksLayer:       getOpsWorksLayer(client),
		opsWorksStack:       getOpsWorksStack(client),
		opsWorksUserProfile: getOpsWorksUserProfile(client),
		opsWorksVolume:      getOpsWorksVolume(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getOpsWorksApp(client *opsworks.OpsWorks) (r resourceSliceError) {
	logInfo("Start fetching OpsWorksApp resources")
	output, err := client.DescribeApps(&opsworks.DescribeAppsInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Apps {
		logDebug("Got OpsWorksApp resource with PhysicalResourceId", *resource.Name)
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getOpsWorksInstance(client *opsworks.OpsWorks) (r resourceSliceError) {
	logInfo("Start fetching OpsWorksInstance resources")
	output, err := client.DescribeInstances(&opsworks.DescribeInstancesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Instances {
		logDebug("Got OpsWorksInstance resource with PhysicalResourceId", *resource.InstanceId)
		r.resources = append(r.resources, *resource.InstanceId)
	}
	return
}

func getOpsWorksLayer(client *opsworks.OpsWorks) (r resourceSliceError) {
	logInfo("Start fetching OpsWorksLayer resources")
	output, err := client.DescribeLayers(&opsworks.DescribeLayersInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Layers {
		logDebug("Got OpsWorksLayer resource with PhysicalResourceId", *resource.LayerId)
		r.resources = append(r.resources, *resource.LayerId)
	}
	return
}

func getOpsWorksStack(client *opsworks.OpsWorks) (r resourceSliceError) {
	logInfo("Start fetching OpsWorksStack resources")
	output, err := client.DescribeStacks(&opsworks.DescribeStacksInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Stacks {
		logDebug("Got OpsWorksStack resource with PhysicalResourceId", *resource.StackId)
		r.resources = append(r.resources, *resource.StackId)
	}
	return
}

func getOpsWorksUserProfile(client *opsworks.OpsWorks) (r resourceSliceError) {
	logInfo("Start fetching OpsWorksUserProfile resources")
	output, err := client.DescribeUserProfiles(&opsworks.DescribeUserProfilesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.UserProfiles {
		logDebug("Got OpsWorksUserProfile resource with PhysicalResourceId", *resource.Name)
		r.resources = append(r.resources, *resource.Name)
	}
	return
}

func getOpsWorksVolume(client *opsworks.OpsWorks) (r resourceSliceError) {
	logInfo("Start fetching OpsWorksVolume resources")
	output, err := client.DescribeVolumes(&opsworks.DescribeVolumesInput{})
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range output.Volumes {
		logDebug("Got OpsWorksVolume resource with PhysicalResourceId", *resource.VolumeId)
		r.resources = append(r.resources, *resource.VolumeId)
	}
	return
}

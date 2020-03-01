package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func getCloudForamtionState(session *session.Session) (resources resourceMap, err error) {
	client := cloudformation.New(session)
	stackIDs, err := getCloudformationActiveStackIDs(client)
	if err != nil {
		return resources, err
	}
	resourcesCloudFormation, err := getCloudformationResources(client, stackIDs)
	if err != nil {
		return resources, err
	}
	resources = resourceMap{}
	for cloudFormationType := range resourcesCloudFormation {
		resourceType, ok := fromCloudFormationType(cloudFormationType)
		if ok {
			resources[resourceType] = resourcesCloudFormation[cloudFormationType]
		}
	}
	return resources, nil
}

func getCloudformationActiveStackIDs(client *cloudformation.CloudFormation) (stackIDs []*string, err error) {
	logDebug("Listing CloudformationActiveStackID resources")
	err = client.ListStacksPages(&cloudformation.ListStacksInput{
		StackStatusFilter: aws.StringSlice([]string{
			"CREATE_FAILED",
			"CREATE_COMPLETE",
			"ROLLBACK_IN_PROGRESS",
			"ROLLBACK_FAILED",
			"ROLLBACK_COMPLETE",
			"DELETE_IN_PROGRESS",
			"DELETE_FAILED",
			"UPDATE_IN_PROGRESS",
			"UPDATE_COMPLETE_CLEANUP_IN_PROGRESS",
			"UPDATE_COMPLETE",
			"UPDATE_ROLLBACK_IN_PROGRESS",
			"UPDATE_ROLLBACK_FAILED",
			"UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS",
			"UPDATE_ROLLBACK_COMPLETE",
			"REVIEW_IN_PROGRESS",
			"IMPORT_IN_PROGRESS",
			"IMPORT_COMPLETE",
			"IMPORT_ROLLBACK_IN_PROGRESS",
			"IMPORT_ROLLBACK_FAILED",
			"IMPORT_ROLLBACK_COMPLETE",
		})},
		func(page *cloudformation.ListStacksOutput, lastPage bool) bool {
			for _, resource := range page.StackSummaries {
				logDebug("Got CloudformationActiveStackID resource with PhysicalResourceId", *resource.StackId)
				stackIDs = append(stackIDs, (resource.StackId))
			}
			return true
		})
	if err != nil {
		return stackIDs, err
	}
	return stackIDs, nil
}

func getCloudformationResources(client *cloudformation.CloudFormation, stackIDs []*string) (resources map[string][]string, err error) {
	logDebug("Listing CloudformationResources resources")
	resources = make(map[string][]string)
	for _, stackID := range stackIDs {
		err := client.ListStackResourcesPages(&cloudformation.ListStackResourcesInput{
			StackName: stackID,
		},
			func(page *cloudformation.ListStackResourcesOutput, lastPage bool) bool {
				for _, resource := range page.StackResourceSummaries {
					logDebug("Got cloudformation resource with ResourceType", *resource.ResourceType, "and PhysicalResourceId", *resource.PhysicalResourceId)
					resources[*resource.ResourceType] = append(resources[*resource.ResourceType], *resource.PhysicalResourceId)
				}
				return true
			})
		if err != nil {
			return resources, err
		}
	}
	return resources, nil
}

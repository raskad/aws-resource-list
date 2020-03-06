package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func getCloudForamtionState(config aws.Config) (resources resourceMap, err error) {
	client := cloudformation.New(config)
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

func getCloudformationActiveStackIDs(client *cloudformation.Client) (stackIDs []*string, err error) {
	logDebug("Listing CloudformationActiveStackID resources")
	req := client.ListStacksRequest(&cloudformation.ListStacksInput{
		StackStatusFilter: []cloudformation.StackStatus{
			cloudformation.StackStatusCreateInProgress,
			cloudformation.StackStatusCreateComplete,
			cloudformation.StackStatusRollbackInProgress,
			cloudformation.StackStatusRollbackFailed,
			cloudformation.StackStatusRollbackComplete,
			cloudformation.StackStatusDeleteInProgress,
			cloudformation.StackStatusDeleteFailed,
			cloudformation.StackStatusUpdateInProgress,
			cloudformation.StackStatusUpdateCompleteCleanupInProgress,
			cloudformation.StackStatusUpdateComplete,
			cloudformation.StackStatusUpdateRollbackInProgress,
			cloudformation.StackStatusUpdateRollbackFailed,
			cloudformation.StackStatusUpdateRollbackCompleteCleanupInProgress,
			cloudformation.StackStatusUpdateRollbackComplete,
			cloudformation.StackStatusReviewInProgress,
			cloudformation.StackStatusImportInProgress,
			cloudformation.StackStatusImportComplete,
			cloudformation.StackStatusImportRollbackInProgress,
			cloudformation.StackStatusImportRollbackFailed,
			cloudformation.StackStatusImportRollbackComplete,
		},
	})
	p := cloudformation.NewListStacksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.StackSummaries {
			logDebug("Got CloudformationActiveStackID resource with PhysicalResourceId", *resource.StackId)
			stackIDs = append(stackIDs, (resource.StackId))
		}
	}
	if err = p.Err(); err != nil {
		return stackIDs, err
	}
	return stackIDs, nil
}

func getCloudformationResources(client *cloudformation.Client, stackIDs []*string) (resources map[string][]string, err error) {
	logDebug("Listing CloudformationResources resources")
	resources = make(map[string][]string)

	for _, stackID := range stackIDs {
		req := client.ListStackResourcesRequest(&cloudformation.ListStackResourcesInput{
			StackName: stackID,
		})
		p := cloudformation.NewListStackResourcesPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.StackResourceSummaries {
				logDebug("Got cloudformation resource with ResourceType", *resource.ResourceType, "and PhysicalResourceId", *resource.PhysicalResourceId)
				resources[*resource.ResourceType] = append(resources[*resource.ResourceType], *resource.PhysicalResourceId)
			}
		}
		if err = p.Err(); err != nil {
			return resources, err
		}
	}
	return resources, nil
}

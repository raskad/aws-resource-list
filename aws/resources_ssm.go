package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func getSsm(config aws.Config) (resources awsResourceMap) {
	client := ssm.New(config)

	ssmActivationIDs := getSsmActivationIDs(client)
	ssmMaintenanceWindowIDs := getSsmMaintenanceWindowIDs(client)
	ssmAssociationIDs := getSsmAssociationIDs(client)
	ssmDocumentNames := getSsmDocumentNames(client)
	ssmMaintenanceWindowTargetIDs := getSsmMaintenanceWindowTargetIDs(client, ssmMaintenanceWindowIDs)
	ssmMaintenanceWindowTaskIDs := getSsmMaintenanceWindowTaskIDs(client, ssmMaintenanceWindowIDs)
	ssmParameterNames := getSsmParameterNames(client)
	ssmPatchBaselineIDs := getSsmPatchBaselineIDs(client)
	ssmPatchGroupNames := getSsmPatchGroupNames(client)
	ssmResourceDataSyncNames := getSsmResourceDataSyncNames(client)

	resources = awsResourceMap{
		ssmActivation:              ssmActivationIDs,
		ssmAssociation:             ssmAssociationIDs,
		ssmDocument:                ssmDocumentNames,
		ssmMaintenanceWindow:       ssmMaintenanceWindowIDs,
		ssmMaintenanceWindowTarget: ssmMaintenanceWindowTargetIDs,
		ssmMaintenanceWindowTask:   ssmMaintenanceWindowTaskIDs,
		ssmParameter:               ssmParameterNames,
		ssmPatchBaseline:           ssmPatchBaselineIDs,
		ssmPatchGroup:              ssmPatchGroupNames,
		ssmResourceDataSync:        ssmResourceDataSyncNames,
	}
	return
}

func getSsmActivationIDs(client *ssm.Client) (resources []string) {
	req := client.DescribeActivationsRequest(&ssm.DescribeActivationsInput{})
	p := ssm.NewDescribeActivationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ActivationList {
			resources = append(resources, *resource.ActivationId)
		}
	}
	return
}

func getSsmAssociationIDs(client *ssm.Client) (resources []string) {
	req := client.ListAssociationsRequest(&ssm.ListAssociationsInput{})
	p := ssm.NewListAssociationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Associations {
			resources = append(resources, *resource.AssociationId)
		}
	}
	return
}

func getSsmDocumentNames(client *ssm.Client) (resources []string) {
	req := client.ListDocumentsRequest(&ssm.ListDocumentsInput{})
	p := ssm.NewListDocumentsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DocumentIdentifiers {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getSsmMaintenanceWindowIDs(client *ssm.Client) (resources []string) {
	input := ssm.DescribeMaintenanceWindowsInput{}
	for {
		page, err := client.DescribeMaintenanceWindowsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.WindowIdentities {
			resources = append(resources, *resource.WindowId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSsmMaintenanceWindowTargetIDs(client *ssm.Client, windowIDs []string) (resources []string) {
	for _, windowID := range windowIDs {
		input := ssm.DescribeMaintenanceWindowTargetsInput{
			WindowId: aws.String(windowID),
		}
		for {
			page, err := client.DescribeMaintenanceWindowTargetsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.Targets {
				resources = append(resources, *resource.WindowTargetId)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getSsmMaintenanceWindowTaskIDs(client *ssm.Client, windowIDs []string) (resources []string) {
	for _, windowID := range windowIDs {
		input := ssm.DescribeMaintenanceWindowTasksInput{
			WindowId: aws.String(windowID),
		}
		for {
			page, err := client.DescribeMaintenanceWindowTasksRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.Tasks {
				resources = append(resources, *resource.WindowTaskId)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getSsmParameterNames(client *ssm.Client) (resources []string) {
	req := client.DescribeParametersRequest(&ssm.DescribeParametersInput{})
	p := ssm.NewDescribeParametersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Parameters {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getSsmPatchBaselineIDs(client *ssm.Client) (resources []string) {
	input := ssm.DescribePatchBaselinesInput{}
	for {
		page, err := client.DescribePatchBaselinesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.BaselineIdentities {
			resources = append(resources, *resource.BaselineId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSsmPatchGroupNames(client *ssm.Client) (resources []string) {
	input := ssm.DescribePatchGroupsInput{}
	for {
		page, err := client.DescribePatchGroupsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Mappings {
			resources = append(resources, *resource.PatchGroup)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSsmResourceDataSyncNames(client *ssm.Client) (resources []string) {
	input := ssm.ListResourceDataSyncInput{}
	for {
		page, err := client.ListResourceDataSyncRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ResourceDataSyncItems {
			resources = append(resources, *resource.SyncName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func getSsm(config aws.Config) (resources resourceMap) {
	client := ssm.New(config)

	ssmMaintenanceWindowResourceMap := getSsmMaintenanceWindow(client).unwrap(ssmMaintenanceWindow)
	ssmMaintenanceWindowIDs := ssmMaintenanceWindowResourceMap[ssmMaintenanceWindow]

	resources = reduce(
		getSsmAssociation(client).unwrap(ssmAssociation),
		getSsmDocument(client).unwrap(ssmDocument),
		ssmMaintenanceWindowResourceMap,
		getSsmMaintenanceWindowTarget(client, ssmMaintenanceWindowIDs).unwrap(ssmMaintenanceWindowTarget),
		getSsmMaintenanceWindowTask(client, ssmMaintenanceWindowIDs).unwrap(ssmMaintenanceWindowTask),
		getSsmParameter(client).unwrap(ssmParameter),
		getSsmPatchBaseline(client).unwrap(ssmPatchBaseline),
		getSsmResourceDataSync(client).unwrap(ssmResourceDataSync),
	)
	return
}

func getSsmAssociation(client *ssm.Client) (r resourceSliceError) {
	req := client.ListAssociationsRequest(&ssm.ListAssociationsInput{})
	p := ssm.NewListAssociationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Associations {
			r.resources = append(r.resources, *resource.AssociationId)
		}
	}
	r.err = p.Err()
	return
}

func getSsmDocument(client *ssm.Client) (r resourceSliceError) {
	req := client.ListDocumentsRequest(&ssm.ListDocumentsInput{})
	p := ssm.NewListDocumentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DocumentIdentifiers {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getSsmMaintenanceWindow(client *ssm.Client) (r resourceSliceError) {
	input := ssm.DescribeMaintenanceWindowsInput{}
	for {
		page, err := client.DescribeMaintenanceWindowsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.WindowIdentities {
			r.resources = append(r.resources, *resource.WindowId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSsmMaintenanceWindowTarget(client *ssm.Client, windowIDs []string) (r resourceSliceError) {
	for _, windowID := range windowIDs {
		input := ssm.DescribeMaintenanceWindowTargetsInput{
			WindowId: aws.String(windowID),
		}
		for {
			page, err := client.DescribeMaintenanceWindowTargetsRequest(&input).Send(context.Background())
			if err != nil {
				r.err = err
				return
			}
			for _, resource := range page.Targets {
				r.resources = append(r.resources, *resource.WindowTargetId)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getSsmMaintenanceWindowTask(client *ssm.Client, windowIDs []string) (r resourceSliceError) {
	for _, windowID := range windowIDs {
		input := ssm.DescribeMaintenanceWindowTasksInput{
			WindowId: aws.String(windowID),
		}
		for {
			page, err := client.DescribeMaintenanceWindowTasksRequest(&input).Send(context.Background())
			if err != nil {
				r.err = err
				return
			}
			for _, resource := range page.Tasks {
				r.resources = append(r.resources, *resource.WindowTaskId)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getSsmParameter(client *ssm.Client) (r resourceSliceError) {
	req := client.DescribeParametersRequest(&ssm.DescribeParametersInput{})
	p := ssm.NewDescribeParametersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Parameters {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getSsmPatchBaseline(client *ssm.Client) (r resourceSliceError) {
	input := ssm.DescribePatchBaselinesInput{}
	for {
		page, err := client.DescribePatchBaselinesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.BaselineIdentities {
			r.resources = append(r.resources, *resource.BaselineId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSsmResourceDataSync(client *ssm.Client) (r resourceSliceError) {
	input := ssm.ListResourceDataSyncInput{}
	for {
		page, err := client.ListResourceDataSyncRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ResourceDataSyncItems {
			r.resources = append(r.resources, *resource.SyncName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

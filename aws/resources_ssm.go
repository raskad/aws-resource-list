package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func getSsm(session *session.Session) (resources resourceMap) {
	client := ssm.New(session)

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
	)
	return
}

func getSsmAssociation(client *ssm.SSM) (r resourceSliceError) {
	r.err = client.ListAssociationsPages(&ssm.ListAssociationsInput{}, func(page *ssm.ListAssociationsOutput, lastPage bool) bool {
		for _, resource := range page.Associations {
			r.resources = append(r.resources, *resource.AssociationId)
		}
		return true
	})
	return
}

func getSsmDocument(client *ssm.SSM) (r resourceSliceError) {
	r.err = client.ListDocumentsPages(&ssm.ListDocumentsInput{}, func(page *ssm.ListDocumentsOutput, lastPage bool) bool {
		for _, resource := range page.DocumentIdentifiers {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getSsmMaintenanceWindow(client *ssm.SSM) (r resourceSliceError) {
	input := ssm.DescribeMaintenanceWindowsInput{}
	for {
		page, err := client.DescribeMaintenanceWindows(&input)
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

func getSsmMaintenanceWindowTarget(client *ssm.SSM, windowIDs []string) (r resourceSliceError) {
	for _, windowID := range windowIDs {
		input := ssm.DescribeMaintenanceWindowTargetsInput{
			WindowId: aws.String(windowID),
		}
		for {
			page, err := client.DescribeMaintenanceWindowTargets(&input)
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

func getSsmMaintenanceWindowTask(client *ssm.SSM, windowIDs []string) (r resourceSliceError) {
	for _, windowID := range windowIDs {
		input := ssm.DescribeMaintenanceWindowTasksInput{
			WindowId: aws.String(windowID),
		}
		for {
			page, err := client.DescribeMaintenanceWindowTasks(&input)
			if err != nil {
				r.err = err
				return
			}
			for _, resource := range page.Tasks {
				r.resources = append(r.resources, *resource.Name)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getSsmParameter(client *ssm.SSM) (r resourceSliceError) {
	r.err = client.DescribeParametersPages(&ssm.DescribeParametersInput{}, func(page *ssm.DescribeParametersOutput, lastPage bool) bool {
		for _, resource := range page.Parameters {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getSsmPatchBaseline(client *ssm.SSM) (r resourceSliceError) {
	input := ssm.DescribePatchBaselinesInput{}
	for {
		page, err := client.DescribePatchBaselines(&input)
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

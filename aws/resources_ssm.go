package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func getSsm(session *session.Session) (resources resourceMap) {
	client := ssm.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		ssmAssociation:             getSsmAssociation(client),
		ssmDocument:                getSsmDocument(client),
		ssmMaintenanceWindow:       getSsmMaintenanceWindow(client),
		ssmMaintenanceWindowTarget: getSsmMaintenanceWindowTarget(client),
		ssmMaintenanceWindowTask:   getSsmMaintenanceWindowTask(client),
		ssmParameter:               getSsmParameter(client),
		ssmPatchBaseline:           getSsmPatchBaseline(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getSsmAssociation(client *ssm.SSM) (r resourceSliceError) {
	logDebug("Listing SsmAssociation resources")
	r.err = client.ListAssociationsPages(&ssm.ListAssociationsInput{}, func(page *ssm.ListAssociationsOutput, lastPage bool) bool {
		for _, resource := range page.Associations {
			logDebug("Got SsmAssociation resource with PhysicalResourceId", *resource.AssociationId)
			r.resources = append(r.resources, *resource.AssociationId)
		}
		return true
	})
	return
}

func getSsmDocument(client *ssm.SSM) (r resourceSliceError) {
	logDebug("Listing SsmDocument resources")
	r.err = client.ListDocumentsPages(&ssm.ListDocumentsInput{}, func(page *ssm.ListDocumentsOutput, lastPage bool) bool {
		for _, resource := range page.DocumentIdentifiers {
			logDebug("Got SsmDocument resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getSsmMaintenanceWindow(client *ssm.SSM) (r resourceSliceError) {
	logDebug("Listing SsmMaintenanceWindow resources")
	input := ssm.DescribeMaintenanceWindowsInput{}
	for {
		page, err := client.DescribeMaintenanceWindows(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.WindowIdentities {
			logDebug("Got SsmMaintenanceWindow resource with PhysicalResourceId", *resource.WindowId)
			r.resources = append(r.resources, *resource.WindowId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSsmMaintenanceWindowTarget(client *ssm.SSM) (r resourceSliceError) {
	logDebug("Listing SsmMaintenanceWindowTarget resources")
	input := ssm.DescribeMaintenanceWindowTargetsInput{}
	for {
		page, err := client.DescribeMaintenanceWindowTargets(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Targets {
			logDebug("Got SsmMaintenanceWindowTarget resource with PhysicalResourceId", *resource.WindowTargetId)
			r.resources = append(r.resources, *resource.WindowTargetId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSsmMaintenanceWindowTask(client *ssm.SSM) (r resourceSliceError) {
	logDebug("Listing SsmMaintenanceWindowTask resources")
	input := ssm.DescribeMaintenanceWindowTasksInput{}
	for {
		page, err := client.DescribeMaintenanceWindowTasks(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Tasks {
			logDebug("Got SsmMaintenanceWindowTask resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getSsmParameter(client *ssm.SSM) (r resourceSliceError) {
	logDebug("Listing SsmParameter resources")
	r.err = client.DescribeParametersPages(&ssm.DescribeParametersInput{}, func(page *ssm.DescribeParametersOutput, lastPage bool) bool {
		for _, resource := range page.Parameters {
			logDebug("Got SsmParameter resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getSsmPatchBaseline(client *ssm.SSM) (r resourceSliceError) {
	logDebug("Listing SsmPatchBaseline resources")
	input := ssm.DescribePatchBaselinesInput{}
	for {
		page, err := client.DescribePatchBaselines(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.BaselineIdentities {
			logDebug("Got SsmPatchBaseline resource with PhysicalResourceId", *resource.BaselineId)
			r.resources = append(r.resources, *resource.BaselineId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

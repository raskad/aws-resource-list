package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
)

func getGuardDuty(config aws.Config) (resources awsResourceMap) {
	client := guardduty.New(config)

	guardDutyDetectorIDs := getGuardDutyDetectorIDs(client)
	guardDutyOrganizationAdminAccountIDs := getGuardDutyOrganizationAdminAccountIDs(client)

	resources = awsResourceMap{
		guardDutyDetector:                 guardDutyDetectorIDs,
		guardDutyOrganizationAdminAccount: guardDutyOrganizationAdminAccountIDs,
	}
	return
}

func getGuardDutyDetectorIDs(client *guardduty.Client) (resources []string) {
	req := client.ListDetectorsRequest(&guardduty.ListDetectorsInput{})
	p := guardduty.NewListDetectorsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.DetectorIds...)
	}
	return
}

func getGuardDutyOrganizationAdminAccountIDs(client *guardduty.Client) (resources []string) {
	req := client.ListOrganizationAdminAccountsRequest(&guardduty.ListOrganizationAdminAccountsInput{})
	p := guardduty.NewListOrganizationAdminAccountsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.AdminAccounts {
			resources = append(resources, *resource.AdminAccountId)
		}
	}
	return
}

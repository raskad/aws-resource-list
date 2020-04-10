package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

func getOrganizations(config aws.Config) (resources awsResourceMap) {
	client := organizations.New(config)

	organizationsRootIDs := getOrganizationsRootIDs(client)
	organizationsAccountIDs := getOrganizationsAccountIDs(client)
	organizationsOrganizationIDs := getOrganizationsOrganizationIDs(client)
	organizationsOrganizationalUnitIDs := getOrganizationsOrganizationalUnitIDs(client, organizationsRootIDs)
	organizationsPolicyIDs := getOrganizationsPolicyIDs(client)

	resources = awsResourceMap{
		organizationsAccount:            organizationsAccountIDs,
		organizationsOrganization:       organizationsOrganizationIDs,
		organizationsOrganizationalUnit: organizationsOrganizationalUnitIDs,
		organizationsPolicy:             organizationsPolicyIDs,
	}
	return
}

func getOrganizationsAccountIDs(client *organizations.Client) (resources []string) {
	input := organizations.ListAccountsInput{}
	for {
		page, err := client.ListAccountsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Accounts {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getOrganizationsOrganizationIDs(client *organizations.Client) (resources []string) {
	page, err := client.DescribeOrganizationRequest(&organizations.DescribeOrganizationInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	resources = append(resources, *page.Organization.Id)
	return
}

func getOrganizationsOrganizationalUnitIDs(client *organizations.Client, organizationsRootIDs []string) (resources []string) {
	for _, organizationsRootID := range organizationsRootIDs {
		page, err := client.ListOrganizationalUnitsForParentRequest(&organizations.ListOrganizationalUnitsForParentInput{
			ParentId: aws.String(organizationsRootID),
		}).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.OrganizationalUnits {
			resources = append(resources, *resource.Id)
		}
	}
	return
}

func getOrganizationsPolicyIDs(client *organizations.Client) (resources []string) {
	policyTypes := []organizations.PolicyType{organizations.PolicyTypeServiceControlPolicy, organizations.PolicyTypeTagPolicy}
	for _, policyType := range policyTypes {
		input := organizations.ListPoliciesInput{
			Filter: policyType,
		}
		for {
			page, err := client.ListPoliciesRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				return
			}
			for _, resource := range page.Policies {
				resources = append(resources, *resource.Id)
			}
			if page.NextToken == nil {
				return
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getOrganizationsRootIDs(client *organizations.Client) (resources []string) {
	input := organizations.ListRootsInput{}
	for {
		page, err := client.ListRootsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Roots {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

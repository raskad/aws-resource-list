package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func getIam(config aws.Config) (resources resourceMap) {
	client := iam.New(config)

	iamUserNames := getIamUserNames(client)
	iamAccessKeyIDs := getIamAccessKeyIDs(client, iamUserNames)
	iamGroupNames := getIamGroupNames(client)
	iamInstanceProfileNames := getIamInstanceProfileNames(client)
	iamPolicyNames := getIamPolicyNames(client)
	iamRoleNames := getIamRoleNames(client)
	iamRolePolicyNames := getIamRolePolicyNames(client, iamRoleNames)
	iamServiceLinkedRoleNames := getIamServiceLinkedRoleNames(client)

	resources = resourceMap{
		iamAccessKey:         iamAccessKeyIDs,
		iamGroup:             iamGroupNames,
		iamInstanceProfile:   iamInstanceProfileNames,
		iamPolicy:            iamPolicyNames,
		iamRole:              iamRoleNames,
		iamRolePolicy:        iamRolePolicyNames,
		iamServiceLinkedRole: iamServiceLinkedRoleNames,
		iamUser:              iamUserNames,
	}
	return
}

func getIamAccessKeyIDs(client *iam.Client, userNames []string) (resources []string) {
	for _, userName := range userNames {
		req := client.ListAccessKeysRequest(&iam.ListAccessKeysInput{
			UserName: aws.String(userName),
		})
		p := iam.NewListAccessKeysPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.AccessKeyMetadata {
				resources = append(resources, *resource.AccessKeyId)
			}
		}
	}
	return
}

func getIamGroupNames(client *iam.Client) (resources []string) {
	req := client.ListGroupsRequest(&iam.ListGroupsInput{})
	p := iam.NewListGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Groups {
			resources = append(resources, *resource.GroupName)
		}
	}
	return
}

func getIamInstanceProfileNames(client *iam.Client) (resources []string) {
	req := client.ListInstanceProfilesRequest(&iam.ListInstanceProfilesInput{})
	p := iam.NewListInstanceProfilesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.InstanceProfiles {
			resources = append(resources, *resource.InstanceProfileName)
		}
	}
	return
}

func getIamPolicyNames(client *iam.Client) (resources []string) {
	req := client.ListPoliciesRequest(&iam.ListPoliciesInput{
		Scope: iam.PolicyScopeTypeLocal,
	})
	p := iam.NewListPoliciesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Policies {
			resources = append(resources, *resource.PolicyName)
		}
	}
	return
}

func getIamRoleNames(client *iam.Client) (resources []string) {
	req := client.ListRolesRequest(&iam.ListRolesInput{})
	p := iam.NewListRolesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Roles {
			if !strings.HasPrefix(*resource.Path, "/aws-service-role/") {
				resources = append(resources, *resource.RoleName)
			}
		}
	}
	return
}

func getIamRolePolicyNames(client *iam.Client, roleNames []string) (resources []string) {
	for _, roleName := range roleNames {
		req := client.ListRolePoliciesRequest(&iam.ListRolePoliciesInput{
			RoleName: aws.String(roleName),
		})
		p := iam.NewListRolePoliciesPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			resources = append(resources, page.PolicyNames...)
		}
	}
	return
}

func getIamServiceLinkedRoleNames(client *iam.Client) (resources []string) {
	req := client.ListRolesRequest(&iam.ListRolesInput{
		PathPrefix: aws.String("/aws-service-role/"),
	})
	p := iam.NewListRolesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Roles {
			resources = append(resources, *resource.RoleName)
		}
	}
	return
}

func getIamUserNames(client *iam.Client) (resources []string) {
	req := client.ListUsersRequest(&iam.ListUsersInput{})
	p := iam.NewListUsersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Users {
			resources = append(resources, *resource.UserName)
		}
	}
	return
}

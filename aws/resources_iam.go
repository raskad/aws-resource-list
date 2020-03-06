package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func getIam(config aws.Config) (resources resourceMap) {
	client := iam.New(config)

	iamUserResourceMap := getIamUser(client).unwrap(iamUser)
	iamUserNames := iamUserResourceMap[iamUser]

	iamRoleResourceMap := getIamRole(client).unwrap(iamRole)
	iamRoleNames := iamRoleResourceMap[iamRole]

	resources = reduce(
		getIamAccessKey(client, iamUserNames).unwrap(iamAccessKey),
		getIamGroup(client).unwrap(iamGroup),
		getIamInstanceProfile(client).unwrap(iamInstanceProfile),
		getIamPolicy(client).unwrap(iamPolicy),
		iamRoleResourceMap,
		getIamRolePolicy(client, iamRoleNames).unwrap(iamRolePolicy),
		getIamServiceLinkedRole(client).unwrap(iamServiceLinkedRole),
		iamUserResourceMap,
	)
	return
}

func getIamAccessKey(client *iam.Client, userNames []string) (r resourceSliceError) {
	for _, userName := range userNames {
		req := client.ListAccessKeysRequest(&iam.ListAccessKeysInput{
			UserName: aws.String(userName),
		})
		p := iam.NewListAccessKeysPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.AccessKeyMetadata {
				r.resources = append(r.resources, *resource.AccessKeyId)
			}
		}
		r.err = p.Err()
	}
	return
}

func getIamGroup(client *iam.Client) (r resourceSliceError) {
	req := client.ListGroupsRequest(&iam.ListGroupsInput{})
	p := iam.NewListGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Groups {
			r.resources = append(r.resources, *resource.GroupName)
		}
	}
	r.err = p.Err()
	return
}

func getIamInstanceProfile(client *iam.Client) (r resourceSliceError) {
	req := client.ListInstanceProfilesRequest(&iam.ListInstanceProfilesInput{})
	p := iam.NewListInstanceProfilesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.InstanceProfiles {
			r.resources = append(r.resources, *resource.InstanceProfileName)
		}
	}
	r.err = p.Err()
	return
}

func getIamPolicy(client *iam.Client) (r resourceSliceError) {
	req := client.ListPoliciesRequest(&iam.ListPoliciesInput{
		Scope: iam.PolicyScopeTypeLocal,
	})
	p := iam.NewListPoliciesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Policies {
			r.resources = append(r.resources, *resource.PolicyName)
		}
	}
	r.err = p.Err()
	return
}

func getIamRole(client *iam.Client) (r resourceSliceError) {
	req := client.ListRolesRequest(&iam.ListRolesInput{})
	p := iam.NewListRolesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Roles {
			if !strings.HasPrefix(*resource.Path, "/aws-service-role/") {
				r.resources = append(r.resources, *resource.RoleName)
			}
		}
	}
	r.err = p.Err()
	return
}

func getIamRolePolicy(client *iam.Client, roleNames []string) (r resourceSliceError) {
	for _, roleName := range roleNames {
		req := client.ListRolePoliciesRequest(&iam.ListRolePoliciesInput{
			RoleName: aws.String(roleName),
		})
		p := iam.NewListRolePoliciesPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.PolicyNames {
				r.resources = append(r.resources, resource)
			}
		}
		r.err = p.Err()
	}
	return
}

func getIamServiceLinkedRole(client *iam.Client) (r resourceSliceError) {
	req := client.ListRolesRequest(&iam.ListRolesInput{
		PathPrefix: aws.String("/aws-service-role/"),
	})
	p := iam.NewListRolesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Roles {
			r.resources = append(r.resources, *resource.RoleName)
		}
	}
	r.err = p.Err()
	return
}

func getIamUser(client *iam.Client) (r resourceSliceError) {
	req := client.ListUsersRequest(&iam.ListUsersInput{})
	p := iam.NewListUsersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Users {
			r.resources = append(r.resources, *resource.UserName)
		}
	}
	r.err = p.Err()
	return
}

package aws

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func getIam(session *session.Session) (resources resourceMap) {
	client := iam.New(session)
	resources = reduce(
		getIamAccessKey(client).unwrap(iamAccessKey),
		getIamGroup(client).unwrap(iamGroup),
		getIamInstanceProfile(client).unwrap(iamInstanceProfile),
		getIamPolicy(client).unwrap(iamPolicy),
		getIamRole(client).unwrap(iamRole),
		getIamRolePolicy(client).unwrap(iamRolePolicy),
		getIamServiceLinkedRole(client).unwrap(iamServiceLinkedRole),
		getIamUser(client).unwrap(iamUser),
	)
	return
}

func getIamAccessKey(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListAccessKeysPages(&iam.ListAccessKeysInput{}, func(page *iam.ListAccessKeysOutput, lastPage bool) bool {
		for _, resource := range page.AccessKeyMetadata {
			r.resources = append(r.resources, *resource.AccessKeyId)
		}
		return true
	})
	return
}

func getIamGroup(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListGroupsPages(&iam.ListGroupsInput{}, func(page *iam.ListGroupsOutput, lastPage bool) bool {
		for _, resource := range page.Groups {
			r.resources = append(r.resources, *resource.GroupName)
		}
		return true
	})
	return
}

func getIamInstanceProfile(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListInstanceProfilesPages(&iam.ListInstanceProfilesInput{}, func(page *iam.ListInstanceProfilesOutput, lastPage bool) bool {
		for _, resource := range page.InstanceProfiles {
			r.resources = append(r.resources, *resource.InstanceProfileName)
		}
		return true
	})
	return
}

func getIamPolicy(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListPoliciesPages(&iam.ListPoliciesInput{
		Scope: aws.String(iam.PolicyScopeTypeLocal),
	}, func(page *iam.ListPoliciesOutput, lastPage bool) bool {
		for _, resource := range page.Policies {
			r.resources = append(r.resources, *resource.PolicyName)
		}
		return true
	})
	return
}

func getIamRole(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListRolesPages(&iam.ListRolesInput{}, func(page *iam.ListRolesOutput, lastPage bool) bool {
		for _, resource := range page.Roles {
			if !strings.HasPrefix(*resource.Path, "/aws-service-role/") {
				r.resources = append(r.resources, *resource.RoleName)
			}
		}
		return true
	})
	return
}

func getIamRolePolicy(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListRolePoliciesPages(&iam.ListRolePoliciesInput{}, func(page *iam.ListRolePoliciesOutput, lastPage bool) bool {
		for _, resource := range page.PolicyNames {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getIamServiceLinkedRole(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListRolesPages(&iam.ListRolesInput{
		PathPrefix: aws.String("/aws-service-role/"),
	}, func(page *iam.ListRolesOutput, lastPage bool) bool {
		for _, resource := range page.Roles {
			r.resources = append(r.resources, *resource.RoleName)
		}
		return true
	})
	return
}

func getIamUser(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListUsersPages(&iam.ListUsersInput{}, func(page *iam.ListUsersOutput, lastPage bool) bool {
		for _, resource := range page.Users {
			r.resources = append(r.resources, *resource.UserName)
		}
		return true
	})
	return
}

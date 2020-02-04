package aws

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func getIam(session *session.Session) (resources resourceMap) {
	client := iam.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		iamAccessKey:         getIamAccessKey(client),
		iamGroup:             getIamGroup(client),
		iamInstanceProfile:   getIamInstanceProfile(client),
		iamPolicy:            getIamPolicy(client),
		iamRole:              getIamRole(client),
		iamRolePolicy:        getIamRolePolicy(client),
		iamServiceLinkedRole: getIamServiceLinkedRole(client),
		iamUser:              getIamUser(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getIamAccessKey(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListAccessKeysPages(&iam.ListAccessKeysInput{}, func(page *iam.ListAccessKeysOutput, lastPage bool) bool {
		logDebug("Listing IamAccessKey resources page. Remaining pages", page.Marker)
		for _, resource := range page.AccessKeyMetadata {
			logDebug("Got IamAccessKey resource with PhysicalResourceId", *resource.AccessKeyId)
			r.resources = append(r.resources, *resource.AccessKeyId)
		}
		return true
	})
	return
}

func getIamGroup(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListGroupsPages(&iam.ListGroupsInput{}, func(page *iam.ListGroupsOutput, lastPage bool) bool {
		logDebug("Listing IamGroup resources page. Remaining pages", page.Marker)
		for _, resource := range page.Groups {
			logDebug("Got IamGroup resource with PhysicalResourceId", *resource.GroupName)
			r.resources = append(r.resources, *resource.GroupName)
		}
		return true
	})
	return
}

func getIamInstanceProfile(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListInstanceProfilesPages(&iam.ListInstanceProfilesInput{}, func(page *iam.ListInstanceProfilesOutput, lastPage bool) bool {
		logDebug("Listing IamInstanceProfile resources page. Remaining pages", page.Marker)
		for _, resource := range page.InstanceProfiles {
			logDebug("Got IamInstanceProfile resource with PhysicalResourceId", *resource.InstanceProfileName)
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
		logDebug("Listing IamManagedPolicy resources page. Remaining pages", page.Marker)
		for _, resource := range page.Policies {
			logDebug("Got IamManagedPolicy resource with PhysicalResourceId", *resource.PolicyName)
			r.resources = append(r.resources, *resource.PolicyName)
		}
		return true
	})
	return
}

func getIamRole(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListRolesPages(&iam.ListRolesInput{}, func(page *iam.ListRolesOutput, lastPage bool) bool {
		logDebug("Listing IamRole resources page. Remaining pages", page.Marker)
		for _, resource := range page.Roles {
			if !strings.HasPrefix(*resource.Path, "/aws-service-role/") {
				logDebug("Got IamRole resource with PhysicalResourceId", *resource.RoleName)
				r.resources = append(r.resources, *resource.RoleName)
			}
		}
		return true
	})
	return
}

func getIamRolePolicy(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListRolePoliciesPages(&iam.ListRolePoliciesInput{}, func(page *iam.ListRolePoliciesOutput, lastPage bool) bool {
		logDebug("Listing IamRolePolicy resources page. Remaining pages", page.Marker)
		for _, resource := range page.PolicyNames {
			logDebug("Got IamRolePolicy resource with PhysicalResourceId", *resource)
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
		logDebug("Listing IamServiceLinkedRole resources page. Remaining pages", page.Marker)
		for _, resource := range page.Roles {
			logDebug("Got IamServiceLinkedRole resource with PhysicalResourceId", *resource.RoleName)
			r.resources = append(r.resources, *resource.RoleName)
		}
		return true
	})
	return
}

func getIamUser(client *iam.IAM) (r resourceSliceError) {
	r.err = client.ListUsersPages(&iam.ListUsersInput{}, func(page *iam.ListUsersOutput, lastPage bool) bool {
		logDebug("Listing IamUser resources page. Remaining pages", page.Marker)
		for _, resource := range page.Users {
			logDebug("Got IamUser resource with PhysicalResourceId", *resource.UserName)
			r.resources = append(r.resources, *resource.UserName)
		}
		return true
	})
	return
}

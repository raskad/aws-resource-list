package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func getIam(config aws.Config) (resources awsResourceMap) {
	client := iam.New(config)

	iamUserNames := getIamUserNames(client)
	iamAccessKeyIDs := getIamAccessKeyIDs(client, iamUserNames)
	iamAccountAliasNames := getIamAccountAliasNames(client)
	iamGroupNames := getIamGroupNames(client)
	iamGroupPolicyNames := getIamGroupPolicyNames(client)
	iamInstanceProfileNames := getIamInstanceProfileNames(client)
	iamOpenidConnectProviderARNs := getIamOpenidConnectProviderARNs(client)
	iamPolicyNames := getIamPolicyNames(client)
	iamRoleNames := getIamRoleNames(client)
	iamRolePolicyNames := getIamRolePolicyNames(client, iamRoleNames)
	iamSamlProviderARNs := getIamSamlProviderARNs(client)
	iamServerCertificateNames := getIamServerCertificateNames(client)
	iamServiceLinkedRoleNames := getIamServiceLinkedRoleNames(client)
	iamUserPolicyNames := getIamUserPolicyNames(client, iamUserNames)
	iamUserSSHKeyIDs := getIamUserSSHKeyIDs(client, iamUserNames)

	resources = awsResourceMap{
		iamAccessKey:             iamAccessKeyIDs,
		iamAccountAlias:          iamAccountAliasNames,
		iamGroup:                 iamGroupNames,
		iamGroupPolicy:           iamGroupPolicyNames,
		iamInstanceProfile:       iamInstanceProfileNames,
		iamOpenidConnectProvider: iamOpenidConnectProviderARNs,
		iamPolicy:                iamPolicyNames,
		iamRole:                  iamRoleNames,
		iamRolePolicy:            iamRolePolicyNames,
		iamSamlProvider:          iamSamlProviderARNs,
		iamServerCertificate:     iamServerCertificateNames,
		iamServiceLinkedRole:     iamServiceLinkedRoleNames,
		iamUser:                  iamUserNames,
		iamUserPolicy:            iamUserPolicyNames,
		iamUserSSHKey:            iamUserSSHKeyIDs,
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
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.AccessKeyMetadata {
				resources = append(resources, *resource.AccessKeyId)
			}
		}
	}
	return
}

func getIamAccountAliasNames(client *iam.Client) (resources []string) {
	req := client.ListAccountAliasesRequest(&iam.ListAccountAliasesInput{})
	p := iam.NewListAccountAliasesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.AccountAliases...)
	}
	return
}

func getIamGroupNames(client *iam.Client) (resources []string) {
	req := client.ListGroupsRequest(&iam.ListGroupsInput{})
	p := iam.NewListGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Groups {
			resources = append(resources, *resource.GroupName)
		}
	}
	return
}

func getIamGroupPolicyNames(client *iam.Client) (resources []string) {
	req := client.ListGroupPoliciesRequest(&iam.ListGroupPoliciesInput{})
	p := iam.NewListGroupPoliciesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.PolicyNames...)
	}
	return
}

func getIamInstanceProfileNames(client *iam.Client) (resources []string) {
	req := client.ListInstanceProfilesRequest(&iam.ListInstanceProfilesInput{})
	p := iam.NewListInstanceProfilesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.InstanceProfiles {
			resources = append(resources, *resource.InstanceProfileName)
		}
	}
	return
}

func getIamOpenidConnectProviderARNs(client *iam.Client) (resources []string) {
	page, err := client.ListOpenIDConnectProvidersRequest(&iam.ListOpenIDConnectProvidersInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.OpenIDConnectProviderList {
		resources = append(resources, *resource.Arn)
	}
	return
}

func getIamPolicyNames(client *iam.Client) (resources []string) {
	req := client.ListPoliciesRequest(&iam.ListPoliciesInput{
		Scope: iam.PolicyScopeTypeLocal,
	})
	p := iam.NewListPoliciesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			resources = append(resources, page.PolicyNames...)
		}
	}
	return
}

func getIamSamlProviderARNs(client *iam.Client) (resources []string) {
	page, err := client.ListSAMLProvidersRequest(&iam.ListSAMLProvidersInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range page.SAMLProviderList {
		resources = append(resources, *resource.Arn)
	}
	return
}

func getIamServerCertificateNames(client *iam.Client) (resources []string) {
	req := client.ListServerCertificatesRequest(&iam.ListServerCertificatesInput{})
	p := iam.NewListServerCertificatesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ServerCertificateMetadataList {
			resources = append(resources, *resource.ServerCertificateName)
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
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Users {
			resources = append(resources, *resource.UserName)
		}
	}
	return
}

func getIamUserPolicyNames(client *iam.Client, iamUserNames []string) (resources []string) {
	for _, iamUserName := range iamUserNames {
		req := client.ListUserPoliciesRequest(&iam.ListUserPoliciesInput{
			UserName: aws.String(iamUserName),
		})
		p := iam.NewListUserPoliciesPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			resources = append(resources, page.PolicyNames...)
		}
	}
	return
}

func getIamUserSSHKeyIDs(client *iam.Client, iamUserNames []string) (resources []string) {
	for _, iamUserName := range iamUserNames {
		req := client.ListSSHPublicKeysRequest(&iam.ListSSHPublicKeysInput{
			UserName: aws.String(iamUserName),
		})
		p := iam.NewListSSHPublicKeysPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.SSHPublicKeys {
				resources = append(resources, *resource.SSHPublicKeyId)
			}
		}
	}
	return
}

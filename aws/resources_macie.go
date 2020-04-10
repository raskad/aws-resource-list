package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/macie"
)

func getMacie(config aws.Config) (resources awsResourceMap) {
	client := macie.New(config)

	macieMemberAccountAssociationIDs := getMacieMemberAccountAssociationIDs(client)
	macieS3BucketAssociationNames := getMacieS3BucketAssociationNames(client)

	resources = awsResourceMap{
		macieMemberAccountAssociation: macieMemberAccountAssociationIDs,
		macieS3BucketAssociation:      macieS3BucketAssociationNames,
	}
	return
}

func getMacieMemberAccountAssociationIDs(client *macie.Client) (resources []string) {
	req := client.ListMemberAccountsRequest(&macie.ListMemberAccountsInput{})
	p := macie.NewListMemberAccountsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.MemberAccounts {
			resources = append(resources, *resource.AccountId)
		}
	}
	return
}

func getMacieS3BucketAssociationNames(client *macie.Client) (resources []string) {
	req := client.ListS3ResourcesRequest(&macie.ListS3ResourcesInput{})
	p := macie.NewListS3ResourcesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.S3Resources {
			resources = append(resources, *resource.BucketName)
		}
	}
	return
}

package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func getKms(config aws.Config) (resources awsResourceMap) {
	client := kms.New(config)

	kmsAliasNames := getKmsAliasNames(client)
	kmsGrantIDs := getKmsGrantIDs(client)
	kmsKeyIDs := getKmsKeyIDs(client)

	resources = awsResourceMap{
		kmsAlias: kmsAliasNames,
		kmsGrant: kmsGrantIDs,
		kmsKey:   kmsKeyIDs,
	}
	return
}

func getKmsAliasNames(client *kms.Client) (resources []string) {
	req := client.ListAliasesRequest(&kms.ListAliasesInput{})
	p := kms.NewListAliasesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Aliases {
			if !strings.HasPrefix(*resource.AliasName, "alias/aws/") {
				resources = append(resources, *resource.AliasName)
			}
		}
	}
	return
}

func getKmsGrantIDs(client *kms.Client) (resources []string) {
	req := client.ListGrantsRequest(&kms.ListGrantsInput{})
	p := kms.NewListGrantsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Grants {
			resources = append(resources, *resource.GrantId)
		}
	}
	return
}

func getKmsKeyIDs(client *kms.Client) (resources []string) {
	req := client.ListKeysRequest(&kms.ListKeysInput{})
	p := kms.NewListKeysPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Keys {
			resources = append(resources, *resource.KeyId)
		}
	}
	return
}

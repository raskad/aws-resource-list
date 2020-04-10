package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
)

func getGlacier(config aws.Config) (resources awsResourceMap) {
	client := glacier.New(config)

	glacierVaultNames := getGlacierVaultNames(client)

	resources = awsResourceMap{
		glacierVault: glacierVaultNames,
	}
	return
}

func getGlacierVaultNames(client *glacier.Client) (resources []string) {
	req := client.ListVaultsRequest(&glacier.ListVaultsInput{
		AccountId: aws.String("-"),
	})
	p := glacier.NewListVaultsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.VaultList {
			resources = append(resources, *resource.VaultName)
		}
	}
	return
}

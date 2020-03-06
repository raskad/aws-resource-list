package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func getKms(config aws.Config) (resources resourceMap) {
	client := kms.New(config)
	resources = reduce(
		getkmsAlias(client).unwrap(kmsAlias),
		getkmsKey(client).unwrap(kmsKey),
	)
	return
}

func getkmsAlias(client *kms.Client) (r resourceSliceError) {
	req := client.ListAliasesRequest(&kms.ListAliasesInput{})
	p := kms.NewListAliasesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Aliases {
			r.resources = append(r.resources, *resource.AliasName)
		}
	}
	r.err = p.Err()
	return
}

func getkmsKey(client *kms.Client) (r resourceSliceError) {
	req := client.ListKeysRequest(&kms.ListKeysInput{})
	p := kms.NewListKeysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Keys {
			r.resources = append(r.resources, *resource.KeyId)
		}
	}
	r.err = p.Err()
	return
}

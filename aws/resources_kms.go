package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func getKms(session *session.Session) (resources resourceMap) {
	client := kms.New(session)
	resources = reduce(
		getkmsAlias(client).unwrap(kmsAlias),
		getkmsKey(client).unwrap(kmsKey),
	)
	return
}

func getkmsAlias(client *kms.KMS) (r resourceSliceError) {
	r.err = client.ListAliasesPages(&kms.ListAliasesInput{}, func(page *kms.ListAliasesOutput, lastPage bool) bool {
		for _, resource := range page.Aliases {
			r.resources = append(r.resources, *resource.AliasName)
		}
		return true
	})
	return
}

func getkmsKey(client *kms.KMS) (r resourceSliceError) {
	r.err = client.ListKeysPages(&kms.ListKeysInput{}, func(page *kms.ListKeysOutput, lastPage bool) bool {
		for _, resource := range page.Keys {
			r.resources = append(r.resources, *resource.KeyId)
		}
		return true
	})
	return
}

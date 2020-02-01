package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func getKms(session *session.Session) (resources resourceMap) {
	client := kms.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		kmsAlias: getkmsAlias(client),
		kmsKey:   getkmsKey(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getkmsAlias(client *kms.KMS) (r resourceSliceError) {
	r.err = client.ListAliasesPages(&kms.ListAliasesInput{}, func(page *kms.ListAliasesOutput, lastPage bool) bool {
		logDebug("Listing kmsAlias resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.Aliases {
			logDebug("Got kmsAlias resource with PhysicalResourceId", *resource.AliasName)
			r.resources = append(r.resources, *resource.AliasName)
		}
		return true
	})
	return
}

func getkmsKey(client *kms.KMS) (r resourceSliceError) {
	r.err = client.ListKeysPages(&kms.ListKeysInput{}, func(page *kms.ListKeysOutput, lastPage bool) bool {
		logDebug("Listing kmsAlias resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.Keys {
			logDebug("Got kmsAlias resource with PhysicalResourceId", *resource.KeyId)
			r.resources = append(r.resources, *resource.KeyId)
		}
		return true
	})
	return
}

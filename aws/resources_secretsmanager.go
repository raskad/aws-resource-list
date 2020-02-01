package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func getSecretsManager(session *session.Session) (resources resourceMap) {
	client := secretsmanager.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		secretsManagerSecret: getSecretsManagerSecret(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getSecretsManagerSecret(client *secretsmanager.SecretsManager) (r resourceSliceError) {
	r.err = client.ListSecretsPages(&secretsmanager.ListSecretsInput{}, func(page *secretsmanager.ListSecretsOutput, lastPage bool) bool {
		logDebug("List SecretsManagerSecret resources page")
		for _, resource := range page.SecretList {
			logDebug("Got SecretsManagerSecret resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

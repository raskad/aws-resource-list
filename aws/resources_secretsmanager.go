package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func getSecretsManager(session *session.Session) (resources resourceMap) {
	client := secretsmanager.New(session)
	resources = reduce(
		getSecretsManagerSecret(client).unwrap(secretsManagerSecret),
	)
	return
}

func getSecretsManagerSecret(client *secretsmanager.SecretsManager) (r resourceSliceError) {
	r.err = client.ListSecretsPages(&secretsmanager.ListSecretsInput{}, func(page *secretsmanager.ListSecretsOutput, lastPage bool) bool {
		for _, resource := range page.SecretList {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

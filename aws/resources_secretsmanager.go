package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func getSecretsManager(config aws.Config) (resources resourceMap) {
	client := secretsmanager.New(config)

	secretsManagerSecretARNs := getSecretsManagerSecretARNs(client)

	resources = resourceMap{
		secretsManagerSecret: secretsManagerSecretARNs,
	}
	return
}

func getSecretsManagerSecretARNs(client *secretsmanager.Client) (resources []string) {
	req := client.ListSecretsRequest(&secretsmanager.ListSecretsInput{})
	p := secretsmanager.NewListSecretsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.SecretList {
			resources = append(resources, *resource.ARN)
		}
	}
	return
}

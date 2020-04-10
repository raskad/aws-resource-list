package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func getSecretsManager(config aws.Config) (resources awsResourceMap) {
	client := secretsmanager.New(config)

	secretsManagerSecretNames := getSecretsManagerSecretNames(client)
	secretsManagerSecretVersionIDs := getSecretsManagerSecretVersionIDs(client)

	resources = awsResourceMap{
		secretsManagerSecret:        secretsManagerSecretNames,
		secretsManagerSecretVersion: secretsManagerSecretVersionIDs,
	}
	return
}

func getSecretsManagerSecretNames(client *secretsmanager.Client) (resources []string) {
	req := client.ListSecretsRequest(&secretsmanager.ListSecretsInput{})
	p := secretsmanager.NewListSecretsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.SecretList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getSecretsManagerSecretVersionIDs(client *secretsmanager.Client) (resources []string) {
	req := client.ListSecretVersionIdsRequest(&secretsmanager.ListSecretVersionIdsInput{})
	p := secretsmanager.NewListSecretVersionIdsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Versions {
			resources = append(resources, *resource.VersionId)
		}
	}
	return
}

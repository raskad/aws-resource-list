package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func getSecretsManager(config aws.Config) (resources resourceMap) {
	client := secretsmanager.New(config)
	resources = reduce(
		getSecretsManagerSecret(client).unwrap(secretsManagerSecret),
	)
	return
}

func getSecretsManagerSecret(client *secretsmanager.Client) (r resourceSliceError) {
	req := client.ListSecretsRequest(&secretsmanager.ListSecretsInput{})
	p := secretsmanager.NewListSecretsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.SecretList {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

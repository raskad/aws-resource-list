package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
)

func getAcmpca(config aws.Config) (resources awsResourceMap) {
	client := acmpca.New(config)

	acmpcaCertificateAuthorityArns := getAcmpcaCertificateAuthorityArns(client)

	resources = awsResourceMap{
		acmpcaCertificateAuthority: acmpcaCertificateAuthorityArns,
	}
	return
}

func getAcmpcaCertificateAuthorityArns(client *acmpca.Client) (resources []string) {
	req := client.ListCertificateAuthoritiesRequest(&acmpca.ListCertificateAuthoritiesInput{})
	p := acmpca.NewListCertificateAuthoritiesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.CertificateAuthorities {
			resources = append(resources, *resource.Arn)
		}
	}
	return
}

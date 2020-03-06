package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
)

func getAcmpca(config aws.Config) (resources resourceMap) {
	client := acmpca.New(config)
	resources = reduce(
		getAcmpcaCertificateAuthority(client).unwrap(acmpcaCertificateAuthority),
	)
	return
}

func getAcmpcaCertificateAuthority(client *acmpca.Client) (r resourceSliceError) {
	req := client.ListCertificateAuthoritiesRequest(&acmpca.ListCertificateAuthoritiesInput{})
	p := acmpca.NewListCertificateAuthoritiesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.CertificateAuthorities {
			r.resources = append(r.resources, *resource.Arn)
		}
	}
	r.err = p.Err()
	return
}

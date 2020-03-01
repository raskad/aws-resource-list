package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acmpca"
)

func getAcmpca(session *session.Session) (resources resourceMap) {
	client := acmpca.New(session)
	resources = reduce(
		getAcmpcaCertificateAuthority(client).unwrap(acmpcaCertificateAuthority),
	)
	return
}

func getAcmpcaCertificateAuthority(client *acmpca.ACMPCA) (r resourceSliceError) {
	r.err = client.ListCertificateAuthoritiesPages(&acmpca.ListCertificateAuthoritiesInput{}, func(page *acmpca.ListCertificateAuthoritiesOutput, lastPage bool) bool {
		for _, resource := range page.CertificateAuthorities {
			r.resources = append(r.resources, *resource.Arn)
		}
		return true
	})
	return
}

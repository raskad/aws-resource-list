package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acmpca"
)

func getAcmpca(session *session.Session) (resources resourceMap) {
	client := acmpca.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		acmpcaCertificateAuthority: getAcmpcaCertificateAuthority(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAcmpcaCertificateAuthority(client *acmpca.ACMPCA) (r resourceSliceError) {
	logDebug("Listing AcmpcaCertificateAuthority resources")
	r.err = client.ListCertificateAuthoritiesPages(&acmpca.ListCertificateAuthoritiesInput{}, func(page *acmpca.ListCertificateAuthoritiesOutput, lastPage bool) bool {
		for _, resource := range page.CertificateAuthorities {
			logDebug("Got AcmpcaCertificateAuthority resource with PhysicalResourceId", *resource.Arn)
			r.resources = append(r.resources, *resource.Arn)
		}
		return true
	})
	return
}

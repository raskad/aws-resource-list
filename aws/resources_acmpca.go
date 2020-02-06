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
	r.err = client.ListCertificateAuthoritiesPages(&acmpca.ListCertificateAuthoritiesInput{}, func(page *acmpca.ListCertificateAuthoritiesOutput, lastPage bool) bool {
		logDebug("Listing AcmpcaCertificateAuthority resources page. Remaining pages", page.NextToken)
		for _, resource := range page.CertificateAuthorities {
			logDebug("Got AcmpcaCertificateAuthority resource with PhysicalResourceId", *resource.Arn)
			r.resources = append(r.resources, *resource.Arn)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
)

func getAcm(session *session.Session) (resources resourceMap) {
	client := acm.New(session)
	resources = reduce(
		getCertificateManagerCertificate(client).unwrap(certificateManagerCertificate),
	)
	return
}

func getCertificateManagerCertificate(client *acm.ACM) (r resourceSliceError) {
	r.err = client.ListCertificatesPages(&acm.ListCertificatesInput{}, func(page *acm.ListCertificatesOutput, lastPage bool) bool {
		for _, resource := range page.CertificateSummaryList {
			r.resources = append(r.resources, *resource.CertificateArn)
		}
		return true
	})
	return
}

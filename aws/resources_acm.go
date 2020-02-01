package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
)

func getAcm(session *session.Session) (resources resourceMap) {
	client := acm.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		certificateManagerCertificate: getCertificateManagerCertificate(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCertificateManagerCertificate(client *acm.ACM) (r resourceSliceError) {
	r.err = client.ListCertificatesPages(&acm.ListCertificatesInput{}, func(page *acm.ListCertificatesOutput, lastPage bool) bool {
		logDebug("List CertificateManagerCertificate resources page. Remaining pages", page.NextToken)
		for _, resource := range page.CertificateSummaryList {
			logDebug("Got CertificateManagerCertificate resource with PhysicalResourceId", *resource.CertificateArn)
			r.resources = append(r.resources, *resource.CertificateArn)
		}
		return true
	})
	return
}

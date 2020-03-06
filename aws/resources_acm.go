package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
)

func getAcm(config aws.Config) (resources resourceMap) {
	client := acm.New(config)
	resources = reduce(
		getCertificateManagerCertificate(client).unwrap(certificateManagerCertificate),
	)
	return
}

func getCertificateManagerCertificate(client *acm.Client) (r resourceSliceError) {
	req := client.ListCertificatesRequest(&acm.ListCertificatesInput{})
	p := acm.NewListCertificatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.CertificateSummaryList {
			r.resources = append(r.resources, *resource.CertificateArn)
		}
	}
	r.err = p.Err()
	return
}

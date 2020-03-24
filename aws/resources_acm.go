package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
)

func getAcm(config aws.Config) (resources resourceMap) {
	client := acm.New(config)

	certificateManagerCertificateArns := getCertificateManagerCertificateArns(client)

	resources = resourceMap{
		certificateManagerCertificate: certificateManagerCertificateArns,
	}
	return
}

func getCertificateManagerCertificateArns(client *acm.Client) (resources []string) {
	req := client.ListCertificatesRequest(&acm.ListCertificatesInput{})
	p := acm.NewListCertificatesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.CertificateSummaryList {
			resources = append(resources, *resource.CertificateArn)
		}
	}
	return
}

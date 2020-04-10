package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
)

func getAcm(config aws.Config) (resources awsResourceMap) {
	client := acm.New(config)

	certificateManagerCertificateArns := getCertificateManagerCertificateArns(client)

	resources = awsResourceMap{
		certificateManagerCertificate: certificateManagerCertificateArns,
	}
	return
}

func getCertificateManagerCertificateArns(client *acm.Client) (resources []string) {
	req := client.ListCertificatesRequest(&acm.ListCertificatesInput{})
	p := acm.NewListCertificatesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.CertificateSummaryList {
			resources = append(resources, *resource.CertificateArn)
		}
	}
	return
}

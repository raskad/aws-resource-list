package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
)

func getCostAndUsageReportService(config aws.Config) (resources awsResourceMap) {
	client := costandusagereportservice.New(config)

	costAndUsageReportServiceReportDefinitionNames := getCostAndUsageReportServiceReportDefinitionNames(client)

	resources = awsResourceMap{
		costAndUsageReportServiceReportDefinition: costAndUsageReportServiceReportDefinitionNames,
	}
	return
}

func getCostAndUsageReportServiceReportDefinitionNames(client *costandusagereportservice.Client) (resources []string) {
	req := client.DescribeReportDefinitionsRequest(&costandusagereportservice.DescribeReportDefinitionsInput{})
	p := costandusagereportservice.NewDescribeReportDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ReportDefinitions {
			resources = append(resources, *resource.ReportName)
		}
	}
	return
}

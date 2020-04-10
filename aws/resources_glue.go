package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func getGlue(config aws.Config) (resources awsResourceMap) {
	client := glue.New(config)

	glueDatabaseNames := getGlueDatabaseNames(client)
	glueConnectionNames := getGlueConnectionNames(client)
	glueCrawlerNames := getGlueCrawlerNames(client)
	glueDevEndpointNames := getGlueDevEndpointNames(client)
	glueJobNames := getGlueJobNames(client)
	glueMLTransformNames := getGlueMLTransformNames(client)
	glueSecurityConfigurationNames := getGlueSecurityConfigurationNames(client)
	glueTableNames := getGlueTableNames(client, glueDatabaseNames)
	glueTriggerNames := getGlueTriggerNames(client)
	glueWorkflowNames := getGlueWorkflowNames(client)

	resources = awsResourceMap{
		glueDatabase:              glueDatabaseNames,
		glueConnection:            glueConnectionNames,
		glueCrawler:               glueCrawlerNames,
		glueDevEndpoint:           glueDevEndpointNames,
		glueJob:                   glueJobNames,
		glueMLTransform:           glueMLTransformNames,
		glueSecurityConfiguration: glueSecurityConfigurationNames,
		glueTable:                 glueTableNames,
		glueTrigger:               glueTriggerNames,
		glueWorkflow:              glueWorkflowNames,
	}
	return
}

func getGlueConnectionNames(client *glue.Client) (resources []string) {
	req := client.GetConnectionsRequest(&glue.GetConnectionsInput{})
	p := glue.NewGetConnectionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ConnectionList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getGlueCrawlerNames(client *glue.Client) (resources []string) {
	req := client.GetCrawlersRequest(&glue.GetCrawlersInput{})
	p := glue.NewGetCrawlersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Crawlers {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getGlueDatabaseNames(client *glue.Client) (resources []string) {
	req := client.GetDatabasesRequest(&glue.GetDatabasesInput{})
	p := glue.NewGetDatabasesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DatabaseList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getGlueDevEndpointNames(client *glue.Client) (resources []string) {
	req := client.GetDevEndpointsRequest(&glue.GetDevEndpointsInput{})
	p := glue.NewGetDevEndpointsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DevEndpoints {
			resources = append(resources, *resource.EndpointName)
		}
	}
	return
}

func getGlueJobNames(client *glue.Client) (resources []string) {
	req := client.GetJobsRequest(&glue.GetJobsInput{})
	p := glue.NewGetJobsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Jobs {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getGlueMLTransformNames(client *glue.Client) (resources []string) {
	req := client.GetMLTransformsRequest(&glue.GetMLTransformsInput{})
	p := glue.NewGetMLTransformsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Transforms {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getGlueSecurityConfigurationNames(client *glue.Client) (resources []string) {
	req := client.GetSecurityConfigurationsRequest(&glue.GetSecurityConfigurationsInput{})
	p := glue.NewGetSecurityConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.SecurityConfigurations {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getGlueTableNames(client *glue.Client, databaseNames []string) (resources []string) {
	for _, databaseName := range databaseNames {
		req := client.GetTablesRequest(&glue.GetTablesInput{
			DatabaseName: aws.String(databaseName),
		})
		p := glue.NewGetTablesPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.TableList {
				resources = append(resources, *resource.Name)
			}
		}
	}
	return
}

func getGlueTriggerNames(client *glue.Client) (resources []string) {
	req := client.GetTriggersRequest(&glue.GetTriggersInput{})
	p := glue.NewGetTriggersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Triggers {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getGlueWorkflowNames(client *glue.Client) (resources []string) {
	req := client.ListWorkflowsRequest(&glue.ListWorkflowsInput{})
	p := glue.NewListWorkflowsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.Workflows...)
	}
	return
}

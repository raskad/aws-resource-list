package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func getGlue(config aws.Config) (resources resourceMap) {
	client := glue.New(config)

	glueDatabaseResourceMap := getGlueDatabase(client).unwrap(glueDatabase)
	glueDatabaseNames := glueDatabaseResourceMap[glueDatabase]

	resources = reduce(
		getGlueConnection(client).unwrap(glueConnection),
		getGlueCrawler(client).unwrap(glueCrawler),
		glueDatabaseResourceMap,
		getGlueDevEndpoint(client).unwrap(glueDevEndpoint),
		getGlueJob(client).unwrap(glueJob),
		getGlueMLTransform(client).unwrap(glueMLTransform),
		getGlueSecurityConfiguration(client).unwrap(glueSecurityConfiguration),
		getGlueTable(client, glueDatabaseNames).unwrap(glueTable),
		getGlueTrigger(client).unwrap(glueTrigger),
		getGlueWorkflow(client).unwrap(glueWorkflow),
	)
	return
}

func getGlueConnection(client *glue.Client) (r resourceSliceError) {
	req := client.GetConnectionsRequest(&glue.GetConnectionsInput{})
	p := glue.NewGetConnectionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ConnectionList {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getGlueCrawler(client *glue.Client) (r resourceSliceError) {
	req := client.GetCrawlersRequest(&glue.GetCrawlersInput{})
	p := glue.NewGetCrawlersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Crawlers {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getGlueDatabase(client *glue.Client) (r resourceSliceError) {
	req := client.GetDatabasesRequest(&glue.GetDatabasesInput{})
	p := glue.NewGetDatabasesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DatabaseList {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getGlueDevEndpoint(client *glue.Client) (r resourceSliceError) {
	req := client.GetDevEndpointsRequest(&glue.GetDevEndpointsInput{})
	p := glue.NewGetDevEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DevEndpoints {
			r.resources = append(r.resources, *resource.EndpointName)
		}
	}
	r.err = p.Err()
	return
}

func getGlueJob(client *glue.Client) (r resourceSliceError) {
	req := client.GetJobsRequest(&glue.GetJobsInput{})
	p := glue.NewGetJobsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Jobs {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getGlueMLTransform(client *glue.Client) (r resourceSliceError) {
	req := client.GetMLTransformsRequest(&glue.GetMLTransformsInput{})
	p := glue.NewGetMLTransformsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Transforms {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getGlueSecurityConfiguration(client *glue.Client) (r resourceSliceError) {
	req := client.GetSecurityConfigurationsRequest(&glue.GetSecurityConfigurationsInput{})
	p := glue.NewGetSecurityConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.SecurityConfigurations {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getGlueTable(client *glue.Client, databaseNames []string) (r resourceSliceError) {
	for _, databaseName := range databaseNames {
		req := client.GetTablesRequest(&glue.GetTablesInput{
			DatabaseName: aws.String(databaseName),
		})
		p := glue.NewGetTablesPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.TableList {
				r.resources = append(r.resources, *resource.Name)
			}
		}
		r.err = p.Err()
	}
	return
}

func getGlueTrigger(client *glue.Client) (r resourceSliceError) {
	req := client.GetTriggersRequest(&glue.GetTriggersInput{})
	p := glue.NewGetTriggersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Triggers {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getGlueWorkflow(client *glue.Client) (r resourceSliceError) {
	req := client.ListWorkflowsRequest(&glue.ListWorkflowsInput{})
	p := glue.NewListWorkflowsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Workflows {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

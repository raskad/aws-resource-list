package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

func getGlue(session *session.Session) (resources resourceMap) {
	client := glue.New(session)
	resources = reduce(
		getGlueConnection(client).unwrap(glueConnection),
		getGlueCrawler(client).unwrap(glueCrawler),
		getGlueDatabase(client).unwrap(glueDatabase),
		getGlueDevEndpoint(client).unwrap(glueDevEndpoint),
		getGlueJob(client).unwrap(glueJob),
		getGlueMLTransform(client).unwrap(glueMLTransform),
		getGlueSecurityConfiguration(client).unwrap(glueSecurityConfiguration),
		getGlueTable(client).unwrap(glueTable),
		getGlueTrigger(client).unwrap(glueTrigger),
		getGlueWorkflow(client).unwrap(glueWorkflow),
	)
	return
}

func getGlueConnection(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetConnectionsPages(&glue.GetConnectionsInput{}, func(page *glue.GetConnectionsOutput, lastPage bool) bool {
		for _, resource := range page.ConnectionList {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueCrawler(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetCrawlersPages(&glue.GetCrawlersInput{}, func(page *glue.GetCrawlersOutput, lastPage bool) bool {
		for _, resource := range page.Crawlers {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueDatabase(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetDatabasesPages(&glue.GetDatabasesInput{}, func(page *glue.GetDatabasesOutput, lastPage bool) bool {
		for _, resource := range page.DatabaseList {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueDevEndpoint(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetDevEndpointsPages(&glue.GetDevEndpointsInput{}, func(page *glue.GetDevEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.DevEndpoints {
			r.resources = append(r.resources, *resource.EndpointName)
		}
		return true
	})
	return
}

func getGlueJob(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetJobsPages(&glue.GetJobsInput{}, func(page *glue.GetJobsOutput, lastPage bool) bool {
		for _, resource := range page.Jobs {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueMLTransform(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetMLTransformsPages(&glue.GetMLTransformsInput{}, func(page *glue.GetMLTransformsOutput, lastPage bool) bool {
		for _, resource := range page.Transforms {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueSecurityConfiguration(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetSecurityConfigurationsPages(&glue.GetSecurityConfigurationsInput{}, func(page *glue.GetSecurityConfigurationsOutput, lastPage bool) bool {
		for _, resource := range page.SecurityConfigurations {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueTable(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetTablesPages(&glue.GetTablesInput{}, func(page *glue.GetTablesOutput, lastPage bool) bool {
		for _, resource := range page.TableList {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueTrigger(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetTriggersPages(&glue.GetTriggersInput{}, func(page *glue.GetTriggersOutput, lastPage bool) bool {
		for _, resource := range page.Triggers {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueWorkflow(client *glue.Glue) (r resourceSliceError) {
	r.err = client.ListWorkflowsPages(&glue.ListWorkflowsInput{}, func(page *glue.ListWorkflowsOutput, lastPage bool) bool {
		for _, resource := range page.Workflows {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

func getGlue(session *session.Session) (resources resourceMap) {
	client := glue.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		glueConnection:            getGlueConnection(client),
		glueCrawler:               getGlueCrawler(client),
		glueDatabase:              getGlueDatabase(client),
		glueDevEndpoint:           getGlueDevEndpoint(client),
		glueJob:                   getGlueJob(client),
		glueMLTransform:           getGlueMLTransform(client),
		glueSecurityConfiguration: getGlueSecurityConfiguration(client),
		glueTable:                 getGlueTable(client),
		glueTrigger:               getGlueTrigger(client),
		glueWorkflow:              getGlueWorkflow(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getGlueConnection(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetConnectionsPages(&glue.GetConnectionsInput{}, func(page *glue.GetConnectionsOutput, lastPage bool) bool {
		logDebug("Listing GlueConnection resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ConnectionList {
			logDebug("Got GlueConnection resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueCrawler(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetCrawlersPages(&glue.GetCrawlersInput{}, func(page *glue.GetCrawlersOutput, lastPage bool) bool {
		logDebug("Listing GlueCrawler resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Crawlers {
			logDebug("Got GlueCrawler resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueDatabase(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetDatabasesPages(&glue.GetDatabasesInput{}, func(page *glue.GetDatabasesOutput, lastPage bool) bool {
		logDebug("Listing GlueDatabase resources page. Remaining pages", page.NextToken)
		for _, resource := range page.DatabaseList {
			logDebug("Got GlueDatabase resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueDevEndpoint(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetDevEndpointsPages(&glue.GetDevEndpointsInput{}, func(page *glue.GetDevEndpointsOutput, lastPage bool) bool {
		logDebug("Listing GlueDevEndpoint resources page. Remaining pages", page.NextToken)
		for _, resource := range page.DevEndpoints {
			logDebug("Got GlueDevEndpoint resource with PhysicalResourceId", *resource.EndpointName)
			r.resources = append(r.resources, *resource.EndpointName)
		}
		return true
	})
	return
}

func getGlueJob(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetJobsPages(&glue.GetJobsInput{}, func(page *glue.GetJobsOutput, lastPage bool) bool {
		logDebug("Listing GlueJob resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Jobs {
			logDebug("Got GlueJob resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueMLTransform(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetMLTransformsPages(&glue.GetMLTransformsInput{}, func(page *glue.GetMLTransformsOutput, lastPage bool) bool {
		logDebug("Listing GlueMLTransform resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Transforms {
			logDebug("Got GlueMLTransform resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueSecurityConfiguration(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetSecurityConfigurationsPages(&glue.GetSecurityConfigurationsInput{}, func(page *glue.GetSecurityConfigurationsOutput, lastPage bool) bool {
		logDebug("Listing GlueSecurityConfiguration resources page. Remaining pages", page.NextToken)
		for _, resource := range page.SecurityConfigurations {
			logDebug("Got GlueSecurityConfiguration resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueTable(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetTablesPages(&glue.GetTablesInput{}, func(page *glue.GetTablesOutput, lastPage bool) bool {
		logDebug("Listing GlueTable resources page. Remaining pages", page.NextToken)
		for _, resource := range page.TableList {
			logDebug("Got GlueTable resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueTrigger(client *glue.Glue) (r resourceSliceError) {
	r.err = client.GetTriggersPages(&glue.GetTriggersInput{}, func(page *glue.GetTriggersOutput, lastPage bool) bool {
		logDebug("Listing GlueTrigger resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Triggers {
			logDebug("Got GlueTrigger resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueWorkflow(client *glue.Glue) (r resourceSliceError) {
	r.err = client.ListWorkflowsPages(&glue.ListWorkflowsInput{}, func(page *glue.ListWorkflowsOutput, lastPage bool) bool {
		logDebug("Listing GlueWorkflow resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Workflows {
			logDebug("Got GlueWorkflow resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

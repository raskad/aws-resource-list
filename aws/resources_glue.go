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
	logDebug("Listing GlueConnection resources")
	r.err = client.GetConnectionsPages(&glue.GetConnectionsInput{}, func(page *glue.GetConnectionsOutput, lastPage bool) bool {
		for _, resource := range page.ConnectionList {
			logDebug("Got GlueConnection resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueCrawler(client *glue.Glue) (r resourceSliceError) {
	logDebug("Listing GlueCrawler resources")
	r.err = client.GetCrawlersPages(&glue.GetCrawlersInput{}, func(page *glue.GetCrawlersOutput, lastPage bool) bool {
		for _, resource := range page.Crawlers {
			logDebug("Got GlueCrawler resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueDatabase(client *glue.Glue) (r resourceSliceError) {
	logDebug("Listing GlueDatabase resources")
	r.err = client.GetDatabasesPages(&glue.GetDatabasesInput{}, func(page *glue.GetDatabasesOutput, lastPage bool) bool {
		for _, resource := range page.DatabaseList {
			logDebug("Got GlueDatabase resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueDevEndpoint(client *glue.Glue) (r resourceSliceError) {
	logDebug("Listing GlueDevEndpoint resources")
	r.err = client.GetDevEndpointsPages(&glue.GetDevEndpointsInput{}, func(page *glue.GetDevEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.DevEndpoints {
			logDebug("Got GlueDevEndpoint resource with PhysicalResourceId", *resource.EndpointName)
			r.resources = append(r.resources, *resource.EndpointName)
		}
		return true
	})
	return
}

func getGlueJob(client *glue.Glue) (r resourceSliceError) {
	logDebug("Listing GlueJob resources")
	r.err = client.GetJobsPages(&glue.GetJobsInput{}, func(page *glue.GetJobsOutput, lastPage bool) bool {
		for _, resource := range page.Jobs {
			logDebug("Got GlueJob resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueMLTransform(client *glue.Glue) (r resourceSliceError) {
	logDebug("Listing GlueMLTransform resources")
	r.err = client.GetMLTransformsPages(&glue.GetMLTransformsInput{}, func(page *glue.GetMLTransformsOutput, lastPage bool) bool {
		for _, resource := range page.Transforms {
			logDebug("Got GlueMLTransform resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueSecurityConfiguration(client *glue.Glue) (r resourceSliceError) {
	logDebug("Listing GlueSecurityConfiguration resources")
	r.err = client.GetSecurityConfigurationsPages(&glue.GetSecurityConfigurationsInput{}, func(page *glue.GetSecurityConfigurationsOutput, lastPage bool) bool {
		for _, resource := range page.SecurityConfigurations {
			logDebug("Got GlueSecurityConfiguration resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueTable(client *glue.Glue) (r resourceSliceError) {
	logDebug("Listing GlueTable resources")
	r.err = client.GetTablesPages(&glue.GetTablesInput{}, func(page *glue.GetTablesOutput, lastPage bool) bool {
		for _, resource := range page.TableList {
			logDebug("Got GlueTable resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueTrigger(client *glue.Glue) (r resourceSliceError) {
	logDebug("Listing GlueTrigger resources")
	r.err = client.GetTriggersPages(&glue.GetTriggersInput{}, func(page *glue.GetTriggersOutput, lastPage bool) bool {
		for _, resource := range page.Triggers {
			logDebug("Got GlueTrigger resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getGlueWorkflow(client *glue.Glue) (r resourceSliceError) {
	logDebug("Listing GlueWorkflow resources")
	r.err = client.ListWorkflowsPages(&glue.ListWorkflowsInput{}, func(page *glue.ListWorkflowsOutput, lastPage bool) bool {
		for _, resource := range page.Workflows {
			logDebug("Got GlueWorkflow resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

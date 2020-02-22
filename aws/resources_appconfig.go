package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appconfig"
)

func getAppConfig(session *session.Session) (resources resourceMap) {
	client := appconfig.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		appConfigApplication:        getAppConfigApplication(client),
		appConfigDeploymentStrategy: getAppConfigDeploymentStrategy(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAppConfigApplication(client *appconfig.AppConfig) (r resourceSliceError) {
	r.err = client.ListApplicationsPages(&appconfig.ListApplicationsInput{}, func(page *appconfig.ListApplicationsOutput, lastPage bool) bool {
		logDebug("Listing AppConfigApplication resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Items {
			logDebug("Got AppConfigApplication resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getAppConfigDeploymentStrategy(client *appconfig.AppConfig) (r resourceSliceError) {
	r.err = client.ListDeploymentStrategiesPages(&appconfig.ListDeploymentStrategiesInput{}, func(page *appconfig.ListDeploymentStrategiesOutput, lastPage bool) bool {
		logDebug("Listing AppConfigDeploymentStrategy resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Items {
			logDebug("Got AppConfigDeploymentStrategy resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

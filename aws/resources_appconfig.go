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
	logDebug("Listing AppConfigApplication resources")
	r.err = client.ListApplicationsPages(&appconfig.ListApplicationsInput{}, func(page *appconfig.ListApplicationsOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			logDebug("Got AppConfigApplication resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getAppConfigDeploymentStrategy(client *appconfig.AppConfig) (r resourceSliceError) {
	logDebug("Listing AppConfigDeploymentStrategy resources")
	r.err = client.ListDeploymentStrategiesPages(&appconfig.ListDeploymentStrategiesInput{}, func(page *appconfig.ListDeploymentStrategiesOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			logDebug("Got AppConfigDeploymentStrategy resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appconfig"
)

func getAppConfig(session *session.Session) (resources resourceMap) {
	client := appconfig.New(session)
	resources = reduce(
		getAppConfigApplication(client).unwrap(appConfigApplication),
		getAppConfigDeploymentStrategy(client).unwrap(appConfigDeploymentStrategy),
	)
	return
}

func getAppConfigApplication(client *appconfig.AppConfig) (r resourceSliceError) {
	r.err = client.ListApplicationsPages(&appconfig.ListApplicationsInput{}, func(page *appconfig.ListApplicationsOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getAppConfigDeploymentStrategy(client *appconfig.AppConfig) (r resourceSliceError) {
	r.err = client.ListDeploymentStrategiesPages(&appconfig.ListDeploymentStrategiesInput{}, func(page *appconfig.ListDeploymentStrategiesOutput, lastPage bool) bool {
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

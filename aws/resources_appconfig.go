package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
)

func getAppConfig(config aws.Config) (resources resourceMap) {
	client := appconfig.New(config)

	appConfigApplicationNames := getAppConfigApplicationNames(client)
	appConfigDeploymentStrategyNames := getAppConfigDeploymentStrategyNames(client)

	resources = resourceMap{
		appConfigApplication:        appConfigApplicationNames,
		appConfigDeploymentStrategy: appConfigDeploymentStrategyNames,
	}
	return
}

func getAppConfigApplicationNames(client *appconfig.Client) (resources []string) {
	req := client.ListApplicationsRequest(&appconfig.ListApplicationsInput{})
	p := appconfig.NewListApplicationsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Items {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getAppConfigDeploymentStrategyNames(client *appconfig.Client) (resources []string) {
	req := client.ListDeploymentStrategiesRequest(&appconfig.ListDeploymentStrategiesInput{})
	p := appconfig.NewListDeploymentStrategiesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Items {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

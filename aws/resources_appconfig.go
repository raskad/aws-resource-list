package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
)

func getAppConfig(config aws.Config) (resources resourceMap) {
	client := appconfig.New(config)
	resources = reduce(
		getAppConfigApplication(client).unwrap(appConfigApplication),
		getAppConfigDeploymentStrategy(client).unwrap(appConfigDeploymentStrategy),
	)
	return
}

func getAppConfigApplication(client *appconfig.Client) (r resourceSliceError) {
	req := client.ListApplicationsRequest(&appconfig.ListApplicationsInput{})
	p := appconfig.NewListApplicationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getAppConfigDeploymentStrategy(client *appconfig.Client) (r resourceSliceError) {
	req := client.ListDeploymentStrategiesRequest(&appconfig.ListDeploymentStrategiesInput{})
	p := appconfig.NewListDeploymentStrategiesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Items {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

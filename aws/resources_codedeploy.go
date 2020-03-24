package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
)

func getCodeDeploy(config aws.Config) (resources resourceMap) {
	client := codedeploy.New(config)

	codeDeployApplicationNames := getCodeDeployApplicationNames(client)
	codeDeployDeploymentConfigNames := getCodeDeployDeploymentConfigNames(client)
	codeDeployDeploymentGroupNames := getCodeDeployDeploymentGroupNames(client, codeDeployApplicationNames)

	resources = resourceMap{
		codeDeployApplication:      codeDeployApplicationNames,
		codeDeployDeploymentConfig: codeDeployDeploymentConfigNames,
		codeDeployDeploymentGroup:  codeDeployDeploymentGroupNames,
	}
	return
}

func getCodeDeployApplicationNames(client *codedeploy.Client) (resources []string) {
	req := client.ListApplicationsRequest(&codedeploy.ListApplicationsInput{})
	p := codedeploy.NewListApplicationsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Applications {
			resources = append(resources, resource)
		}
	}
	return
}

func getCodeDeployDeploymentConfigNames(client *codedeploy.Client) (resources []string) {
	req := client.ListDeploymentConfigsRequest(&codedeploy.ListDeploymentConfigsInput{})
	p := codedeploy.NewListDeploymentConfigsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.DeploymentConfigsList {
			resources = append(resources, resource)
		}
	}
	return
}

func getCodeDeployDeploymentGroupNames(client *codedeploy.Client, applicationNames []string) (resources []string) {
	for _, applicationName := range applicationNames {
		req := client.ListDeploymentGroupsRequest(&codedeploy.ListDeploymentGroupsInput{
			ApplicationName: aws.String(applicationName),
		})
		p := codedeploy.NewListDeploymentGroupsPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.DeploymentGroups {
				resources = append(resources, resource)
			}
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
)

func getCodeDeploy(config aws.Config) (resources resourceMap) {
	client := codedeploy.New(config)

	codeDeployApplicationResourceMap := getCodeDeployApplication(client).unwrap(codeDeployApplication)
	codeDeployApplicationNames := codeDeployApplicationResourceMap[codeDeployApplication]

	resources = reduce(
		codeDeployApplicationResourceMap,
		getCodeDeployDeploymentConfig(client).unwrap(codeDeployDeploymentConfig),
		getCodeDeployDeploymentGroup(client, codeDeployApplicationNames).unwrap(codeDeployDeploymentGroup),
	)
	return
}

func getCodeDeployApplication(client *codedeploy.Client) (r resourceSliceError) {
	req := client.ListApplicationsRequest(&codedeploy.ListApplicationsInput{})
	p := codedeploy.NewListApplicationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Applications {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

func getCodeDeployDeploymentConfig(client *codedeploy.Client) (r resourceSliceError) {
	req := client.ListDeploymentConfigsRequest(&codedeploy.ListDeploymentConfigsInput{})
	p := codedeploy.NewListDeploymentConfigsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DeploymentConfigsList {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

func getCodeDeployDeploymentGroup(client *codedeploy.Client, applicationNames []string) (r resourceSliceError) {
	for _, applicationName := range applicationNames {
		req := client.ListDeploymentGroupsRequest(&codedeploy.ListDeploymentGroupsInput{
			ApplicationName: aws.String(applicationName),
		})
		p := codedeploy.NewListDeploymentGroupsPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.DeploymentGroups {
				r.resources = append(r.resources, resource)
			}
		}
		r.err = p.Err()
	}
	return
}

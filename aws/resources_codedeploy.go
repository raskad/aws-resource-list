package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codedeploy"
)

func getCodeDeploy(session *session.Session) (resources resourceMap) {
	client := codedeploy.New(session)

	codeDeployApplicationResourceMap := getCodeDeployApplication(client).unwrap(codeDeployApplication)
	codeDeployApplicationNames := codeDeployApplicationResourceMap[codeDeployApplication]

	resources = reduce(
		codeDeployApplicationResourceMap,
		getCodeDeployDeploymentConfig(client).unwrap(codeDeployDeploymentConfig),
		getCodeDeployDeploymentGroup(client, codeDeployApplicationNames).unwrap(codeDeployDeploymentGroup),
	)
	return
}

func getCodeDeployApplication(client *codedeploy.CodeDeploy) (r resourceSliceError) {
	r.err = client.ListApplicationsPages(&codedeploy.ListApplicationsInput{}, func(page *codedeploy.ListApplicationsOutput, lastPage bool) bool {
		for _, resource := range page.Applications {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getCodeDeployDeploymentConfig(client *codedeploy.CodeDeploy) (r resourceSliceError) {
	r.err = client.ListDeploymentConfigsPages(&codedeploy.ListDeploymentConfigsInput{}, func(page *codedeploy.ListDeploymentConfigsOutput, lastPage bool) bool {
		for _, resource := range page.DeploymentConfigsList {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getCodeDeployDeploymentGroup(client *codedeploy.CodeDeploy, applicationNames []string) (r resourceSliceError) {
	for _, applicationName := range applicationNames {
		r.err = client.ListDeploymentGroupsPages(&codedeploy.ListDeploymentGroupsInput{
			ApplicationName: aws.String(applicationName),
		}, func(page *codedeploy.ListDeploymentGroupsOutput, lastPage bool) bool {
			for _, resource := range page.DeploymentGroups {
				r.resources = append(r.resources, *resource)
			}
			return true
		})
	}
	return
}

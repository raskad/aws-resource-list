package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codedeploy"
)

func getCodeDeploy(session *session.Session) (resources resourceMap) {
	client := codedeploy.New(session)
	resources = reduce(
		getCodeDeployApplication(client).unwrap(codeDeployApplication),
		getCodeDeployDeploymentConfig(client).unwrap(codeDeployDeploymentConfig),
		getCodeDeployDeploymentGroup(client).unwrap(codeDeployDeploymentGroup),
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

func getCodeDeployDeploymentGroup(client *codedeploy.CodeDeploy) (r resourceSliceError) {
	r.err = client.ListDeploymentGroupsPages(&codedeploy.ListDeploymentGroupsInput{}, func(page *codedeploy.ListDeploymentGroupsOutput, lastPage bool) bool {
		for _, resource := range page.DeploymentGroups {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

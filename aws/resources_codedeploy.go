package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codedeploy"
)

func getCodeDeploy(session *session.Session) (resources resourceMap) {
	client := codedeploy.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		codeDeployApplication:      getCodeDeployApplication(client),
		codeDeployDeploymentConfig: getCodeDeployDeploymentConfig(client),
		codeDeployDeploymentGroup:  getCodeDeployDeploymentGroup(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCodeDeployApplication(client *codedeploy.CodeDeploy) (r resourceSliceError) {
	r.err = client.ListApplicationsPages(&codedeploy.ListApplicationsInput{}, func(page *codedeploy.ListApplicationsOutput, lastPage bool) bool {
		logDebug("List CodeDeployApplication resources page")
		for _, resource := range page.Applications {
			logDebug("Got CodeDeployApplication resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getCodeDeployDeploymentConfig(client *codedeploy.CodeDeploy) (r resourceSliceError) {
	r.err = client.ListDeploymentConfigsPages(&codedeploy.ListDeploymentConfigsInput{}, func(page *codedeploy.ListDeploymentConfigsOutput, lastPage bool) bool {
		logDebug("List CodeDeployDeploymentConfig resources page")
		for _, resource := range page.DeploymentConfigsList {
			logDebug("Got CodeDeployDeploymentConfig resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getCodeDeployDeploymentGroup(client *codedeploy.CodeDeploy) (r resourceSliceError) {
	r.err = client.ListDeploymentGroupsPages(&codedeploy.ListDeploymentGroupsInput{}, func(page *codedeploy.ListDeploymentGroupsOutput, lastPage bool) bool {
		logDebug("List CodeDeployDeploymentGroup resources page")
		for _, resource := range page.DeploymentGroups {
			logDebug("Got CodeDeployDeploymentGroup resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

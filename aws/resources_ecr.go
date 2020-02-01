package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func getEcr(session *session.Session) (resources resourceMap) {
	client := ecr.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		ecrRepository: getEcrRepository(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getEcrRepository(client *ecr.ECR) (r resourceSliceError) {
	r.err = client.DescribeRepositoriesPages(&ecr.DescribeRepositoriesInput{}, func(page *ecr.DescribeRepositoriesOutput, lastPage bool) bool {
		logDebug("Listing EcrRepository resources page. Remaining pages", page.Repositories)
		for _, resource := range page.Repositories {
			logDebug("Got EcrRepository resource with PhysicalResourceId", *resource.RepositoryName)
			r.resources = append(r.resources, *resource.RepositoryName)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func getEcr(session *session.Session) (resources resourceMap) {
	client := ecr.New(session)
	resources = reduce(
		getEcrRepository(client).unwrap(ecrRepository),
	)
	return
}

func getEcrRepository(client *ecr.ECR) (r resourceSliceError) {
	r.err = client.DescribeRepositoriesPages(&ecr.DescribeRepositoriesInput{}, func(page *ecr.DescribeRepositoriesOutput, lastPage bool) bool {
		for _, resource := range page.Repositories {
			r.resources = append(r.resources, *resource.RepositoryName)
		}
		return true
	})
	return
}

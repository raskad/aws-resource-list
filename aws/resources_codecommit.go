package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codecommit"
)

func getCodeCommit(session *session.Session) (resources resourceMap) {
	client := codecommit.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		codeCommitRepository: getCodeCommitRepository(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getCodeCommitRepository(client *codecommit.CodeCommit) (r resourceSliceError) {
	r.err = client.ListRepositoriesPages(&codecommit.ListRepositoriesInput{}, func(page *codecommit.ListRepositoriesOutput, lastPage bool) bool {
		logDebug("List CodeCommitRepository resources page")
		for _, resource := range page.Repositories {
			logDebug("Got CodeCommitRepository resource with PhysicalResourceId", *resource.RepositoryId)
			r.resources = append(r.resources, *resource.RepositoryId)
		}
		return true
	})
	return
}

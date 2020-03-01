package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codecommit"
)

func getCodeCommit(session *session.Session) (resources resourceMap) {
	client := codecommit.New(session)
	resources = reduce(
		getCodeCommitRepository(client).unwrap(codeCommitRepository),
	)
	return
}

func getCodeCommitRepository(client *codecommit.CodeCommit) (r resourceSliceError) {
	r.err = client.ListRepositoriesPages(&codecommit.ListRepositoriesInput{}, func(page *codecommit.ListRepositoriesOutput, lastPage bool) bool {
		for _, resource := range page.Repositories {
			r.resources = append(r.resources, *resource.RepositoryId)
		}
		return true
	})
	return
}

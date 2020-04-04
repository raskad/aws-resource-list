package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

func getCodeCommit(config aws.Config) (resources awsResourceMap) {
	client := codecommit.New(config)

	codeCommitRepositoryNames := getCodeCommitRepositoryNames(client)
	codeCommitTriggerNames := getCodeCommitTriggerNames(client, codeCommitRepositoryNames)

	resources = awsResourceMap{
		codeCommitRepository: codeCommitRepositoryNames,
		codeCommitTrigger:    codeCommitTriggerNames,
	}
	return
}

func getCodeCommitRepositoryNames(client *codecommit.Client) (resources []string) {
	req := client.ListRepositoriesRequest(&codecommit.ListRepositoriesInput{})
	p := codecommit.NewListRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Repositories {
			resources = append(resources, *resource.RepositoryName)
		}
	}
	return
}

func getCodeCommitTriggerNames(client *codecommit.Client, codeCommitRepositoryNames []string) (resources []string) {
	for _, codeCommitRepositoryName := range codeCommitRepositoryNames {
		page, err := client.GetRepositoryTriggersRequest(&codecommit.GetRepositoryTriggersInput{
			RepositoryName: &codeCommitRepositoryName,
		}).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Triggers {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

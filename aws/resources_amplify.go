package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

func getAmplify(config aws.Config) (resources awsResourceMap) {
	client := amplify.New(config)

	amplifyAppIDs := getAmplifyAppIDs(client)
	amplifyBranchNames := getAmplifyBranchNames(client, amplifyAppIDs)
	amplifyDomainNames := getAmplifyDomainNames(client, amplifyAppIDs)

	resources = awsResourceMap{
		amplifyApp:    amplifyAppIDs,
		amplifyBranch: amplifyBranchNames,
		amplifyDomain: amplifyDomainNames,
	}
	return
}

func getAmplifyAppIDs(client *amplify.Client) (resources []string) {
	input := amplify.ListAppsInput{}
	for {
		page, err := client.ListAppsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Apps {
			resources = append(resources, *resource.AppId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAmplifyBranchNames(client *amplify.Client, appIDs []string) (resources []string) {
	for _, appID := range appIDs {
		input := amplify.ListBranchesInput{
			AppId: aws.String(appID),
		}
		for {
			page, err := client.ListBranchesRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.Branches {
				resources = append(resources, *resource.BranchName)
			}
			if page.NextToken == nil {
				break
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getAmplifyDomainNames(client *amplify.Client, appIDs []string) (resources []string) {
	for _, appID := range appIDs {
		input := amplify.ListDomainAssociationsInput{
			AppId: aws.String(appID),
		}
		for {
			page, err := client.ListDomainAssociationsRequest(&input).Send(context.Background())
			if err != nil {
				logErr(err)
				break
			}
			for _, resource := range page.DomainAssociations {
				resources = append(resources, *resource.DomainName)
			}
			if page.NextToken == nil {
				break
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

func getAmplify(config aws.Config) (resources resourceMap) {
	client := amplify.New(config)

	amplifyAppResourceMap := getAmplifyApp(client).unwrap(amplifyApp)
	amplifyAppIDs := amplifyAppResourceMap[amplifyApp]

	resources = reduce(
		amplifyAppResourceMap,
		getAmplifyBranch(client, amplifyAppIDs).unwrap(amplifyBranch),
		getAmplifyDomain(client, amplifyAppIDs).unwrap(amplifyDomain),
	)
	return
}

func getAmplifyApp(client *amplify.Client) (r resourceSliceError) {
	input := amplify.ListAppsInput{}
	for {
		page, err := client.ListAppsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Apps {
			r.resources = append(r.resources, *resource.AppId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getAmplifyBranch(client *amplify.Client, appIDs []string) (r resourceSliceError) {
	for _, appID := range appIDs {
		input := amplify.ListBranchesInput{
			AppId: aws.String(appID),
		}
		for {
			page, err := client.ListBranchesRequest(&input).Send(context.Background())
			if err != nil {
				r.err = err
				return
			}
			for _, resource := range page.Branches {
				r.resources = append(r.resources, *resource.BranchName)
			}
			if page.NextToken == nil {
				break
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

func getAmplifyDomain(client *amplify.Client, appIDs []string) (r resourceSliceError) {
	for _, appID := range appIDs {
		input := amplify.ListDomainAssociationsInput{
			AppId: aws.String(appID),
		}
		for {
			page, err := client.ListDomainAssociationsRequest(&input).Send(context.Background())
			if err != nil {
				r.err = err
				return
			}
			for _, resource := range page.DomainAssociations {
				r.resources = append(r.resources, *resource.DomainName)
			}
			if page.NextToken == nil {
				break
			}
			input.NextToken = page.NextToken
		}
	}
	return
}

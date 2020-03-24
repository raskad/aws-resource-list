package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func getGameLift(config aws.Config) (resources resourceMap) {
	client := gamelift.New(config)

	getGameLiftAliasIDs := getGameLiftAliasIDs(client)
	getGameLiftBuildIDs := getGameLiftBuildIDs(client)
	getGameLiftFleetIDs := getGameLiftFleetIDs(client)
	getGameLiftGameSessionQueueNames := getGameLiftGameSessionQueueNames(client)
	getGameLiftMatchmakingConfigurationNames := getGameLiftMatchmakingConfigurationNames(client)
	getGameLiftMatchmakingRuleSetNames := getGameLiftMatchmakingRuleSetNames(client)
	getGameLiftScriptIDs := getGameLiftScriptIDs(client)

	resources = resourceMap{
		gameLiftAlias:                    getGameLiftAliasIDs,
		gameLiftBuild:                    getGameLiftBuildIDs,
		gameLiftFleet:                    getGameLiftFleetIDs,
		gameLiftGameSessionQueue:         getGameLiftGameSessionQueueNames,
		gameLiftMatchmakingConfiguration: getGameLiftMatchmakingConfigurationNames,
		gameLiftMatchmakingRuleSet:       getGameLiftMatchmakingRuleSetNames,
		gameLiftScript:                   getGameLiftScriptIDs,
	}
	return
}

func getGameLiftAliasIDs(client *gamelift.Client) (resources []string) {
	input := gamelift.ListAliasesInput{}
	for {
		page, err := client.ListAliasesRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Aliases {
			resources = append(resources, *resource.AliasId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftBuildIDs(client *gamelift.Client) (resources []string) {
	input := gamelift.ListBuildsInput{}
	for {
		page, err := client.ListBuildsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Builds {
			resources = append(resources, *resource.BuildId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftFleetIDs(client *gamelift.Client) (resources []string) {
	input := gamelift.ListFleetsInput{}
	for {
		page, err := client.ListFleetsRequest(&input).Send(context.Background())
		logErr(err)
		resources = append(resources, page.FleetIds...)
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftGameSessionQueueNames(client *gamelift.Client) (resources []string) {
	input := gamelift.DescribeGameSessionQueuesInput{}
	for {
		page, err := client.DescribeGameSessionQueuesRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.GameSessionQueues {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftMatchmakingConfigurationNames(client *gamelift.Client) (resources []string) {
	input := gamelift.DescribeMatchmakingConfigurationsInput{}
	for {
		page, err := client.DescribeMatchmakingConfigurationsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Configurations {
			resources = append(resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftMatchmakingRuleSetNames(client *gamelift.Client) (resources []string) {
	input := gamelift.DescribeMatchmakingRuleSetsInput{}
	for {
		page, err := client.DescribeMatchmakingRuleSetsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.RuleSets {
			resources = append(resources, *resource.RuleSetName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftScriptIDs(client *gamelift.Client) (resources []string) {
	input := gamelift.ListScriptsInput{}
	for {
		page, err := client.ListScriptsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Scripts {
			resources = append(resources, *resource.ScriptId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

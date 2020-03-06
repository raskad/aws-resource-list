package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func getGameLift(config aws.Config) (resources resourceMap) {
	client := gamelift.New(config)
	resources = reduce(
		getGameLiftAlias(client).unwrap(gameLiftAlias),
		getGameLiftBuild(client).unwrap(gameLiftBuild),
		getGameLiftFleet(client).unwrap(gameLiftFleet),
		getGameLiftGameSessionQueue(client).unwrap(gameLiftGameSessionQueue),
		getGameLiftMatchmakingConfiguration(client).unwrap(gameLiftMatchmakingConfiguration),
		getGameLiftMatchmakingRuleSet(client).unwrap(gameLiftMatchmakingRuleSet),
		getGameLiftScript(client).unwrap(gameLiftScript),
	)
	return
}

func getGameLiftAlias(client *gamelift.Client) (r resourceSliceError) {
	input := gamelift.ListAliasesInput{}
	for {
		page, err := client.ListAliasesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Aliases {
			r.resources = append(r.resources, *resource.AliasId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftBuild(client *gamelift.Client) (r resourceSliceError) {
	input := gamelift.ListBuildsInput{}
	for {
		page, err := client.ListBuildsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Builds {
			r.resources = append(r.resources, *resource.BuildId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftFleet(client *gamelift.Client) (r resourceSliceError) {
	input := gamelift.ListFleetsInput{}
	for {
		page, err := client.ListFleetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.FleetIds {
			r.resources = append(r.resources, resource)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftGameSessionQueue(client *gamelift.Client) (r resourceSliceError) {
	input := gamelift.DescribeGameSessionQueuesInput{}
	for {
		page, err := client.DescribeGameSessionQueuesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.GameSessionQueues {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftMatchmakingConfiguration(client *gamelift.Client) (r resourceSliceError) {
	input := gamelift.DescribeMatchmakingConfigurationsInput{}
	for {
		page, err := client.DescribeMatchmakingConfigurationsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Configurations {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftMatchmakingRuleSet(client *gamelift.Client) (r resourceSliceError) {
	input := gamelift.DescribeMatchmakingRuleSetsInput{}
	for {
		page, err := client.DescribeMatchmakingRuleSetsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.RuleSets {
			r.resources = append(r.resources, *resource.RuleSetName)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftScript(client *gamelift.Client) (r resourceSliceError) {
	input := gamelift.ListScriptsInput{}
	for {
		page, err := client.ListScriptsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Scripts {
			r.resources = append(r.resources, *resource.ScriptId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/gamelift"
)

func getGameLift(session *session.Session) (resources resourceMap) {
	client := gamelift.New(session)
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

func getGameLiftAlias(client *gamelift.GameLift) (r resourceSliceError) {
	input := gamelift.ListAliasesInput{}
	for {
		page, err := client.ListAliases(&input)
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

func getGameLiftBuild(client *gamelift.GameLift) (r resourceSliceError) {
	input := gamelift.ListBuildsInput{}
	for {
		page, err := client.ListBuilds(&input)
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

func getGameLiftFleet(client *gamelift.GameLift) (r resourceSliceError) {
	input := gamelift.ListFleetsInput{}
	for {
		page, err := client.ListFleets(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.FleetIds {
			r.resources = append(r.resources, *resource)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getGameLiftGameSessionQueue(client *gamelift.GameLift) (r resourceSliceError) {
	input := gamelift.DescribeGameSessionQueuesInput{}
	for {
		page, err := client.DescribeGameSessionQueues(&input)
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

func getGameLiftMatchmakingConfiguration(client *gamelift.GameLift) (r resourceSliceError) {
	input := gamelift.DescribeMatchmakingConfigurationsInput{}
	for {
		page, err := client.DescribeMatchmakingConfigurations(&input)
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

func getGameLiftMatchmakingRuleSet(client *gamelift.GameLift) (r resourceSliceError) {
	input := gamelift.DescribeMatchmakingRuleSetsInput{}
	for {
		page, err := client.DescribeMatchmakingRuleSets(&input)
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

func getGameLiftScript(client *gamelift.GameLift) (r resourceSliceError) {
	input := gamelift.ListScriptsInput{}
	for {
		page, err := client.ListScripts(&input)
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

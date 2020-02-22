package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/gamelift"
)

func getGameLift(session *session.Session) (resources resourceMap) {
	client := gamelift.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		gameLiftAlias:                    getGameLiftAlias(client),
		gameLiftBuild:                    getGameLiftBuild(client),
		gameLiftFleet:                    getGameLiftFleet(client),
		gameLiftGameSessionQueue:         getGameLiftGameSessionQueue(client),
		gameLiftMatchmakingConfiguration: getGameLiftMatchmakingConfiguration(client),
		gameLiftMatchmakingRuleSet:       getGameLiftMatchmakingRuleSet(client),
		gameLiftScript:                   getGameLiftScript(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
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
		logDebug("Listing GameLiftAlias resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Aliases {
			logDebug("Got GameLiftAlias resource with PhysicalResourceId", *resource.AliasId)
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
		logDebug("Listing GameLiftBuild resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Builds {
			logDebug("Got GameLiftBuild resource with PhysicalResourceId", *resource.BuildId)
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
		logDebug("Listing GameLiftFleet resources page. Remaining pages", page.NextToken)
		for _, resource := range page.FleetIds {
			logDebug("Got GameLiftFleet resource with PhysicalResourceId", *resource)
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
		logDebug("Listing GameLiftGameSessionQueue resources page. Remaining pages", page.NextToken)
		for _, resource := range page.GameSessionQueues {
			logDebug("Got GameLiftGameSessionQueue resource with PhysicalResourceId", *resource.Name)
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
		logDebug("Listing GameLiftMatchmakingConfiguration resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Configurations {
			logDebug("Got GameLiftMatchmakingConfiguration resource with PhysicalResourceId", *resource.Name)
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
		logDebug("Listing GameLiftMatchmakingRuleSet resources page. Remaining pages", page.NextToken)
		for _, resource := range page.RuleSets {
			logDebug("Got GameLiftMatchmakingRuleSet resource with PhysicalResourceId", *resource.RuleSetName)
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
		logDebug("Listing GameLiftScript resources page. Remaining pages", page.NextToken)
		for _, resource := range page.Scripts {
			logDebug("Got GameLiftScript resource with PhysicalResourceId", *resource.ScriptId)
			r.resources = append(r.resources, *resource.ScriptId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

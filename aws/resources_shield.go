package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
)

func getShield(config aws.Config) (resources awsResourceMap) {
	client := shield.New(config)

	shieldProtectionIDs := getShieldProtectionIDs(client)

	resources = awsResourceMap{
		shieldProtection: shieldProtectionIDs,
	}
	return
}

func getShieldProtectionIDs(client *shield.Client) (resources []string) {
	input := shield.ListProtectionsInput{}
	for {
		page, err := client.ListProtectionsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Protections {
			resources = append(resources, *resource.Id)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

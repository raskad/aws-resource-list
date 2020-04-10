package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness"
)

func getAlexaForBusiness(config aws.Config) (resources awsResourceMap) {
	client := alexaforbusiness.New(config)

	alexaAskSkillNames := getAlexaAskSkillNames(client)

	resources = awsResourceMap{
		alexaAskSkill: alexaAskSkillNames,
	}
	return
}

func getAlexaAskSkillNames(client *alexaforbusiness.Client) (resources []string) {
	req := client.ListSkillsRequest(&alexaforbusiness.ListSkillsInput{})
	p := alexaforbusiness.NewListSkillsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.SkillSummaries {
			resources = append(resources, *resource.SkillName)
		}
	}
	return
}

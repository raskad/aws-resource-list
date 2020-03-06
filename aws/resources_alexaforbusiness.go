package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness"
)

func getAlexaForBusiness(config aws.Config) (resources resourceMap) {
	client := alexaforbusiness.New(config)
	resources = reduce(
		getAlexaAskSkill(client).unwrap(alexaAskSkill),
	)
	return
}

func getAlexaAskSkill(client *alexaforbusiness.Client) (r resourceSliceError) {
	req := client.ListSkillsRequest(&alexaforbusiness.ListSkillsInput{})
	p := alexaforbusiness.NewListSkillsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.SkillSummaries {
			r.resources = append(r.resources, *resource.SkillName)
		}
	}
	r.err = p.Err()
	return
}

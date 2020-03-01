package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
)

func getAlexaForBusiness(session *session.Session) (resources resourceMap) {
	client := alexaforbusiness.New(session)
	resources = reduce(
		getAlexaAskSkill(client).unwrap(alexaAskSkill),
	)
	return
}

func getAlexaAskSkill(client *alexaforbusiness.AlexaForBusiness) (r resourceSliceError) {
	r.err = client.ListSkillsPages(&alexaforbusiness.ListSkillsInput{}, func(page *alexaforbusiness.ListSkillsOutput, lastPage bool) bool {
		for _, resource := range page.SkillSummaries {
			r.resources = append(r.resources, *resource.SkillName)
		}
		return true
	})
	return
}

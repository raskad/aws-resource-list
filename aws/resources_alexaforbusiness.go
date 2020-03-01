package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
)

func getAlexaForBusiness(session *session.Session) (resources resourceMap) {
	client := alexaforbusiness.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		alexaAskSkill: getAlexaAskSkill(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getAlexaAskSkill(client *alexaforbusiness.AlexaForBusiness) (r resourceSliceError) {
	logDebug("Listing AlexaAskSkill resources")
	r.err = client.ListSkillsPages(&alexaforbusiness.ListSkillsInput{}, func(page *alexaforbusiness.ListSkillsOutput, lastPage bool) bool {
		for _, resource := range page.SkillSummaries {
			logDebug("Got AlexaAskSkill resource with PhysicalResourceId", *resource.SkillName)
			r.resources = append(r.resources, *resource.SkillName)
		}
		return true
	})
	return
}

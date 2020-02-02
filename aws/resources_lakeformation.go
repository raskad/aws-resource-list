package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lakeformation"
)

func getLakeFormation(session *session.Session) (resources resourceMap) {
	client := lakeformation.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		lakeFormationResource: getLakeFormationResource(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getLakeFormationResource(client *lakeformation.LakeFormation) (r resourceSliceError) {
	r.err = client.ListResourcesPages(&lakeformation.ListResourcesInput{}, func(page *lakeformation.ListResourcesOutput, lastPage bool) bool {
		logDebug("Listing LakeFormationResource resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ResourceInfoList {
			logDebug("Got LakeFormationResource resource with PhysicalResourceId", *resource.ResourceArn)
			r.resources = append(r.resources, *resource.ResourceArn)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lakeformation"
)

func getLakeFormation(session *session.Session) (resources resourceMap) {
	client := lakeformation.New(session)
	resources = reduce(
		getLakeFormationResource(client).unwrap(lakeFormationResource),
	)
	return
}

func getLakeFormationResource(client *lakeformation.LakeFormation) (r resourceSliceError) {
	r.err = client.ListResourcesPages(&lakeformation.ListResourcesInput{}, func(page *lakeformation.ListResourcesOutput, lastPage bool) bool {
		for _, resource := range page.ResourceInfoList {
			r.resources = append(r.resources, *resource.ResourceArn)
		}
		return true
	})
	return
}

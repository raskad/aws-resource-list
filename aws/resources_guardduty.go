package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/guardduty"
)

func getGuardDuty(session *session.Session) (resources resourceMap) {
	client := guardduty.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		guardDutyDetector: getGuardDutyDetector(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getGuardDutyDetector(client *guardduty.GuardDuty) (r resourceSliceError) {
	logDebug("Listing GuardDutyDetector resources")
	r.err = client.ListDetectorsPages(&guardduty.ListDetectorsInput{}, func(page *guardduty.ListDetectorsOutput, lastPage bool) bool {
		for _, resource := range page.DetectorIds {
			logDebug("Got GuardDutyDetector resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

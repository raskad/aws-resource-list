package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/guardduty"
)

func getGuardDuty(session *session.Session) (resources resourceMap) {
	client := guardduty.New(session)
	resources = reduce(
		getGuardDutyDetector(client).unwrap(guardDutyDetector),
	)
	return
}

func getGuardDutyDetector(client *guardduty.GuardDuty) (r resourceSliceError) {
	r.err = client.ListDetectorsPages(&guardduty.ListDetectorsInput{}, func(page *guardduty.ListDetectorsOutput, lastPage bool) bool {
		for _, resource := range page.DetectorIds {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

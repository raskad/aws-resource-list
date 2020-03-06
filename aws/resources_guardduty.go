package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
)

func getGuardDuty(config aws.Config) (resources resourceMap) {
	client := guardduty.New(config)
	resources = reduce(
		getGuardDutyDetector(client).unwrap(guardDutyDetector),
	)
	return
}

func getGuardDutyDetector(client *guardduty.Client) (r resourceSliceError) {
	req := client.ListDetectorsRequest(&guardduty.ListDetectorsInput{})
	p := guardduty.NewListDetectorsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.DetectorIds {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

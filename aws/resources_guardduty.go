package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
)

func getGuardDuty(config aws.Config) (resources resourceMap) {
	client := guardduty.New(config)

	guardDutyDetectorIDs := getGuardDutyDetectorIDs(client)

	resources = resourceMap{
		guardDutyDetector: guardDutyDetectorIDs,
	}
	return
}

func getGuardDutyDetectorIDs(client *guardduty.Client) (resources []string) {
	req := client.ListDetectorsRequest(&guardduty.ListDetectorsInput{})
	p := guardduty.NewListDetectorsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		resources = append(resources, page.DetectorIds...)
	}
	return
}

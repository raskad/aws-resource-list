package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
)

func getGuardDuty(config aws.Config) (resources awsResourceMap) {
	client := guardduty.New(config)

	guardDutyDetectorIDs := getGuardDutyDetectorIDs(client)

	resources = awsResourceMap{
		guardDutyDetector: guardDutyDetectorIDs,
	}
	return
}

func getGuardDutyDetectorIDs(client *guardduty.Client) (resources []string) {
	req := client.ListDetectorsRequest(&guardduty.ListDetectorsInput{})
	p := guardduty.NewListDetectorsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.DetectorIds...)
	}
	return
}

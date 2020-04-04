package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/worklink"
)

func getWorkLink(config aws.Config) (resources awsResourceMap) {
	client := worklink.New(config)

	workLinkFleetNames := getWorkLinkFleetNames(client)

	resources = awsResourceMap{
		workLinkFleet: workLinkFleetNames,
	}

	return
}

func getWorkLinkFleetNames(client *worklink.Client) (resources []string) {
	req := client.ListFleetsRequest(&worklink.ListFleetsInput{})
	p := worklink.NewListFleetsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.FleetSummaryList {
			resources = append(resources, *resource.FleetName)
		}
	}
	return
}

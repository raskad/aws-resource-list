package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/detective"
)

func getDetective(config aws.Config) (resources awsResourceMap) {
	client := detective.New(config)

	detectiveInvitationIDs := getDetectiveInvitationIDs(client)
	detectiveGraphARNs := getDetectiveGraphARNs(client)

	resources = awsResourceMap{
		detectiveInvitation: detectiveInvitationIDs,
		detectiveGraph:      detectiveGraphARNs,
	}
	return
}

func getDetectiveInvitationIDs(client *detective.Client) (resources []string) {
	req := client.ListInvitationsRequest(&detective.ListInvitationsInput{})
	p := detective.NewListInvitationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Invitations {
			resources = append(resources, *resource.AccountId)
		}
	}
	return
}

func getDetectiveGraphARNs(client *detective.Client) (resources []string) {
	req := client.ListGraphsRequest(&detective.ListGraphsInput{})
	p := detective.NewListGraphsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.GraphList {
			resources = append(resources, *resource.Arn)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
)

func getAthena(config aws.Config) (resources awsResourceMap) {
	client := athena.New(config)

	athenaNamedQueryIDs := getAthenaNamedQueryIDs(client)
	athenaWorkGroupNames := getAthenaWorkGroupNames(client)

	resources = awsResourceMap{
		athenaNamedQuery: athenaNamedQueryIDs,
		athenaWorkGroup:  athenaWorkGroupNames,
	}
	return
}

func getAthenaNamedQueryIDs(client *athena.Client) (resources []string) {
	req := client.ListNamedQueriesRequest(&athena.ListNamedQueriesInput{})
	p := athena.NewListNamedQueriesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.NamedQueryIds...)
	}
	return
}

func getAthenaWorkGroupNames(client *athena.Client) (resources []string) {
	req := client.ListWorkGroupsRequest(&athena.ListWorkGroupsInput{})
	p := athena.NewListWorkGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.WorkGroups {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

func getDataSync(config aws.Config) (resources awsResourceMap) {
	client := datasync.New(config)

	dataSyncAgentIDs := getDataSyncAgentIDs(client)
	dataSyncLocationARNs := getDataSyncLocationARNs(client)
	dataSyncTaskARNs := getDataSyncTaskARNs(client)

	resources = awsResourceMap{
		dataSyncAgent:    dataSyncAgentIDs,
		dataSyncLocation: dataSyncLocationARNs,
		dataSyncTask:     dataSyncTaskARNs,
	}
	return
}

func getDataSyncAgentIDs(client *datasync.Client) (resources []string) {
	req := client.ListAgentsRequest(&datasync.ListAgentsInput{})
	p := datasync.NewListAgentsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Agents {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getDataSyncLocationARNs(client *datasync.Client) (resources []string) {
	req := client.ListLocationsRequest(&datasync.ListLocationsInput{})
	p := datasync.NewListLocationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Locations {
			resources = append(resources, *resource.LocationArn)
		}
	}
	return
}

func getDataSyncTaskARNs(client *datasync.Client) (resources []string) {
	req := client.ListTasksRequest(&datasync.ListTasksInput{})
	p := datasync.NewListTasksPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Tasks {
			resources = append(resources, *resource.TaskArn)
		}
	}
	return
}

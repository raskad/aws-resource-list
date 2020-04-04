package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
)

func getDeviceFarm(config aws.Config) (resources awsResourceMap) {
	client := devicefarm.New(config)

	deviceFarmProjectNames := getDeviceFarmProjectNames(client)

	resources = awsResourceMap{
		deviceFarmProject: deviceFarmProjectNames,
	}
	return
}

func getDeviceFarmProjectNames(client *devicefarm.Client) (resources []string) {
	req := client.ListProjectsRequest(&devicefarm.ListProjectsInput{})
	p := devicefarm.NewListProjectsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Projects {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
)

func getCloud9(config aws.Config) (resources awsResourceMap) {
	client := cloud9.New(config)

	cloud9EnvironmentEC2IDs := getCloud9EnvironmentEC2IDs(client)

	resources = awsResourceMap{
		cloud9EnvironmentEC2: cloud9EnvironmentEC2IDs,
	}
	return
}

func getCloud9EnvironmentEC2IDs(client *cloud9.Client) (resources []string) {
	req := client.ListEnvironmentsRequest(&cloud9.ListEnvironmentsInput{})
	p := cloud9.NewListEnvironmentsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.EnvironmentIds...)
	}
	return
}

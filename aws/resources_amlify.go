package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

func getAmplify(config aws.Config) (resources resourceMap) {
	client := amplify.New(config)
	resources = reduce(
		getAmplifyApp(client).unwrap(amplifyApp),
	)
	return
}

func getAmplifyApp(client *amplify.Client) (r resourceSliceError) {
	input := amplify.ListAppsInput{}
	for {
		page, err := client.ListAppsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Apps {
			r.resources = append(r.resources, *resource.Name)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

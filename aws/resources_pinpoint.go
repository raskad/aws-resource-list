package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

func getPinpoint(config aws.Config) (resources resourceMap) {
	client := pinpoint.New(config)
	resources = reduce(
		getPinpointApp(client).unwrap(pinpointApp),
		getPinpointEmailTemplate(client).unwrap(pinpointEmailTemplate),
		getPinpointPushTemplate(client).unwrap(pinpointPushTemplate),
		getPinpointSmsTemplate(client).unwrap(pinpointSmsTemplate),
	)
	return
}

func getPinpointApp(client *pinpoint.Client) (r resourceSliceError) {
	input := pinpoint.GetAppsInput{}
	for {
		page, err := client.GetAppsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ApplicationsResponse.Item {
			r.resources = append(r.resources, *resource.Id)
		}
		if page.ApplicationsResponse.NextToken == nil {
			return
		}
		input.Token = page.ApplicationsResponse.NextToken
	}
}

func getPinpointEmailTemplate(client *pinpoint.Client) (r resourceSliceError) {
	input := pinpoint.ListTemplatesInput{}
	for {
		page, err := client.ListTemplatesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.TemplatesResponse.Item {
			if resource.TemplateType == pinpoint.TemplateTypeEmail {
				r.resources = append(r.resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

func getPinpointPushTemplate(client *pinpoint.Client) (r resourceSliceError) {
	input := pinpoint.ListTemplatesInput{}
	for {
		page, err := client.ListTemplatesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.TemplatesResponse.Item {
			if resource.TemplateType == pinpoint.TemplateTypePush {
				r.resources = append(r.resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

func getPinpointSmsTemplate(client *pinpoint.Client) (r resourceSliceError) {
	input := pinpoint.ListTemplatesInput{}
	for {
		page, err := client.ListTemplatesRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.TemplatesResponse.Item {
			if resource.TemplateType == pinpoint.TemplateTypeSms {
				r.resources = append(r.resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

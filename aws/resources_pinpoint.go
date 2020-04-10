package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

func getPinpoint(config aws.Config) (resources awsResourceMap) {
	client := pinpoint.New(config)

	pinpointAppIDs := getPinpointAppIDs(client)
	pinpointEmailTemplateNames := getPinpointEmailTemplateNames(client)
	pinpointPushTemplateNames := getPinpointPushTemplateNames(client)
	pinpointSmsTemplateNames := getPinpointSmsTemplateNames(client)

	resources = awsResourceMap{
		pinpointApp:           pinpointAppIDs,
		pinpointEmailTemplate: pinpointEmailTemplateNames,
		pinpointPushTemplate:  pinpointPushTemplateNames,
		pinpointSmsTemplate:   pinpointSmsTemplateNames,
	}
	return
}

func getPinpointAppIDs(client *pinpoint.Client) (resources []string) {
	input := pinpoint.GetAppsInput{}
	for {
		page, err := client.GetAppsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.ApplicationsResponse.Item {
			resources = append(resources, *resource.Id)
		}
		if page.ApplicationsResponse.NextToken == nil {
			return
		}
		input.Token = page.ApplicationsResponse.NextToken
	}
}

func getPinpointEmailTemplateNames(client *pinpoint.Client) (resources []string) {
	input := pinpoint.ListTemplatesInput{}
	for {
		page, err := client.ListTemplatesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.TemplatesResponse.Item {
			if resource.TemplateType == pinpoint.TemplateTypeEmail {
				resources = append(resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

func getPinpointPushTemplateNames(client *pinpoint.Client) (resources []string) {
	input := pinpoint.ListTemplatesInput{}
	for {
		page, err := client.ListTemplatesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.TemplatesResponse.Item {
			if resource.TemplateType == pinpoint.TemplateTypePush {
				resources = append(resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

func getPinpointSmsTemplateNames(client *pinpoint.Client) (resources []string) {
	input := pinpoint.ListTemplatesInput{}
	for {
		page, err := client.ListTemplatesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.TemplatesResponse.Item {
			if resource.TemplateType == pinpoint.TemplateTypeSms {
				resources = append(resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

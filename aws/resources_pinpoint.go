package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpoint"
)

func getPinpoint(session *session.Session) (resources resourceMap) {
	client := pinpoint.New(session)
	resources = reduce(
		getPinpointApp(client).unwrap(pinpointApp),
		getPinpointEmailTemplate(client).unwrap(pinpointEmailTemplate),
		getPinpointPushTemplate(client).unwrap(pinpointPushTemplate),
		getPinpointSmsTemplate(client).unwrap(pinpointSmsTemplate),
	)
	return
}

func getPinpointApp(client *pinpoint.Pinpoint) (r resourceSliceError) {
	input := pinpoint.GetAppsInput{}
	for {
		page, err := client.GetApps(&input)
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

func getPinpointEmailTemplate(client *pinpoint.Pinpoint) (r resourceSliceError) {
	input := pinpoint.ListTemplatesInput{}
	for {
		page, err := client.ListTemplates(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.TemplatesResponse.Item {
			if *resource.TemplateType == pinpoint.TemplateTypeEmail {
				r.resources = append(r.resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

func getPinpointPushTemplate(client *pinpoint.Pinpoint) (r resourceSliceError) {
	input := pinpoint.ListTemplatesInput{}
	for {
		page, err := client.ListTemplates(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.TemplatesResponse.Item {
			if *resource.TemplateType == pinpoint.TemplateTypePush {
				r.resources = append(r.resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

func getPinpointSmsTemplate(client *pinpoint.Pinpoint) (r resourceSliceError) {
	input := pinpoint.ListTemplatesInput{}
	for {
		page, err := client.ListTemplates(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.TemplatesResponse.Item {
			if *resource.TemplateType == pinpoint.TemplateTypeSms {
				r.resources = append(r.resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

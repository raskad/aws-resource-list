package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpoint"
)

func getPinpoint(session *session.Session) (resources resourceMap) {
	client := pinpoint.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		pinpointApp:           getPinpointApp(client),
		pinpointEmailTemplate: getPinpointEmailTemplate(client),
		pinpointPushTemplate:  getPinpointPushTemplate(client),
		pinpointSmsTemplate:   getPinpointSmsTemplate(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
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
		logDebug("Listing PinpointApp resources page. Remaining pages", page.ApplicationsResponse.NextToken)
		for _, resource := range page.ApplicationsResponse.Item {
			logDebug("Got PinpointApp resource with PhysicalResourceId", *resource.Id)
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
		logDebug("Listing PinpointEmailTemplate resources page. Remaining pages", page.TemplatesResponse.NextToken)
		for _, resource := range page.TemplatesResponse.Item {
			if *resource.TemplateType == pinpoint.TemplateTypeEmail {
				logDebug("Got PinpointEmailTemplate resource with PhysicalResourceId", *resource.TemplateName)
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
		logDebug("Listing PinpointPushTemplate resources page. Remaining pages", page.TemplatesResponse.NextToken)
		for _, resource := range page.TemplatesResponse.Item {
			if *resource.TemplateType == pinpoint.TemplateTypePush {
				logDebug("Got PinpointPushTemplate resource with PhysicalResourceId", *resource.TemplateName)
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
		logDebug("Listing PinpointSmsTemplate resources page. Remaining pages", page.TemplatesResponse.NextToken)
		for _, resource := range page.TemplatesResponse.Item {
			if *resource.TemplateType == pinpoint.TemplateTypeSms {
				logDebug("Got PinpointSmsTemplate resource with PhysicalResourceId", *resource.TemplateName)
				r.resources = append(r.resources, *resource.TemplateName)
			}
		}
		if page.TemplatesResponse.NextToken == nil {
			return
		}
		input.NextToken = page.TemplatesResponse.NextToken
	}
}

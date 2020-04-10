package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func getLambda(config aws.Config) (resources awsResourceMap) {
	client := lambda.New(config)

	lambdaFunctionNames := getLambdaFunctionNames(client)
	lambdaLayerNames := getLambdaLayerNames(client)
	lambdaAliasNames := getLambdaAliasNames(client, lambdaFunctionNames)
	lambdaEventSourceMappingIDs := getLambdaEventSourceMappingIDs(client)
	lambdaLayerVersionARNs := getLambdaLayerVersionARNs(client, lambdaLayerNames)

	resources = awsResourceMap{
		lambdaAlias:              lambdaAliasNames,
		lambdaEventSourceMapping: lambdaEventSourceMappingIDs,
		lambdaFunction:           lambdaFunctionNames,
		lambdaLayer:              lambdaLayerNames,
		lambdaLayerVersion:       lambdaLayerVersionARNs,
	}
	return
}

func getLambdaAliasNames(client *lambda.Client, lambdaFunctionNames []string) (resources []string) {
	for _, lambdaFunctionName := range lambdaFunctionNames {
		req := client.ListAliasesRequest(&lambda.ListAliasesInput{
			FunctionName: aws.String(lambdaFunctionName),
		})
		p := lambda.NewListAliasesPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.Aliases {
				resources = append(resources, *resource.Name)
			}
		}
	}
	return
}

func getLambdaEventSourceMappingIDs(client *lambda.Client) (resources []string) {
	req := client.ListEventSourceMappingsRequest(&lambda.ListEventSourceMappingsInput{})
	p := lambda.NewListEventSourceMappingsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.EventSourceMappings {
			resources = append(resources, *resource.UUID)
		}
	}
	return
}

func getLambdaFunctionNames(client *lambda.Client) (resources []string) {
	req := client.ListFunctionsRequest(&lambda.ListFunctionsInput{})
	p := lambda.NewListFunctionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Functions {
			resources = append(resources, *resource.FunctionName)
		}
	}
	return
}

func getLambdaLayerNames(client *lambda.Client) (resources []string) {
	req := client.ListLayersRequest(&lambda.ListLayersInput{})
	p := lambda.NewListLayersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Layers {
			resources = append(resources, *resource.LayerName)
		}
	}
	return
}

func getLambdaLayerVersionARNs(client *lambda.Client, lambdaLayerNames []string) (resources []string) {
	for _, lambdaLayerName := range lambdaLayerNames {
		req := client.ListLayerVersionsRequest(&lambda.ListLayerVersionsInput{
			LayerName: aws.String(lambdaLayerName),
		})
		p := lambda.NewListLayerVersionsPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.LayerVersions {
				resources = append(resources, *resource.LayerVersionArn)
			}
		}
	}
	return
}

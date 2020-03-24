package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func getLambda(config aws.Config) (resources resourceMap) {
	client := lambda.New(config)

	lambdaFunctionNames := getLambdaFunctionNames(client)
	lambdaLayerNames := getLambdaLayerNames(client)
	lambdaAliasNames := getLambdaAliasNames(client, lambdaFunctionNames)
	lambdaLayerVersionARNs := getLambdaLayerVersionARNs(client, lambdaLayerNames)

	resources = resourceMap{
		lambdaFunction:     lambdaFunctionNames,
		lambdaLayer:        lambdaLayerNames,
		lambdaAlias:        lambdaAliasNames,
		lambdaLayerVersion: lambdaLayerVersionARNs,
	}
	return
}

func getLambdaFunctionNames(client *lambda.Client) (resources []string) {
	req := client.ListFunctionsRequest(&lambda.ListFunctionsInput{})
	p := lambda.NewListFunctionsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Functions {
			resources = append(resources, *resource.FunctionName)
		}
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
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.Aliases {
				resources = append(resources, *resource.Name)
			}
		}
		return
	}
	return
}

func getLambdaLayerNames(client *lambda.Client) (resources []string) {
	req := client.ListLayersRequest(&lambda.ListLayersInput{})
	p := lambda.NewListLayersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
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
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.LayerVersions {
				resources = append(resources, *resource.LayerVersionArn)
			}
		}
		return
	}
	return
}

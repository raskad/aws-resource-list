package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func getLambda(config aws.Config) (resources resourceMap) {
	client := lambda.New(config)

	lambdaFunctionResourceMap := getLambdaFunction(client).unwrap(lambdaFunction)
	lambdaFunctionNames := lambdaFunctionResourceMap[lambdaFunction]
	lambdaLayerResourcesMap := getLambdaLayer(client).unwrap(lambdaLayer)
	lambdaLayerNames := lambdaLayerResourcesMap[lambdaLayer]

	resources = reduce(
		lambdaFunctionResourceMap,
		getLambdaAlias(client, lambdaFunctionNames).unwrap(lambdaAlias),
		getLambdaLayerVersion(client, lambdaLayerNames).unwrap(lambdaLayerVersion),
	)
	return
}

func getLambdaFunction(client *lambda.Client) (r resourceSliceError) {
	req := client.ListFunctionsRequest(&lambda.ListFunctionsInput{})
	p := lambda.NewListFunctionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Functions {
			r.resources = append(r.resources, *resource.FunctionName)
		}
	}
	r.err = p.Err()
	return
}

func getLambdaAlias(client *lambda.Client, lambdaFunctionNames []string) (r resourceSliceError) {
	for _, lambdaFunctionName := range lambdaFunctionNames {
		req := client.ListAliasesRequest(&lambda.ListAliasesInput{
			FunctionName: aws.String(lambdaFunctionName),
		})
		p := lambda.NewListAliasesPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.Aliases {
				r.resources = append(r.resources, *resource.Name)
			}
		}
		r.err = p.Err()
		return
	}
	return
}

func getLambdaLayer(client *lambda.Client) (r resourceSliceError) {
	req := client.ListLayersRequest(&lambda.ListLayersInput{})
	p := lambda.NewListLayersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Layers {
			r.resources = append(r.resources, *resource.LayerName)
		}
	}
	r.err = p.Err()
	return
}

func getLambdaLayerVersion(client *lambda.Client, lambdaLayerNames []string) (r resourceSliceError) {
	for _, lambdaLayerName := range lambdaLayerNames {
		req := client.ListLayerVersionsRequest(&lambda.ListLayerVersionsInput{
			LayerName: aws.String(lambdaLayerName),
		})
		p := lambda.NewListLayerVersionsPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.LayerVersions {
				r.resources = append(r.resources, *resource.LayerVersionArn)
			}
		}
		r.err = p.Err()
		return
	}
	return
}

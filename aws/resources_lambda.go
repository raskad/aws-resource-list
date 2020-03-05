package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func getLambda(session *session.Session) (resources resourceMap) {
	client := lambda.New(session)

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

func getLambdaFunction(client *lambda.Lambda) (r resourceSliceError) {
	r.err = client.ListFunctionsPages(&lambda.ListFunctionsInput{}, func(page *lambda.ListFunctionsOutput, lastPage bool) bool {
		for _, resource := range page.Functions {
			r.resources = append(r.resources, *resource.FunctionName)
		}
		return true
	})
	return
}

func getLambdaAlias(client *lambda.Lambda, lambdaFunctionNames []string) (r resourceSliceError) {
	for _, lambdaFunctionName := range lambdaFunctionNames {
		r.err = client.ListAliasesPages(&lambda.ListAliasesInput{
			FunctionName: aws.String(lambdaFunctionName),
		}, func(page *lambda.ListAliasesOutput, lastPage bool) bool {
			for _, resource := range page.Aliases {
				r.resources = append(r.resources, *resource.Name)
			}
			return true
		})
	}
	return
}

func getLambdaLayer(client *lambda.Lambda) (r resourceSliceError) {
	r.err = client.ListLayersPages(&lambda.ListLayersInput{}, func(page *lambda.ListLayersOutput, lastPage bool) bool {
		for _, resource := range page.Layers {
			r.resources = append(r.resources, *resource.LayerName)
		}
		return true
	})
	return
}

func getLambdaLayerVersion(client *lambda.Lambda, lambdaLayerNames []string) (r resourceSliceError) {
	for _, lambdaLayerName := range lambdaLayerNames {
		r.err = client.ListLayerVersionsPages(&lambda.ListLayerVersionsInput{
			LayerName: aws.String(lambdaLayerName),
		}, func(page *lambda.ListLayerVersionsOutput, lastPage bool) bool {
			for _, resource := range page.LayerVersions {
				r.resources = append(r.resources, *resource.LayerVersionArn)
			}
			return true
		})
	}
	return
}

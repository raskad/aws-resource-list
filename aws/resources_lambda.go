package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func getLambda(session *session.Session) (resources resourceMap) {
	client := lambda.New(session)
	resources = reduce(
		getLambdaAlias(client).unwrap(lambdaAlias),
		getLambdaFunction(client).unwrap(lambdaFunction),
		getLambdaLayerVersion(client).unwrap(lambdaLayerVersion),
	)
	return
}

func getLambdaAlias(client *lambda.Lambda) (r resourceSliceError) {
	r.err = client.ListAliasesPages(&lambda.ListAliasesInput{}, func(page *lambda.ListAliasesOutput, lastPage bool) bool {
		for _, resource := range page.Aliases {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
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

func getLambdaLayerVersion(client *lambda.Lambda) (r resourceSliceError) {
	r.err = client.ListLayerVersionsPages(&lambda.ListLayerVersionsInput{}, func(page *lambda.ListLayerVersionsOutput, lastPage bool) bool {
		for _, resource := range page.LayerVersions {
			r.resources = append(r.resources, *resource.LayerVersionArn)
		}
		return true
	})
	return
}

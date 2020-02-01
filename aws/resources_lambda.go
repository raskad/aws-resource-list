package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func getLambda(session *session.Session) (resources resourceMap) {
	client := lambda.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		lambdaAlias:        getLambdaAlias(client),
		lambdaFunction:     getLambdaFunction(client),
		lambdaLayerVersion: getLambdaLayerVersion(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getLambdaAlias(client *lambda.Lambda) (r resourceSliceError) {
	r.err = client.ListAliasesPages(&lambda.ListAliasesInput{}, func(page *lambda.ListAliasesOutput, lastPage bool) bool {
		logDebug("Listing LambdaAlias resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.Aliases {
			logDebug("Got LambdaAlias resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getLambdaFunction(client *lambda.Lambda) (r resourceSliceError) {
	r.err = client.ListFunctionsPages(&lambda.ListFunctionsInput{}, func(page *lambda.ListFunctionsOutput, lastPage bool) bool {
		logDebug("Listing LambdaFunction resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.Functions {
			logDebug("Got LambdaFunction resource with PhysicalResourceId", *resource.FunctionName)
			r.resources = append(r.resources, *resource.FunctionName)
		}
		return true
	})
	return
}

func getLambdaLayerVersion(client *lambda.Lambda) (r resourceSliceError) {
	r.err = client.ListLayerVersionsPages(&lambda.ListLayerVersionsInput{}, func(page *lambda.ListLayerVersionsOutput, lastPage bool) bool {
		logDebug("Listing LambdaLayerVersion resources page. Remaining pages", page.NextMarker)
		for _, resource := range page.LayerVersions {
			logDebug("Got LambdaLayerVersion resource with PhysicalResourceId", *resource.LayerVersionArn)
			r.resources = append(r.resources, *resource.LayerVersionArn)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemaker"
)

func getSageMaker(session *session.Session) (resources resourceMap) {
	client := sagemaker.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		sageMakerCodeRepository:                  getSageMakerCodeRepository(client),
		sageMakerEndpoint:                        getSageMakerEndpoint(client),
		sageMakerEndpointConfig:                  getSageMakerEndpointConfig(client),
		sageMakerModel:                           getSageMakerModel(client),
		sageMakerNotebookInstance:                getSageMakerNotebookInstance(client),
		sageMakerNotebookInstanceLifecycleConfig: getSageMakerNotebookInstanceLifecycleConfig(client),
		sageMakerWorkteam:                        getSageMakerWorkteam(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getSageMakerCodeRepository(client *sagemaker.SageMaker) (r resourceSliceError) {
	logDebug("Listing SageMakerCodeRepository resources")
	r.err = client.ListCodeRepositoriesPages(&sagemaker.ListCodeRepositoriesInput{}, func(page *sagemaker.ListCodeRepositoriesOutput, lastPage bool) bool {
		for _, resource := range page.CodeRepositorySummaryList {
			logDebug("Got SageMakerCodeRepository resource with PhysicalResourceId", *resource.CodeRepositoryName)
			r.resources = append(r.resources, *resource.CodeRepositoryName)
		}
		return true
	})
	return
}

func getSageMakerEndpoint(client *sagemaker.SageMaker) (r resourceSliceError) {
	logDebug("Listing SageMakerEndpoint resources")
	r.err = client.ListEndpointsPages(&sagemaker.ListEndpointsInput{}, func(page *sagemaker.ListEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.Endpoints {
			logDebug("Got SageMakerEndpoint resource with PhysicalResourceId", *resource.EndpointName)
			r.resources = append(r.resources, *resource.EndpointName)
		}
		return true
	})
	return
}

func getSageMakerEndpointConfig(client *sagemaker.SageMaker) (r resourceSliceError) {
	logDebug("Listing SageMakerEndpointConfig resources")
	r.err = client.ListEndpointConfigsPages(&sagemaker.ListEndpointConfigsInput{}, func(page *sagemaker.ListEndpointConfigsOutput, lastPage bool) bool {
		for _, resource := range page.EndpointConfigs {
			logDebug("Got SageMakerEndpointConfig resource with PhysicalResourceId", *resource.EndpointConfigName)
			r.resources = append(r.resources, *resource.EndpointConfigName)
		}
		return true
	})
	return
}

func getSageMakerModel(client *sagemaker.SageMaker) (r resourceSliceError) {
	logDebug("Listing SageMakerModel resources")
	r.err = client.ListModelsPages(&sagemaker.ListModelsInput{}, func(page *sagemaker.ListModelsOutput, lastPage bool) bool {
		for _, resource := range page.Models {
			logDebug("Got SageMakerModel resource with PhysicalResourceId", *resource.ModelName)
			r.resources = append(r.resources, *resource.ModelName)
		}
		return true
	})
	return
}

func getSageMakerNotebookInstance(client *sagemaker.SageMaker) (r resourceSliceError) {
	logDebug("Listing SageMakerNotebookInstance resources")
	r.err = client.ListNotebookInstancesPages(&sagemaker.ListNotebookInstancesInput{}, func(page *sagemaker.ListNotebookInstancesOutput, lastPage bool) bool {
		for _, resource := range page.NotebookInstances {
			logDebug("Got SageMakerNotebookInstance resource with PhysicalResourceId", *resource.NotebookInstanceName)
			r.resources = append(r.resources, *resource.NotebookInstanceName)
		}
		return true
	})
	return
}

func getSageMakerNotebookInstanceLifecycleConfig(client *sagemaker.SageMaker) (r resourceSliceError) {
	logDebug("Listing SageMakerNotebookInstanceLifecycleConfig resources")
	r.err = client.ListNotebookInstanceLifecycleConfigsPages(&sagemaker.ListNotebookInstanceLifecycleConfigsInput{}, func(page *sagemaker.ListNotebookInstanceLifecycleConfigsOutput, lastPage bool) bool {
		for _, resource := range page.NotebookInstanceLifecycleConfigs {
			logDebug("Got SageMakerNotebookInstanceLifecycleConfig resource with PhysicalResourceId", *resource.NotebookInstanceLifecycleConfigName)
			r.resources = append(r.resources, *resource.NotebookInstanceLifecycleConfigName)
		}
		return true
	})
	return
}

func getSageMakerWorkteam(client *sagemaker.SageMaker) (r resourceSliceError) {
	logDebug("Listing SageMakerWorkteam resources")
	r.err = client.ListWorkteamsPages(&sagemaker.ListWorkteamsInput{}, func(page *sagemaker.ListWorkteamsOutput, lastPage bool) bool {
		for _, resource := range page.Workteams {
			logDebug("Got SageMakerWorkteam resource with PhysicalResourceId", *resource.WorkteamName)
			r.resources = append(r.resources, *resource.WorkteamName)
		}
		return true
	})
	return
}

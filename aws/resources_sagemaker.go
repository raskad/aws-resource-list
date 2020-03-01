package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemaker"
)

func getSageMaker(session *session.Session) (resources resourceMap) {
	client := sagemaker.New(session)
	resources = reduce(
		getSageMakerCodeRepository(client).unwrap(sageMakerCodeRepository),
		getSageMakerEndpoint(client).unwrap(sageMakerEndpoint),
		getSageMakerEndpointConfig(client).unwrap(sageMakerEndpointConfig),
		getSageMakerModel(client).unwrap(sageMakerModel),
		getSageMakerNotebookInstance(client).unwrap(sageMakerNotebookInstance),
		getSageMakerNotebookInstanceLifecycleConfig(client).unwrap(sageMakerNotebookInstanceLifecycleConfig),
		getSageMakerWorkteam(client).unwrap(sageMakerWorkteam),
	)
	return
}

func getSageMakerCodeRepository(client *sagemaker.SageMaker) (r resourceSliceError) {
	r.err = client.ListCodeRepositoriesPages(&sagemaker.ListCodeRepositoriesInput{}, func(page *sagemaker.ListCodeRepositoriesOutput, lastPage bool) bool {
		for _, resource := range page.CodeRepositorySummaryList {
			r.resources = append(r.resources, *resource.CodeRepositoryName)
		}
		return true
	})
	return
}

func getSageMakerEndpoint(client *sagemaker.SageMaker) (r resourceSliceError) {
	r.err = client.ListEndpointsPages(&sagemaker.ListEndpointsInput{}, func(page *sagemaker.ListEndpointsOutput, lastPage bool) bool {
		for _, resource := range page.Endpoints {
			r.resources = append(r.resources, *resource.EndpointName)
		}
		return true
	})
	return
}

func getSageMakerEndpointConfig(client *sagemaker.SageMaker) (r resourceSliceError) {
	r.err = client.ListEndpointConfigsPages(&sagemaker.ListEndpointConfigsInput{}, func(page *sagemaker.ListEndpointConfigsOutput, lastPage bool) bool {
		for _, resource := range page.EndpointConfigs {
			r.resources = append(r.resources, *resource.EndpointConfigName)
		}
		return true
	})
	return
}

func getSageMakerModel(client *sagemaker.SageMaker) (r resourceSliceError) {
	r.err = client.ListModelsPages(&sagemaker.ListModelsInput{}, func(page *sagemaker.ListModelsOutput, lastPage bool) bool {
		for _, resource := range page.Models {
			r.resources = append(r.resources, *resource.ModelName)
		}
		return true
	})
	return
}

func getSageMakerNotebookInstance(client *sagemaker.SageMaker) (r resourceSliceError) {
	r.err = client.ListNotebookInstancesPages(&sagemaker.ListNotebookInstancesInput{}, func(page *sagemaker.ListNotebookInstancesOutput, lastPage bool) bool {
		for _, resource := range page.NotebookInstances {
			r.resources = append(r.resources, *resource.NotebookInstanceName)
		}
		return true
	})
	return
}

func getSageMakerNotebookInstanceLifecycleConfig(client *sagemaker.SageMaker) (r resourceSliceError) {
	r.err = client.ListNotebookInstanceLifecycleConfigsPages(&sagemaker.ListNotebookInstanceLifecycleConfigsInput{}, func(page *sagemaker.ListNotebookInstanceLifecycleConfigsOutput, lastPage bool) bool {
		for _, resource := range page.NotebookInstanceLifecycleConfigs {
			r.resources = append(r.resources, *resource.NotebookInstanceLifecycleConfigName)
		}
		return true
	})
	return
}

func getSageMakerWorkteam(client *sagemaker.SageMaker) (r resourceSliceError) {
	r.err = client.ListWorkteamsPages(&sagemaker.ListWorkteamsInput{}, func(page *sagemaker.ListWorkteamsOutput, lastPage bool) bool {
		for _, resource := range page.Workteams {
			r.resources = append(r.resources, *resource.WorkteamName)
		}
		return true
	})
	return
}

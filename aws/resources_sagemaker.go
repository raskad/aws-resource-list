package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func getSageMaker(config aws.Config) (resources resourceMap) {
	client := sagemaker.New(config)
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

func getSageMakerCodeRepository(client *sagemaker.Client) (r resourceSliceError) {
	req := client.ListCodeRepositoriesRequest(&sagemaker.ListCodeRepositoriesInput{})
	p := sagemaker.NewListCodeRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.CodeRepositorySummaryList {
			r.resources = append(r.resources, *resource.CodeRepositoryName)
		}
	}
	r.err = p.Err()
	return
}

func getSageMakerEndpoint(client *sagemaker.Client) (r resourceSliceError) {
	req := client.ListEndpointsRequest(&sagemaker.ListEndpointsInput{})
	p := sagemaker.NewListEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Endpoints {
			r.resources = append(r.resources, *resource.EndpointName)
		}
	}
	r.err = p.Err()
	return
}

func getSageMakerEndpointConfig(client *sagemaker.Client) (r resourceSliceError) {
	req := client.ListEndpointConfigsRequest(&sagemaker.ListEndpointConfigsInput{})
	p := sagemaker.NewListEndpointConfigsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.EndpointConfigs {
			r.resources = append(r.resources, *resource.EndpointConfigName)
		}
	}
	r.err = p.Err()
	return
}

func getSageMakerModel(client *sagemaker.Client) (r resourceSliceError) {
	req := client.ListModelsRequest(&sagemaker.ListModelsInput{})
	p := sagemaker.NewListModelsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Models {
			r.resources = append(r.resources, *resource.ModelName)
		}
	}
	r.err = p.Err()
	return
}

func getSageMakerNotebookInstance(client *sagemaker.Client) (r resourceSliceError) {
	req := client.ListNotebookInstancesRequest(&sagemaker.ListNotebookInstancesInput{})
	p := sagemaker.NewListNotebookInstancesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.NotebookInstances {
			r.resources = append(r.resources, *resource.NotebookInstanceName)
		}
	}
	r.err = p.Err()
	return
}

func getSageMakerNotebookInstanceLifecycleConfig(client *sagemaker.Client) (r resourceSliceError) {
	req := client.ListNotebookInstanceLifecycleConfigsRequest(&sagemaker.ListNotebookInstanceLifecycleConfigsInput{})
	p := sagemaker.NewListNotebookInstanceLifecycleConfigsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.NotebookInstanceLifecycleConfigs {
			r.resources = append(r.resources, *resource.NotebookInstanceLifecycleConfigName)
		}
	}
	r.err = p.Err()
	return
}

func getSageMakerWorkteam(client *sagemaker.Client) (r resourceSliceError) {
	req := client.ListWorkteamsRequest(&sagemaker.ListWorkteamsInput{})
	p := sagemaker.NewListWorkteamsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Workteams {
			r.resources = append(r.resources, *resource.WorkteamName)
		}
	}
	r.err = p.Err()
	return
}

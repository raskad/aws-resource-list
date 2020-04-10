package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func getSageMaker(config aws.Config) (resources awsResourceMap) {
	client := sagemaker.New(config)

	sageMakerCodeRepositoryNames := getSageMakerCodeRepositoryNames(client)
	sageMakerEndpointNames := getSageMakerEndpointNames(client)
	sageMakerEndpointConfigNames := getSageMakerEndpointConfigNames(client)
	sageMakerModelNames := getSageMakerModelNames(client)
	sageMakerNotebookInstanceNames := getSageMakerNotebookInstanceNames(client)
	sageMakerNotebookInstanceLifecycleConfigNames := getSageMakerNotebookInstanceLifecycleConfigNames(client)
	sageMakerWorkteamNames := getSageMakerWorkteamNames(client)

	resources = awsResourceMap{
		sageMakerCodeRepository:                  sageMakerCodeRepositoryNames,
		sageMakerEndpoint:                        sageMakerEndpointNames,
		sageMakerEndpointConfig:                  sageMakerEndpointConfigNames,
		sageMakerModel:                           sageMakerModelNames,
		sageMakerNotebookInstance:                sageMakerNotebookInstanceNames,
		sageMakerNotebookInstanceLifecycleConfig: sageMakerNotebookInstanceLifecycleConfigNames,
		sageMakerWorkteam:                        sageMakerWorkteamNames,
	}
	return
}

func getSageMakerCodeRepositoryNames(client *sagemaker.Client) (resources []string) {
	req := client.ListCodeRepositoriesRequest(&sagemaker.ListCodeRepositoriesInput{})
	p := sagemaker.NewListCodeRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.CodeRepositorySummaryList {
			resources = append(resources, *resource.CodeRepositoryName)
		}
	}
	return
}

func getSageMakerEndpointNames(client *sagemaker.Client) (resources []string) {
	req := client.ListEndpointsRequest(&sagemaker.ListEndpointsInput{})
	p := sagemaker.NewListEndpointsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Endpoints {
			resources = append(resources, *resource.EndpointName)
		}
	}
	return
}

func getSageMakerEndpointConfigNames(client *sagemaker.Client) (resources []string) {
	req := client.ListEndpointConfigsRequest(&sagemaker.ListEndpointConfigsInput{})
	p := sagemaker.NewListEndpointConfigsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.EndpointConfigs {
			resources = append(resources, *resource.EndpointConfigName)
		}
	}
	return
}

func getSageMakerModelNames(client *sagemaker.Client) (resources []string) {
	req := client.ListModelsRequest(&sagemaker.ListModelsInput{})
	p := sagemaker.NewListModelsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Models {
			resources = append(resources, *resource.ModelName)
		}
	}
	return
}

func getSageMakerNotebookInstanceNames(client *sagemaker.Client) (resources []string) {
	req := client.ListNotebookInstancesRequest(&sagemaker.ListNotebookInstancesInput{})
	p := sagemaker.NewListNotebookInstancesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.NotebookInstances {
			resources = append(resources, *resource.NotebookInstanceName)
		}
	}
	return
}

func getSageMakerNotebookInstanceLifecycleConfigNames(client *sagemaker.Client) (resources []string) {
	req := client.ListNotebookInstanceLifecycleConfigsRequest(&sagemaker.ListNotebookInstanceLifecycleConfigsInput{})
	p := sagemaker.NewListNotebookInstanceLifecycleConfigsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.NotebookInstanceLifecycleConfigs {
			resources = append(resources, *resource.NotebookInstanceLifecycleConfigName)
		}
	}
	return
}

func getSageMakerWorkteamNames(client *sagemaker.Client) (resources []string) {
	req := client.ListWorkteamsRequest(&sagemaker.ListWorkteamsInput{})
	p := sagemaker.NewListWorkteamsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.Workteams {
			resources = append(resources, *resource.WorkteamName)
		}
	}
	return
}

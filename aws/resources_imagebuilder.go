package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
)

func getImageBuilder(config aws.Config) (resources awsResourceMap) {
	client := imagebuilder.New(config)

	imageBuilderComponentNames := getImageBuilderComponentNames(client)
	imageBuilderDistributionConfigurationNames := getImageBuilderDistributionConfigurationNames(client)
	imageBuilderImageARNs := getImageBuilderImageARNs(client)
	imageBuilderImagePipelineNames := getImageBuilderImagePipelineNames(client)
	imageBuilderImageRecipeNames := getImageBuilderImageRecipeNames(client)
	imageBuilderInfrastructureConfigurationNames := getImageBuilderInfrastructureConfigurationNames(client)

	resources = awsResourceMap{
		imageBuilderComponent:                   imageBuilderComponentNames,
		imageBuilderDistributionConfiguration:   imageBuilderDistributionConfigurationNames,
		imageBuilderImage:                       imageBuilderImageARNs,
		imageBuilderImagePipeline:               imageBuilderImagePipelineNames,
		imageBuilderImageRecipe:                 imageBuilderImageRecipeNames,
		imageBuilderInfrastructureConfiguration: imageBuilderInfrastructureConfigurationNames,
	}
	return
}

func getImageBuilderComponentNames(client *imagebuilder.Client) (resources []string) {
	req := client.ListComponentsRequest(&imagebuilder.ListComponentsInput{})
	p := imagebuilder.NewListComponentsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ComponentVersionList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getImageBuilderDistributionConfigurationNames(client *imagebuilder.Client) (resources []string) {
	req := client.ListDistributionConfigurationsRequest(&imagebuilder.ListDistributionConfigurationsInput{})
	p := imagebuilder.NewListDistributionConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.DistributionConfigurationSummaryList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getImageBuilderImageARNs(client *imagebuilder.Client) (resources []string) {
	req := client.ListImagesRequest(&imagebuilder.ListImagesInput{})
	p := imagebuilder.NewListImagesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ImageVersionList {
			resources = append(resources, *resource.Arn)
		}
	}
	return
}

func getImageBuilderInfrastructureConfigurationNames(client *imagebuilder.Client) (resources []string) {
	req := client.ListInfrastructureConfigurationsRequest(&imagebuilder.ListInfrastructureConfigurationsInput{})
	p := imagebuilder.NewListInfrastructureConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.InfrastructureConfigurationSummaryList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getImageBuilderImageRecipeNames(client *imagebuilder.Client) (resources []string) {
	req := client.ListImageRecipesRequest(&imagebuilder.ListImageRecipesInput{})
	p := imagebuilder.NewListImageRecipesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ImageRecipeSummaryList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getImageBuilderImagePipelineNames(client *imagebuilder.Client) (resources []string) {
	req := client.ListImagePipelinesRequest(&imagebuilder.ListImagePipelinesInput{})
	p := imagebuilder.NewListImagePipelinesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ImagePipelineList {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

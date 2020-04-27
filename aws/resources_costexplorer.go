package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

func getCostExplorer(config aws.Config) (resources awsResourceMap) {
	client := costexplorer.New(config)

	costExplorerCostCategoryNames := getCostExplorerCostCategoryNames(client)

	resources = awsResourceMap{
		costExplorerCostCategory: costExplorerCostCategoryNames,
	}
	return
}

func getCostExplorerCostCategoryNames(client *costexplorer.Client) (resources []string) {
	req := client.ListCostCategoryDefinitionsRequest(&costexplorer.ListCostCategoryDefinitionsInput{})
	p := costexplorer.NewListCostCategoryDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.CostCategoryReferences {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

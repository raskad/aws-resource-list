package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
)

func getEmr(config aws.Config) (resources resourceMap) {
	client := emr.New(config)

	emrClusterNames := getEmrClusterNames(client)
	emrSecurityConfigurationNames := getEmrSecurityConfigurationNames(client)

	resources = resourceMap{
		emrCluster:               emrClusterNames,
		emrSecurityConfiguration: emrSecurityConfigurationNames,
	}
	return
}

func getEmrClusterNames(client *emr.Client) (resources []string) {
	req := client.ListClustersRequest(&emr.ListClustersInput{})
	p := emr.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.Clusters {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getEmrSecurityConfigurationNames(client *emr.Client) (resources []string) {
	req := client.ListSecurityConfigurationsRequest(&emr.ListSecurityConfigurationsInput{})
	p := emr.NewListSecurityConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.SecurityConfigurations {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

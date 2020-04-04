package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
)

func getEmr(config aws.Config) (resources awsResourceMap) {
	client := emr.New(config)

	emrClusterNames := getEmrClusterNames(client)
	emrSecurityConfigurationNames := getEmrSecurityConfigurationNames(client)
	emrInstanceGroupNames := getEmrInstanceGroupNames(client)

	resources = awsResourceMap{
		emrCluster:               emrClusterNames,
		emrSecurityConfiguration: emrSecurityConfigurationNames,
		emrInstanceGroup:         emrInstanceGroupNames,
	}
	return
}

func getEmrClusterNames(client *emr.Client) (resources []string) {
	req := client.ListClustersRequest(&emr.ListClustersInput{})
	p := emr.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
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
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.SecurityConfigurations {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

func getEmrInstanceGroupNames(client *emr.Client) (resources []string) {
	req := client.ListInstanceGroupsRequest(&emr.ListInstanceGroupsInput{})
	p := emr.NewListInstanceGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.InstanceGroups {
			resources = append(resources, *resource.Name)
		}
	}
	return
}

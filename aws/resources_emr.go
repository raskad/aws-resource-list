package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
)

func getEmr(config aws.Config) (resources resourceMap) {
	client := emr.New(config)
	resources = reduce(
		getEmrCluster(client).unwrap(emrCluster),
		getEmrSecurityConfiguration(client).unwrap(emrSecurityConfiguration),
	)
	return
}

func getEmrCluster(client *emr.Client) (r resourceSliceError) {
	req := client.ListClustersRequest(&emr.ListClustersInput{})
	p := emr.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.Clusters {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

func getEmrSecurityConfiguration(client *emr.Client) (r resourceSliceError) {
	req := client.ListSecurityConfigurationsRequest(&emr.ListSecurityConfigurationsInput{})
	p := emr.NewListSecurityConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.SecurityConfigurations {
			r.resources = append(r.resources, *resource.Name)
		}
	}
	r.err = p.Err()
	return
}

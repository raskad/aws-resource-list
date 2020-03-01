package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/emr"
)

func getEmr(session *session.Session) (resources resourceMap) {
	client := emr.New(session)
	resources = reduce(
		getEmrCluster(client).unwrap(emrCluster),
		getEmrSecurityConfiguration(client).unwrap(emrSecurityConfiguration),
	)
	return
}

func getEmrCluster(client *emr.EMR) (r resourceSliceError) {
	r.err = client.ListClustersPages(&emr.ListClustersInput{}, func(page *emr.ListClustersOutput, lastPage bool) bool {
		for _, resource := range page.Clusters {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getEmrSecurityConfiguration(client *emr.EMR) (r resourceSliceError) {
	r.err = client.ListSecurityConfigurationsPages(&emr.ListSecurityConfigurationsInput{}, func(page *emr.ListSecurityConfigurationsOutput, lastPage bool) bool {
		for _, resource := range page.SecurityConfigurations {
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

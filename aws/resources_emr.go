package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/emr"
)

func getEmr(session *session.Session) (resources resourceMap) {
	client := emr.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		emrCluster:               getEmrCluster(client),
		emrSecurityConfiguration: getEmrSecurityConfiguration(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getEmrCluster(client *emr.EMR) (r resourceSliceError) {
	logDebug("Listing EmrCluster resources")
	r.err = client.ListClustersPages(&emr.ListClustersInput{}, func(page *emr.ListClustersOutput, lastPage bool) bool {
		for _, resource := range page.Clusters {
			logDebug("Got EmrCluster resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

func getEmrSecurityConfiguration(client *emr.EMR) (r resourceSliceError) {
	logDebug("Listing EmrSecurityConfiguration resources")
	r.err = client.ListSecurityConfigurationsPages(&emr.ListSecurityConfigurationsInput{}, func(page *emr.ListSecurityConfigurationsOutput, lastPage bool) bool {
		for _, resource := range page.SecurityConfigurations {
			logDebug("Got EmrSecurityConfiguration resource with PhysicalResourceId", *resource.Name)
			r.resources = append(r.resources, *resource.Name)
		}
		return true
	})
	return
}

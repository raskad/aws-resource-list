package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func getEcs(session *session.Session) (resources resourceMap) {
	client := ecs.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		ecsCluster:        getEcsCluster(client),
		ecsService:        getEcsService(client),
		ecsTaskDefinition: getEcsTaskDefinition(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getEcsCluster(client *ecs.ECS) (r resourceSliceError) {
	r.err = client.ListClustersPages(&ecs.ListClustersInput{}, func(page *ecs.ListClustersOutput, lastPage bool) bool {
		logDebug("Listing EcsCluster resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ClusterArns {
			logDebug("Got EcsCluster resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getEcsService(client *ecs.ECS) (r resourceSliceError) {
	r.err = client.ListServicesPages(&ecs.ListServicesInput{}, func(page *ecs.ListServicesOutput, lastPage bool) bool {
		logDebug("Listing EcsService resources page. Remaining pages", page.NextToken)
		for _, resource := range page.ServiceArns {
			logDebug("Got EcsService resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getEcsTaskDefinition(client *ecs.ECS) (r resourceSliceError) {
	r.err = client.ListTaskDefinitionsPages(&ecs.ListTaskDefinitionsInput{}, func(page *ecs.ListTaskDefinitionsOutput, lastPage bool) bool {
		logDebug("Listing EcsTaskDefinition resources page. Remaining pages", page.NextToken)
		for _, resource := range page.TaskDefinitionArns {
			logDebug("Got EcsTaskDefinition resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

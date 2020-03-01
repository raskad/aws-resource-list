package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func getEcs(session *session.Session) (resources resourceMap) {
	client := ecs.New(session)
	resources = reduce(
		getEcsCluster(client).unwrap(ecsCluster),
		getEcsService(client).unwrap(ecsService),
		getEcsTaskDefinition(client).unwrap(ecsTaskDefinition),
	)
	return
}

func getEcsCluster(client *ecs.ECS) (r resourceSliceError) {
	r.err = client.ListClustersPages(&ecs.ListClustersInput{}, func(page *ecs.ListClustersOutput, lastPage bool) bool {
		for _, resource := range page.ClusterArns {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getEcsService(client *ecs.ECS) (r resourceSliceError) {
	r.err = client.ListServicesPages(&ecs.ListServicesInput{}, func(page *ecs.ListServicesOutput, lastPage bool) bool {
		for _, resource := range page.ServiceArns {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

func getEcsTaskDefinition(client *ecs.ECS) (r resourceSliceError) {
	r.err = client.ListTaskDefinitionsPages(&ecs.ListTaskDefinitionsInput{}, func(page *ecs.ListTaskDefinitionsOutput, lastPage bool) bool {
		for _, resource := range page.TaskDefinitionArns {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

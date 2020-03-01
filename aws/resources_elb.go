package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

func getElb(session *session.Session) (resources resourceMap) {
	client := elb.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		elasticLoadBalancingLoadBalancer: getElasticLoadBalancingLoadBalancer(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getElasticLoadBalancingLoadBalancer(client *elb.ELB) (r resourceSliceError) {
	logDebug("Listing ElasticLoadBalancingLoadBalancer resources")
	r.err = client.DescribeLoadBalancersPages(&elb.DescribeLoadBalancersInput{}, func(page *elb.DescribeLoadBalancersOutput, lastPage bool) bool {
		for _, resource := range page.LoadBalancerDescriptions {
			logDebug("Got ElasticLoadBalancingLoadBalancer resource with PhysicalResourceId", *resource.LoadBalancerName)
			r.resources = append(r.resources, *resource.LoadBalancerName)
		}
		return true
	})
	return
}

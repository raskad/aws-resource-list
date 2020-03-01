package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

func getElb(session *session.Session) (resources resourceMap) {
	client := elb.New(session)
	resources = reduce(
		getElasticLoadBalancingLoadBalancer(client).unwrap(elasticLoadBalancingLoadBalancer),
	)
	return
}

func getElasticLoadBalancingLoadBalancer(client *elb.ELB) (r resourceSliceError) {
	r.err = client.DescribeLoadBalancersPages(&elb.DescribeLoadBalancersInput{}, func(page *elb.DescribeLoadBalancersOutput, lastPage bool) bool {
		for _, resource := range page.LoadBalancerDescriptions {
			r.resources = append(r.resources, *resource.LoadBalancerName)
		}
		return true
	})
	return
}

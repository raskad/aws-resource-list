package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

func getElasticLoadBalancing(config aws.Config) (resources resourceMap) {
	client := elasticloadbalancing.New(config)
	resources = reduce(
		getElasticLoadBalancingLoadBalancer(client).unwrap(elasticLoadBalancingLoadBalancer),
	)
	return
}

func getElasticLoadBalancingLoadBalancer(client *elasticloadbalancing.Client) (r resourceSliceError) {
	req := client.DescribeLoadBalancersRequest(&elasticloadbalancing.DescribeLoadBalancersInput{})
	p := elasticloadbalancing.NewDescribeLoadBalancersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.LoadBalancerDescriptions {
			r.resources = append(r.resources, *resource.LoadBalancerName)
		}
	}
	r.err = p.Err()
	return
}

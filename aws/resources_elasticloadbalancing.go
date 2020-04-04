package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

func getElasticLoadBalancing(config aws.Config) (resources awsResourceMap) {
	client := elasticloadbalancing.New(config)

	elasticLoadBalancingLoadBalancerNames := getElasticLoadBalancingLoadBalancerNames(client)

	resources = awsResourceMap{
		elasticLoadBalancingLoadBalancer: elasticLoadBalancingLoadBalancerNames,
	}
	return
}

func getElasticLoadBalancingLoadBalancerNames(client *elasticloadbalancing.Client) (resources []string) {
	req := client.DescribeLoadBalancersRequest(&elasticloadbalancing.DescribeLoadBalancersInput{})
	p := elasticloadbalancing.NewDescribeLoadBalancersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.LoadBalancerDescriptions {
			resources = append(resources, *resource.LoadBalancerName)
		}
	}
	return
}

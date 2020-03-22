package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func getElasticLoadBalancingV2(config aws.Config) (resources resourceMap) {
	client := elasticloadbalancingv2.New(config)

	LoadBalancerResourceMap := getElasticLoadBalancingV2LoadBalancer(client).unwrap(elasticLoadBalancingV2LoadBalancer)
	LoadBalancerArns := LoadBalancerResourceMap[elasticLoadBalancingV2LoadBalancer]

	ListenerResourceMap := getElasticLoadBalancingV2Listener(client, LoadBalancerArns).unwrap(elasticLoadBalancingV2Listener)
	ListenerArns := ListenerResourceMap[elasticLoadBalancingV2Listener]

	resources = reduce(
		getElasticLoadBalancingV2Listener(client, LoadBalancerArns).unwrap(elasticLoadBalancingV2Listener),
		getElasticLoadBalancingV2ListenerRule(client, ListenerArns).unwrap(elasticLoadBalancingV2ListenerRule),
		LoadBalancerResourceMap,
		getElasticLoadBalancingV2TargetGroup(client).unwrap(elasticLoadBalancingV2TargetGroup),
	)
	return
}

func getElasticLoadBalancingV2Listener(client *elasticloadbalancingv2.Client, loadBalancerArns []string) (r resourceSliceError) {
	for _, loadBalancerArn := range loadBalancerArns {
		req := client.DescribeListenersRequest(&elasticloadbalancingv2.DescribeListenersInput{
			LoadBalancerArn: aws.String(loadBalancerArn),
		})
		p := elasticloadbalancingv2.NewDescribeListenersPaginator(req)
		for p.Next(context.Background()) {
			page := p.CurrentPage()
			for _, resource := range page.Listeners {
				r.resources = append(r.resources, *resource.ListenerArn)
			}
		}
		r.err = p.Err()
	}
	return
}

func getElasticLoadBalancingV2ListenerRule(client *elasticloadbalancingv2.Client, ListenerArns []string) (r resourceSliceError) {
	for _, ListenerArn := range ListenerArns {
		input := elasticloadbalancingv2.DescribeRulesInput{
			ListenerArn: aws.String(ListenerArn),
		}
		for {
			page, err := client.DescribeRulesRequest(&input).Send(context.Background())
			if err != nil {
				r.err = err
				return
			}
			for _, resource := range page.Rules {
				r.resources = append(r.resources, *resource.RuleArn)
			}
			if page.NextMarker == nil {
				return
			}
			input.Marker = page.NextMarker
		}
	}
	return
}

func getElasticLoadBalancingV2LoadBalancer(client *elasticloadbalancingv2.Client) (r resourceSliceError) {
	req := client.DescribeLoadBalancersRequest(&elasticloadbalancingv2.DescribeLoadBalancersInput{})
	p := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.LoadBalancers {
			r.resources = append(r.resources, *resource.LoadBalancerArn)
		}
	}
	r.err = p.Err()
	return
}

func getElasticLoadBalancingV2TargetGroup(client *elasticloadbalancingv2.Client) (r resourceSliceError) {
	req := client.DescribeTargetGroupsRequest(&elasticloadbalancingv2.DescribeTargetGroupsInput{})
	p := elasticloadbalancingv2.NewDescribeTargetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TargetGroups {
			r.resources = append(r.resources, *resource.TargetGroupArn)
		}
	}
	r.err = p.Err()
	return
}

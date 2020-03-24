package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func getElasticLoadBalancingV2(config aws.Config) (resources resourceMap) {
	client := elasticloadbalancingv2.New(config)

	elasticLoadBalancingV2LoadBalancerARNs := getElasticLoadBalancingV2LoadBalancerARNs(client)
	elasticLoadBalancingV2ListenerARNs := getElasticLoadBalancingV2ListenerARNs(client, elasticLoadBalancingV2LoadBalancerARNs)
	elasticLoadBalancingV2ListenerRuleARNs := getElasticLoadBalancingV2ListenerRuleARNs(client, elasticLoadBalancingV2ListenerARNs)
	elasticLoadBalancingV2TargetGroupARNs := getElasticLoadBalancingV2TargetGroupARNs(client)

	resources = resourceMap{
		elasticLoadBalancingV2LoadBalancer: elasticLoadBalancingV2LoadBalancerARNs,
		elasticLoadBalancingV2Listener:     elasticLoadBalancingV2ListenerARNs,
		elasticLoadBalancingV2ListenerRule: elasticLoadBalancingV2ListenerRuleARNs,
		elasticLoadBalancingV2TargetGroup:  elasticLoadBalancingV2TargetGroupARNs,
	}
	return
}

func getElasticLoadBalancingV2ListenerARNs(client *elasticloadbalancingv2.Client, loadBalancerArns []string) (resources []string) {
	for _, loadBalancerArn := range loadBalancerArns {
		req := client.DescribeListenersRequest(&elasticloadbalancingv2.DescribeListenersInput{
			LoadBalancerArn: aws.String(loadBalancerArn),
		})
		p := elasticloadbalancingv2.NewDescribeListenersPaginator(req)
		for p.Next(context.Background()) {
			logErr(p.Err())
			page := p.CurrentPage()
			for _, resource := range page.Listeners {
				resources = append(resources, *resource.ListenerArn)
			}
		}
	}
	return
}

func getElasticLoadBalancingV2ListenerRuleARNs(client *elasticloadbalancingv2.Client, ListenerArns []string) (resources []string) {
	for _, ListenerArn := range ListenerArns {
		input := elasticloadbalancingv2.DescribeRulesInput{
			ListenerArn: aws.String(ListenerArn),
		}
		for {
			page, err := client.DescribeRulesRequest(&input).Send(context.Background())
			logErr(err)
			for _, resource := range page.Rules {
				resources = append(resources, *resource.RuleArn)
			}
			if page.NextMarker == nil {
				return
			}
			input.Marker = page.NextMarker
		}
	}
	return
}

func getElasticLoadBalancingV2LoadBalancerARNs(client *elasticloadbalancingv2.Client) (resources []string) {
	req := client.DescribeLoadBalancersRequest(&elasticloadbalancingv2.DescribeLoadBalancersInput{})
	p := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.LoadBalancers {
			resources = append(resources, *resource.LoadBalancerArn)
		}
	}
	return
}

func getElasticLoadBalancingV2TargetGroupARNs(client *elasticloadbalancingv2.Client) (resources []string) {
	req := client.DescribeTargetGroupsRequest(&elasticloadbalancingv2.DescribeTargetGroupsInput{})
	p := elasticloadbalancingv2.NewDescribeTargetGroupsPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.TargetGroups {
			resources = append(resources, *resource.TargetGroupArn)
		}
	}
	return
}

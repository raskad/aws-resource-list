package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

func getElasticLoadBalancingV2(session *session.Session) (resources resourceMap) {
	client := elbv2.New(session)

	LoadBalancerResourceMap := getElasticLoadBalancingV2LoadBalancer(client).unwrap(elasticLoadBalancingV2LoadBalancer)
	LoadBalancerArns := LoadBalancerResourceMap[elasticLoadBalancingV2LoadBalancer]

	ListenerResourceMap := getElasticLoadBalancingV2Listener(client, LoadBalancerArns).unwrap(elasticLoadBalancingV2Listener)
	ListenerArns := ListenerResourceMap[elasticLoadBalancingV2Listener]

	resources = reduce(
		getElasticLoadBalancingV2Listener(client, LoadBalancerArns).unwrap(elasticLoadBalancingV2Listener),
		getElasticLoadBalancingV2ListenerCertificate(client, ListenerArns).unwrap(elasticLoadBalancingV2ListenerCertificate),
		getElasticLoadBalancingV2ListenerRule(client, ListenerArns).unwrap(elasticLoadBalancingV2ListenerRule),
		LoadBalancerResourceMap,
		getElasticLoadBalancingV2TargetGroup(client).unwrap(elasticLoadBalancingV2TargetGroup),
	)
	return
}

func getElasticLoadBalancingV2Listener(client *elbv2.ELBV2, LoadBalancerArns []string) (r resourceSliceError) {
	for _, LoadBalancerArn := range LoadBalancerArns {
		r.err = client.DescribeListenersPages(&elbv2.DescribeListenersInput{
			LoadBalancerArn: aws.String(LoadBalancerArn),
		}, func(page *elbv2.DescribeListenersOutput, lastPage bool) bool {
			for _, resource := range page.Listeners {
				r.resources = append(r.resources, *resource.ListenerArn)
			}
			return true
		})
	}
	return
}

func getElasticLoadBalancingV2ListenerCertificate(client *elbv2.ELBV2, ListenerArns []string) (r resourceSliceError) {
	for _, ListenerArn := range ListenerArns {
		input := elbv2.DescribeListenerCertificatesInput{
			ListenerArn: aws.String(ListenerArn),
		}
		for {
			page, err := client.DescribeListenerCertificates(&input)
			if err != nil {
				r.err = err
				return
			}
			for _, resource := range page.Certificates {
				r.resources = append(r.resources, *resource.CertificateArn)
			}
			if page.NextMarker == nil {
				return
			}
			input.Marker = page.NextMarker
		}
	}
	return
}

func getElasticLoadBalancingV2ListenerRule(client *elbv2.ELBV2, ListenerArns []string) (r resourceSliceError) {
	for _, ListenerArn := range ListenerArns {
		input := elbv2.DescribeRulesInput{
			ListenerArn: aws.String(ListenerArn),
		}
		for {
			page, err := client.DescribeRules(&input)
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

func getElasticLoadBalancingV2LoadBalancer(client *elbv2.ELBV2) (r resourceSliceError) {
	r.err = client.DescribeLoadBalancersPages(&elbv2.DescribeLoadBalancersInput{}, func(page *elbv2.DescribeLoadBalancersOutput, lastPage bool) bool {
		for _, resource := range page.LoadBalancers {
			r.resources = append(r.resources, *resource.LoadBalancerArn)
		}
		return true
	})
	return
}

func getElasticLoadBalancingV2TargetGroup(client *elbv2.ELBV2) (r resourceSliceError) {
	r.err = client.DescribeTargetGroupsPages(&elbv2.DescribeTargetGroupsInput{}, func(page *elbv2.DescribeTargetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.TargetGroups {
			r.resources = append(r.resources, *resource.TargetGroupName)
		}
		return true
	})
	return
}

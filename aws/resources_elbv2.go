package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

func getElasticLoadBalancingV2(session *session.Session) (resources resourceMap) {
	client := elbv2.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		elasticLoadBalancingV2Listener:            getElasticLoadBalancingV2Listener(client),
		elasticLoadBalancingV2ListenerCertificate: getElasticLoadBalancingV2ListenerCertificate(client),
		elasticLoadBalancingV2ListenerRule:        getElasticLoadBalancingV2ListenerRule(client),
		elasticLoadBalancingV2LoadBalancer:        getElasticLoadBalancingV2LoadBalancer(client),
		elasticLoadBalancingV2TargetGroup:         getElasticLoadBalancingV2TargetGroup(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getElasticLoadBalancingV2Listener(client *elbv2.ELBV2) (r resourceSliceError) {
	logDebug("Listing ElasticLoadBalancingV2Listener resources")
	r.err = client.DescribeListenersPages(&elbv2.DescribeListenersInput{}, func(page *elbv2.DescribeListenersOutput, lastPage bool) bool {
		for _, resource := range page.Listeners {
			logDebug("Got ElasticLoadBalancingV2Listener resource with PhysicalResourceId", *resource.ListenerArn)
			r.resources = append(r.resources, *resource.ListenerArn)
		}
		return true
	})
	return
}

func getElasticLoadBalancingV2ListenerCertificate(client *elbv2.ELBV2) (r resourceSliceError) {
	logDebug("Listing ElasticLoadBalancingV2ListenerCertificate resources")
	input := elbv2.DescribeListenerCertificatesInput{}
	for {
		page, err := client.DescribeListenerCertificates(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Certificates {
			logDebug("Got ElasticLoadBalancingV2ListenerCertificate resource with PhysicalResourceId", *resource.CertificateArn)
			r.resources = append(r.resources, *resource.CertificateArn)
		}
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

func getElasticLoadBalancingV2ListenerRule(client *elbv2.ELBV2) (r resourceSliceError) {
	logDebug("Listing ElasticLoadBalancingV2ListenerRule resources")
	input := elbv2.DescribeRulesInput{}
	for {
		page, err := client.DescribeRules(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Rules {
			logDebug("Got ElasticLoadBalancingV2ListenerRule resource with PhysicalResourceId", *resource.RuleArn)
			r.resources = append(r.resources, *resource.RuleArn)
		}
		if page.NextMarker == nil {
			return
		}
		input.Marker = page.NextMarker
	}
}

func getElasticLoadBalancingV2LoadBalancer(client *elbv2.ELBV2) (r resourceSliceError) {
	logDebug("Listing ElasticLoadBalancingV2LoadBalancer resources")
	r.err = client.DescribeLoadBalancersPages(&elbv2.DescribeLoadBalancersInput{}, func(page *elbv2.DescribeLoadBalancersOutput, lastPage bool) bool {
		for _, resource := range page.LoadBalancers {
			logDebug("Got ElasticLoadBalancingV2LoadBalancer resource with PhysicalResourceId", *resource.LoadBalancerName)
			r.resources = append(r.resources, *resource.LoadBalancerName)
		}
		return true
	})
	return
}

func getElasticLoadBalancingV2TargetGroup(client *elbv2.ELBV2) (r resourceSliceError) {
	logDebug("Listing ElasticLoadBalancingV2TargetGroup resources")
	r.err = client.DescribeTargetGroupsPages(&elbv2.DescribeTargetGroupsInput{}, func(page *elbv2.DescribeTargetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.TargetGroups {
			logDebug("Got ElasticLoadBalancingV2TargetGroup resource with PhysicalResourceId", *resource.TargetGroupName)
			r.resources = append(r.resources, *resource.TargetGroupName)
		}
		return true
	})
	return
}

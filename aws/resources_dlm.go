package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
)

func getDLM(config aws.Config) (resources resourceMap) {
	client := dlm.New(config)

	dlmLifecyclePolicyIDs := getDlmLifecyclePolicyIDs(client)

	resources = resourceMap{
		dlmLifecyclePolicy: dlmLifecyclePolicyIDs,
	}
	return
}

func getDlmLifecyclePolicyIDs(client *dlm.Client) (resources []string) {
	buckets, err := client.GetLifecyclePoliciesRequest(&dlm.GetLifecyclePoliciesInput{}).Send(context.Background())
	logErr(err)
	for _, resource := range buckets.Policies {
		resources = append(resources, *resource.PolicyId)
	}
	return
}

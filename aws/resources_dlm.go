package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
)

func getDLM(config aws.Config) (resources awsResourceMap) {
	client := dlm.New(config)

	dlmLifecyclePolicyIDs := getDlmLifecyclePolicyIDs(client)

	resources = awsResourceMap{
		dlmLifecyclePolicy: dlmLifecyclePolicyIDs,
	}
	return
}

func getDlmLifecyclePolicyIDs(client *dlm.Client) (resources []string) {
	buckets, err := client.GetLifecyclePoliciesRequest(&dlm.GetLifecyclePoliciesInput{}).Send(context.Background())
	if err != nil {
		logErr(err)
		return
	}
	for _, resource := range buckets.Policies {
		resources = append(resources, *resource.PolicyId)
	}
	return
}

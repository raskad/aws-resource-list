package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
)

func getDLM(config aws.Config) (resources resourceMap) {
	client := dlm.New(config)
	resources = reduce(
		getDlmLifecyclePolicy(client).unwrap(dlmLifecyclePolicy),
	)
	return
}

func getDlmLifecyclePolicy(client *dlm.Client) (r resourceSliceError) {
	buckets, err := client.GetLifecyclePoliciesRequest(&dlm.GetLifecyclePoliciesInput{}).Send(context.Background())
	for _, resource := range buckets.Policies {
		r.resources = append(r.resources, *resource.PolicyId)
	}
	r.err = err
	return
}

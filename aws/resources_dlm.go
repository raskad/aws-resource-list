package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dlm"
)

func getDLM(session *session.Session) (resources resourceMap) {
	client := dlm.New(session)
	resources = reduce(
		getDlmLifecyclePolicy(client).unwrap(dlmLifecyclePolicy),
	)
	return
}

func getDlmLifecyclePolicy(client *dlm.DLM) (r resourceSliceError) {
	buckets, err := client.GetLifecyclePolicies(&dlm.GetLifecyclePoliciesInput{})
	for _, resource := range buckets.Policies {
		r.resources = append(r.resources, *resource.PolicyId)
	}
	r.err = err
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dlm"
)

func getDLM(session *session.Session) (resources resourceMap) {
	client := dlm.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		dlmLifecyclePolicy: getDlmLifecyclePolicy(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getDlmLifecyclePolicy(client *dlm.DLM) (r resourceSliceError) {
	logInfo("Start fetching DlmLifecyclePolicy resources")
	buckets, err := client.GetLifecyclePolicies(&dlm.GetLifecyclePoliciesInput{})
	for _, resource := range buckets.Policies {
		logDebug("Got DlmLifecyclePolicy resource with PhysicalResourceId", *resource.PolicyId)
		r.resources = append(r.resources, *resource.PolicyId)
	}
	r.err = err
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
)

func getCodeGuruProfiler(config aws.Config) (resources awsResourceMap) {
	client := codeguruprofiler.New(config)

	codeGuruProfilerProfilingGroupNames := getCodeGuruProfilerProfilingGroupNames(client)

	resources = awsResourceMap{
		codeGuruProfilerProfilingGroup: codeGuruProfilerProfilingGroupNames,
	}
	return
}

func getCodeGuruProfilerProfilingGroupNames(client *codeguruprofiler.Client) (resources []string) {
	req := client.ListProfilingGroupsRequest(&codeguruprofiler.ListProfilingGroupsInput{})
	p := codeguruprofiler.NewListProfilingGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.ProfilingGroupNames...)
	}
	return
}

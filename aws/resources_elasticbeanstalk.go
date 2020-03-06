package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func getElasticBeanstalk(config aws.Config) (resources resourceMap) {
	client := elasticbeanstalk.New(config)
	resources = reduce(
		getElasticBeanstalkApplication(client).unwrap(elasticBeanstalkApplication),
		getElasticBeanstalkApplicationVersion(client).unwrap(elasticBeanstalkApplicationVersion),
		getElasticBeanstalkConfigurationTemplate(client).unwrap(elasticBeanstalkConfigurationTemplate),
		getElasticBeanstalkEnvironment(client).unwrap(elasticBeanstalkEnvironment),
	)
	return
}

func getElasticBeanstalkApplication(client *elasticbeanstalk.Client) (r resourceSliceError) {
	input := elasticbeanstalk.DescribeApplicationsInput{}
	page, err := client.DescribeApplicationsRequest(&input).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.Applications {
		r.resources = append(r.resources, *resource.ApplicationName)
	}
	return
}

func getElasticBeanstalkApplicationVersion(client *elasticbeanstalk.Client) (r resourceSliceError) {
	input := elasticbeanstalk.DescribeApplicationVersionsInput{}
	for {
		page, err := client.DescribeApplicationVersionsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ApplicationVersions {
			r.resources = append(r.resources, *resource.VersionLabel)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getElasticBeanstalkConfigurationTemplate(client *elasticbeanstalk.Client) (r resourceSliceError) {
	input := elasticbeanstalk.DescribeApplicationsInput{}
	page, err := client.DescribeApplicationsRequest(&input).Send(context.Background())
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.Applications {
		for _, resource := range resource.ConfigurationTemplates {
			r.resources = append(r.resources, resource)
		}
	}
	return
}

func getElasticBeanstalkEnvironment(client *elasticbeanstalk.Client) (r resourceSliceError) {
	input := elasticbeanstalk.DescribeEnvironmentsInput{}
	for {
		page, err := client.DescribeEnvironmentsRequest(&input).Send(context.Background())
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Environments {
			r.resources = append(r.resources, *resource.EnvironmentId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func getElasticBeanstalk(config aws.Config) (resources resourceMap) {
	client := elasticbeanstalk.New(config)

	elasticBeanstalkApplicationNames := getElasticBeanstalkApplicationNames(client)
	elasticBeanstalkApplicationVersionARNs := getElasticBeanstalkApplicationVersionARNs(client)
	elasticBeanstalkConfigurationTemplateNames := getElasticBeanstalkConfigurationTemplateNames(client)
	elasticBeanstalkEnvironmentIDs := getElasticBeanstalkEnvironmentIDs(client)

	resources = resourceMap{
		elasticBeanstalkApplication:           elasticBeanstalkApplicationNames,
		elasticBeanstalkApplicationVersion:    elasticBeanstalkApplicationVersionARNs,
		elasticBeanstalkConfigurationTemplate: elasticBeanstalkConfigurationTemplateNames,
		elasticBeanstalkEnvironment:           elasticBeanstalkEnvironmentIDs,
	}
	return
}

func getElasticBeanstalkApplicationNames(client *elasticbeanstalk.Client) (resources []string) {
	input := elasticbeanstalk.DescribeApplicationsInput{}
	page, err := client.DescribeApplicationsRequest(&input).Send(context.Background())
	logErr(err)
	for _, resource := range page.Applications {
		resources = append(resources, *resource.ApplicationName)
	}
	return
}

func getElasticBeanstalkApplicationVersionARNs(client *elasticbeanstalk.Client) (resources []string) {
	input := elasticbeanstalk.DescribeApplicationVersionsInput{}
	for {
		page, err := client.DescribeApplicationVersionsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.ApplicationVersions {
			resources = append(resources, *resource.ApplicationVersionArn)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getElasticBeanstalkConfigurationTemplateNames(client *elasticbeanstalk.Client) (resources []string) {
	input := elasticbeanstalk.DescribeApplicationsInput{}
	page, err := client.DescribeApplicationsRequest(&input).Send(context.Background())
	logErr(err)
	for _, resource := range page.Applications {
		resources = append(resources, resource.ConfigurationTemplates...)
	}
	return
}

func getElasticBeanstalkEnvironmentIDs(client *elasticbeanstalk.Client) (resources []string) {
	input := elasticbeanstalk.DescribeEnvironmentsInput{}
	for {
		page, err := client.DescribeEnvironmentsRequest(&input).Send(context.Background())
		logErr(err)
		for _, resource := range page.Environments {
			resources = append(resources, *resource.EnvironmentId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

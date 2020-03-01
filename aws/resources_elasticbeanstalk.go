package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

func getElasticBeanstalk(session *session.Session) (resources resourceMap) {
	client := elasticbeanstalk.New(session)
	resources = reduce(
		getElasticBeanstalkApplication(client).unwrap(elasticBeanstalkApplication),
		getElasticBeanstalkApplicationVersion(client).unwrap(elasticBeanstalkApplicationVersion),
		getElasticBeanstalkConfigurationTemplate(client).unwrap(elasticBeanstalkConfigurationTemplate),
		getElasticBeanstalkEnvironment(client).unwrap(elasticBeanstalkEnvironment),
	)
	return
}

func getElasticBeanstalkApplication(client *elasticbeanstalk.ElasticBeanstalk) (r resourceSliceError) {
	input := elasticbeanstalk.DescribeApplicationsInput{}
	page, err := client.DescribeApplications(&input)
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.Applications {
		r.resources = append(r.resources, *resource.ApplicationName)
	}
	return
}

func getElasticBeanstalkApplicationVersion(client *elasticbeanstalk.ElasticBeanstalk) (r resourceSliceError) {
	input := elasticbeanstalk.DescribeApplicationVersionsInput{}
	for {
		page, err := client.DescribeApplicationVersions(&input)
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

func getElasticBeanstalkConfigurationTemplate(client *elasticbeanstalk.ElasticBeanstalk) (r resourceSliceError) {
	input := elasticbeanstalk.DescribeApplicationsInput{}
	page, err := client.DescribeApplications(&input)
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.Applications {
		for _, resource := range resource.ConfigurationTemplates {
			r.resources = append(r.resources, *resource)
		}
	}
	return
}

func getElasticBeanstalkEnvironment(client *elasticbeanstalk.ElasticBeanstalk) (r resourceSliceError) {
	input := elasticbeanstalk.DescribeEnvironmentsInput{}
	for {
		page, err := client.DescribeEnvironments(&input)
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

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

func getElasticBeanstalk(session *session.Session) (resources resourceMap) {
	client := elasticbeanstalk.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		elasticBeanstalkApplication:           getElasticBeanstalkApplication(client),
		elasticBeanstalkApplicationVersion:    getElasticBeanstalkApplicationVersion(client),
		elasticBeanstalkConfigurationTemplate: getElasticBeanstalkConfigurationTemplate(client),
		elasticBeanstalkEnvironment:           getElasticBeanstalkEnvironment(client),
	}

	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getElasticBeanstalkApplication(client *elasticbeanstalk.ElasticBeanstalk) (r resourceSliceError) {
	logDebug("Listing ElasticBeanstalkApplication resources")
	input := elasticbeanstalk.DescribeApplicationsInput{}
	page, err := client.DescribeApplications(&input)
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.Applications {
		logDebug("Got ElasticBeanstalkApplication resource with PhysicalResourceId", *resource.ApplicationName)
		r.resources = append(r.resources, *resource.ApplicationName)
	}
	return
}

func getElasticBeanstalkApplicationVersion(client *elasticbeanstalk.ElasticBeanstalk) (r resourceSliceError) {
	logDebug("Listing ElasticBeanstalkApplicationVersion resources")
	input := elasticbeanstalk.DescribeApplicationVersionsInput{}
	for {
		page, err := client.DescribeApplicationVersions(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.ApplicationVersions {
			logDebug("Got ElasticBeanstalkApplicationVersion resource with PhysicalResourceId", *resource.VersionLabel)
			r.resources = append(r.resources, *resource.VersionLabel)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

func getElasticBeanstalkConfigurationTemplate(client *elasticbeanstalk.ElasticBeanstalk) (r resourceSliceError) {
	logDebug("Listing ElasticBeanstalkConfigurationTemplate resources")
	input := elasticbeanstalk.DescribeApplicationsInput{}
	page, err := client.DescribeApplications(&input)
	if err != nil {
		r.err = err
		return
	}
	for _, resource := range page.Applications {
		for _, resource := range resource.ConfigurationTemplates {
			logDebug("Got ElasticBeanstalkConfigurationTemplate resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
	}
	return
}

func getElasticBeanstalkEnvironment(client *elasticbeanstalk.ElasticBeanstalk) (r resourceSliceError) {
	logDebug("Listing ElasticBeanstalkEnvironment resources")
	input := elasticbeanstalk.DescribeEnvironmentsInput{}
	for {
		page, err := client.DescribeEnvironments(&input)
		if err != nil {
			r.err = err
			return
		}
		for _, resource := range page.Environments {
			logDebug("Got ElasticBeanstalkEnvironment resource with PhysicalResourceId", *resource.EnvironmentId)
			r.resources = append(r.resources, *resource.EnvironmentId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

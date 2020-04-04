package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
)

func getLicenseManager(config aws.Config) (resources awsResourceMap) {
	client := licensemanager.New(config)

	licenseManagerLicenseConfigurationIDs := getLicenseManagerLicenseConfigurationIDs(client)

	resources = awsResourceMap{
		licenseManagerLicenseConfiguration: licenseManagerLicenseConfigurationIDs,
	}
	return
}

func getLicenseManagerLicenseConfigurationIDs(client *licensemanager.Client) (resources []string) {
	input := licensemanager.ListLicenseConfigurationsInput{}
	for {
		page, err := client.ListLicenseConfigurationsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.LicenseConfigurations {
			resources = append(resources, *resource.LicenseConfigurationId)
		}
		if page.NextToken == nil {
			return
		}
		input.NextToken = page.NextToken
	}
}

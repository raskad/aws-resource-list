package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func getLightsail(config aws.Config) (resources awsResourceMap) {
	client := lightsail.New(config)

	lightsailDomainNames := getLightsailDomainNames(client)
	lightsailInstanceNames := getLightsailInstanceNames(client)
	lightsailKeyPairNames := getLightsailKeyPairNames(client)
	lightsailStaticIPNames := getLightsailStaticIPNames(client)

	resources = awsResourceMap{
		lightsailDomain:   lightsailDomainNames,
		lightsailInstance: lightsailInstanceNames,
		lightsailKeyPair:  lightsailKeyPairNames,
		lightsailStaticIP: lightsailStaticIPNames,
	}
	return
}

func getLightsailDomainNames(client *lightsail.Client) (resources []string) {
	input := lightsail.GetDomainsInput{}
	for {
		page, err := client.GetDomainsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Domains {
			resources = append(resources, *resource.Name)
		}
		if page.NextPageToken == nil {
			return
		}
		input.PageToken = page.NextPageToken
	}
}

func getLightsailInstanceNames(client *lightsail.Client) (resources []string) {
	input := lightsail.GetInstancesInput{}
	for {
		page, err := client.GetInstancesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.Instances {
			resources = append(resources, *resource.Name)
		}
		if page.NextPageToken == nil {
			return
		}
		input.PageToken = page.NextPageToken
	}
}

func getLightsailKeyPairNames(client *lightsail.Client) (resources []string) {
	input := lightsail.GetKeyPairsInput{}
	for {
		page, err := client.GetKeyPairsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.KeyPairs {
			resources = append(resources, *resource.Name)
		}
		if page.NextPageToken == nil {
			return
		}
		input.PageToken = page.NextPageToken
	}
}

func getLightsailStaticIPNames(client *lightsail.Client) (resources []string) {
	input := lightsail.GetStaticIpsInput{}
	for {
		page, err := client.GetStaticIpsRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.StaticIps {
			resources = append(resources, *resource.Name)
		}
		if page.NextPageToken == nil {
			return
		}
		input.PageToken = page.NextPageToken
	}
}

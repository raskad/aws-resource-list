package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
)

func getNetworkManager(config aws.Config) (resources awsResourceMap) {
	client := networkmanager.New(config)

	networkManagerGlobalNetworkIDs := getNetworkManagerGlobalNetworkIDs(client)
	networkManagerDeviceIDs := getNetworkManagerDeviceIDs(client, networkManagerGlobalNetworkIDs)
	networkManagerLinkIDs := getNetworkManagerLinkIDs(client, networkManagerGlobalNetworkIDs)
	networkManagerSiteIDs := getNetworkManagerSiteIDs(client, networkManagerGlobalNetworkIDs)

	resources = awsResourceMap{
		networkManagerDevice:        networkManagerDeviceIDs,
		networkManagerGlobalNetwork: networkManagerGlobalNetworkIDs,
		networkManagerLink:          networkManagerLinkIDs,
		networkManagerSite:          networkManagerSiteIDs,
	}
	return
}

func getNetworkManagerDeviceIDs(client *networkmanager.Client, networkManagerGlobalNetworkIDs []string) (resources []string) {
	for _, networkManagerGlobalNetworkID := range networkManagerGlobalNetworkIDs {
		req := client.GetDevicesRequest(&networkmanager.GetDevicesInput{
			GlobalNetworkId: aws.String(networkManagerGlobalNetworkID),
		})
		p := networkmanager.NewGetDevicesPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.Devices {
				resources = append(resources, *resource.DeviceId)
			}
		}
	}
	return
}

func getNetworkManagerGlobalNetworkIDs(client *networkmanager.Client) (resources []string) {
	req := client.DescribeGlobalNetworksRequest(&networkmanager.DescribeGlobalNetworksInput{})
	p := networkmanager.NewDescribeGlobalNetworksPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.GlobalNetworks {
			resources = append(resources, *resource.GlobalNetworkId)
		}
	}
	return
}

func getNetworkManagerLinkIDs(client *networkmanager.Client, networkManagerGlobalNetworkIDs []string) (resources []string) {
	for _, networkManagerGlobalNetworkID := range networkManagerGlobalNetworkIDs {
		req := client.GetLinksRequest(&networkmanager.GetLinksInput{
			GlobalNetworkId: aws.String(networkManagerGlobalNetworkID),
		})
		p := networkmanager.NewGetLinksPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.Links {
				resources = append(resources, *resource.LinkId)
			}
		}
	}
	return
}

func getNetworkManagerSiteIDs(client *networkmanager.Client, networkManagerGlobalNetworkIDs []string) (resources []string) {
	for _, networkManagerGlobalNetworkID := range networkManagerGlobalNetworkIDs {
		req := client.GetSitesRequest(&networkmanager.GetSitesInput{
			GlobalNetworkId: aws.String(networkManagerGlobalNetworkID),
		})
		p := networkmanager.NewGetSitesPaginator(req)
		for p.Next(context.Background()) {
			if p.Err() != nil {
				logErr(p.Err())
				return
			}
			page := p.CurrentPage()
			for _, resource := range page.Sites {
				resources = append(resources, *resource.SiteId)
			}
		}
	}
	return
}

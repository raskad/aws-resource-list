package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

func getElasticache(config aws.Config) (resources awsResourceMap) {
	client := elasticache.New(config)

	elastiCacheCacheClusterIDs := getElastiCacheCacheClusterIDs(client)
	elastiCacheParameterGroupNames := getElastiCacheParameterGroupNames(client)
	elastiCacheReplicationGroupIDs := getElastiCacheReplicationGroupIDs(client)
	elastiCacheSecurityGroupNames := getElastiCacheSecurityGroupNames(client)
	elastiCacheSubnetGroupNames := getElastiCacheSubnetGroupNames(client)

	resources = awsResourceMap{
		elastiCacheCacheCluster:     elastiCacheCacheClusterIDs,
		elastiCacheParameterGroup:   elastiCacheParameterGroupNames,
		elastiCacheReplicationGroup: elastiCacheReplicationGroupIDs,
		elastiCacheSecurityGroup:    elastiCacheSecurityGroupNames,
		elastiCacheSubnetGroup:      elastiCacheSubnetGroupNames,
	}
	return
}

func getElastiCacheCacheClusterIDs(client *elasticache.Client) (resources []string) {
	req := client.DescribeCacheClustersRequest(&elasticache.DescribeCacheClustersInput{})
	p := elasticache.NewDescribeCacheClustersPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.CacheClusters {
			resources = append(resources, *resource.CacheClusterId)
		}
	}
	return
}

func getElastiCacheParameterGroupNames(client *elasticache.Client) (resources []string) {
	req := client.DescribeCacheParameterGroupsRequest(&elasticache.DescribeCacheParameterGroupsInput{})
	p := elasticache.NewDescribeCacheParameterGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.CacheParameterGroups {
			resources = append(resources, *resource.CacheParameterGroupName)
		}
	}
	return
}

func getElastiCacheReplicationGroupIDs(client *elasticache.Client) (resources []string) {
	req := client.DescribeReplicationGroupsRequest(&elasticache.DescribeReplicationGroupsInput{})
	p := elasticache.NewDescribeReplicationGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.ReplicationGroups {
			resources = append(resources, *resource.ReplicationGroupId)
		}
	}
	return
}

func getElastiCacheSecurityGroupNames(client *elasticache.Client) (resources []string) {
	req := client.DescribeCacheSecurityGroupsRequest(&elasticache.DescribeCacheSecurityGroupsInput{})
	p := elasticache.NewDescribeCacheSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.CacheSecurityGroups {
			resources = append(resources, *resource.CacheSecurityGroupName)
		}
	}
	return
}

func getElastiCacheSubnetGroupNames(client *elasticache.Client) (resources []string) {
	req := client.DescribeCacheSubnetGroupsRequest(&elasticache.DescribeCacheSubnetGroupsInput{})
	p := elasticache.NewDescribeCacheSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		for _, resource := range page.CacheSubnetGroups {
			resources = append(resources, *resource.CacheSubnetGroupName)
		}
	}
	return
}

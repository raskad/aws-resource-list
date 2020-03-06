package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

func getElasticache(config aws.Config) (resources resourceMap) {
	client := elasticache.New(config)
	resources = reduce(
		getElastiCacheCacheCluster(client).unwrap(elastiCacheCacheCluster),
		getElastiCacheParameterGroup(client).unwrap(elastiCacheParameterGroup),
		getElastiCacheReplicationGroup(client).unwrap(elastiCacheReplicationGroup),
		getElastiCacheSecurityGroup(client).unwrap(elastiCacheSecurityGroup),
		getElastiCacheSubnetGroup(client).unwrap(elastiCacheSubnetGroup),
	)
	return
}

func getElastiCacheCacheCluster(client *elasticache.Client) (r resourceSliceError) {
	req := client.DescribeCacheClustersRequest(&elasticache.DescribeCacheClustersInput{})
	p := elasticache.NewDescribeCacheClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.CacheClusters {
			r.resources = append(r.resources, *resource.CacheClusterId)
		}
	}
	r.err = p.Err()
	return
}

func getElastiCacheParameterGroup(client *elasticache.Client) (r resourceSliceError) {
	req := client.DescribeCacheParameterGroupsRequest(&elasticache.DescribeCacheParameterGroupsInput{})
	p := elasticache.NewDescribeCacheParameterGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.CacheParameterGroups {
			r.resources = append(r.resources, *resource.CacheParameterGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getElastiCacheReplicationGroup(client *elasticache.Client) (r resourceSliceError) {
	req := client.DescribeReplicationGroupsRequest(&elasticache.DescribeReplicationGroupsInput{})
	p := elasticache.NewDescribeReplicationGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.ReplicationGroups {
			r.resources = append(r.resources, *resource.ReplicationGroupId)
		}
	}
	r.err = p.Err()
	return
}

func getElastiCacheSecurityGroup(client *elasticache.Client) (r resourceSliceError) {
	req := client.DescribeCacheSecurityGroupsRequest(&elasticache.DescribeCacheSecurityGroupsInput{})
	p := elasticache.NewDescribeCacheSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.CacheSecurityGroups {
			r.resources = append(r.resources, *resource.CacheSecurityGroupName)
		}
	}
	r.err = p.Err()
	return
}

func getElastiCacheSubnetGroup(client *elasticache.Client) (r resourceSliceError) {
	req := client.DescribeCacheSubnetGroupsRequest(&elasticache.DescribeCacheSubnetGroupsInput{})
	p := elasticache.NewDescribeCacheSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.CacheSubnetGroups {
			r.resources = append(r.resources, *resource.CacheSubnetGroupName)
		}
	}
	r.err = p.Err()
	return
}

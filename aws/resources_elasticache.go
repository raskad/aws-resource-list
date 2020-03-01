package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

func getElasticache(session *session.Session) (resources resourceMap) {
	client := elasticache.New(session)
	resources = reduce(
		getElastiCacheCacheCluster(client).unwrap(elastiCacheCacheCluster),
		getElastiCacheParameterGroup(client).unwrap(elastiCacheParameterGroup),
		getElastiCacheReplicationGroup(client).unwrap(elastiCacheReplicationGroup),
		getElastiCacheSecurityGroup(client).unwrap(elastiCacheSecurityGroup),
		getElastiCacheSubnetGroup(client).unwrap(elastiCacheSubnetGroup),
	)
	return
}

func getElastiCacheCacheCluster(client *elasticache.ElastiCache) (r resourceSliceError) {
	r.err = client.DescribeCacheClustersPages(&elasticache.DescribeCacheClustersInput{}, func(page *elasticache.DescribeCacheClustersOutput, lastPage bool) bool {
		for _, resource := range page.CacheClusters {
			r.resources = append(r.resources, *resource.CacheClusterId)
		}
		return true
	})
	return
}

func getElastiCacheParameterGroup(client *elasticache.ElastiCache) (r resourceSliceError) {
	r.err = client.DescribeCacheParameterGroupsPages(&elasticache.DescribeCacheParameterGroupsInput{}, func(page *elasticache.DescribeCacheParameterGroupsOutput, lastPage bool) bool {
		for _, resource := range page.CacheParameterGroups {
			r.resources = append(r.resources, *resource.CacheParameterGroupName)
		}
		return true
	})
	return
}

func getElastiCacheReplicationGroup(client *elasticache.ElastiCache) (r resourceSliceError) {
	r.err = client.DescribeReplicationGroupsPages(&elasticache.DescribeReplicationGroupsInput{}, func(page *elasticache.DescribeReplicationGroupsOutput, lastPage bool) bool {
		for _, resource := range page.ReplicationGroups {
			r.resources = append(r.resources, *resource.ReplicationGroupId)
		}
		return true
	})
	return
}

func getElastiCacheSecurityGroup(client *elasticache.ElastiCache) (r resourceSliceError) {
	r.err = client.DescribeCacheSecurityGroupsPages(&elasticache.DescribeCacheSecurityGroupsInput{}, func(page *elasticache.DescribeCacheSecurityGroupsOutput, lastPage bool) bool {
		for _, resource := range page.CacheSecurityGroups {
			r.resources = append(r.resources, *resource.CacheSecurityGroupName)
		}
		return true
	})
	return
}

func getElastiCacheSubnetGroup(client *elasticache.ElastiCache) (r resourceSliceError) {
	r.err = client.DescribeCacheSubnetGroupsPages(&elasticache.DescribeCacheSubnetGroupsInput{}, func(page *elasticache.DescribeCacheSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.CacheSubnetGroups {
			r.resources = append(r.resources, *resource.CacheSubnetGroupName)
		}
		return true
	})
	return
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

func getElasticache(session *session.Session) (resources resourceMap) {
	client := elasticache.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		elastiCacheCacheCluster:     getElastiCacheCacheCluster(client),
		elastiCacheParameterGroup:   getElastiCacheParameterGroup(client),
		elastiCacheReplicationGroup: getElastiCacheReplicationGroup(client),
		elastiCacheSecurityGroup:    getElastiCacheSecurityGroup(client),
		elastiCacheSubnetGroup:      getElastiCacheSubnetGroup(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getElastiCacheCacheCluster(client *elasticache.ElastiCache) (r resourceSliceError) {
	logDebug("Listing ElastiCacheCacheCluster resources")
	r.err = client.DescribeCacheClustersPages(&elasticache.DescribeCacheClustersInput{}, func(page *elasticache.DescribeCacheClustersOutput, lastPage bool) bool {
		for _, resource := range page.CacheClusters {
			logDebug("Got ElastiCacheCacheCluster resource with PhysicalResourceId", *resource.CacheClusterId)
			r.resources = append(r.resources, *resource.CacheClusterId)
		}
		return true
	})
	return
}

func getElastiCacheParameterGroup(client *elasticache.ElastiCache) (r resourceSliceError) {
	logDebug("Listing ElastiCacheParameterGroup resources")
	r.err = client.DescribeCacheParameterGroupsPages(&elasticache.DescribeCacheParameterGroupsInput{}, func(page *elasticache.DescribeCacheParameterGroupsOutput, lastPage bool) bool {
		for _, resource := range page.CacheParameterGroups {
			logDebug("Got ElastiCacheParameterGroup resource with PhysicalResourceId", *resource.CacheParameterGroupName)
			r.resources = append(r.resources, *resource.CacheParameterGroupName)
		}
		return true
	})
	return
}

func getElastiCacheReplicationGroup(client *elasticache.ElastiCache) (r resourceSliceError) {
	logDebug("Listing ElastiCacheReplicationGroup resources")
	r.err = client.DescribeReplicationGroupsPages(&elasticache.DescribeReplicationGroupsInput{}, func(page *elasticache.DescribeReplicationGroupsOutput, lastPage bool) bool {
		for _, resource := range page.ReplicationGroups {
			logDebug("Got ElastiCacheReplicationGroup resource with PhysicalResourceId", *resource.ReplicationGroupId)
			r.resources = append(r.resources, *resource.ReplicationGroupId)
		}
		return true
	})
	return
}

func getElastiCacheSecurityGroup(client *elasticache.ElastiCache) (r resourceSliceError) {
	logDebug("Listing ElastiCacheSecurityGroup resources")
	r.err = client.DescribeCacheSecurityGroupsPages(&elasticache.DescribeCacheSecurityGroupsInput{}, func(page *elasticache.DescribeCacheSecurityGroupsOutput, lastPage bool) bool {
		for _, resource := range page.CacheSecurityGroups {
			logDebug("Got ElastiCacheSecurityGroup resource with PhysicalResourceId", *resource.CacheSecurityGroupName)
			r.resources = append(r.resources, *resource.CacheSecurityGroupName)
		}
		return true
	})
	return
}

func getElastiCacheSubnetGroup(client *elasticache.ElastiCache) (r resourceSliceError) {
	logDebug("Listing ElastiCacheSubnetGroup resources")
	r.err = client.DescribeCacheSubnetGroupsPages(&elasticache.DescribeCacheSubnetGroupsInput{}, func(page *elasticache.DescribeCacheSubnetGroupsOutput, lastPage bool) bool {
		for _, resource := range page.CacheSubnetGroups {
			logDebug("Got ElastiCacheSubnetGroup resource with PhysicalResourceId", *resource.CacheSubnetGroupName)
			r.resources = append(r.resources, *resource.CacheSubnetGroupName)
		}
		return true
	})
	return
}

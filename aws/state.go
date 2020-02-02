package aws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type state map[resourceSource]resourceMap

func (state state) writeToDisk() (err error) {
	file, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile("aws-resource-list.json", file, 0644)
	return
}

func (state state) readFromDisk() (err error) {
	jsonFile, err := os.Open("aws-resource-list.json")
	if err != nil {
		return
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(byteValue, &state)
	return
}

type resourceSource string

const (
	real resourceSource = "real"
	cfn  resourceSource = "cfn"
	tf   resourceSource = "tf"
)

type resourceMap map[resourceType][]string

func (resourceMap resourceMap) print() {
	for key, valueList := range resourceMap {
		fmt.Println(" -", key)
		for _, value := range valueList {
			fmt.Println("     -", value)
		}
	}
}

func reduce(maps ...resourceMap) (result resourceMap) {
	result = resourceMap{}
	for _, m := range maps {
		for key, value := range m {
			result[key] = value
		}
	}
	return
}

type resourceSliceError struct {
	resources []string
	err       error
}

type resourceSliceErrorMap map[resourceType]resourceSliceError

func (rSliceErrorMap resourceSliceErrorMap) unwrap() (rMap resourceMap) {
	rMap = resourceMap{}
	for rtype, rSliceError := range rSliceErrorMap {
		if rSliceError.err != nil {
			logError("Cloud not get resources of type", rtype)
		}
		rMap[rtype] = rSliceError.resources
	}
	return
}

type resourceType string

const (
	accessAnalyzerAnalyzer                     resourceType = "accessAnalyzerAnalyzer"
	alexaAskSkill                              resourceType = "alexaAskSkill"
	amazonMQBroker                             resourceType = "amazonMQBroker"
	amazonMQConfiguration                      resourceType = "amazonMQConfiguration"
	amazonMQConfigurationAssociation           resourceType = "amazonMQConfigurationAssociation"
	amplifyApp                                 resourceType = "amplifyApp"
	amplifyBranch                              resourceType = "amplifyBranch"
	amplifyDomain                              resourceType = "amplifyDomain"
	apiGatewayAccount                          resourceType = "apiGatewayAccount"
	apiGatewayAPIKey                           resourceType = "apiGatewayAPIKey"
	apiGatewayAuthorizer                       resourceType = "apiGatewayAuthorizer"
	apiGatewayBasePathMapping                  resourceType = "apiGatewayBasePathMapping"
	apiGatewayClientCertificate                resourceType = "apiGatewayClientCertificate"
	apiGatewayDeployment                       resourceType = "apiGatewayDeployment"
	apiGatewayDocumentationPart                resourceType = "apiGatewayDocumentationPart"
	apiGatewayDocumentationVersion             resourceType = "apiGatewayDocumentationVersion"
	apiGatewayDomainName                       resourceType = "apiGatewayDomainName"
	apiGatewayGatewayResponse                  resourceType = "apiGatewayGatewayResponse"
	apiGatewayMethod                           resourceType = "apiGatewayMethod"
	apiGatewayModel                            resourceType = "apiGatewayModel"
	apiGatewayRequestValidator                 resourceType = "apiGatewayRequestValidator"
	apiGatewayResource                         resourceType = "apiGatewayResource"
	apiGatewayRestAPI                          resourceType = "apiGatewayRestAPI"
	apiGatewayStage                            resourceType = "apiGatewayStage"
	apiGatewayUsagePlan                        resourceType = "apiGatewayUsagePlan"
	apiGatewayUsagePlanKey                     resourceType = "apiGatewayUsagePlanKey"
	apiGatewayVpcLink                          resourceType = "apiGatewayVpcLink"
	apiGatewayV2Api                            resourceType = "apiGatewayV2Api"
	apiGatewayV2ApiMapping                     resourceType = "apiGatewayV2ApiMapping"
	apiGatewayV2Authorizer                     resourceType = "apiGatewayV2Authorizer"
	apiGatewayV2Deployment                     resourceType = "apiGatewayV2Deployment"
	apiGatewayV2DomainName                     resourceType = "apiGatewayV2DomainName"
	apiGatewayV2Integration                    resourceType = "apiGatewayV2Integration"
	apiGatewayV2IntegrationResponse            resourceType = "apiGatewayV2IntegrationResponse"
	apiGatewayV2Model                          resourceType = "apiGatewayV2Model"
	apiGatewayV2Route                          resourceType = "apiGatewayV2Route"
	apiGatewayV2RouteResponse                  resourceType = "apiGatewayV2RouteResponse"
	apiGatewayV2Stage                          resourceType = "apiGatewayV2Stage"
	applicationAutoScalingScalableTarget       resourceType = "applicationAutoScalingScalableTarget"
	applicationAutoScalingScalingPolicy        resourceType = "applicationAutoScalingScalingPolicy"
	appMeshMesh                                resourceType = "appMeshMesh"
	appMeshRoute                               resourceType = "appMeshRoute"
	appMeshVirtualNode                         resourceType = "appMeshVirtualNode"
	appMeshVirtualRouter                       resourceType = "appMeshVirtualRouter"
	appMeshVirtualService                      resourceType = "appMeshVirtualService"
	appStreamDirectoryConfig                   resourceType = "appStreamDirectoryConfig"
	appStreamFleet                             resourceType = "appStreamFleet"
	appStreamImageBuilder                      resourceType = "appStreamImageBuilder"
	appStreamStack                             resourceType = "appStreamStack"
	appStreamStackFleetAssociation             resourceType = "appStreamStackFleetAssociation"
	appStreamStackUserAssociation              resourceType = "appStreamStackUserAssociation"
	appStreamUser                              resourceType = "appStreamUser"
	appSyncAPICache                            resourceType = "appSyncAPICache"
	appSyncAPIKey                              resourceType = "appSyncAPIKey"
	appSyncDataSource                          resourceType = "appSyncDataSource"
	appSyncFunctionConfiguration               resourceType = "appSyncFunctionConfiguration"
	appSyncGraphQLApi                          resourceType = "appSyncGraphQLApi"
	appSyncGraphQLSchema                       resourceType = "appSyncGraphQLSchema"
	appSyncResolver                            resourceType = "appSyncResolver"
	athenaNamedQuery                           resourceType = "athenaNamedQuery"
	autoScalingPlansScalingPlan                resourceType = "autoScalingPlansScalingPlan"
	autoScalingAutoScalingGroup                resourceType = "autoScalingAutoScalingGroup"
	autoScalingLaunchConfiguration             resourceType = "autoScalingLaunchConfiguration"
	autoScalingScalingPolicy                   resourceType = "autoScalingScalingPolicy"
	autoScalingScheduledAction                 resourceType = "autoScalingScheduledAction"
	backupBackupPlan                           resourceType = "backupBackupPlan"
	backupBackupSelection                      resourceType = "backupBackupSelection"
	backupBackupVault                          resourceType = "backupBackupVault"
	batchComputeEnvironment                    resourceType = "batchComputeEnvironment"
	batchJobDefinition                         resourceType = "batchJobDefinition"
	batchJobQueue                              resourceType = "batchJobQueue"
	budgetsBudget                              resourceType = "budgetsBudget"
	certificateManagerCertificate              resourceType = "certificateManagerCertificate"
	cloud9EnvironmentEC2                       resourceType = "cloud9EnvironmentEC2"
	cloudFrontCloudFrontOriginAccessIdentity   resourceType = "cloudFrontCloudFrontOriginAccessIdentity"
	cloudFrontDistribution                     resourceType = "cloudFrontDistribution"
	cloudFrontStreamingDistribution            resourceType = "cloudFrontStreamingDistribution"
	serviceDiscoveryHTTPNamespace              resourceType = "serviceDiscoveryHTTPNamespace"
	serviceDiscoveryInstance                   resourceType = "serviceDiscoveryInstance"
	serviceDiscoveryPrivateDNSNamespace        resourceType = "serviceDiscoveryPrivateDNSNamespace"
	serviceDiscoveryPublicDNSNamespace         resourceType = "serviceDiscoveryPublicDNSNamespace"
	serviceDiscoveryService                    resourceType = "serviceDiscoveryService"
	cloudTrailTrail                            resourceType = "cloudTrailTrail"
	cloudWatchAlarm                            resourceType = "cloudWatchAlarm"
	cloudWatchDashboard                        resourceType = "cloudWatchDashboard"
	cloudWatchInsightRule                      resourceType = "cloudWatchInsightRule"
	logsDestination                            resourceType = "logsDestination"
	logsLogGroup                               resourceType = "logsLogGroup"
	logsLogStream                              resourceType = "logsLogStream"
	logsMetricFilter                           resourceType = "logsMetricFilter"
	logsSubscriptionFilter                     resourceType = "logsSubscriptionFilter"
	eventsEventBus                             resourceType = "eventsEventBus"
	eventsRule                                 resourceType = "eventsRule"
	codeBuildProject                           resourceType = "codeBuildProject"
	codeBuildReportGroup                       resourceType = "codeBuildReportGroup"
	codeBuildSourceCredential                  resourceType = "codeBuildSourceCredential"
	codeCommitRepository                       resourceType = "codeCommitRepository"
	codeDeployApplication                      resourceType = "codeDeployApplication"
	codeDeployDeploymentConfig                 resourceType = "codeDeployDeploymentConfig"
	codeDeployDeploymentGroup                  resourceType = "codeDeployDeploymentGroup"
	codePipelineCustomActionType               resourceType = "codePipelineCustomActionType"
	codePipelinePipeline                       resourceType = "codePipelinePipeline"
	codePipelineWebhook                        resourceType = "codePipelineWebhook"
	codeStarGitHubRepository                   resourceType = "codeStarGitHubRepository"
	codeStarNotificationsNotificationRule      resourceType = "codeStarNotificationsNotificationRule"
	cognitoIdentityPool                        resourceType = "cognitoIdentityPool"
	cognitoIdentityPoolRoleAttachment          resourceType = "cognitoIdentityPoolRoleAttachment"
	cognitoUserPool                            resourceType = "cognitoUserPool"
	cognitoUserPoolClient                      resourceType = "cognitoUserPoolClient"
	cognitoUserPoolDomain                      resourceType = "cognitoUserPoolDomain"
	cognitoUserPoolGroup                       resourceType = "cognitoUserPoolGroup"
	cognitoUserPoolIdentityProvider            resourceType = "cognitoUserPoolIdentityProvider"
	cognitoUserPoolResourceServer              resourceType = "cognitoUserPoolResourceServer"
	cognitoUserPoolRiskConfigurationAttachment resourceType = "cognitoUserPoolRiskConfigurationAttachment"
	cognitoUserPoolUICustomizationAttachment   resourceType = "cognitoUserPoolUICustomizationAttachment"
	cognitoUserPoolUser                        resourceType = "cognitoUserPoolUser"
	cognitoUserPoolUserToGroupAttachment       resourceType = "cognitoUserPoolUserToGroupAttachment"
	configAggregationAuthorization             resourceType = "configAggregationAuthorization"
	configConfigRule                           resourceType = "configConfigRule"
	configConfigurationAggregator              resourceType = "configConfigurationAggregator"
	configConfigurationRecorder                resourceType = "configConfigurationRecorder"
	configDeliveryChannel                      resourceType = "configDeliveryChannel"
	configOrganizationConfigRule               resourceType = "configOrganizationConfigRule"
	configRemediationConfiguration             resourceType = "configRemediationConfiguration"
	dataPipelinePipeline                       resourceType = "dataPipelinePipeline"
	daxCluster                                 resourceType = "daxCluster"
	daxParameterGroup                          resourceType = "daxParameterGroup"
	daxSubnetGroup                             resourceType = "daxSubnetGroup"
	directoryServiceMicrosoftAD                resourceType = "directoryServiceMicrosoftAD"
	directoryServiceSimpleAD                   resourceType = "directoryServiceSimpleAD"
	dlmLifecyclePolicy                         resourceType = "dlmLifecyclePolicy"
	dmsCertificate                             resourceType = "dmsCertificate"
	dmsEndpoint                                resourceType = "dmsEndpoint"
	dmsEventSubscription                       resourceType = "dmsEventSubscription"
	dmsReplicationInstance                     resourceType = "dmsReplicationInstance"
	dmsReplicationSubnetGroup                  resourceType = "dmsReplicationSubnetGroup"
	dmsReplicationTask                         resourceType = "dmsReplicationTask"
	docDBDBCluster                             resourceType = "docDBDBCluster"
	docDBDBClusterParameterGroup               resourceType = "docDBDBClusterParameterGroup"
	docDBDBInstance                            resourceType = "docDBDBInstance"
	docDBDBSubnetGroup                         resourceType = "docDBDBSubnetGroup"
	dynamoDBTable                              resourceType = "dynamoDBTable"
	ec2CapacityReservation                     resourceType = "ec2CapacityReservation"
	ec2ClientVpnEndpoint                       resourceType = "ec2ClientVpnEndpoint"
	ec2CustomerGateway                         resourceType = "ec2CustomerGateway"
	ec2DHCPOptions                             resourceType = "ec2DHCPOptions"
	ec2EC2Fleet                                resourceType = "ec2EC2Fleet"
	ec2EgressOnlyInternetGateway               resourceType = "ec2EgressOnlyInternetGateway"
	ec2EIP                                     resourceType = "ec2EIP"
	ec2EIPAssociation                          resourceType = "ec2EIPAssociation"
	ec2FlowLog                                 resourceType = "ec2FlowLog"
	ec2Host                                    resourceType = "ec2Host"
	ec2Instance                                resourceType = "ec2Instance"
	ec2InternetGateway                         resourceType = "ec2InternetGateway"
	ec2LaunchTemplate                          resourceType = "ec2LaunchTemplate"
	ec2NatGateway                              resourceType = "ec2NatGateway"
	ec2NetworkACL                              resourceType = "ec2NetworkACL"
	ec2NetworkInterface                        resourceType = "ec2NetworkInterface"
	ec2NetworkInterfaceAttachment              resourceType = "ec2NetworkInterfaceAttachment"
	ec2NetworkInterfacePermission              resourceType = "ec2NetworkInterfacePermission"
	ec2PlacementGroup                          resourceType = "ec2PlacementGroup"
	ec2RouteTable                              resourceType = "ec2RouteTable"
	ec2SecurityGroup                           resourceType = "ec2SecurityGroup"
	ec2SpotFleet                               resourceType = "ec2SpotFleet"
	ec2Subnet                                  resourceType = "ec2Subnet"
	ec2NetworkACLSubnetAssociation             resourceType = "ec2NetworkACLSubnetAssociation"
	ec2RouteTableSubnetAssociation             resourceType = "ec2RouteTableSubnetAssociation"
	ec2TrafficMirrorFilter                     resourceType = "ec2TrafficMirrorFilter"
	ec2TrafficMirrorFilterRule                 resourceType = "ec2TrafficMirrorFilterRule"
	ec2TrafficMirrorSession                    resourceType = "ec2TrafficMirrorSession"
	ec2TrafficMirrorTarget                     resourceType = "ec2TrafficMirrorTarget"
	ec2TransitGateway                          resourceType = "ec2TransitGateway"
	ec2TransitGatewayAttachment                resourceType = "ec2TransitGatewayAttachment"
	ec2TransitGatewayRouteTable                resourceType = "ec2TransitGatewayRouteTable"
	ec2Volume                                  resourceType = "ec2Volume"
	ec2VPC                                     resourceType = "ec2VPC"
	ec2VPCCidrBlock                            resourceType = "ec2VPCCidrBlock"
	ec2VPCDHCPOptionsAssociation               resourceType = "ec2VPCDHCPOptionsAssociation"
	ec2VPCEndpoint                             resourceType = "ec2VPCEndpoint"
	ec2VPCEndpointConnectionNotification       resourceType = "ec2VPCEndpointConnectionNotification"
	ec2VPCEndpointService                      resourceType = "ec2VPCEndpointService"
	ec2VPCPeeringConnection                    resourceType = "ec2VPCPeeringConnection"
	ec2VPNConnection                           resourceType = "ec2VPNConnection"
	ec2VPNGateway                              resourceType = "ec2VPNGateway"
	ecrRepository                              resourceType = "ecrRepository"
	ecsCluster                                 resourceType = "ecsCluster"
	ecsService                                 resourceType = "ecsService"
	ecsTaskDefinition                          resourceType = "ecsTaskDefinition"
	efsFileSystem                              resourceType = "efsFileSystem"
	efsMountTarget                             resourceType = "efsMountTarget"
	eksCluster                                 resourceType = "eksCluster"
	eksNodegroup                               resourceType = "eksNodegroup"
	elastiCacheCacheCluster                    resourceType = "elastiCacheCacheCluster"
	elastiCacheParameterGroup                  resourceType = "elastiCacheParameterGroup"
	elastiCacheReplicationGroup                resourceType = "elastiCacheReplicationGroup"
	elastiCacheSecurityGroup                   resourceType = "elastiCacheSecurityGroup"
	elastiCacheSubnetGroup                     resourceType = "elastiCacheSubnetGroup"
	elasticsearchDomain                        resourceType = "elasticsearchDomain"
	elasticBeanstalkApplication                resourceType = "elasticBeanstalkApplication"
	elasticBeanstalkApplicationVersion         resourceType = "elasticBeanstalkApplicationVersion"
	elasticBeanstalkConfigurationTemplate      resourceType = "elasticBeanstalkConfigurationTemplate"
	elasticBeanstalkEnvironment                resourceType = "elasticBeanstalkEnvironment"
	elasticLoadBalancingLoadBalancer           resourceType = "elasticLoadBalancingLoadBalancer"
	elasticLoadBalancingV2Listener             resourceType = "elasticLoadBalancingV2Listener"
	elasticLoadBalancingV2ListenerCertificate  resourceType = "elasticLoadBalancingV2ListenerCertificate"
	elasticLoadBalancingV2ListenerRule         resourceType = "elasticLoadBalancingV2ListenerRule"
	elasticLoadBalancingV2LoadBalancer         resourceType = "elasticLoadBalancingV2LoadBalancer"
	elasticLoadBalancingV2TargetGroup          resourceType = "elasticLoadBalancingV2TargetGroup"
	emrCluster                                 resourceType = "emrCluster"
	emrInstanceFleetConfig                     resourceType = "emrInstanceFleetConfig"
	emrInstanceGroupConfig                     resourceType = "emrInstanceGroupConfig"
	emrSecurityConfiguration                   resourceType = "emrSecurityConfiguration"
	emrStep                                    resourceType = "emrStep"
	eventSchemasDiscoverer                     resourceType = "eventSchemasDiscoverer"
	eventSchemasRegistry                       resourceType = "eventSchemasRegistry"
	eventSchemasSchema                         resourceType = "eventSchemasSchema"
	fsxFileSystem                              resourceType = "fsxFileSystem"
	gameLiftAlias                              resourceType = "gameLiftAlias"
	gameLiftBuild                              resourceType = "gameLiftBuild"
	gameLiftFleet                              resourceType = "gameLiftFleet"
	gameLiftGameSessionQueue                   resourceType = "gameLiftGameSessionQueue"
	gameLiftMatchmakingConfiguration           resourceType = "gameLiftMatchmakingConfiguration"
	gameLiftMatchmakingRuleSet                 resourceType = "gameLiftMatchmakingRuleSet"
	gameLiftScript                             resourceType = "gameLiftScript"
	glueConnection                             resourceType = "glueConnection"
	glueCrawler                                resourceType = "glueCrawler"
	glueDatabase                               resourceType = "glueDatabase"
	glueDevEndpoint                            resourceType = "glueDevEndpoint"
	glueJob                                    resourceType = "glueJob"
	glueMLTransform                            resourceType = "glueMLTransform"
	glueSecurityConfiguration                  resourceType = "glueSecurityConfiguration"
	glueTable                                  resourceType = "glueTable"
	glueTrigger                                resourceType = "glueTrigger"
	glueWorkflow                               resourceType = "glueWorkflow"
	guardDutyDetector                          resourceType = "guardDutyDetector"
	iamAccessKey                               resourceType = "iamAccessKey"
	iamGroup                                   resourceType = "iamGroup"
	iamInstanceProfile                         resourceType = "iamInstanceProfile"
	iamPolicy                                  resourceType = "iamPolicy"
	iamRole                                    resourceType = "iamRole"
	iamRolePolicy                              resourceType = "iamRolePolicy"
	iamServiceLinkedRole                       resourceType = "iamServiceLinkedRole"
	iamUser                                    resourceType = "iamUser"
	inspectorAssessmentTarget                  resourceType = "inspectorAssessmentTarget"
	inspectorAssessmentTemplate                resourceType = "inspectorAssessmentTemplate"
	inspectorResourceGroup                     resourceType = "inspectorResourceGroup"
	ioTCertificate                             resourceType = "ioTCertificate"
	ioTPolicy                                  resourceType = "ioTPolicy"
	ioTPolicyPrincipalAttachment               resourceType = "ioTPolicyPrincipalAttachment"
	ioTThing                                   resourceType = "ioTThing"
	ioTThingPrincipalAttachment                resourceType = "ioTThingPrincipalAttachment"
	ioTTopicRule                               resourceType = "ioTTopicRule"
	ioT1ClickDevice                            resourceType = "ioT1ClickDevice"
	ioT1ClickPlacement                         resourceType = "ioT1ClickPlacement"
	ioT1ClickProject                           resourceType = "ioT1ClickProject"
	ioTAnalyticsChannel                        resourceType = "ioTAnalyticsChannel"
	ioTAnalyticsDataset                        resourceType = "ioTAnalyticsDataset"
	ioTAnalyticsDatastore                      resourceType = "ioTAnalyticsDatastore"
	ioTAnalyticsPipeline                       resourceType = "ioTAnalyticsPipeline"
	ioTEventsDetectorModel                     resourceType = "ioTEventsDetectorModel"
	ioTEventsInput                             resourceType = "ioTEventsInput"
	greengrassConnectorDefinition              resourceType = "greengrassConnectorDefinition"
	greengrassConnectorDefinitionVersion       resourceType = "greengrassConnectorDefinitionVersion"
	greengrassCoreDefinition                   resourceType = "greengrassCoreDefinition"
	greengrassCoreDefinitionVersion            resourceType = "greengrassCoreDefinitionVersion"
	greengrassDeviceDefinition                 resourceType = "greengrassDeviceDefinition"
	greengrassDeviceDefinitionVersion          resourceType = "greengrassDeviceDefinitionVersion"
	greengrassFunctionDefinition               resourceType = "greengrassFunctionDefinition"
	greengrassFunctionDefinitionVersion        resourceType = "greengrassFunctionDefinitionVersion"
	greengrassGroup                            resourceType = "greengrassGroup"
	greengrassGroupVersion                     resourceType = "greengrassGroupVersion"
	greengrassLoggerDefinition                 resourceType = "greengrassLoggerDefinition"
	greengrassLoggerDefinitionVersion          resourceType = "greengrassLoggerDefinitionVersion"
	greengrassResourceDefinition               resourceType = "greengrassResourceDefinition"
	greengrassResourceDefinitionVersion        resourceType = "greengrassResourceDefinitionVersion"
	greengrassSubscriptionDefinition           resourceType = "greengrassSubscriptionDefinition"
	greengrassSubscriptionDefinitionVersion    resourceType = "greengrassSubscriptionDefinitionVersion"
	ioTThingsGraphFlowTemplate                 resourceType = "ioTThingsGraphFlowTemplate"
	kinesisStream                              resourceType = "kinesisStream"
	kinesisStreamConsumer                      resourceType = "kinesisStreamConsumer"
	kinesisAnalyticsApplication                resourceType = "kinesisAnalyticsApplication"
	kinesisAnalyticsV2Application              resourceType = "kinesisAnalyticsV2Application"
	kinesisFirehoseDeliveryStream              resourceType = "kinesisFirehoseDeliveryStream"
	kmsAlias                                   resourceType = "kmsAlias"
	kmsKey                                     resourceType = "kmsKey"
	lakeFormationDataLakeSettings              resourceType = "lakeFormationDataLakeSettings"
	lakeFormationPermissions                   resourceType = "lakeFormationPermissions"
	lakeFormationResource                      resourceType = "lakeFormationResource"
	lambdaAlias                                resourceType = "lambdaAlias"
	lambdaFunction                             resourceType = "lambdaFunction"
	lambdaLayerVersion                         resourceType = "lambdaLayerVersion"
	lambdaPermission                           resourceType = "lambdaPermission"
	managedBlockchainMember                    resourceType = "managedBlockchainMember"
	managedBlockchainNode                      resourceType = "managedBlockchainNode"
	mediaConvertJobTemplate                    resourceType = "mediaConvertJobTemplate"
	mediaConvertPreset                         resourceType = "mediaConvertPreset"
	mediaConvertQueue                          resourceType = "mediaConvertQueue"
	mediaLiveChannel                           resourceType = "mediaLiveChannel"
	mediaLiveInput                             resourceType = "mediaLiveInput"
	mediaLiveInputSecurityGroup                resourceType = "mediaLiveInputSecurityGroup"
	mediaStoreContainer                        resourceType = "mediaStoreContainer"
	mskCluster                                 resourceType = "mskCluster"
	neptuneDBCluster                           resourceType = "neptuneDBCluster"
	neptuneDBClusterParameterGroup             resourceType = "neptuneDBClusterParameterGroup"
	neptuneDBInstance                          resourceType = "neptuneDBInstance"
	neptuneDBParameterGroup                    resourceType = "neptuneDBParameterGroup"
	neptuneDBSubnetGroup                       resourceType = "neptuneDBSubnetGroup"
	opsWorksApp                                resourceType = "opsWorksApp"
	opsWorksElasticLoadBalancerAttachment      resourceType = "opsWorksElasticLoadBalancerAttachment"
	opsWorksInstance                           resourceType = "opsWorksInstance"
	opsWorksLayer                              resourceType = "opsWorksLayer"
	opsWorksStack                              resourceType = "opsWorksStack"
	opsWorksUserProfile                        resourceType = "opsWorksUserProfile"
	opsWorksVolume                             resourceType = "opsWorksVolume"
	opsWorksCMServer                           resourceType = "opsWorksCMServer"
	qLDBLedger                                 resourceType = "qLDBLedger"
	ramResourceShare                           resourceType = "ramResourceShare"
	rdsDBCluster                               resourceType = "rdsDBCluster"
	rdsDBClusterParameterGroup                 resourceType = "rdsDBClusterParameterGroup"
	rdsDBInstance                              resourceType = "rdsDBInstance"
	rdsDBParameterGroup                        resourceType = "rdsDBParameterGroup"
	rdsDBSecurityGroup                         resourceType = "rdsDBSecurityGroup"
	rdsDBSubnetGroup                           resourceType = "rdsDBSubnetGroup"
	rdsEventSubscription                       resourceType = "rdsEventSubscription"
	rdsOptionGroup                             resourceType = "rdsOptionGroup"
	redshiftCluster                            resourceType = "redshiftCluster"
	redshiftClusterParameterGroup              resourceType = "redshiftClusterParameterGroup"
	redshiftClusterSecurityGroup               resourceType = "redshiftClusterSecurityGroup"
	redshiftClusterSubnetGroup                 resourceType = "redshiftClusterSubnetGroup"
	roboMakerFleet                             resourceType = "roboMakerFleet"
	roboMakerRobot                             resourceType = "roboMakerRobot"
	roboMakerRobotApplication                  resourceType = "roboMakerRobotApplication"
	roboMakerRobotApplicationVersion           resourceType = "roboMakerRobotApplicationVersion"
	roboMakerSimulationApplication             resourceType = "roboMakerSimulationApplication"
	roboMakerSimulationApplicationVersion      resourceType = "roboMakerSimulationApplicationVersion"
	route53HealthCheck                         resourceType = "route53HealthCheck"
	route53HostedZone                          resourceType = "route53HostedZone"
	route53RecordSet                           resourceType = "route53RecordSet"
	route53ResolverResolverEndpoint            resourceType = "route53ResolverResolverEndpoint"
	route53ResolverResolverRule                resourceType = "route53ResolverResolverRule"
	route53ResolverResolverRuleAssociation     resourceType = "route53ResolverResolverRuleAssociation"
	s3AccessPoint                              resourceType = "s3AccessPoint"
	s3Bucket                                   resourceType = "s3Bucket"
	s3BucketPolicy                             resourceType = "s3BucketPolicy"
	sageMakerCodeRepository                    resourceType = "sageMakerCodeRepository"
	sageMakerEndpoint                          resourceType = "sageMakerEndpoint"
	sageMakerEndpointConfig                    resourceType = "sageMakerEndpointConfig"
	sageMakerModel                             resourceType = "sageMakerModel"
	sageMakerNotebookInstance                  resourceType = "sageMakerNotebookInstance"
	sageMakerNotebookInstanceLifecycleConfig   resourceType = "sageMakerNotebookInstanceLifecycleConfig"
	sageMakerWorkteam                          resourceType = "sageMakerWorkteam"
	secretsManagerSecret                       resourceType = "secretsManagerSecret"
	securityHubHub                             resourceType = "securityHubHub"
	sesConfigurationSet                        resourceType = "sesConfigurationSet"
	sesReceiptFilter                           resourceType = "sesReceiptFilter"
	sesReceiptRuleSet                          resourceType = "sesReceiptRuleSet"
	sesTemplate                                resourceType = "sesTemplate"
	sdbDomain                                  resourceType = "sdbDomain"
	snsSubscription                            resourceType = "snsSubscription"
	snsTopic                                   resourceType = "snsTopic"
	sqsQueue                                   resourceType = "sqsQueue"
	stepFunctionsActivity                      resourceType = "stepFunctionsActivity"
	stepFunctionsStateMachine                  resourceType = "stepFunctionsStateMachine"
	ssmAssociation                             resourceType = "ssmAssociation"
	ssmDocument                                resourceType = "ssmDocument"
	ssmMaintenanceWindow                       resourceType = "ssmMaintenanceWindow"
	ssmMaintenanceWindowTarget                 resourceType = "ssmMaintenanceWindowTarget"
	ssmMaintenanceWindowTask                   resourceType = "ssmMaintenanceWindowTask"
	ssmParameter                               resourceType = "ssmParameter"
	ssmPatchBaseline                           resourceType = "ssmPatchBaseline"
	transferServer                             resourceType = "transferServer"
	transferUser                               resourceType = "transferUser"
	wafByteMatchSet                            resourceType = "wafByteMatchSet"
	wafIPSet                                   resourceType = "wafIPSet"
	wafRule                                    resourceType = "wafRule"
	wafSizeConstraintSet                       resourceType = "wafSizeConstraintSet"
	wafSQLInjectionMatchSet                    resourceType = "wafSQLInjectionMatchSet"
	wafWebACL                                  resourceType = "wafWebACL"
	wafXSSMatchSet                             resourceType = "wafXSSMatchSet"
	wafv2IPSet                                 resourceType = "wafv2IPSet"
	wafv2RegexPatternSet                       resourceType = "wafv2RegexPatternSet"
	wafv2RuleGroup                             resourceType = "wafv2RuleGroup"
	wafv2WebACL                                resourceType = "wafv2WebACL"
	wafRegionalByteMatchSet                    resourceType = "wafRegionalByteMatchSet"
	wafRegionalGeoMatchSet                     resourceType = "wafRegionalGeoMatchSet"
	wafRegionalIPSet                           resourceType = "wafRegionalIPSet"
	wafRegionalRateBasedRule                   resourceType = "wafRegionalRateBasedRule"
	wafRegionalRegexPatternSet                 resourceType = "wafRegionalRegexPatternSet"
	wafRegionalRule                            resourceType = "wafRegionalRule"
	wafRegionalSizeConstraintSet               resourceType = "wafRegionalSizeConstraintSet"
	wafRegionalSQLInjectionMatchSet            resourceType = "wafRegionalSQLInjectionMatchSet"
	wafRegionalWebACL                          resourceType = "wafRegionalWebACL"
	wafRegionalXSSMatchSet                     resourceType = "wafRegionalXSSMatchSet"
	workSpacesWorkspace                        resourceType = "workSpacesWorkspace"
)

func fromCloudFormationType(cloudFormationType string) (resourceType, bool) {
	cfn := map[string]resourceType{
		"AWS::AccessAnalyzer::Analyzer":                     accessAnalyzerAnalyzer,
		"Alexa::ASK::Skill":                                 alexaAskSkill,
		"AWS::AmazonMQ::Broker":                             amazonMQBroker,
		"AWS::AmazonMQ::Configuration":                      amazonMQConfiguration,
		"AWS::AmazonMQ::ConfigurationAssociation":           amazonMQConfigurationAssociation,
		"AWS::Amplify::App":                                 amplifyApp,
		"AWS::Amplify::Branch":                              amplifyBranch,
		"AWS::Amplify::Domain":                              amplifyDomain,
		"AWS::ApiGateway::Account":                          apiGatewayAccount,
		"AWS::ApiGateway::ApiKey":                           apiGatewayAPIKey,
		"AWS::ApiGateway::Authorizer":                       apiGatewayAuthorizer,
		"AWS::ApiGateway::BasePathMapping":                  apiGatewayBasePathMapping,
		"AWS::ApiGateway::ClientCertificate":                apiGatewayClientCertificate,
		"AWS::ApiGateway::Deployment":                       apiGatewayDeployment,
		"AWS::ApiGateway::DocumentationPart":                apiGatewayDocumentationPart,
		"AWS::ApiGateway::DocumentationVersion":             apiGatewayDocumentationVersion,
		"AWS::ApiGateway::DomainName":                       apiGatewayDomainName,
		"AWS::ApiGateway::GatewayResponse":                  apiGatewayGatewayResponse,
		"AWS::ApiGateway::Method":                           apiGatewayMethod,
		"AWS::ApiGateway::Model":                            apiGatewayModel,
		"AWS::ApiGateway::RequestValidator":                 apiGatewayRequestValidator,
		"AWS::ApiGateway::Resource":                         apiGatewayResource,
		"AWS::ApiGateway::RestApi":                          apiGatewayRestAPI,
		"AWS::ApiGateway::Stage":                            apiGatewayStage,
		"AWS::ApiGateway::UsagePlan":                        apiGatewayUsagePlan,
		"AWS::ApiGateway::UsagePlanKey":                     apiGatewayUsagePlanKey,
		"AWS::ApiGateway::VpcLink":                          apiGatewayVpcLink,
		"AWS::ApiGatewayV2::Api":                            apiGatewayV2Api,
		"AWS::ApiGatewayV2::ApiMapping":                     apiGatewayV2ApiMapping,
		"AWS::ApiGatewayV2::Authorizer":                     apiGatewayV2Authorizer,
		"AWS::ApiGatewayV2::Deployment":                     apiGatewayV2Deployment,
		"AWS::ApiGatewayV2::DomainName":                     apiGatewayV2DomainName,
		"AWS::ApiGatewayV2::Integration":                    apiGatewayV2Integration,
		"AWS::ApiGatewayV2::IntegrationResponse":            apiGatewayV2IntegrationResponse,
		"AWS::ApiGatewayV2::Model":                          apiGatewayV2Model,
		"AWS::ApiGatewayV2::Route":                          apiGatewayV2Route,
		"AWS::ApiGatewayV2::RouteResponse":                  apiGatewayV2RouteResponse,
		"AWS::ApiGatewayV2::Stage":                          apiGatewayV2Stage,
		"AWS::ApplicationAutoScaling::ScalableTarget":       applicationAutoScalingScalableTarget,
		"AWS::ApplicationAutoScaling::ScalingPolicy":        applicationAutoScalingScalingPolicy,
		"AWS::AppMesh::Mesh":                                appMeshMesh,
		"AWS::AppMesh::Route":                               appMeshRoute,
		"AWS::AppMesh::VirtualNode":                         appMeshVirtualNode,
		"AWS::AppMesh::VirtualRouter":                       appMeshVirtualRouter,
		"AWS::AppMesh::VirtualService":                      appMeshVirtualService,
		"AWS::AppStream::DirectoryConfig":                   appStreamDirectoryConfig,
		"AWS::AppStream::Fleet":                             appStreamFleet,
		"AWS::AppStream::ImageBuilder":                      appStreamImageBuilder,
		"AWS::AppStream::Stack":                             appStreamStack,
		"AWS::AppStream::StackFleetAssociation":             appStreamStackFleetAssociation,
		"AWS::AppStream::StackUserAssociation":              appStreamStackUserAssociation,
		"AWS::AppStream::User":                              appStreamUser,
		"AWS::AppSync::ApiCache":                            appSyncAPICache,
		"AWS::AppSync::ApiKey":                              appSyncAPIKey,
		"AWS::AppSync::DataSource":                          appSyncDataSource,
		"AWS::AppSync::FunctionConfiguration":               appSyncFunctionConfiguration,
		"AWS::AppSync::GraphQLApi":                          appSyncGraphQLApi,
		"AWS::AppSync::GraphQLSchema":                       appSyncGraphQLSchema,
		"AWS::AppSync::Resolver":                            appSyncResolver,
		"AWS::Athena::NamedQuery":                           athenaNamedQuery,
		"AWS::AutoScalingPlans::ScalingPlan":                autoScalingPlansScalingPlan,
		"AWS::AutoScaling::AutoScalingGroup":                autoScalingAutoScalingGroup,
		"AWS::AutoScaling::LaunchConfiguration":             autoScalingLaunchConfiguration,
		"AWS::AutoScaling::ScalingPolicy":                   autoScalingScalingPolicy,
		"AWS::AutoScaling::ScheduledAction":                 autoScalingScheduledAction,
		"AWS::Backup::BackupPlan":                           backupBackupPlan,
		"AWS::Backup::BackupSelection":                      backupBackupSelection,
		"AWS::Backup::BackupVault":                          backupBackupVault,
		"AWS::Batch::ComputeEnvironment":                    batchComputeEnvironment,
		"AWS::Batch::JobDefinition":                         batchJobDefinition,
		"AWS::Batch::JobQueue":                              batchJobQueue,
		"AWS::Budgets::Budget":                              budgetsBudget,
		"AWS::CertificateManager::Certificate":              certificateManagerCertificate,
		"AWS::Cloud9::EnvironmentEC2":                       cloud9EnvironmentEC2,
		"AWS::CloudFront::CloudFrontOriginAccessIdentity":   cloudFrontCloudFrontOriginAccessIdentity,
		"AWS::CloudFront::Distribution":                     cloudFrontDistribution,
		"AWS::CloudFront::StreamingDistribution":            cloudFrontStreamingDistribution,
		"AWS::ServiceDiscovery::HttpNamespace":              serviceDiscoveryHTTPNamespace,
		"AWS::ServiceDiscovery::Instance":                   serviceDiscoveryInstance,
		"AWS::ServiceDiscovery::PrivateDnsNamespace":        serviceDiscoveryPrivateDNSNamespace,
		"AWS::ServiceDiscovery::PublicDnsNamespace":         serviceDiscoveryPublicDNSNamespace,
		"AWS::ServiceDiscovery::Service":                    serviceDiscoveryService,
		"AWS::CloudTrail::Trail":                            cloudTrailTrail,
		"AWS::CloudWatch::Alarm":                            cloudWatchAlarm,
		"AWS::CloudWatch::Dashboard":                        cloudWatchDashboard,
		"AWS::CloudWatch::InsightRule":                      cloudWatchInsightRule,
		"AWS::Logs::Destination":                            logsDestination,
		"AWS::Logs::LogGroup":                               logsLogGroup,
		"AWS::Logs::LogStream":                              logsLogStream,
		"AWS::Logs::MetricFilter":                           logsMetricFilter,
		"AWS::Logs::SubscriptionFilter":                     logsSubscriptionFilter,
		"AWS::Events::EventBus":                             eventsEventBus,
		"AWS::Events::Rule":                                 eventsRule,
		"AWS::CodeBuild::Project":                           codeBuildProject,
		"AWS::CodeBuild::ReportGroup":                       codeBuildReportGroup,
		"AWS::CodeBuild::SourceCredential":                  codeBuildSourceCredential,
		"AWS::CodeCommit::Repository":                       codeCommitRepository,
		"AWS::CodeDeploy::Application":                      codeDeployApplication,
		"AWS::CodeDeploy::DeploymentConfig":                 codeDeployDeploymentConfig,
		"AWS::CodeDeploy::DeploymentGroup":                  codeDeployDeploymentGroup,
		"AWS::CodePipeline::CustomActionType":               codePipelineCustomActionType,
		"AWS::CodePipeline::Pipeline":                       codePipelinePipeline,
		"AWS::CodePipeline::Webhook":                        codePipelineWebhook,
		"AWS::CodeStar::GitHubRepository":                   codeStarGitHubRepository,
		"AWS::CodeStarNotifications::NotificationRule":      codeStarNotificationsNotificationRule,
		"AWS::Cognito::IdentityPool":                        cognitoIdentityPool,
		"AWS::Cognito::IdentityPoolRoleAttachment":          cognitoIdentityPoolRoleAttachment,
		"AWS::Cognito::UserPool":                            cognitoUserPool,
		"AWS::Cognito::UserPoolClient":                      cognitoUserPoolClient,
		"AWS::Cognito::UserPoolDomain":                      cognitoUserPoolDomain,
		"AWS::Cognito::UserPoolGroup":                       cognitoUserPoolGroup,
		"AWS::Cognito::UserPoolIdentityProvider":            cognitoUserPoolIdentityProvider,
		"AWS::Cognito::UserPoolResourceServer":              cognitoUserPoolResourceServer,
		"AWS::Cognito::UserPoolRiskConfigurationAttachment": cognitoUserPoolRiskConfigurationAttachment,
		"AWS::Cognito::UserPoolUICustomizationAttachment":   cognitoUserPoolUICustomizationAttachment,
		"AWS::Cognito::UserPoolUser":                        cognitoUserPoolUser,
		"AWS::Cognito::UserPoolUserToGroupAttachment":       cognitoUserPoolUserToGroupAttachment,
		"AWS::Config::AggregationAuthorization":             configAggregationAuthorization,
		"AWS::Config::ConfigRule":                           configConfigRule,
		"AWS::Config::ConfigurationAggregator":              configConfigurationAggregator,
		"AWS::Config::ConfigurationRecorder":                configConfigurationRecorder,
		"AWS::Config::DeliveryChannel":                      configDeliveryChannel,
		"AWS::Config::OrganizationConfigRule":               configOrganizationConfigRule,
		"AWS::Config::RemediationConfiguration":             configRemediationConfiguration,
		"AWS::DataPipeline::Pipeline":                       dataPipelinePipeline,
		"AWS::DAX::Cluster":                                 daxCluster,
		"AWS::DAX::ParameterGroup":                          daxParameterGroup,
		"AWS::DAX::SubnetGroup":                             daxSubnetGroup,
		"AWS::DirectoryService::MicrosoftAD":                directoryServiceMicrosoftAD,
		"AWS::DirectoryService::SimpleAD":                   directoryServiceSimpleAD,
		"AWS::DLM::LifecyclePolicy":                         dlmLifecyclePolicy,
		"AWS::DMS::Certificate":                             dmsCertificate,
		"AWS::DMS::Endpoint":                                dmsEndpoint,
		"AWS::DMS::EventSubscription":                       dmsEventSubscription,
		"AWS::DMS::ReplicationInstance":                     dmsReplicationInstance,
		"AWS::DMS::ReplicationSubnetGroup":                  dmsReplicationSubnetGroup,
		"AWS::DMS::ReplicationTask":                         dmsReplicationTask,
		"AWS::DocDB::DBCluster":                             docDBDBCluster,
		"AWS::DocDB::DBClusterParameterGroup":               docDBDBClusterParameterGroup,
		"AWS::DocDB::DBInstance":                            docDBDBInstance,
		"AWS::DocDB::DBSubnetGroup":                         docDBDBSubnetGroup,
		"AWS::DynamoDB::Table":                              dynamoDBTable,
		"AWS::EC2::CapacityReservation":                     ec2CapacityReservation,
		"AWS::EC2::ClientVpnEndpoint":                       ec2ClientVpnEndpoint,
		"AWS::EC2::CustomerGateway":                         ec2CustomerGateway,
		"AWS::EC2::DHCPOptions":                             ec2DHCPOptions,
		"AWS::EC2::EC2Fleet":                                ec2EC2Fleet,
		"AWS::EC2::EgressOnlyInternetGateway":               ec2EgressOnlyInternetGateway,
		"AWS::EC2::EIP":                                     ec2EIP,
		"AWS::EC2::EIPAssociation":                          ec2EIPAssociation,
		"AWS::EC2::FlowLog":                                 ec2FlowLog,
		"AWS::EC2::Host":                                    ec2Host,
		"AWS::EC2::Instance":                                ec2Instance,
		"AWS::EC2::InternetGateway":                         ec2InternetGateway,
		"AWS::EC2::LaunchTemplate":                          ec2LaunchTemplate,
		"AWS::EC2::NatGateway":                              ec2NatGateway,
		"AWS::EC2::NetworkAcl":                              ec2NetworkACL,
		"AWS::EC2::NetworkInterface":                        ec2NetworkInterface,
		"AWS::EC2::NetworkInterfaceAttachment":              ec2NetworkInterfaceAttachment,
		"AWS::EC2::NetworkInterfacePermission":              ec2NetworkInterfacePermission,
		"AWS::EC2::PlacementGroup":                          ec2PlacementGroup,
		"AWS::EC2::RouteTable":                              ec2RouteTable,
		"AWS::EC2::SecurityGroup":                           ec2SecurityGroup,
		"AWS::EC2::SpotFleet":                               ec2SpotFleet,
		"AWS::EC2::Subnet":                                  ec2Subnet,
		"AWS::EC2::SubnetNetworkAclAssociation":             ec2NetworkACLSubnetAssociation,
		"AWS::EC2::SubnetRouteTableAssociation":             ec2RouteTableSubnetAssociation,
		"AWS::EC2::TrafficMirrorFilter":                     ec2TrafficMirrorFilter,
		"AWS::EC2::TrafficMirrorFilterRule":                 ec2TrafficMirrorFilterRule,
		"AWS::EC2::TrafficMirrorSession":                    ec2TrafficMirrorSession,
		"AWS::EC2::TrafficMirrorTarget":                     ec2TrafficMirrorTarget,
		"AWS::EC2::TransitGateway":                          ec2TransitGateway,
		"AWS::EC2::TransitGatewayAttachment":                ec2TransitGatewayAttachment,
		"AWS::EC2::TransitGatewayRouteTable":                ec2TransitGatewayRouteTable,
		"AWS::EC2::Volume":                                  ec2Volume,
		"AWS::EC2::VPC":                                     ec2VPC,
		"AWS::EC2::VPCCidrBlock":                            ec2VPCCidrBlock,
		"AWS::EC2::VPCEndpoint":                             ec2VPCEndpoint,
		"AWS::EC2::VPCEndpointConnectionNotification":       ec2VPCEndpointConnectionNotification,
		"AWS::EC2::VPCEndpointService":                      ec2VPCEndpointService,
		"AWS::EC2::VPCPeeringConnection":                    ec2VPCPeeringConnection,
		"AWS::EC2::VPNConnection":                           ec2VPNConnection,
		"AWS::EC2::VPNGateway":                              ec2VPNGateway,
		"AWS::ECR::Repository":                              ecrRepository,
		"AWS::ECS::Cluster":                                 ecsCluster,
		"AWS::ECS::Service":                                 ecsService,
		"AWS::ECS::TaskDefinition":                          ecsTaskDefinition,
		"AWS::EFS::FileSystem":                              efsFileSystem,
		"AWS::EFS::MountTarget":                             efsMountTarget,
		"AWS::EKS::Cluster":                                 eksCluster,
		"AWS::EKS::Nodegroup":                               eksNodegroup,
		"AWS::ElastiCache::CacheCluster":                    elastiCacheCacheCluster,
		"AWS::ElastiCache::ParameterGroup":                  elastiCacheParameterGroup,
		"AWS::ElastiCache::ReplicationGroup":                elastiCacheReplicationGroup,
		"AWS::ElastiCache::SecurityGroup":                   elastiCacheSecurityGroup,
		"AWS::ElastiCache::SubnetGroup":                     elastiCacheSubnetGroup,
		"AWS::Elasticsearch::Domain":                        elasticsearchDomain,
		"AWS::ElasticBeanstalk::Application":                elasticBeanstalkApplication,
		"AWS::ElasticBeanstalk::ApplicationVersion":         elasticBeanstalkApplicationVersion,
		"AWS::ElasticBeanstalk::ConfigurationTemplate":      elasticBeanstalkConfigurationTemplate,
		"AWS::ElasticBeanstalk::Environment":                elasticBeanstalkEnvironment,
		"AWS::ElasticLoadBalancing::LoadBalancer":           elasticLoadBalancingLoadBalancer,
		"AWS::ElasticLoadBalancingV2::Listener":             elasticLoadBalancingV2Listener,
		"AWS::ElasticLoadBalancingV2::ListenerCertificate":  elasticLoadBalancingV2ListenerCertificate,
		"AWS::ElasticLoadBalancingV2::ListenerRule":         elasticLoadBalancingV2ListenerRule,
		"AWS::ElasticLoadBalancingV2::LoadBalancer":         elasticLoadBalancingV2LoadBalancer,
		"AWS::ElasticLoadBalancingV2::TargetGroup":          elasticLoadBalancingV2TargetGroup,
		"AWS::EMR::Cluster":                                 emrCluster,
		"AWS::EMR::InstanceFleetConfig":                     emrInstanceFleetConfig,
		"AWS::EMR::InstanceGroupConfig":                     emrInstanceGroupConfig,
		"AWS::EMR::SecurityConfiguration":                   emrSecurityConfiguration,
		"AWS::EMR::Step":                                    emrStep,
		"AWS::EventSchemas::Discoverer":                     eventSchemasDiscoverer,
		"AWS::EventSchemas::Registry":                       eventSchemasRegistry,
		"AWS::EventSchemas::Schema":                         eventSchemasSchema,
		"AWS::FSx::FileSystem":                              fsxFileSystem,
		"AWS::GameLift::Alias":                              gameLiftAlias,
		"AWS::GameLift::Build":                              gameLiftBuild,
		"AWS::GameLift::Fleet":                              gameLiftFleet,
		"AWS::GameLift::GameSessionQueue":                   gameLiftGameSessionQueue,
		"AWS::GameLift::MatchmakingConfiguration":           gameLiftMatchmakingConfiguration,
		"AWS::GameLift::MatchmakingRuleSet":                 gameLiftMatchmakingRuleSet,
		"AWS::GameLift::Script":                             gameLiftScript,
		"AWS::Glue::Connection":                             glueConnection,
		"AWS::Glue::Crawler":                                glueCrawler,
		"AWS::Glue::Database":                               glueDatabase,
		"AWS::Glue::DevEndpoint":                            glueDevEndpoint,
		"AWS::Glue::Job":                                    glueJob,
		"AWS::Glue::MLTransform":                            glueMLTransform,
		"AWS::Glue::SecurityConfiguration":                  glueSecurityConfiguration,
		"AWS::Glue::Table":                                  glueTable,
		"AWS::Glue::Trigger":                                glueTrigger,
		"AWS::Glue::Workflow":                               glueWorkflow,
		"AWS::GuardDuty::Detector":                          guardDutyDetector,
		"AWS::IAM::AccessKey":                               iamAccessKey,
		"AWS::IAM::Group":                                   iamGroup,
		"AWS::IAM::InstanceProfile":                         iamInstanceProfile,
		"AWS::IAM::Policy":                                  iamPolicy,
		"AWS::IAM::Role":                                    iamRole,
		"AWS::IAM::ManagedPolicy":                           iamRolePolicy,
		"AWS::IAM::ServiceLinkedRole":                       iamServiceLinkedRole,
		"AWS::IAM::User":                                    iamUser,
		"AWS::Inspector::AssessmentTarget":                  inspectorAssessmentTarget,
		"AWS::Inspector::AssessmentTemplate":                inspectorAssessmentTemplate,
		"AWS::Inspector::ResourceGroup":                     inspectorResourceGroup,
		"AWS::IoT::Certificate":                             ioTCertificate,
		"AWS::IoT::Policy":                                  ioTPolicy,
		"AWS::IoT::PolicyPrincipalAttachment":               ioTPolicyPrincipalAttachment,
		"AWS::IoT::Thing":                                   ioTThing,
		"AWS::IoT::ThingPrincipalAttachment":                ioTThingPrincipalAttachment,
		"AWS::IoT::TopicRule":                               ioTTopicRule,
		"AWS::IoT1Click::Device":                            ioT1ClickDevice,
		"AWS::IoT1Click::Placement":                         ioT1ClickPlacement,
		"AWS::IoT1Click::Project":                           ioT1ClickProject,
		"AWS::IoTAnalytics::Channel":                        ioTAnalyticsChannel,
		"AWS::IoTAnalytics::Dataset":                        ioTAnalyticsDataset,
		"AWS::IoTAnalytics::Datastore":                      ioTAnalyticsDatastore,
		"AWS::IoTAnalytics::Pipeline":                       ioTAnalyticsPipeline,
		"AWS::IoTEvents::DetectorModel":                     ioTEventsDetectorModel,
		"AWS::IoTEvents::Input":                             ioTEventsInput,
		"AWS::Greengrass::ConnectorDefinition":              greengrassConnectorDefinition,
		"AWS::Greengrass::ConnectorDefinitionVersion":       greengrassConnectorDefinitionVersion,
		"AWS::Greengrass::CoreDefinition":                   greengrassCoreDefinition,
		"AWS::Greengrass::CoreDefinitionVersion":            greengrassCoreDefinitionVersion,
		"AWS::Greengrass::DeviceDefinition":                 greengrassDeviceDefinition,
		"AWS::Greengrass::DeviceDefinitionVersion":          greengrassDeviceDefinitionVersion,
		"AWS::Greengrass::FunctionDefinition":               greengrassFunctionDefinition,
		"AWS::Greengrass::FunctionDefinitionVersion":        greengrassFunctionDefinitionVersion,
		"AWS::Greengrass::Group":                            greengrassGroup,
		"AWS::Greengrass::GroupVersion":                     greengrassGroupVersion,
		"AWS::Greengrass::LoggerDefinition":                 greengrassLoggerDefinition,
		"AWS::Greengrass::LoggerDefinitionVersion":          greengrassLoggerDefinitionVersion,
		"AWS::Greengrass::ResourceDefinition":               greengrassResourceDefinition,
		"AWS::Greengrass::ResourceDefinitionVersion":        greengrassResourceDefinitionVersion,
		"AWS::Greengrass::SubscriptionDefinition":           greengrassSubscriptionDefinition,
		"AWS::Greengrass::SubscriptionDefinitionVersion":    greengrassSubscriptionDefinitionVersion,
		"AWS::IoTThingsGraph::FlowTemplate":                 ioTThingsGraphFlowTemplate,
		"AWS::Kinesis::Stream":                              kinesisStream,
		"AWS::Kinesis::StreamConsumer":                      kinesisStreamConsumer,
		"AWS::KinesisAnalytics::Application":                kinesisAnalyticsApplication,
		"AWS::KinesisAnalyticsV2::Application":              kinesisAnalyticsV2Application,
		"AWS::KinesisFirehose::DeliveryStream":              kinesisFirehoseDeliveryStream,
		"AWS::KMS::Alias":                                   kmsAlias,
		"AWS::KMS::Key":                                     kmsKey,
		"AWS::LakeFormation::DataLakeSettings":              lakeFormationDataLakeSettings,
		"AWS::LakeFormation::Permissions":                   lakeFormationPermissions,
		"AWS::LakeFormation::Resource":                      lakeFormationResource,
		"AWS::Lambda::Alias":                                lambdaAlias,
		"AWS::Lambda::Function":                             lambdaFunction,
		"AWS::Lambda::LayerVersion":                         lambdaLayerVersion,
		"AWS::Lambda::Permission":                           lambdaPermission,
		"AWS::ManagedBlockchain::Member":                    managedBlockchainMember,
		"AWS::ManagedBlockchain::Node":                      managedBlockchainNode,
		"AWS::MediaConvert::JobTemplate":                    mediaConvertJobTemplate,
		"AWS::MediaConvert::Preset":                         mediaConvertPreset,
		"AWS::MediaConvert::Queue":                          mediaConvertQueue,
		"AWS::MediaLive::Channel":                           mediaLiveChannel,
		"AWS::MediaLive::Input":                             mediaLiveInput,
		"AWS::MediaLive::InputSecurityGroup":                mediaLiveInputSecurityGroup,
		"AWS::MediaStore::Container":                        mediaStoreContainer,
		"AWS::MSK::Cluster":                                 mskCluster,
		"AWS::Neptune::DBCluster":                           neptuneDBCluster,
		"AWS::Neptune::DBClusterParameterGroup":             neptuneDBClusterParameterGroup,
		"AWS::Neptune::DBInstance":                          neptuneDBInstance,
		"AWS::Neptune::DBParameterGroup":                    neptuneDBParameterGroup,
		"AWS::Neptune::DBSubnetGroup":                       neptuneDBSubnetGroup,
		"AWS::OpsWorks::App":                                opsWorksApp,
		"AWS::OpsWorks::ElasticLoadBalancerAttachment":      opsWorksElasticLoadBalancerAttachment,
		"AWS::OpsWorks::Instance":                           opsWorksInstance,
		"AWS::OpsWorks::Layer":                              opsWorksLayer,
		"AWS::OpsWorks::Stack":                              opsWorksStack,
		"AWS::OpsWorks::UserProfile":                        opsWorksUserProfile,
		"AWS::OpsWorks::Volume":                             opsWorksVolume,
		"AWS::OpsWorksCM::Server":                           opsWorksCMServer,
		"AWS::QLDB::Ledger":                                 qLDBLedger,
		"AWS::RAM::ResourceShare":                           ramResourceShare,
		"AWS::RDS::DBCluster":                               rdsDBCluster,
		"AWS::RDS::DBClusterParameterGroup":                 rdsDBClusterParameterGroup,
		"AWS::RDS::DBInstance":                              rdsDBInstance,
		"AWS::RDS::DBParameterGroup":                        rdsDBParameterGroup,
		"AWS::RDS::DBSecurityGroup":                         rdsDBSecurityGroup,
		"AWS::RDS::DBSubnetGroup":                           rdsDBSubnetGroup,
		"AWS::RDS::EventSubscription":                       rdsEventSubscription,
		"AWS::RDS::OptionGroup":                             rdsOptionGroup,
		"AWS::Redshift::Cluster":                            redshiftCluster,
		"AWS::Redshift::ClusterParameterGroup":              redshiftClusterParameterGroup,
		"AWS::Redshift::ClusterSecurityGroup":               redshiftClusterSecurityGroup,
		"AWS::Redshift::ClusterSubnetGroup":                 redshiftClusterSubnetGroup,
		"AWS::RoboMaker::Fleet":                             roboMakerFleet,
		"AWS::RoboMaker::Robot":                             roboMakerRobot,
		"AWS::RoboMaker::RobotApplication":                  roboMakerRobotApplication,
		"AWS::RoboMaker::RobotApplicationVersion":           roboMakerRobotApplicationVersion,
		"AWS::RoboMaker::SimulationApplication":             roboMakerSimulationApplication,
		"AWS::RoboMaker::SimulationApplicationVersion":      roboMakerSimulationApplicationVersion,
		"AWS::Route53::HealthCheck":                         route53HealthCheck,
		"AWS::Route53::HostedZone":                          route53HostedZone,
		"AWS::Route53::RecordSet":                           route53RecordSet,
		"AWS::Route53Resolver::ResolverEndpoint":            route53ResolverResolverEndpoint,
		"AWS::Route53Resolver::ResolverRule":                route53ResolverResolverRule,
		"AWS::Route53Resolver::ResolverRuleAssociation":     route53ResolverResolverRuleAssociation,
		"AWS::S3::AccessPoint":                              s3AccessPoint,
		"AWS::S3::Bucket":                                   s3Bucket,
		"AWS::S3::BucketPolicy":                             s3BucketPolicy,
		"AWS::SageMaker::CodeRepository":                    sageMakerCodeRepository,
		"AWS::SageMaker::Endpoint":                          sageMakerEndpoint,
		"AWS::SageMaker::EndpointConfig":                    sageMakerEndpointConfig,
		"AWS::SageMaker::Model":                             sageMakerModel,
		"AWS::SageMaker::NotebookInstance":                  sageMakerNotebookInstance,
		"AWS::SageMaker::NotebookInstanceLifecycleConfig":   sageMakerNotebookInstanceLifecycleConfig,
		"AWS::SageMaker::Workteam":                          sageMakerWorkteam,
		"AWS::SecretsManager::Secret":                       secretsManagerSecret,
		"AWS::SecurityHub::Hub":                             securityHubHub,
		"AWS::SES::ConfigurationSet":                        sesConfigurationSet,
		"AWS::SES::ReceiptFilter":                           sesReceiptFilter,
		"AWS::SES::ReceiptRuleSet":                          sesReceiptRuleSet,
		"AWS::SES::Template":                                sesTemplate,
		"AWS::SDB::Domain":                                  sdbDomain,
		"AWS::SNS::Subscription":                            snsSubscription,
		"AWS::SNS::Topic":                                   snsTopic,
		"AWS::SQS::Queue":                                   sqsQueue,
		"AWS::StepFunctions::Activity":                      stepFunctionsActivity,
		"AWS::StepFunctions::StateMachine":                  stepFunctionsStateMachine,
		"AWS::SSM::Association":                             ssmAssociation,
		"AWS::SSM::Document":                                ssmDocument,
		"AWS::SSM::MaintenanceWindow":                       ssmMaintenanceWindow,
		"AWS::SSM::MaintenanceWindowTarget":                 ssmMaintenanceWindowTarget,
		"AWS::SSM::MaintenanceWindowTask":                   ssmMaintenanceWindowTask,
		"AWS::SSM::Parameter":                               ssmParameter,
		"AWS::SSM::PatchBaseline":                           ssmPatchBaseline,
		"AWS::Transfer::Server":                             transferServer,
		"AWS::Transfer::User":                               transferUser,
		"AWS::WAF::ByteMatchSet":                            wafByteMatchSet,
		"AWS::WAF::IPSet":                                   wafIPSet,
		"AWS::WAF::Rule":                                    wafRule,
		"AWS::WAF::SizeConstraintSet":                       wafSizeConstraintSet,
		"AWS::WAF::SqlInjectionMatchSet":                    wafSQLInjectionMatchSet,
		"AWS::WAF::WebACL":                                  wafWebACL,
		"AWS::WAF::XssMatchSet":                             wafXSSMatchSet,
		"AWS::WAFv2::IPSet":                                 wafv2IPSet,
		"AWS::WAFv2::RegexPatternSet":                       wafv2RegexPatternSet,
		"AWS::WAFv2::RuleGroup":                             wafv2RuleGroup,
		"AWS::WAFv2::WebACL":                                wafv2WebACL,
		"AWS::WAFRegional::ByteMatchSet":                    wafRegionalByteMatchSet,
		"AWS::WAFRegional::GeoMatchSet":                     wafRegionalGeoMatchSet,
		"AWS::WAFRegional::IPSet":                           wafRegionalIPSet,
		"AWS::WAFRegional::RateBasedRule":                   wafRegionalRateBasedRule,
		"AWS::WAFRegional::RegexPatternSet":                 wafRegionalRegexPatternSet,
		"AWS::WAFRegional::Rule":                            wafRegionalRule,
		"AWS::WAFRegional::SizeConstraintSet":               wafRegionalSizeConstraintSet,
		"AWS::WAFRegional::SqlInjectionMatchSet":            wafRegionalSQLInjectionMatchSet,
		"AWS::WAFRegional::WebACL":                          wafRegionalWebACL,
		"AWS::WAFRegional::XssMatchSet":                     wafRegionalXSSMatchSet,
		"AWS::WorkSpaces::Workspace":                        workSpacesWorkspace,
	}
	value, ok := cfn[cloudFormationType]
	return value, ok
}

func (state state) filter(from resourceSource, to resourceSource) (r resourceMap) {
	logInfo("Filter all resources that are in the", from, "state out of the", to, "state")
	r = make(map[resourceType][]string)
	for resourceType, resourceIDs := range state[from] {
		for _, resourceID := range resourceIDs {
			if !contains(state[to][resourceType], resourceID) {
				logDebug("State", "does not contain resource of type", resourceType, "with identifier", resourceID)
				r[resourceType] = append(r[resourceType], resourceID)
			}
			logDebug("State", to, "contains resource of type", resourceType, "with identifier", resourceID)
		}
	}
	return
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

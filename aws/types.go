package aws

type resourceType string

const (
	accessAnalyzerAnalyzer                    resourceType = "accessAnalyzerAnalyzer"
	acmpcaCertificateAuthority                resourceType = "acmpcaCertificateAuthority"
	alexaAskSkill                             resourceType = "alexaAskSkill"
	amazonMQBroker                            resourceType = "amazonMQBroker"
	amazonMQConfiguration                     resourceType = "amazonMQConfiguration"
	amplifyApp                                resourceType = "amplifyApp"
	amplifyBranch                             resourceType = "amplifyBranch"
	amplifyDomain                             resourceType = "amplifyDomain"
	apiGatewayAPIKey                          resourceType = "apiGatewayAPIKey"
	apiGatewayClientCertificate               resourceType = "apiGatewayClientCertificate"
	apiGatewayDomainName                      resourceType = "apiGatewayDomainName"
	apiGatewayRestAPI                         resourceType = "apiGatewayRestAPI"
	apiGatewayUsagePlan                       resourceType = "apiGatewayUsagePlan"
	apiGatewayVpcLink                         resourceType = "apiGatewayVpcLink"
	apiGatewayV2Api                           resourceType = "apiGatewayV2Api"
	apiGatewayV2DomainName                    resourceType = "apiGatewayV2DomainName"
	appConfigApplication                      resourceType = "appConfigApplication"
	appConfigDeploymentStrategy               resourceType = "appConfigDeploymentStrategy"
	applicationAutoScalingScheduledAction     resourceType = "applicationAutoScalingScheduledAction"
	appMeshMesh                               resourceType = "appMeshMesh"
	appStreamDirectoryConfig                  resourceType = "appStreamDirectoryConfig"
	appStreamFleet                            resourceType = "appStreamFleet"
	appStreamImageBuilder                     resourceType = "appStreamImageBuilder"
	appStreamStack                            resourceType = "appStreamStack"
	appSyncFunctions                          resourceType = "appSyncFunctions"
	appSyncGraphQLApi                         resourceType = "appSyncGraphQLApi"
	athenaNamedQuery                          resourceType = "athenaNamedQuery"
	athenaWorkGroup                           resourceType = "athenaWorkGroup"
	autoScalingPlansScalingPlan               resourceType = "autoScalingPlansScalingPlan"
	autoScalingAutoScalingGroup               resourceType = "autoScalingAutoScalingGroup"
	autoScalingLaunchConfiguration            resourceType = "autoScalingLaunchConfiguration"
	autoScalingScalingPolicy                  resourceType = "autoScalingScalingPolicy"
	autoScalingScheduledAction                resourceType = "autoScalingScheduledAction"
	backupBackupPlan                          resourceType = "backupBackupPlan"
	backupBackupSelection                     resourceType = "backupBackupSelection"
	backupBackupVault                         resourceType = "backupBackupVault"
	batchComputeEnvironment                   resourceType = "batchComputeEnvironment"
	batchJobDefinition                        resourceType = "batchJobDefinition"
	batchJobQueue                             resourceType = "batchJobQueue"
	budgetsBudget                             resourceType = "budgetsBudget"
	certificateManagerCertificate             resourceType = "certificateManagerCertificate"
	cloud9EnvironmentEC2                      resourceType = "cloud9EnvironmentEC2"
	cloudFrontCloudFrontOriginAccessIdentity  resourceType = "cloudFrontCloudFrontOriginAccessIdentity"
	cloudFrontDistribution                    resourceType = "cloudFrontDistribution"
	cloudFrontStreamingDistribution           resourceType = "cloudFrontStreamingDistribution"
	cloudFrontPublicKey                       resourceType = "cloudFrontPublicKey"
	cloudHSMV2Cluster                         resourceType = "cloudHSMV2Cluster"
	cloudHSMV2HSM                             resourceType = "cloudHSMV2HSM"
	serviceDiscoveryHTTPNamespace             resourceType = "serviceDiscoveryHTTPNamespace"
	serviceDiscoveryPrivateDNSNamespace       resourceType = "serviceDiscoveryPrivateDNSNamespace"
	serviceDiscoveryPublicDNSNamespace        resourceType = "serviceDiscoveryPublicDNSNamespace"
	serviceDiscoveryService                   resourceType = "serviceDiscoveryService"
	cloudTrailTrail                           resourceType = "cloudTrailTrail"
	cloudWatchAlarm                           resourceType = "cloudWatchAlarm"
	cloudWatchDashboard                       resourceType = "cloudWatchDashboard"
	cloudWatchInsightRule                     resourceType = "cloudWatchInsightRule"
	logsDestination                           resourceType = "logsDestination"
	logsLogGroup                              resourceType = "logsLogGroup"
	logsMetricFilter                          resourceType = "logsMetricFilter"
	logsSubscriptionFilter                    resourceType = "logsSubscriptionFilter"
	logsResourcePolicy                        resourceType = "logsResourcePolicy"
	macieMemberAccountAssociation             resourceType = "macieMemberAccountAssociation"
	macieS3BucketAssociation                  resourceType = "macieS3BucketAssociation"
	eventsEventBus                            resourceType = "eventsEventBus"
	eventsRule                                resourceType = "eventsRule"
	eventsTarget                              resourceType = "eventsTarget"
	codeBuildProject                          resourceType = "codeBuildProject"
	codeBuildReportGroup                      resourceType = "codeBuildReportGroup"
	codeBuildSourceCredential                 resourceType = "codeBuildSourceCredential"
	codeCommitRepository                      resourceType = "codeCommitRepository"
	codeCommitTrigger                         resourceType = "codeCommitTrigger"
	codeDeployApplication                     resourceType = "codeDeployApplication"
	codeDeployDeploymentConfig                resourceType = "codeDeployDeploymentConfig"
	codeDeployDeploymentGroup                 resourceType = "codeDeployDeploymentGroup"
	codeGuruProfilerProfilingGroup            resourceType = "codeGuruProfilerProfilingGroup"
	codePipelinePipeline                      resourceType = "codePipelinePipeline"
	codePipelineWebhook                       resourceType = "codePipelineWebhook"
	codeStarGitHubRepository                  resourceType = "codeStarGitHubRepository"
	codeStarNotificationsNotificationRule     resourceType = "codeStarNotificationsNotificationRule"
	cognitoIdentityPool                       resourceType = "cognitoIdentityPool"
	cognitoUserPool                           resourceType = "cognitoUserPool"
	cognitoUserPoolClient                     resourceType = "cognitoUserPoolClient"
	cognitoUserPoolGroup                      resourceType = "cognitoUserPoolGroup"
	cognitoUserPoolIdentityProvider           resourceType = "cognitoUserPoolIdentityProvider"
	cognitoUserPoolResourceServer             resourceType = "cognitoUserPoolResourceServer"
	cognitoUserPoolUser                       resourceType = "cognitoUserPoolUser"
	configAggregationAuthorization            resourceType = "configAggregationAuthorization"
	configConfigRule                          resourceType = "configConfigRule"
	configConfigurationAggregator             resourceType = "configConfigurationAggregator"
	configConfigurationRecorder               resourceType = "configConfigurationRecorder"
	configConformancePack                     resourceType = "configConformancePack"
	configDeliveryChannel                     resourceType = "configDeliveryChannel"
	configOrganizationConfigRule              resourceType = "configOrganizationConfigRule"
	configOrganizationConformancePack         resourceType = "configOrganizationConformancePack"
	configRemediationConfiguration            resourceType = "configRemediationConfiguration"
	costAndUsageReportServiceReportDefinition resourceType = "costAndUsageReportServiceReportDefinition"
	dataPipelinePipeline                      resourceType = "dataPipelinePipeline"
	dataSyncAgent                             resourceType = "dataSyncAgent"
	dataSyncLocation                          resourceType = "dataSyncLocation"
	dataSyncTask                              resourceType = "dataSyncTask"
	daxCluster                                resourceType = "daxCluster"
	daxParameterGroup                         resourceType = "daxParameterGroup"
	daxSubnetGroup                            resourceType = "daxSubnetGroup"
	deviceFarmProject                         resourceType = "deviceFarmProject"
	directConnectConnection                   resourceType = "directConnectConnection"
	directConnectGateway                      resourceType = "directConnectGateway"
	directConnectGatewayAssociation           resourceType = "directConnectGatewayAssociation"
	directConnectGatewayAssociationProposal   resourceType = "directConnectGatewayAssociationProposal"
	directConnectLAG                          resourceType = "directConnectLAG"
	directConnectVirtualInterface             resourceType = "directConnectVirtualInterface"
	directoryServiceDirectory                 resourceType = "directoryServiceDirectory"
	dlmLifecyclePolicy                        resourceType = "dlmLifecyclePolicy"
	dmsCertificate                            resourceType = "dmsCertificate"
	dmsEndpoint                               resourceType = "dmsEndpoint"
	dmsEventSubscription                      resourceType = "dmsEventSubscription"
	dmsReplicationInstance                    resourceType = "dmsReplicationInstance"
	dmsReplicationSubnetGroup                 resourceType = "dmsReplicationSubnetGroup"
	dmsReplicationTask                        resourceType = "dmsReplicationTask"
	docDBDBCluster                            resourceType = "docDBDBCluster"
	docDBDBClusterParameterGroup              resourceType = "docDBDBClusterParameterGroup"
	docDBDBInstance                           resourceType = "docDBDBInstance"
	docDBDBSubnetGroup                        resourceType = "docDBDBSubnetGroup"
	docDBDBClusterSnapshot                    resourceType = "docDBDBClusterSnapshot"
	dynamoDBTable                             resourceType = "dynamoDBTable"
	dynamoDBGlobalTable                       resourceType = "dynamoDBGlobalTable"
	ec2CapacityReservation                    resourceType = "ec2CapacityReservation"
	ec2ClientVpnEndpoint                      resourceType = "ec2ClientVpnEndpoint"
	ec2CustomerGateway                        resourceType = "ec2CustomerGateway"
	ec2DHCPOptions                            resourceType = "ec2DHCPOptions"
	ec2EC2Fleet                               resourceType = "ec2EC2Fleet"
	ec2EgressOnlyInternetGateway              resourceType = "ec2EgressOnlyInternetGateway"
	ec2EIP                                    resourceType = "ec2EIP"
	ec2EIPAssociation                         resourceType = "ec2EIPAssociation"
	ec2FlowLog                                resourceType = "ec2FlowLog"
	ec2Host                                   resourceType = "ec2Host"
	ec2Image                                  resourceType = "ec2Image"
	ec2Instance                               resourceType = "ec2Instance"
	ec2InternetGateway                        resourceType = "ec2InternetGateway"
	ec2LaunchTemplate                         resourceType = "ec2LaunchTemplate"
	ec2NatGateway                             resourceType = "ec2NatGateway"
	ec2NetworkACL                             resourceType = "ec2NetworkACL"
	ec2NetworkInterface                       resourceType = "ec2NetworkInterface"
	ec2NetworkInterfaceAttachment             resourceType = "ec2NetworkInterfaceAttachment"
	ec2NetworkInterfacePermission             resourceType = "ec2NetworkInterfacePermission"
	ec2PlacementGroup                         resourceType = "ec2PlacementGroup"
	ec2RouteTable                             resourceType = "ec2RouteTable"
	ec2SecurityGroup                          resourceType = "ec2SecurityGroup"
	ec2Snapshot                               resourceType = "ec2Snapshot"
	ec2SpotFleet                              resourceType = "ec2SpotFleet"
	ec2Subnet                                 resourceType = "ec2Subnet"
	ec2NetworkACLSubnetAssociation            resourceType = "ec2NetworkACLSubnetAssociation"
	ec2RouteTableSubnetAssociation            resourceType = "ec2RouteTableSubnetAssociation"
	ec2TrafficMirrorFilter                    resourceType = "ec2TrafficMirrorFilter"
	ec2TrafficMirrorFilterRule                resourceType = "ec2TrafficMirrorFilterRule"
	ec2TrafficMirrorSession                   resourceType = "ec2TrafficMirrorSession"
	ec2TrafficMirrorTarget                    resourceType = "ec2TrafficMirrorTarget"
	ec2TransitGateway                         resourceType = "ec2TransitGateway"
	ec2TransitGatewayAttachment               resourceType = "ec2TransitGatewayAttachment"
	ec2TransitGatewayRouteTable               resourceType = "ec2TransitGatewayRouteTable"
	ec2Volume                                 resourceType = "ec2Volume"
	ec2VPC                                    resourceType = "ec2VPC"
	ec2VPCCidrBlock                           resourceType = "ec2VPCCidrBlock"
	ec2VPCEndpoint                            resourceType = "ec2VPCEndpoint"
	ec2VPCEndpointConnectionNotification      resourceType = "ec2VPCEndpointConnectionNotification"
	ec2VPCEndpointService                     resourceType = "ec2VPCEndpointService"
	ec2VPCPeeringConnection                   resourceType = "ec2VPCPeeringConnection"
	ec2VPNConnection                          resourceType = "ec2VPNConnection"
	ec2VPNGateway                             resourceType = "ec2VPNGateway"
	ec2KeyPair                                resourceType = "ec2KeyPair"
	ec2SpotInstanceRequest                    resourceType = "ec2SpotInstanceRequest"
	ecrRepository                             resourceType = "ecrRepository"
	ecsCluster                                resourceType = "ecsCluster"
	ecsService                                resourceType = "ecsService"
	ecsTaskDefinition                         resourceType = "ecsTaskDefinition"
	ecsCapacityProvider                       resourceType = "ecsCapacityProvider"
	efsFileSystem                             resourceType = "efsFileSystem"
	efsMountTarget                            resourceType = "efsMountTarget"
	eksCluster                                resourceType = "eksCluster"
	eksNodegroup                              resourceType = "eksNodegroup"
	eksFargateProfile                         resourceType = "eksFargateProfile"
	elastiCacheCacheCluster                   resourceType = "elastiCacheCacheCluster"
	elastiCacheParameterGroup                 resourceType = "elastiCacheParameterGroup"
	elastiCacheReplicationGroup               resourceType = "elastiCacheReplicationGroup"
	elastiCacheSecurityGroup                  resourceType = "elastiCacheSecurityGroup"
	elastiCacheSubnetGroup                    resourceType = "elastiCacheSubnetGroup"
	elasticsearchDomain                       resourceType = "elasticsearchDomain"
	elasticTranscoderPipeline                 resourceType = "elasticTranscoderPipeline"
	elasticTranscoderPreset                   resourceType = "elasticTranscoderPreset"
	elasticBeanstalkApplication               resourceType = "elasticBeanstalkApplication"
	elasticBeanstalkApplicationVersion        resourceType = "elasticBeanstalkApplicationVersion"
	elasticBeanstalkConfigurationTemplate     resourceType = "elasticBeanstalkConfigurationTemplate"
	elasticBeanstalkEnvironment               resourceType = "elasticBeanstalkEnvironment"
	elasticLoadBalancingLoadBalancer          resourceType = "elasticLoadBalancingLoadBalancer"
	elasticLoadBalancingV2Listener            resourceType = "elasticLoadBalancingV2Listener"
	elasticLoadBalancingV2ListenerRule        resourceType = "elasticLoadBalancingV2ListenerRule"
	elasticLoadBalancingV2LoadBalancer        resourceType = "elasticLoadBalancingV2LoadBalancer"
	elasticLoadBalancingV2TargetGroup         resourceType = "elasticLoadBalancingV2TargetGroup"
	emrCluster                                resourceType = "emrCluster"
	emrSecurityConfiguration                  resourceType = "emrSecurityConfiguration"
	emrInstanceGroup                          resourceType = "emrInstanceGroup"
	eventSchemasDiscoverer                    resourceType = "eventSchemasDiscoverer"
	eventSchemasRegistry                      resourceType = "eventSchemasRegistry"
	fsxFileSystem                             resourceType = "fsxFileSystem"
	gameLiftAlias                             resourceType = "gameLiftAlias"
	gameLiftBuild                             resourceType = "gameLiftBuild"
	gameLiftFleet                             resourceType = "gameLiftFleet"
	gameLiftGameSessionQueue                  resourceType = "gameLiftGameSessionQueue"
	gameLiftMatchmakingConfiguration          resourceType = "gameLiftMatchmakingConfiguration"
	gameLiftMatchmakingRuleSet                resourceType = "gameLiftMatchmakingRuleSet"
	gameLiftScript                            resourceType = "gameLiftScript"
	glueConnection                            resourceType = "glueConnection"
	glueCrawler                               resourceType = "glueCrawler"
	glueDatabase                              resourceType = "glueDatabase"
	glueDevEndpoint                           resourceType = "glueDevEndpoint"
	glueJob                                   resourceType = "glueJob"
	glueMLTransform                           resourceType = "glueMLTransform"
	glueSecurityConfiguration                 resourceType = "glueSecurityConfiguration"
	glueTable                                 resourceType = "glueTable"
	glueTrigger                               resourceType = "glueTrigger"
	glueWorkflow                              resourceType = "glueWorkflow"
	glacierVault                              resourceType = "glacierVault"
	globalAcceleratorAccelerator              resourceType = "globalAcceleratorAccelerator"
	globalAcceleratorEndpointGroup            resourceType = "globalAcceleratorEndpointGroup"
	globalAcceleratorListener                 resourceType = "globalAcceleratorListener"
	groundStationConfig                       resourceType = "groundStationConfig"
	groundStationDataflowEndpointGroup        resourceType = "groundStationDataflowEndpointGroup"
	groundStationMissionProfile               resourceType = "groundStationMissionProfile"
	guardDutyDetector                         resourceType = "guardDutyDetector"
	iamAccessKey                              resourceType = "iamAccessKey"
	iamAccountAlias                           resourceType = "iamAccountAlias"
	iamGroup                                  resourceType = "iamGroup"
	iamGroupPolicy                            resourceType = "iamGroupPolicy"
	iamInstanceProfile                        resourceType = "iamInstanceProfile"
	iamOpenidConnectProvider                  resourceType = "iamOpenidConnectProvider"
	iamPolicy                                 resourceType = "iamPolicy"
	iamRole                                   resourceType = "iamRole"
	iamRolePolicy                             resourceType = "iamRolePolicy"
	iamSamlProvider                           resourceType = "iamSamlProvider"
	iamServerCertificate                      resourceType = "iamServerCertificate"
	iamServiceLinkedRole                      resourceType = "iamServiceLinkedRole"
	iamUser                                   resourceType = "iamUser"
	iamUserPolicy                             resourceType = "iamUserPolicy"
	iamUserSSHKey                             resourceType = "iamUserSSHKey"
	inspectorAssessmentTarget                 resourceType = "inspectorAssessmentTarget"
	inspectorAssessmentTemplate               resourceType = "inspectorAssessmentTemplate"
	ioTCertificate                            resourceType = "ioTCertificate"
	ioTPolicy                                 resourceType = "ioTPolicy"
	iotRoleAlias                              resourceType = "iotRoleAlias"
	ioTThing                                  resourceType = "ioTThing"
	iotThingType                              resourceType = "iotThingType"
	ioTTopicRule                              resourceType = "ioTTopicRule"
	ioT1ClickDevice                           resourceType = "ioT1ClickDevice"
	ioT1ClickProject                          resourceType = "ioT1ClickProject"
	ioTAnalyticsChannel                       resourceType = "ioTAnalyticsChannel"
	ioTAnalyticsDataset                       resourceType = "ioTAnalyticsDataset"
	ioTAnalyticsDatastore                     resourceType = "ioTAnalyticsDatastore"
	ioTAnalyticsPipeline                      resourceType = "ioTAnalyticsPipeline"
	ioTEventsDetectorModel                    resourceType = "ioTEventsDetectorModel"
	ioTEventsInput                            resourceType = "ioTEventsInput"
	greengrassConnectorDefinition             resourceType = "greengrassConnectorDefinition"
	greengrassConnectorDefinitionVersion      resourceType = "greengrassConnectorDefinitionVersion"
	greengrassCoreDefinition                  resourceType = "greengrassCoreDefinition"
	greengrassCoreDefinitionVersion           resourceType = "greengrassCoreDefinitionVersion"
	greengrassDeviceDefinition                resourceType = "greengrassDeviceDefinition"
	greengrassDeviceDefinitionVersion         resourceType = "greengrassDeviceDefinitionVersion"
	greengrassFunctionDefinition              resourceType = "greengrassFunctionDefinition"
	greengrassFunctionDefinitionVersion       resourceType = "greengrassFunctionDefinitionVersion"
	greengrassGroup                           resourceType = "greengrassGroup"
	greengrassGroupVersion                    resourceType = "greengrassGroupVersion"
	greengrassLoggerDefinition                resourceType = "greengrassLoggerDefinition"
	greengrassLoggerDefinitionVersion         resourceType = "greengrassLoggerDefinitionVersion"
	greengrassResourceDefinition              resourceType = "greengrassResourceDefinition"
	greengrassResourceDefinitionVersion       resourceType = "greengrassResourceDefinitionVersion"
	greengrassSubscriptionDefinition          resourceType = "greengrassSubscriptionDefinition"
	greengrassSubscriptionDefinitionVersion   resourceType = "greengrassSubscriptionDefinitionVersion"
	ioTThingsGraphFlowTemplate                resourceType = "ioTThingsGraphFlowTemplate"
	kinesisStream                             resourceType = "kinesisStream"
	kinesisStreamConsumer                     resourceType = "kinesisStreamConsumer"
	kinesisAnalyticsApplication               resourceType = "kinesisAnalyticsApplication"
	kinesisAnalyticsV2Application             resourceType = "kinesisAnalyticsV2Application"
	kinesisVideoStream                        resourceType = "kinesisVideoStream"
	kinesisFirehoseDeliveryStream             resourceType = "kinesisFirehoseDeliveryStream"
	kmsAlias                                  resourceType = "kmsAlias"
	kmsGrant                                  resourceType = "kmsGrant"
	kmsKey                                    resourceType = "kmsKey"
	lakeFormationResource                     resourceType = "lakeFormationResource"
	lambdaAlias                               resourceType = "lambdaAlias"
	lambdaEventSourceMapping                  resourceType = "lambdaEventSourceMapping"
	lambdaFunction                            resourceType = "lambdaFunction"
	lambdaLayer                               resourceType = "lambdaLayer"
	lambdaLayerVersion                        resourceType = "lambdaLayerVersion"
	licenseManagerLicenseConfiguration        resourceType = "licenseManagerLicenseConfiguration"
	lightsailDomain                           resourceType = "lightsailDomain"
	lightsailInstance                         resourceType = "lightsailInstance"
	lightsailKeyPair                          resourceType = "lightsailKeyPair"
	lightsailStaticIP                         resourceType = "lightsailStaticIP"
	mediaConvertJobTemplate                   resourceType = "mediaConvertJobTemplate"
	mediaConvertPreset                        resourceType = "mediaConvertPreset"
	mediaConvertQueue                         resourceType = "mediaConvertQueue"
	mediaLiveChannel                          resourceType = "mediaLiveChannel"
	mediaLiveInput                            resourceType = "mediaLiveInput"
	mediaLiveInputSecurityGroup               resourceType = "mediaLiveInputSecurityGroup"
	mediaStoreContainer                       resourceType = "mediaStoreContainer"
	mskCluster                                resourceType = "mskCluster"
	mskConfiguration                          resourceType = "mskConfiguration"
	neptuneDBCluster                          resourceType = "neptuneDBCluster"
	neptuneDBClusterParameterGroup            resourceType = "neptuneDBClusterParameterGroup"
	neptuneDBClusterSnapshot                  resourceType = "neptuneDBClusterSnapshot"
	neptuneDBEventSubscription                resourceType = "neptuneDBEventSubscription"
	neptuneDBInstance                         resourceType = "neptuneDBInstance"
	neptuneDBParameterGroup                   resourceType = "neptuneDBParameterGroup"
	neptuneDBSubnetGroup                      resourceType = "neptuneDBSubnetGroup"
	networkManagerDevice                      resourceType = "networkManagerDevice"
	networkManagerGlobalNetwork               resourceType = "networkManagerGlobalNetwork"
	networkManagerLink                        resourceType = "networkManagerLink"
	networkManagerSite                        resourceType = "networkManagerSite"
	opsWorksApp                               resourceType = "opsWorksApp"
	opsWorksInstance                          resourceType = "opsWorksInstance"
	opsWorksLayer                             resourceType = "opsWorksLayer"
	opsWorkRdsDbInstance                      resourceType = "opsWorkRdsDbInstance"
	opsWorksStack                             resourceType = "opsWorksStack"
	opsWorksUserProfile                       resourceType = "opsWorksUserProfile"
	opsWorksVolume                            resourceType = "opsWorksVolume"
	organizationsAccount                      resourceType = "organizationsAccount"
	organizationsOrganization                 resourceType = "organizationsOrganization"
	organizationsOrganizationalUnit           resourceType = "organizationsOrganizationalUnit"
	organizationsPolicy                       resourceType = "organizationsPolicy"
	pinpointApp                               resourceType = "pinpointApp"
	pinpointEmailTemplate                     resourceType = "pinpointEmailTemplate"
	pinpointPushTemplate                      resourceType = "pinpointPushTemplate"
	pinpointSmsTemplate                       resourceType = "pinpointSmsTemplate"
	qLDBLedger                                resourceType = "qLDBLedger"
	quickSightGroup                           resourceType = "quickSightGroup"
	quickSightUser                            resourceType = "quickSightUser"
	rdsDBCluster                              resourceType = "rdsDBCluster"
	rdsDBClusterEndpoint                      resourceType = "rdsDBClusterEndpoint"
	rdsDBClusterParameterGroup                resourceType = "rdsDBClusterParameterGroup"
	rdsDBClusterSnapshot                      resourceType = "rdsDBClusterSnapshot"
	rdsDBInstance                             resourceType = "rdsDBInstance"
	rdsDBParameterGroup                       resourceType = "rdsDBParameterGroup"
	rdsDBSecurityGroup                        resourceType = "rdsDBSecurityGroup"
	rdsDBSnapshot                             resourceType = "rdsDBSnapshot"
	rdsDBSubnetGroup                          resourceType = "rdsDBSubnetGroup"
	rdsEventSubscription                      resourceType = "rdsEventSubscription"
	rdsGlobalCluster                          resourceType = "rdsGlobalCluster"
	rdsOptionGroup                            resourceType = "rdsOptionGroup"
	redshiftCluster                           resourceType = "redshiftCluster"
	redshiftClusterParameterGroup             resourceType = "redshiftClusterParameterGroup"
	redshiftClusterSecurityGroup              resourceType = "redshiftClusterSecurityGroup"
	redshiftClusterSubnetGroup                resourceType = "redshiftClusterSubnetGroup"
	redshiftEventSubscription                 resourceType = "redshiftEventSubscription"
	redshiftSnapshotCopyGrant                 resourceType = "redshiftSnapshotCopyGrant"
	redshiftSnapshotSchedule                  resourceType = "redshiftSnapshotSchedule"
	resourceGroupsGroup                       resourceType = "resourceGroupsGroup"
	roboMakerFleet                            resourceType = "roboMakerFleet"
	roboMakerRobot                            resourceType = "roboMakerRobot"
	roboMakerRobotApplication                 resourceType = "roboMakerRobotApplication"
	roboMakerSimulationApplication            resourceType = "roboMakerSimulationApplication"
	route53DelegationSet                      resourceType = "route53DelegationSet"
	route53HealthCheck                        resourceType = "route53HealthCheck"
	route53HostedZone                         resourceType = "route53HostedZone"
	route53RecordSet                          resourceType = "route53RecordSet"
	route53QueryLog                           resourceType = "route53QueryLog"
	route53ResolverResolverEndpoint           resourceType = "route53ResolverResolverEndpoint"
	route53ResolverResolverRule               resourceType = "route53ResolverResolverRule"
	route53ResolverResolverRuleAssociation    resourceType = "route53ResolverResolverRuleAssociation"
	s3AccessPoint                             resourceType = "s3AccessPoint"
	s3Bucket                                  resourceType = "s3Bucket"
	sageMakerCodeRepository                   resourceType = "sageMakerCodeRepository"
	sageMakerEndpoint                         resourceType = "sageMakerEndpoint"
	sageMakerEndpointConfig                   resourceType = "sageMakerEndpointConfig"
	sageMakerModel                            resourceType = "sageMakerModel"
	sageMakerNotebookInstance                 resourceType = "sageMakerNotebookInstance"
	sageMakerNotebookInstanceLifecycleConfig  resourceType = "sageMakerNotebookInstanceLifecycleConfig"
	sageMakerWorkteam                         resourceType = "sageMakerWorkteam"
	secretsManagerSecret                      resourceType = "secretsManagerSecret"
	secretsManagerSecretVersion               resourceType = "secretsManagerSecretVersion"
	sesConfigurationSet                       resourceType = "sesConfigurationSet"
	sesDomainIdentity                         resourceType = "sesDomainIdentity"
	sesEmailIdentity                          resourceType = "sesEmailIdentity"
	sesReceiptFilter                          resourceType = "sesReceiptFilter"
	sesReceiptRuleSet                         resourceType = "sesReceiptRuleSet"
	sesTemplate                               resourceType = "sesTemplate"
	sdbDomain                                 resourceType = "sdbDomain"
	snsPlatformApplication                    resourceType = "snsPlatformApplication"
	snsSubscription                           resourceType = "snsSubscription"
	snsTopic                                  resourceType = "snsTopic"
	sqsQueue                                  resourceType = "sqsQueue"
	stepFunctionsActivity                     resourceType = "stepFunctionsActivity"
	stepFunctionsStateMachine                 resourceType = "stepFunctionsStateMachine"
	shieldProtection                          resourceType = "shieldProtection"
	ssmActivation                             resourceType = "ssmActivation"
	ssmAssociation                            resourceType = "ssmAssociation"
	ssmDocument                               resourceType = "ssmDocument"
	ssmMaintenanceWindow                      resourceType = "ssmMaintenanceWindow"
	ssmMaintenanceWindowTarget                resourceType = "ssmMaintenanceWindowTarget"
	ssmMaintenanceWindowTask                  resourceType = "ssmMaintenanceWindowTask"
	ssmParameter                              resourceType = "ssmParameter"
	ssmPatchBaseline                          resourceType = "ssmPatchBaseline"
	ssmPatchGroup                             resourceType = "ssmPatchGroup"
	ssmResourceDataSync                       resourceType = "ssmResourceDataSync"
	storageGatewayCachedISCSIVolume           resourceType = "storageGatewayCachedISCSIVolume"
	storageGatewayGateway                     resourceType = "storageGatewayGateway"
	storageGatewayNFSFileShare                resourceType = "storageGatewayNFSFileShare"
	storageGatewaySMBFileShare                resourceType = "storageGatewaySMBFileShare"
	swfDomain                                 resourceType = "swfDomain"
	transferServer                            resourceType = "transferServer"
	transferUser                              resourceType = "transferUser"
	wafByteMatchSet                           resourceType = "wafByteMatchSet"
	wafGeoMatchSet                            resourceType = "wafGeoMatchSet"
	wafIPSet                                  resourceType = "wafIPSet"
	wafRateBasedRule                          resourceType = "wafRateBasedRule"
	wafRegexMatchSet                          resourceType = "wafRegexMatchSet"
	wafRegexPatternSet                        resourceType = "wafRegexPatternSet"
	wafRule                                   resourceType = "wafRule"
	wafRuleGroup                              resourceType = "wafRuleGroup"
	wafSizeConstraintSet                      resourceType = "wafSizeConstraintSet"
	wafSQLInjectionMatchSet                   resourceType = "wafSQLInjectionMatchSet"
	wafWebACL                                 resourceType = "wafWebACL"
	wafXSSMatchSet                            resourceType = "wafXSSMatchSet"
	wafv2IPSet                                resourceType = "wafv2IPSet"
	wafv2RegexPatternSet                      resourceType = "wafv2RegexPatternSet"
	wafv2RuleGroup                            resourceType = "wafv2RuleGroup"
	wafv2WebACL                               resourceType = "wafv2WebACL"
	wafRegionalByteMatchSet                   resourceType = "wafRegionalByteMatchSet"
	wafRegionalGeoMatchSet                    resourceType = "wafRegionalGeoMatchSet"
	wafRegionalIPSet                          resourceType = "wafRegionalIPSet"
	wafRegionalRateBasedRule                  resourceType = "wafRegionalRateBasedRule"
	wafregionalRegexMatchSet                  resourceType = "wafregionalRegexMatchSet"
	wafRegionalRegexPatternSet                resourceType = "wafRegionalRegexPatternSet"
	wafRegionalRule                           resourceType = "wafRegionalRule"
	wafregionalRuleGroup                      resourceType = "wafregionalRuleGroup"
	wafRegionalSizeConstraintSet              resourceType = "wafRegionalSizeConstraintSet"
	wafRegionalSQLInjectionMatchSet           resourceType = "wafRegionalSQLInjectionMatchSet"
	wafRegionalWebACL                         resourceType = "wafRegionalWebACL"
	wafRegionalXSSMatchSet                    resourceType = "wafRegionalXSSMatchSet"
	workLinkFleet                             resourceType = "workLinkFleet"
	workSpacesDirectory                       resourceType = "workSpacesDirectory"
	workspacesIPGroup                         resourceType = "workspacesIPGroup"
	workSpacesWorkspace                       resourceType = "workSpacesWorkspace"
	xraySamplingRule                          resourceType = "xraySamplingRule"
)

var cloudformationTypeMap = map[string]resourceType{
	"AWS::AccessAnalyzer::Analyzer":                   accessAnalyzerAnalyzer,
	"AWS::ACMPCA::CertificateAuthority":               acmpcaCertificateAuthority,
	"Alexa::ASK::Skill":                               alexaAskSkill,
	"AWS::AmazonMQ::Broker":                           amazonMQBroker,
	"AWS::AmazonMQ::Configuration":                    amazonMQConfiguration,
	"AWS::Amplify::App":                               amplifyApp,
	"AWS::Amplify::Branch":                            amplifyBranch,
	"AWS::Amplify::Domain":                            amplifyDomain,
	"AWS::ApiGateway::ApiKey":                         apiGatewayAPIKey,
	"AWS::ApiGateway::ClientCertificate":              apiGatewayClientCertificate,
	"AWS::ApiGateway::DomainName":                     apiGatewayDomainName,
	"AWS::ApiGateway::RestApi":                        apiGatewayRestAPI,
	"AWS::ApiGateway::UsagePlan":                      apiGatewayUsagePlan,
	"AWS::ApiGateway::VpcLink":                        apiGatewayVpcLink,
	"AWS::ApiGatewayV2::Api":                          apiGatewayV2Api,
	"AWS::ApiGatewayV2::DomainName":                   apiGatewayV2DomainName,
	"AWS::AppConfig::Application":                     appConfigApplication,
	"AWS::AppConfig::DeploymentStrategy":              appConfigDeploymentStrategy,
	"AWS::AppMesh::Mesh":                              appMeshMesh,
	"AWS::AppStream::DirectoryConfig":                 appStreamDirectoryConfig,
	"AWS::AppStream::Fleet":                           appStreamFleet,
	"AWS::AppStream::ImageBuilder":                    appStreamImageBuilder,
	"AWS::AppStream::Stack":                           appStreamStack,
	"AWS::AppSync::GraphQLApi":                        appSyncGraphQLApi,
	"AWS::Athena::NamedQuery":                         athenaNamedQuery,
	"AWS::Athena::WorkGroup":                          athenaWorkGroup,
	"AWS::AutoScalingPlans::ScalingPlan":              autoScalingPlansScalingPlan,
	"AWS::AutoScaling::AutoScalingGroup":              autoScalingAutoScalingGroup,
	"AWS::AutoScaling::LaunchConfiguration":           autoScalingLaunchConfiguration,
	"AWS::AutoScaling::ScalingPolicy":                 autoScalingScalingPolicy,
	"AWS::AutoScaling::ScheduledAction":               autoScalingScheduledAction,
	"AWS::Backup::BackupPlan":                         backupBackupPlan,
	"AWS::Backup::BackupSelection":                    backupBackupSelection,
	"AWS::Backup::BackupVault":                        backupBackupVault,
	"AWS::Batch::ComputeEnvironment":                  batchComputeEnvironment,
	"AWS::Batch::JobDefinition":                       batchJobDefinition,
	"AWS::Batch::JobQueue":                            batchJobQueue,
	"AWS::Budgets::Budget":                            budgetsBudget,
	"AWS::CertificateManager::Certificate":            certificateManagerCertificate,
	"AWS::Cloud9::EnvironmentEC2":                     cloud9EnvironmentEC2,
	"AWS::CloudFront::CloudFrontOriginAccessIdentity": cloudFrontCloudFrontOriginAccessIdentity,
	"AWS::CloudFront::Distribution":                   cloudFrontDistribution,
	"AWS::CloudFront::StreamingDistribution":          cloudFrontStreamingDistribution,
	"AWS::ServiceDiscovery::HttpNamespace":            serviceDiscoveryHTTPNamespace,
	"AWS::ServiceDiscovery::PrivateDnsNamespace":      serviceDiscoveryPrivateDNSNamespace,
	"AWS::ServiceDiscovery::PublicDnsNamespace":       serviceDiscoveryPublicDNSNamespace,
	"AWS::ServiceDiscovery::Service":                  serviceDiscoveryService,
	"AWS::CloudTrail::Trail":                          cloudTrailTrail,
	"AWS::CloudWatch::Alarm":                          cloudWatchAlarm,
	"AWS::CloudWatch::Dashboard":                      cloudWatchDashboard,
	"AWS::CloudWatch::InsightRule":                    cloudWatchInsightRule,
	"AWS::Logs::Destination":                          logsDestination,
	"AWS::Logs::LogGroup":                             logsLogGroup,
	"AWS::Logs::MetricFilter":                         logsMetricFilter,
	"AWS::Logs::SubscriptionFilter":                   logsSubscriptionFilter,
	"AWS::Events::EventBus":                           eventsEventBus,
	"AWS::Events::Rule":                               eventsRule,
	"AWS::CodeBuild::Project":                         codeBuildProject,
	"AWS::CodeBuild::ReportGroup":                     codeBuildReportGroup,
	"AWS::CodeBuild::SourceCredential":                codeBuildSourceCredential,
	"AWS::CodeCommit::Repository":                     codeCommitRepository,
	"AWS::CodeDeploy::Application":                    codeDeployApplication,
	"AWS::CodeDeploy::DeploymentConfig":               codeDeployDeploymentConfig,
	"AWS::CodeDeploy::DeploymentGroup":                codeDeployDeploymentGroup,
	"AWS::CodeGuruProfiler::ProfilingGroup":           codeGuruProfilerProfilingGroup,
	"AWS::CodePipeline::Pipeline":                     codePipelinePipeline,
	"AWS::CodePipeline::Webhook":                      codePipelineWebhook,
	"AWS::Cognito::IdentityPool":                      cognitoIdentityPool,
	"AWS::Cognito::UserPool":                          cognitoUserPool,
	"AWS::Cognito::UserPoolClient":                    cognitoUserPoolClient,
	"AWS::Cognito::UserPoolGroup":                     cognitoUserPoolGroup,
	"AWS::Cognito::UserPoolIdentityProvider":          cognitoUserPoolIdentityProvider,
	"AWS::Cognito::UserPoolResourceServer":            cognitoUserPoolResourceServer,
	"AWS::Cognito::UserPoolUser":                      cognitoUserPoolUser,
	"AWS::Config::AggregationAuthorization":           configAggregationAuthorization,
	"AWS::Config::ConfigRule":                         configConfigRule,
	"AWS::Config::ConfigurationAggregator":            configConfigurationAggregator,
	"AWS::Config::ConfigurationRecorder":              configConfigurationRecorder,
	"AWS::Config::ConformancePack":                    configConformancePack,
	"AWS::Config::DeliveryChannel":                    configDeliveryChannel,
	"AWS::Config::OrganizationConfigRule":             configOrganizationConfigRule,
	"AWS::Config::OrganizationConformancePack":        configOrganizationConformancePack,
	"AWS::Config::RemediationConfiguration":           configRemediationConfiguration,
	"AWS::DataPipeline::Pipeline":                     dataPipelinePipeline,
	"AWS::DAX::Cluster":                               daxCluster,
	"AWS::DAX::ParameterGroup":                        daxParameterGroup,
	"AWS::DAX::SubnetGroup":                           daxSubnetGroup,
	"AWS::DirectoryService::MicrosoftAD":              directoryServiceDirectory,
	"AWS::DirectoryService::SimpleAD":                 directoryServiceDirectory,
	"AWS::DLM::LifecyclePolicy":                       dlmLifecyclePolicy,
	"AWS::DMS::Certificate":                           dmsCertificate,
	"AWS::DMS::Endpoint":                              dmsEndpoint,
	"AWS::DMS::EventSubscription":                     dmsEventSubscription,
	"AWS::DMS::ReplicationInstance":                   dmsReplicationInstance,
	"AWS::DMS::ReplicationSubnetGroup":                dmsReplicationSubnetGroup,
	"AWS::DMS::ReplicationTask":                       dmsReplicationTask,
	"AWS::DocDB::DBCluster":                           docDBDBCluster,
	"AWS::DocDB::DBClusterParameterGroup":             docDBDBClusterParameterGroup,
	"AWS::DocDB::DBInstance":                          docDBDBInstance,
	"AWS::DocDB::DBSubnetGroup":                       docDBDBSubnetGroup,
	"AWS::DynamoDB::Table":                            dynamoDBTable,
	"AWS::EC2::CapacityReservation":                   ec2CapacityReservation,
	"AWS::EC2::ClientVpnEndpoint":                     ec2ClientVpnEndpoint,
	"AWS::EC2::CustomerGateway":                       ec2CustomerGateway,
	"AWS::EC2::DHCPOptions":                           ec2DHCPOptions,
	"AWS::EC2::EC2Fleet":                              ec2EC2Fleet,
	"AWS::EC2::EgressOnlyInternetGateway":             ec2EgressOnlyInternetGateway,
	"AWS::EC2::EIP":                                   ec2EIP,
	"AWS::EC2::EIPAssociation":                        ec2EIPAssociation,
	"AWS::EC2::FlowLog":                               ec2FlowLog,
	"AWS::EC2::Host":                                  ec2Host,
	"AWS::EC2::Instance":                              ec2Instance,
	"AWS::EC2::InternetGateway":                       ec2InternetGateway,
	"AWS::EC2::LaunchTemplate":                        ec2LaunchTemplate,
	"AWS::EC2::NatGateway":                            ec2NatGateway,
	"AWS::EC2::NetworkAcl":                            ec2NetworkACL,
	"AWS::EC2::NetworkInterface":                      ec2NetworkInterface,
	"AWS::EC2::NetworkInterfaceAttachment":            ec2NetworkInterfaceAttachment,
	"AWS::EC2::NetworkInterfacePermission":            ec2NetworkInterfacePermission,
	"AWS::EC2::PlacementGroup":                        ec2PlacementGroup,
	"AWS::EC2::RouteTable":                            ec2RouteTable,
	"AWS::EC2::SecurityGroup":                         ec2SecurityGroup,
	"AWS::EC2::SpotFleet":                             ec2SpotFleet,
	"AWS::EC2::Subnet":                                ec2Subnet,
	"AWS::EC2::SubnetNetworkAclAssociation":           ec2NetworkACLSubnetAssociation,
	"AWS::EC2::SubnetRouteTableAssociation":           ec2RouteTableSubnetAssociation,
	"AWS::EC2::TrafficMirrorFilter":                   ec2TrafficMirrorFilter,
	"AWS::EC2::TrafficMirrorFilterRule":               ec2TrafficMirrorFilterRule,
	"AWS::EC2::TrafficMirrorSession":                  ec2TrafficMirrorSession,
	"AWS::EC2::TrafficMirrorTarget":                   ec2TrafficMirrorTarget,
	"AWS::EC2::TransitGateway":                        ec2TransitGateway,
	"AWS::EC2::TransitGatewayAttachment":              ec2TransitGatewayAttachment,
	"AWS::EC2::TransitGatewayRouteTable":              ec2TransitGatewayRouteTable,
	"AWS::EC2::Volume":                                ec2Volume,
	"AWS::EC2::VPC":                                   ec2VPC,
	"AWS::EC2::VPCCidrBlock":                          ec2VPCCidrBlock,
	"AWS::EC2::VPCEndpoint":                           ec2VPCEndpoint,
	"AWS::EC2::VPCEndpointConnectionNotification":     ec2VPCEndpointConnectionNotification,
	"AWS::EC2::VPCEndpointService":                    ec2VPCEndpointService,
	"AWS::EC2::VPCPeeringConnection":                  ec2VPCPeeringConnection,
	"AWS::EC2::VPNConnection":                         ec2VPNConnection,
	"AWS::EC2::VPNGateway":                            ec2VPNGateway,
	"AWS::ECR::Repository":                            ecrRepository,
	"AWS::ECS::Cluster":                               ecsCluster,
	"AWS::ECS::Service":                               ecsService,
	"AWS::ECS::TaskDefinition":                        ecsTaskDefinition,
	"AWS::EFS::FileSystem":                            efsFileSystem,
	"AWS::EFS::MountTarget":                           efsMountTarget,
	"AWS::EKS::Cluster":                               eksCluster,
	"AWS::EKS::Nodegroup":                             eksNodegroup,
	"AWS::ElastiCache::CacheCluster":                  elastiCacheCacheCluster,
	"AWS::ElastiCache::ParameterGroup":                elastiCacheParameterGroup,
	"AWS::ElastiCache::ReplicationGroup":              elastiCacheReplicationGroup,
	"AWS::ElastiCache::SecurityGroup":                 elastiCacheSecurityGroup,
	"AWS::ElastiCache::SubnetGroup":                   elastiCacheSubnetGroup,
	"AWS::Elasticsearch::Domain":                      elasticsearchDomain,
	"AWS::ElasticBeanstalk::Application":              elasticBeanstalkApplication,
	"AWS::ElasticBeanstalk::ApplicationVersion":       elasticBeanstalkApplicationVersion,
	"AWS::ElasticBeanstalk::ConfigurationTemplate":    elasticBeanstalkConfigurationTemplate,
	"AWS::ElasticBeanstalk::Environment":              elasticBeanstalkEnvironment,
	"AWS::ElasticLoadBalancing::LoadBalancer":         elasticLoadBalancingLoadBalancer,
	"AWS::ElasticLoadBalancingV2::Listener":           elasticLoadBalancingV2Listener,
	"AWS::ElasticLoadBalancingV2::ListenerRule":       elasticLoadBalancingV2ListenerRule,
	"AWS::ElasticLoadBalancingV2::LoadBalancer":       elasticLoadBalancingV2LoadBalancer,
	"AWS::ElasticLoadBalancingV2::TargetGroup":        elasticLoadBalancingV2TargetGroup,
	"AWS::EMR::Cluster":                               emrCluster,
	"AWS::EMR::SecurityConfiguration":                 emrSecurityConfiguration,
	"AWS::EventSchemas::Discoverer":                   eventSchemasDiscoverer,
	"AWS::EventSchemas::Registry":                     eventSchemasRegistry,
	"AWS::FSx::FileSystem":                            fsxFileSystem,
	"AWS::GameLift::Alias":                            gameLiftAlias,
	"AWS::GameLift::Build":                            gameLiftBuild,
	"AWS::GameLift::Fleet":                            gameLiftFleet,
	"AWS::GameLift::GameSessionQueue":                 gameLiftGameSessionQueue,
	"AWS::GameLift::MatchmakingConfiguration":         gameLiftMatchmakingConfiguration,
	"AWS::GameLift::MatchmakingRuleSet":               gameLiftMatchmakingRuleSet,
	"AWS::GameLift::Script":                           gameLiftScript,
	"AWS::Glue::Connection":                           glueConnection,
	"AWS::Glue::Crawler":                              glueCrawler,
	"AWS::Glue::Database":                             glueDatabase,
	"AWS::Glue::DevEndpoint":                          glueDevEndpoint,
	"AWS::Glue::Job":                                  glueJob,
	"AWS::Glue::MLTransform":                          glueMLTransform,
	"AWS::Glue::SecurityConfiguration":                glueSecurityConfiguration,
	"AWS::Glue::Table":                                glueTable,
	"AWS::Glue::Trigger":                              glueTrigger,
	"AWS::Glue::Workflow":                             glueWorkflow,
	"AWS::GroundStation::Config":                      groundStationConfig,
	"AWS::GroundStation::DataflowEndpointGroup":       groundStationDataflowEndpointGroup,
	"AWS::GroundStation::MissionProfile":              groundStationMissionProfile,
	"AWS::GuardDuty::Detector":                        guardDutyDetector,
	"AWS::IAM::AccessKey":                             iamAccessKey,
	"AWS::IAM::Group":                                 iamGroup,
	"AWS::IAM::InstanceProfile":                       iamInstanceProfile,
	"AWS::IAM::ManagedPolicy":                         iamPolicy,
	"AWS::IAM::Role":                                  iamRole,
	"AWS::IAM::Policy":                                iamRolePolicy,
	"AWS::IAM::ServiceLinkedRole":                     iamServiceLinkedRole,
	"AWS::IAM::User":                                  iamUser,
	"AWS::Inspector::AssessmentTarget":                inspectorAssessmentTarget,
	"AWS::Inspector::AssessmentTemplate":              inspectorAssessmentTemplate,
	"AWS::IoT::Certificate":                           ioTCertificate,
	"AWS::IoT::Policy":                                ioTPolicy,
	"AWS::IoT::Thing":                                 ioTThing,
	"AWS::IoT::TopicRule":                             ioTTopicRule,
	"AWS::IoT1Click::Device":                          ioT1ClickDevice,
	"AWS::IoT1Click::Project":                         ioT1ClickProject,
	"AWS::IoTAnalytics::Channel":                      ioTAnalyticsChannel,
	"AWS::IoTAnalytics::Dataset":                      ioTAnalyticsDataset,
	"AWS::IoTAnalytics::Datastore":                    ioTAnalyticsDatastore,
	"AWS::IoTAnalytics::Pipeline":                     ioTAnalyticsPipeline,
	"AWS::IoTEvents::DetectorModel":                   ioTEventsDetectorModel,
	"AWS::IoTEvents::Input":                           ioTEventsInput,
	"AWS::Greengrass::ConnectorDefinition":            greengrassConnectorDefinition,
	"AWS::Greengrass::ConnectorDefinitionVersion":     greengrassConnectorDefinitionVersion,
	"AWS::Greengrass::CoreDefinition":                 greengrassCoreDefinition,
	"AWS::Greengrass::CoreDefinitionVersion":          greengrassCoreDefinitionVersion,
	"AWS::Greengrass::DeviceDefinition":               greengrassDeviceDefinition,
	"AWS::Greengrass::DeviceDefinitionVersion":        greengrassDeviceDefinitionVersion,
	"AWS::Greengrass::FunctionDefinition":             greengrassFunctionDefinition,
	"AWS::Greengrass::FunctionDefinitionVersion":      greengrassFunctionDefinitionVersion,
	"AWS::Greengrass::Group":                          greengrassGroup,
	"AWS::Greengrass::GroupVersion":                   greengrassGroupVersion,
	"AWS::Greengrass::LoggerDefinition":               greengrassLoggerDefinition,
	"AWS::Greengrass::LoggerDefinitionVersion":        greengrassLoggerDefinitionVersion,
	"AWS::Greengrass::ResourceDefinition":             greengrassResourceDefinition,
	"AWS::Greengrass::ResourceDefinitionVersion":      greengrassResourceDefinitionVersion,
	"AWS::Greengrass::SubscriptionDefinition":         greengrassSubscriptionDefinition,
	"AWS::Greengrass::SubscriptionDefinitionVersion":  greengrassSubscriptionDefinitionVersion,
	"AWS::IoTThingsGraph::FlowTemplate":               ioTThingsGraphFlowTemplate,
	"AWS::Kinesis::Stream":                            kinesisStream,
	"AWS::Kinesis::StreamConsumer":                    kinesisStreamConsumer,
	"AWS::KinesisAnalytics::Application":              kinesisAnalyticsApplication,
	"AWS::KinesisAnalyticsV2::Application":            kinesisAnalyticsV2Application,
	"AWS::KinesisFirehose::DeliveryStream":            kinesisFirehoseDeliveryStream,
	"AWS::KMS::Alias":                                 kmsAlias,
	"AWS::KMS::Key":                                   kmsKey,
	"AWS::LakeFormation::Resource":                    lakeFormationResource,
	"AWS::Lambda::Alias":                              lambdaAlias,
	"AWS::Lambda::Function":                           lambdaFunction,
	"AWS::Lambda::LayerVersion":                       lambdaLayerVersion,
	"AWS::MediaConvert::JobTemplate":                  mediaConvertJobTemplate,
	"AWS::MediaConvert::Preset":                       mediaConvertPreset,
	"AWS::MediaConvert::Queue":                        mediaConvertQueue,
	"AWS::MediaLive::Channel":                         mediaLiveChannel,
	"AWS::MediaLive::Input":                           mediaLiveInput,
	"AWS::MediaLive::InputSecurityGroup":              mediaLiveInputSecurityGroup,
	"AWS::MediaStore::Container":                      mediaStoreContainer,
	"AWS::MSK::Cluster":                               mskCluster,
	"AWS::Neptune::DBCluster":                         neptuneDBCluster,
	"AWS::Neptune::DBClusterParameterGroup":           neptuneDBClusterParameterGroup,
	"AWS::Neptune::DBInstance":                        neptuneDBInstance,
	"AWS::Neptune::DBParameterGroup":                  neptuneDBParameterGroup,
	"AWS::Neptune::DBSubnetGroup":                     neptuneDBSubnetGroup,
	"AWS::NetworkManager::Device":                     networkManagerDevice,
	"AWS::NetworkManager::GlobalNetwork":              networkManagerGlobalNetwork,
	"AWS::NetworkManager::Link":                       networkManagerLink,
	"AWS::NetworkManager::Site":                       networkManagerSite,
	"AWS::OpsWorks::App":                              opsWorksApp,
	"AWS::OpsWorks::Instance":                         opsWorksInstance,
	"AWS::OpsWorks::Layer":                            opsWorksLayer,
	"AWS::OpsWorks::Stack":                            opsWorksStack,
	"AWS::OpsWorks::UserProfile":                      opsWorksUserProfile,
	"AWS::OpsWorks::Volume":                           opsWorksVolume,
	"AWS::Pinpoint::App":                              pinpointApp,
	"AWS::Pinpoint::EmailTemplate":                    pinpointEmailTemplate,
	"AWS::Pinpoint::PushTemplate":                     pinpointPushTemplate,
	"AWS::Pinpoint::SmsTemplate":                      pinpointSmsTemplate,
	"AWS::QLDB::Ledger":                               qLDBLedger,
	"AWS::RDS::DBCluster":                             rdsDBCluster,
	"AWS::RDS::DBClusterParameterGroup":               rdsDBClusterParameterGroup,
	"AWS::RDS::DBInstance":                            rdsDBInstance,
	"AWS::RDS::DBParameterGroup":                      rdsDBParameterGroup,
	"AWS::RDS::DBSecurityGroup":                       rdsDBSecurityGroup,
	"AWS::RDS::DBSubnetGroup":                         rdsDBSubnetGroup,
	"AWS::RDS::EventSubscription":                     rdsEventSubscription,
	"AWS::RDS::OptionGroup":                           rdsOptionGroup,
	"AWS::Redshift::Cluster":                          redshiftCluster,
	"AWS::Redshift::ClusterParameterGroup":            redshiftClusterParameterGroup,
	"AWS::Redshift::ClusterSecurityGroup":             redshiftClusterSecurityGroup,
	"AWS::Redshift::ClusterSubnetGroup":               redshiftClusterSubnetGroup,
	"AWS::RoboMaker::Fleet":                           roboMakerFleet,
	"AWS::RoboMaker::Robot":                           roboMakerRobot,
	"AWS::RoboMaker::RobotApplication":                roboMakerRobotApplication,
	"AWS::RoboMaker::SimulationApplication":           roboMakerSimulationApplication,
	"AWS::ResourceGroups::Group":                      resourceGroupsGroup,
	"AWS::Route53::HealthCheck":                       route53HealthCheck,
	"AWS::Route53::HostedZone":                        route53HostedZone,
	"AWS::Route53::RecordSet":                         route53RecordSet,
	"AWS::Route53Resolver::ResolverEndpoint":          route53ResolverResolverEndpoint,
	"AWS::Route53Resolver::ResolverRule":              route53ResolverResolverRule,
	"AWS::Route53Resolver::ResolverRuleAssociation":   route53ResolverResolverRuleAssociation,
	"AWS::S3::AccessPoint":                            s3AccessPoint,
	"AWS::S3::Bucket":                                 s3Bucket,
	"AWS::SageMaker::CodeRepository":                  sageMakerCodeRepository,
	"AWS::SageMaker::Endpoint":                        sageMakerEndpoint,
	"AWS::SageMaker::EndpointConfig":                  sageMakerEndpointConfig,
	"AWS::SageMaker::Model":                           sageMakerModel,
	"AWS::SageMaker::NotebookInstance":                sageMakerNotebookInstance,
	"AWS::SageMaker::NotebookInstanceLifecycleConfig": sageMakerNotebookInstanceLifecycleConfig,
	"AWS::SageMaker::Workteam":                        sageMakerWorkteam,
	"AWS::SecretsManager::Secret":                     secretsManagerSecret,
	"AWS::SES::ConfigurationSet":                      sesConfigurationSet,
	"AWS::SES::ReceiptFilter":                         sesReceiptFilter,
	"AWS::SES::ReceiptRuleSet":                        sesReceiptRuleSet,
	"AWS::SES::Template":                              sesTemplate,
	"AWS::SDB::Domain":                                sdbDomain,
	"AWS::SNS::Subscription":                          snsSubscription,
	"AWS::SNS::Topic":                                 snsTopic,
	"AWS::SQS::Queue":                                 sqsQueue,
	"AWS::StepFunctions::Activity":                    stepFunctionsActivity,
	"AWS::StepFunctions::StateMachine":                stepFunctionsStateMachine,
	"AWS::SSM::Association":                           ssmAssociation,
	"AWS::SSM::Document":                              ssmDocument,
	"AWS::SSM::MaintenanceWindow":                     ssmMaintenanceWindow,
	"AWS::SSM::MaintenanceWindowTarget":               ssmMaintenanceWindowTarget,
	"AWS::SSM::MaintenanceWindowTask":                 ssmMaintenanceWindowTask,
	"AWS::SSM::Parameter":                             ssmParameter,
	"AWS::SSM::PatchBaseline":                         ssmPatchBaseline,
	"AWS::SSM::ResourceDataSync":                      ssmResourceDataSync,
	"AWS::Transfer::Server":                           transferServer,
	"AWS::Transfer::User":                             transferUser,
	"AWS::WAF::ByteMatchSet":                          wafByteMatchSet,
	"AWS::WAF::IPSet":                                 wafIPSet,
	"AWS::WAF::Rule":                                  wafRule,
	"AWS::WAF::SizeConstraintSet":                     wafSizeConstraintSet,
	"AWS::WAF::SqlInjectionMatchSet":                  wafSQLInjectionMatchSet,
	"AWS::WAF::WebACL":                                wafWebACL,
	"AWS::WAF::XssMatchSet":                           wafXSSMatchSet,
	"AWS::WAFv2::IPSet":                               wafv2IPSet,
	"AWS::WAFv2::RegexPatternSet":                     wafv2RegexPatternSet,
	"AWS::WAFv2::RuleGroup":                           wafv2RuleGroup,
	"AWS::WAFv2::WebACL":                              wafv2WebACL,
	"AWS::WAFRegional::ByteMatchSet":                  wafRegionalByteMatchSet,
	"AWS::WAFRegional::GeoMatchSet":                   wafRegionalGeoMatchSet,
	"AWS::WAFRegional::IPSet":                         wafRegionalIPSet,
	"AWS::WAFRegional::RateBasedRule":                 wafRegionalRateBasedRule,
	"AWS::WAFRegional::RegexPatternSet":               wafRegionalRegexPatternSet,
	"AWS::WAFRegional::Rule":                          wafRegionalRule,
	"AWS::WAFRegional::SizeConstraintSet":             wafRegionalSizeConstraintSet,
	"AWS::WAFRegional::SqlInjectionMatchSet":          wafRegionalSQLInjectionMatchSet,
	"AWS::WAFRegional::WebACL":                        wafRegionalWebACL,
	"AWS::WAFRegional::XssMatchSet":                   wafRegionalXSSMatchSet,
	"AWS::WorkSpaces::Workspace":                      workSpacesWorkspace,
}

func fromCloudFormationType(cloudFormationType string) (resourceType, bool) {
	value, ok := cloudformationTypeMap[cloudFormationType]
	return value, ok
}

var terraformTypeMap = map[string]resourceType{
	"aws_accessanalyzer_analyzer":                             accessAnalyzerAnalyzer,
	"aws_acmpca_certificate_authority":                        acmpcaCertificateAuthority,
	"aws_mq_broker":                                           amazonMQBroker,
	"aws_mq_configuration":                                    amazonMQConfiguration,
	"aws_api_gateway_api_key":                                 apiGatewayAPIKey,
	"aws_api_gateway_client_certificate":                      apiGatewayClientCertificate,
	"aws_api_gateway_domain_name":                             apiGatewayDomainName,
	"aws_api_gateway_rest_api":                                apiGatewayRestAPI,
	"aws_api_gateway_usage_plan":                              apiGatewayUsagePlan,
	"aws_api_gateway_vpc_link":                                apiGatewayVpcLink,
	"aws_apigatewayv2_api":                                    apiGatewayV2Api,
	"aws_appautoscaling_scheduled_action":                     applicationAutoScalingScheduledAction,
	"aws_appmesh_mesh":                                        appMeshMesh,
	"aws_appsync_function":                                    appSyncFunctions,
	"aws_appsync_graphql_api":                                 appSyncGraphQLApi,
	"aws_athena_named_query":                                  athenaNamedQuery,
	"aws_athena_workgroup":                                    athenaWorkGroup,
	"aws_autoscaling_group":                                   autoScalingAutoScalingGroup,
	"aws_launch_configuration":                                autoScalingLaunchConfiguration,
	"aws_autoscaling_policy":                                  autoScalingScalingPolicy,
	"aws_autoscaling_schedule":                                autoScalingScheduledAction,
	"aws_backup_plan":                                         backupBackupPlan,
	"aws_backup_selection":                                    backupBackupSelection,
	"aws_backup_vault":                                        backupBackupVault,
	"aws_batch_compute_environment":                           batchComputeEnvironment,
	"aws_batch_job_definition":                                batchJobDefinition,
	"aws_batch_job_queue":                                     batchJobQueue,
	"aws_acm_certificate":                                     certificateManagerCertificate,
	"aws_cloud9_environment_ec2":                              cloud9EnvironmentEC2,
	"aws_cloudfront_origin_access_identity":                   cloudFrontCloudFrontOriginAccessIdentity,
	"aws_cloudfront_distribution":                             cloudFrontDistribution,
	"aws_cloudfront_public_key":                               cloudFrontPublicKey,
	"aws_cloudhsm_v2_cluster":                                 cloudHSMV2Cluster,
	"aws_cloudhsm_v2_hsm":                                     cloudHSMV2HSM,
	"aws_service_discovery_http_namespace":                    serviceDiscoveryHTTPNamespace,
	"aws_service_discovery_private_dns_namespace":             serviceDiscoveryPrivateDNSNamespace,
	"aws_service_discovery_public_dns_namespace":              serviceDiscoveryPublicDNSNamespace,
	"aws_service_discovery_service":                           serviceDiscoveryService,
	"aws_cloudtrail":                                          cloudTrailTrail,
	"aws_cloudwatch_metric_alarm":                             cloudWatchAlarm,
	"aws_cloudwatch_dashboard":                                cloudWatchDashboard,
	"aws_cloudwatch_log_destination":                          logsDestination,
	"aws_cloudwatch_log_group":                                logsLogGroup,
	"aws_cloudwatch_log_metric_filter":                        logsMetricFilter,
	"aws_cloudwatch_log_subscription_filter":                  logsSubscriptionFilter,
	"aws_cloudwatch_log_resource_policy":                      logsResourcePolicy,
	"aws_macie_member_account_association":                    macieMemberAccountAssociation,
	"aws_macie_s3_bucket_association":                         macieS3BucketAssociation,
	"aws_cloudwatch_event_rule":                               eventsRule,
	"aws_cloudwatch_event_target":                             eventsTarget,
	"aws_codebuild_project":                                   codeBuildProject,
	"aws_codebuild_source_credential":                         codeBuildSourceCredential,
	"aws_codecommit_repository":                               codeCommitRepository,
	"aws_codecommit_trigger":                                  codeCommitTrigger,
	"aws_codedeploy_app":                                      codeDeployApplication,
	"aws_codedeploy_deployment_config":                        codeDeployDeploymentConfig,
	"aws_codedeploy_deployment_group":                         codeDeployDeploymentGroup,
	"aws_codepipeline":                                        codePipelinePipeline,
	"aws_codepipeline_webhook":                                codePipelineWebhook,
	"aws_cognito_identity_pool":                               cognitoIdentityPool,
	"aws_cognito_user_pool":                                   cognitoUserPool,
	"aws_cognito_user_pool_client":                            cognitoUserPoolClient,
	"aws_cognito_user_group":                                  cognitoUserPoolGroup,
	"aws_cognito_identity_provider":                           cognitoUserPoolIdentityProvider,
	"aws_cognito_resource_server":                             cognitoUserPoolResourceServer,
	"aws_config_aggregate_authorization":                      configAggregationAuthorization,
	"aws_config_config_rule":                                  configConfigRule,
	"aws_config_configuration_aggregator":                     configConfigurationAggregator,
	"aws_config_configuration_recorder":                       configConfigurationRecorder,
	"aws_config_delivery_channel":                             configDeliveryChannel,
	"aws_config_organization_managed_rule":                    configOrganizationConfigRule,
	"aws_config_organization_custom_rule":                     configOrganizationConfigRule,
	"aws_cur_report_definition":                               costAndUsageReportServiceReportDefinition,
	"aws_datapipeline_pipeline":                               dataPipelinePipeline,
	"aws_datasync_agent":                                      dataSyncAgent,
	"aws_datasync_location_efs":                               dataSyncLocation,
	"aws_datasync_location_nfs":                               dataSyncLocation,
	"aws_datasync_location_s3":                                dataSyncLocation,
	"aws_datasync_location_smb":                               dataSyncLocation,
	"aws_datasync_task":                                       dataSyncTask,
	"aws_dax_cluster":                                         daxCluster,
	"aws_dax_parameter_group":                                 daxParameterGroup,
	"aws_dax_subnet_group":                                    daxSubnetGroup,
	"aws_devicefarm_project":                                  deviceFarmProject,
	"aws_dx_connection":                                       directConnectConnection,
	"aws_dx_gateway":                                          directConnectGateway,
	"aws_dx_gateway_association":                              directConnectGatewayAssociation,
	"aws_dx_gateway_association_proposal":                     directConnectGatewayAssociationProposal,
	"aws_dx_lag":                                              directConnectLAG,
	"aws_dx_hosted_private_virtual_interface":                 directConnectVirtualInterface,
	"aws_dx_hosted_public_virtual_interface":                  directConnectVirtualInterface,
	"aws_dx_hosted_transit_virtual_interface":                 directConnectVirtualInterface,
	"aws_dx_private_virtual_interface":                        directConnectVirtualInterface,
	"aws_dx_public_virtual_interface":                         directConnectVirtualInterface,
	"aws_dx_transit_virtual_interface":                        directConnectVirtualInterface,
	"aws_directory_service_directory":                         directoryServiceDirectory,
	"aws_dlm_lifecycle_policy":                                dlmLifecyclePolicy,
	"aws_dms_certificate":                                     dmsCertificate,
	"aws_dms_endpoint":                                        dmsEndpoint,
	"aws_dms_replication_instance":                            dmsReplicationInstance,
	"aws_dms_replication_subnet_group":                        dmsReplicationSubnetGroup,
	"aws_dms_replication_task":                                dmsReplicationTask,
	"aws_docdb_cluster":                                       docDBDBCluster,
	"aws_docdb_cluster_parameter_group":                       docDBDBClusterParameterGroup,
	"aws_docdb_cluster_instance":                              docDBDBInstance,
	"aws_docdb_subnet_group":                                  docDBDBSubnetGroup,
	"aws_docdb_cluster_snapshot":                              docDBDBClusterSnapshot,
	"aws_dynamodb_table":                                      dynamoDBTable,
	"aws_dynamodb_global_table":                               dynamoDBGlobalTable,
	"aws_ec2_capacity_reservation":                            ec2CapacityReservation,
	"aws_ec2_client_vpn_endpoint":                             ec2ClientVpnEndpoint,
	"aws_customer_gateway":                                    ec2CustomerGateway,
	"aws_vpc_dhcp_options":                                    ec2DHCPOptions,
	"aws_ec2_fleet":                                           ec2EC2Fleet,
	"aws_egress_only_internet_gateway":                        ec2EgressOnlyInternetGateway,
	"aws_eip":                                                 ec2EIP,
	"aws_eip_association":                                     ec2EIPAssociation,
	"aws_flow_log":                                            ec2FlowLog,
	"aws_ami":                                                 ec2Image,
	"aws_ami_copy":                                            ec2Image,
	"aws_ami_from_instance":                                   ec2Image,
	"aws_instance":                                            ec2Instance,
	"aws_internet_gateway":                                    ec2InternetGateway,
	"aws_launch_template":                                     ec2LaunchTemplate,
	"aws_nat_gateway":                                         ec2NatGateway,
	"aws_network_acl":                                         ec2NetworkACL,
	"aws_network_interface":                                   ec2NetworkInterface,
	"aws_network_interface_attachment":                        ec2NetworkInterfaceAttachment,
	"aws_placement_group":                                     ec2PlacementGroup,
	"aws_route_table":                                         ec2RouteTable,
	"aws_security_group":                                      ec2SecurityGroup,
	"aws_ebs_snapshot":                                        ec2Snapshot,
	"aws_ebs_snapshot_copy":                                   ec2Snapshot,
	"aws_AWS::EC2::SpotFleet":                                 ec2SpotFleet,
	"aws_subnet":                                              ec2Subnet,
	"aws_route_table_association":                             ec2RouteTableSubnetAssociation,
	"aws_ec2_traffic_mirror_filter":                           ec2TrafficMirrorFilter,
	"aws_ec2_traffic_mirror_filter_rule":                      ec2TrafficMirrorFilterRule,
	"aws_ec2_traffic_mirror_session":                          ec2TrafficMirrorSession,
	"aws_ec2_traffic_mirror_target":                           ec2TrafficMirrorTarget,
	"aws_ec2_transit_gateway":                                 ec2TransitGateway,
	"aws_ec2_transit_gateway_vpc_attachment":                  ec2TransitGatewayAttachment,
	"aws_ec2_transit_gateway_route_table":                     ec2TransitGatewayRouteTable,
	"aws_ebs_volume":                                          ec2Volume,
	"aws_vpc":                                                 ec2VPC,
	"aws_vpc_ipv4_cidr_block_association":                     ec2VPCCidrBlock,
	"aws_vpc_endpoint":                                        ec2VPCEndpoint,
	"aws_vpc_endpoint_connection_notification":                ec2VPCEndpointConnectionNotification,
	"aws_vpc_endpoint_service":                                ec2VPCEndpointService,
	"aws_vpc_peering_connection":                              ec2VPCPeeringConnection,
	"aws_vpn_connection":                                      ec2VPNConnection,
	"aws_vpn_gateway":                                         ec2VPNGateway,
	"aws_key_pair":                                            ec2KeyPair,
	"aws_spot_instance_request":                               ec2SpotInstanceRequest,
	"aws_ecr_repository":                                      ecrRepository,
	"aws_ecs_cluster":                                         ecsCluster,
	"aws_ecs_service":                                         ecsService,
	"aws_ecs_task_definition":                                 ecsTaskDefinition,
	"aws_ecs_capacity_provider":                               ecsCapacityProvider,
	"aws_efs_file_system":                                     efsFileSystem,
	"aws_efs_mount_target":                                    efsMountTarget,
	"aws_eks_cluster":                                         eksCluster,
	"aws_eks_node_group":                                      eksNodegroup,
	"aws_eks_fargate_profile":                                 eksFargateProfile,
	"aws_elasticache_cluster":                                 elastiCacheCacheCluster,
	"aws_elasticache_parameter_group":                         elastiCacheParameterGroup,
	"aws_elasticache_replication_group":                       elastiCacheReplicationGroup,
	"aws_elasticache_security_group":                          elastiCacheSecurityGroup,
	"aws_elasticache_subnet_group":                            elastiCacheSubnetGroup,
	"aws_elasticsearch_domain":                                elasticsearchDomain,
	"aws_elastictranscoder_pipeline":                          elasticTranscoderPipeline,
	"aws_elastictranscoder_preset":                            elasticTranscoderPreset,
	"aws_elastic_beanstalk_application":                       elasticBeanstalkApplication,
	"aws_elastic_beanstalk_application_version":               elasticBeanstalkApplicationVersion,
	"aws_elastic_beanstalk_configuration_template":            elasticBeanstalkConfigurationTemplate,
	"aws_elastic_beanstalk_environment":                       elasticBeanstalkEnvironment,
	"aws_elb":                                                 elasticLoadBalancingLoadBalancer,
	"aws_lb_listener":                                         elasticLoadBalancingV2Listener,
	"aws_lb_listener_rule":                                    elasticLoadBalancingV2ListenerRule,
	"aws_lb":                                                  elasticLoadBalancingV2LoadBalancer,
	"aws_lb_target_group":                                     elasticLoadBalancingV2TargetGroup,
	"aws_emr_cluster":                                         emrCluster,
	"aws_emr_security_configuration":                          emrSecurityConfiguration,
	"aws_emr_instance_group":                                  emrInstanceGroup,
	"aws_fsx_lustre_file_system":                              fsxFileSystem,
	"aws_fsx_windows_file_system":                             fsxFileSystem,
	"aws_gamelift_alias":                                      gameLiftAlias,
	"aws_gamelift_build":                                      gameLiftBuild,
	"aws_gamelift_fleet":                                      gameLiftFleet,
	"aws_gamelift_game_session_queue":                         gameLiftGameSessionQueue,
	"aws_glue_connection":                                     glueConnection,
	"aws_glue_crawler":                                        glueCrawler,
	"aws_glue_catalog_database":                               glueDatabase,
	"aws_glue_job":                                            glueJob,
	"aws_glue_security_configuration":                         glueSecurityConfiguration,
	"aws_glue_catalog_table":                                  glueTable,
	"aws_glue_trigger":                                        glueTrigger,
	"aws_glue_workflow":                                       glueWorkflow,
	"aws_glacier_vault":                                       glacierVault,
	"aws_globalaccelerator_accelerator":                       globalAcceleratorAccelerator,
	"aws_globalaccelerator_endpoint_group":                    globalAcceleratorEndpointGroup,
	"aws_globalaccelerator_listener":                          globalAcceleratorListener,
	"aws_guardduty_detector":                                  guardDutyDetector,
	"aws_iam_access_key":                                      iamAccessKey,
	"aws_iam_account_alias":                                   iamAccountAlias,
	"aws_iam_group":                                           iamGroup,
	"aws_iam_group_policy":                                    iamGroupPolicy,
	"aws_iam_instance_profile":                                iamInstanceProfile,
	"aws_iam_openid_connect_provider":                         iamOpenidConnectProvider,
	"aws_iam_policy":                                          iamPolicy,
	"aws_iam_role":                                            iamRole,
	"aws_iam_role_policy":                                     iamRolePolicy,
	"aws_iam_saml_provider":                                   iamSamlProvider,
	"aws_iam_server_certificate":                              iamServerCertificate,
	"aws_iam_service_linked_role":                             iamServiceLinkedRole,
	"aws_iam_user":                                            iamUser,
	"aws_iam_user_policy":                                     iamUserPolicy,
	"aws_iam_user_ssh_key":                                    iamUserSSHKey,
	"aws_inspector_assessment_target":                         inspectorAssessmentTarget,
	"aws_inspector_assessment_template":                       inspectorAssessmentTemplate,
	"aws_iot_certificate":                                     ioTCertificate,
	"aws_iot_policy":                                          ioTPolicy,
	"aws_iot_role_alias":                                      iotRoleAlias,
	"aws_iot_thing":                                           ioTThing,
	"aws_iot_thing_type":                                      iotThingType,
	"aws_iot_topic_rule":                                      ioTTopicRule,
	"aws_kinesis_stream":                                      kinesisStream,
	"aws_kinesis_analytics_application":                       kinesisAnalyticsApplication,
	"aws_kinesis_video_stream":                                kinesisVideoStream,
	"aws_kinesis_firehose_delivery_stream":                    kinesisFirehoseDeliveryStream,
	"aws_kms_alias":                                           kmsAlias,
	"aws_kms_external_key":                                    kmsKey,
	"aws_kms_grant":                                           kmsGrant,
	"aws_kms_key":                                             kmsKey,
	"aws_lambda_alias":                                        lambdaAlias,
	"aws_lambda_event_source_mapping":                         lambdaEventSourceMapping,
	"aws_lambda_function":                                     lambdaFunction,
	"aws_lambda_layer_version":                                lambdaLayerVersion,
	"aws_licensemanager_license_configuration":                licenseManagerLicenseConfiguration,
	"aws_lightsail_domain":                                    lightsailDomain,
	"aws_lightsail_instance":                                  lightsailInstance,
	"aws_lightsail_key_pair":                                  lightsailKeyPair,
	"aws_lightsail_static_ip":                                 lightsailStaticIP,
	"aws_media_convert_queue":                                 mediaConvertQueue,
	"aws_media_package_channel":                               mediaLiveChannel,
	"aws_media_store_container":                               mediaStoreContainer,
	"aws_msk_cluster":                                         mskCluster,
	"aws_msk_configuration":                                   mskConfiguration,
	"aws_pinpoint_app":                                        pinpointApp,
	"aws_organizations_account":                               organizationsAccount,
	"aws_organizations_organization":                          organizationsOrganization,
	"aws_organizations_organizational_unit":                   organizationsOrganizationalUnit,
	"aws_organizations_policy":                                organizationsPolicy,
	"aws_neptune_cluster":                                     neptuneDBCluster,
	"aws_neptune_cluster_parameter_group":                     neptuneDBClusterParameterGroup,
	"aws_neptune_cluster_snapshot":                            neptuneDBClusterSnapshot,
	"aws_neptune_event_subscription":                          neptuneDBEventSubscription,
	"aws_neptune_cluster_instance":                            neptuneDBInstance,
	"aws_neptune_parameter_group":                             neptuneDBParameterGroup,
	"aws_neptune_subnet_group":                                neptuneDBSubnetGroup,
	"aws_opsworks_application":                                opsWorksApp,
	"aws_opsworks_instance":                                   opsWorksInstance,
	"aws_opsworks_custom_layer":                               opsWorksLayer,
	"aws_opsworks_ganglia_layer":                              opsWorksLayer,
	"aws_opsworks_haproxy_layer":                              opsWorksLayer,
	"aws_opsworks_java_app_layer":                             opsWorksLayer,
	"aws_opsworks_memcached_layer":                            opsWorksLayer,
	"aws_opsworks_mysql_layer":                                opsWorksLayer,
	"aws_opsworks_nodejs_app_layer":                           opsWorksLayer,
	"aws_opsworks_php_app_layer":                              opsWorksLayer,
	"aws_opsworks_rails_app_layer":                            opsWorksLayer,
	"aws_opsworks_static_web_layer":                           opsWorksLayer,
	"aws_opsworks_rds_db_instance":                            opsWorkRdsDbInstance,
	"aws_opsworks_stack":                                      opsWorksStack,
	"aws_opsworks_user_profile":                               opsWorksUserProfile,
	"aws_qldb_ledger":                                         qLDBLedger,
	"aws_quicksight_group":                                    quickSightGroup,
	"aws_quicksight_user":                                     quickSightUser,
	"aws_rds_cluster":                                         rdsDBCluster,
	"aws_rds_cluster_endpoint":                                rdsDBClusterEndpoint,
	"aws_rds_cluster_parameter_group":                         rdsDBClusterParameterGroup,
	"aws_db_cluster_snapshot":                                 rdsDBClusterSnapshot,
	"aws_db_instance":                                         rdsDBInstance,
	"aws_rds_cluster_instance":                                rdsDBInstance,
	"aws_db_parameter_group":                                  rdsDBParameterGroup,
	"aws_db_security_group":                                   rdsDBSecurityGroup,
	"aws_db_snapshot":                                         rdsDBSnapshot,
	"aws_db_subnet_group":                                     rdsDBSubnetGroup,
	"aws_db_event_subscription":                               rdsEventSubscription,
	"aws_rds_global_cluster":                                  rdsGlobalCluster,
	"aws_db_option_group":                                     rdsOptionGroup,
	"aws_redshift_cluster":                                    redshiftCluster,
	"aws_redshift_parameter_group":                            redshiftClusterParameterGroup,
	"aws_redshift_security_group":                             redshiftClusterSecurityGroup,
	"aws_redshift_subnet_group":                               redshiftClusterSubnetGroup,
	"aws_redshift_event_subscription":                         redshiftEventSubscription,
	"aws_redshift_snapshot_copy_grant":                        redshiftSnapshotCopyGrant,
	"aws_redshift_snapshot_schedule":                          redshiftSnapshotSchedule,
	"aws_resourcegroups_group":                                resourceGroupsGroup,
	"aws_route53_delegation_set":                              route53DelegationSet,
	"aws_route53_health_check":                                route53HealthCheck,
	"aws_route53_zone":                                        route53HostedZone,
	"aws_route53_record":                                      route53RecordSet,
	"aws_route53_query_log":                                   route53QueryLog,
	"aws_route53_resolver_endpoint":                           route53ResolverResolverEndpoint,
	"aws_route53_resolver_rule":                               route53ResolverResolverRule,
	"aws_route53_resolver_rule_association":                   route53ResolverResolverRuleAssociation,
	"aws_s3_access_point":                                     s3AccessPoint,
	"aws_s3_bucket":                                           s3Bucket,
	"aws_sagemaker_endpoint":                                  sageMakerEndpoint,
	"aws_sagemaker_endpoint_configuration":                    sageMakerEndpointConfig,
	"aws_sagemaker_model":                                     sageMakerModel,
	"aws_sagemaker_notebook_instance":                         sageMakerNotebookInstance,
	"aws_sagemaker_notebook_instance_lifecycle_configuration": sageMakerNotebookInstanceLifecycleConfig,
	"aws_secretsmanager_secret":                               secretsManagerSecret,
	"aws_secretsmanager_secret_version":                       secretsManagerSecretVersion,
	"aws_ses_configuration_set":                               sesConfigurationSet,
	"aws_ses_domain_identity":                                 sesDomainIdentity,
	"aws_ses_email_identity":                                  sesEmailIdentity,
	"aws_ses_receipt_filter":                                  sesReceiptFilter,
	"aws_ses_receipt_rule_set":                                sesReceiptRuleSet,
	"aws_ses_active_receipt_rule_set":                         sesReceiptRuleSet,
	"aws_ses_template":                                        sesTemplate,
	"aws_simpledb_domain":                                     sdbDomain,
	"aws_sns_platform_application":                            snsPlatformApplication,
	"aws_sns_topic_subscription":                              snsSubscription,
	"aws_sns_topic":                                           snsTopic,
	"aws_sqs_queue":                                           sqsQueue,
	"aws_sfn_activity":                                        stepFunctionsActivity,
	"aws_sfn_state_machine":                                   stepFunctionsStateMachine,
	"aws_shield_protection":                                   shieldProtection,
	"aws_ssm_activation":                                      ssmActivation,
	"aws_ssm_association":                                     ssmAssociation,
	"aws_ssm_document":                                        ssmDocument,
	"aws_ssm_maintenance_window":                              ssmMaintenanceWindow,
	"aws_ssm_maintenance_window_target":                       ssmMaintenanceWindowTarget,
	"aws_ssm_maintenance_window_task":                         ssmMaintenanceWindowTask,
	"aws_ssm_parameter":                                       ssmParameter,
	"aws_ssm_patch_baseline":                                  ssmPatchBaseline,
	"aws_ssm_patch_group":                                     ssmPatchGroup,
	"aws_ssm_resource_data_sync":                              ssmResourceDataSync,
	"aws_storagegateway_cached_iscsi_volume":                  storageGatewayCachedISCSIVolume,
	"aws_storagegateway_gateway":                              storageGatewayGateway,
	"aws_storagegateway_nfs_file_share":                       storageGatewayNFSFileShare,
	"aws_storagegateway_smb_file_share":                       storageGatewaySMBFileShare,
	"aws_swf_domain":                                          swfDomain,
	"aws_transfer_server":                                     transferServer,
	"aws_transfer_user":                                       transferUser,
	"aws_waf_byte_match_set":                                  wafByteMatchSet,
	"aws_waf_geo_match_set":                                   wafGeoMatchSet,
	"aws_waf_ipset":                                           wafIPSet,
	"aws_waf_rate_based_rule":                                 wafRateBasedRule,
	"aws_waf_regex_match_set":                                 wafRegexMatchSet,
	"aws_waf_regex_pattern_set":                               wafRegexPatternSet,
	"aws_waf_rule":                                            wafRule,
	"aws_waf_rule_group":                                      wafRuleGroup,
	"aws_waf_size_constraint_set":                             wafSizeConstraintSet,
	"aws_waf_sql_injection_match_set":                         wafSQLInjectionMatchSet,
	"aws_waf_web_acl":                                         wafWebACL,
	"aws_waf_xss_match_set":                                   wafXSSMatchSet,
	"aws_wafregional_byte_match_set":                          wafRegionalByteMatchSet,
	"aws_wafregional_geo_match_set":                           wafRegionalGeoMatchSet,
	"aws_wafregional_ipset":                                   wafRegionalIPSet,
	"aws_wafregional_rate_based_rule":                         wafRegionalRateBasedRule,
	"aws_wafregional_regex_match_set":                         wafregionalRegexMatchSet,
	"aws_wafregional_regex_pattern_set":                       wafRegionalRegexPatternSet,
	"aws_wafregional_rule":                                    wafRegionalRule,
	"aws_wafregional_rule_group":                              wafregionalRuleGroup,
	"aws_wafregional_size_constraint_set":                     wafRegionalSizeConstraintSet,
	"aws_wafregional_sql_injection_match_set":                 wafRegionalSQLInjectionMatchSet,
	"aws_wafregional_web_acl":                                 wafRegionalWebACL,
	"aws_wafregional_xss_match_set":                           wafRegionalXSSMatchSet,
	"aws_worklink_fleet":                                      workLinkFleet,
	"aws_workspaces_ip_group":                                 workspacesIPGroup,
	"aws_workspaces_directory":                                workSpacesDirectory,
	"aws_xray_sampling_rule":                                  xraySamplingRule,
}

func fromTerraformType(terraformType string) (resourceType, bool) {
	value, ok := terraformTypeMap[terraformType]
	return value, ok
}

var terraformPhysicalResourceIDs = map[resourceType]string{
	accessAnalyzerAnalyzer:                    "analyzer_name",
	acmpcaCertificateAuthority:                "arn",
	amazonMQBroker:                            "id",
	amazonMQConfiguration:                     "id",
	apiGatewayAPIKey:                          "id",
	apiGatewayClientCertificate:               "id",
	apiGatewayDomainName:                      "domain_name",
	apiGatewayRestAPI:                         "id",
	apiGatewayUsagePlan:                       "id",
	apiGatewayVpcLink:                         "id",
	apiGatewayV2Api:                           "id",
	applicationAutoScalingScheduledAction:     "name",
	appMeshMesh:                               "name",
	appSyncFunctions:                          "function_id",
	appSyncGraphQLApi:                         "id",
	athenaNamedQuery:                          "name",
	athenaWorkGroup:                           "name",
	autoScalingAutoScalingGroup:               "name",
	autoScalingLaunchConfiguration:            "name",
	autoScalingScalingPolicy:                  "name",
	autoScalingScheduledAction:                "scheduled_action_name",
	backupBackupPlan:                          "id",
	backupBackupSelection:                     "id",
	backupBackupVault:                         "id",
	batchComputeEnvironment:                   "compute_environment_name",
	batchJobDefinition:                        "arn",
	batchJobQueue:                             "arn",
	certificateManagerCertificate:             "arn",
	cloud9EnvironmentEC2:                      "id",
	cloudFrontCloudFrontOriginAccessIdentity:  "id",
	cloudFrontDistribution:                    "id",
	cloudFrontPublicKey:                       "id",
	cloudHSMV2Cluster:                         "cluster_id",
	cloudHSMV2HSM:                             "hsm_id",
	serviceDiscoveryHTTPNamespace:             "name",
	serviceDiscoveryPrivateDNSNamespace:       "name",
	serviceDiscoveryPublicDNSNamespace:        "name",
	serviceDiscoveryService:                   "id",
	cloudTrailTrail:                           "name",
	cloudWatchAlarm:                           "alarm_name",
	cloudWatchDashboard:                       "dashboard_name",
	logsDestination:                           "name",
	logsLogGroup:                              "name",
	logsMetricFilter:                          "name",
	logsSubscriptionFilter:                    "name",
	logsResourcePolicy:                        "policy_name",
	macieMemberAccountAssociation:             "member_account_id",
	macieS3BucketAssociation:                  "bucket_name",
	eventsRule:                                "name",
	eventsTarget:                              "target_id",
	codeBuildProject:                          "name",
	codeBuildSourceCredential:                 "arn",
	codeCommitRepository:                      "repository_name",
	codeCommitTrigger:                         "name",
	codeDeployApplication:                     "name",
	codeDeployDeploymentConfig:                "deployment_config_name",
	codeDeployDeploymentGroup:                 "app_name",
	codePipelinePipeline:                      "name",
	codePipelineWebhook:                       "id",
	cognitoIdentityPool:                       "identity_pool_name",
	cognitoUserPool:                           "id",
	cognitoUserPoolClient:                     "name",
	cognitoUserPoolGroup:                      "name",
	cognitoUserPoolIdentityProvider:           "provider_name",
	cognitoUserPoolResourceServer:             "name",
	configAggregationAuthorization:            "arn",
	configConfigRule:                          "name",
	configConfigurationAggregator:             "name",
	configConfigurationRecorder:               "name",
	configDeliveryChannel:                     "name",
	configOrganizationConfigRule:              "name",
	costAndUsageReportServiceReportDefinition: "report_name",
	dataPipelinePipeline:                      "id",
	dataSyncAgent:                             "name",
	dataSyncLocation:                          "arn",
	dataSyncTask:                              "arn",
	daxCluster:                                "cluster_name",
	daxParameterGroup:                         "name",
	daxSubnetGroup:                            "name",
	deviceFarmProject:                         "name",
	directConnectConnection:                   "id",
	directConnectGateway:                      "id",
	directConnectGatewayAssociation:           "id",
	directConnectGatewayAssociationProposal:   "id",
	directConnectLAG:                          "id",
	directConnectVirtualInterface:             "id",
	directoryServiceDirectory:                 "id",
	dlmLifecyclePolicy:                        "id",
	dmsCertificate:                            "certificate_id",
	dmsEndpoint:                               "endpoint_id",
	dmsReplicationInstance:                    "replication_instance_id",
	dmsReplicationSubnetGroup:                 "replication_subnet_group_id",
	dmsReplicationTask:                        "replication_task_id",
	docDBDBCluster:                            "cluster_identifier",
	docDBDBClusterParameterGroup:              "name",
	docDBDBInstance:                           "identifier",
	docDBDBSubnetGroup:                        "name",
	docDBDBClusterSnapshot:                    "db_cluster_snapshot_identifier",
	dynamoDBTable:                             "name",
	dynamoDBGlobalTable:                       "name",
	ec2CapacityReservation:                    "id",
	ec2ClientVpnEndpoint:                      "id",
	ec2CustomerGateway:                        "id",
	ec2DHCPOptions:                            "id",
	ec2EC2Fleet:                               "id",
	ec2EgressOnlyInternetGateway:              "id",
	ec2EIP:                                    "id",
	ec2EIPAssociation:                         "association_id",
	ec2FlowLog:                                "id",
	ec2Image:                                  "id",
	ec2Instance:                               "id",
	ec2InternetGateway:                        "id",
	ec2LaunchTemplate:                         "id",
	ec2NatGateway:                             "id",
	ec2NetworkACL:                             "id",
	ec2NetworkInterface:                       "id",
	ec2NetworkInterfaceAttachment:             "attachment_id",
	ec2PlacementGroup:                         "id",
	ec2RouteTable:                             "id",
	ec2SecurityGroup:                          "id",
	ec2Snapshot:                               "id",
	ec2SpotFleet:                              "id",
	ec2Subnet:                                 "id",
	ec2RouteTableSubnetAssociation:            "id",
	ec2TrafficMirrorFilter:                    "id",
	ec2TrafficMirrorFilterRule:                "id",
	ec2TrafficMirrorSession:                   "id",
	ec2TrafficMirrorTarget:                    "id",
	ec2TransitGateway:                         "id",
	ec2TransitGatewayAttachment:               "id",
	ec2TransitGatewayRouteTable:               "id",
	ec2Volume:                                 "id",
	ec2VPC:                                    "id",
	ec2VPCCidrBlock:                           "id",
	ec2VPCEndpoint:                            "id",
	ec2VPCEndpointConnectionNotification:      "id",
	ec2VPCEndpointService:                     "id",
	ec2VPCPeeringConnection:                   "id",
	ec2VPNConnection:                          "id",
	ec2VPNGateway:                             "id",
	ec2KeyPair:                                "key_pair_id",
	ec2SpotInstanceRequest:                    "id",
	ecrRepository:                             "name",
	ecsCluster:                                "arn",
	ecsService:                                "id",
	ecsTaskDefinition:                         "arn",
	ecsCapacityProvider:                       "arn",
	efsFileSystem:                             "id",
	efsMountTarget:                            "id",
	eksCluster:                                "name",
	eksNodegroup:                              "node_group_name",
	eksFargateProfile:                         "fargate_profile_name",
	elastiCacheCacheCluster:                   "cluster_id",
	elastiCacheParameterGroup:                 "name",
	elastiCacheReplicationGroup:               "replication_group_id",
	elastiCacheSecurityGroup:                  "name",
	elastiCacheSubnetGroup:                    "name",
	elasticsearchDomain:                       "domain_name",
	elasticTranscoderPipeline:                 "id",
	elasticTranscoderPreset:                   "id",
	elasticBeanstalkApplication:               "name",
	elasticBeanstalkApplicationVersion:        "arn",
	elasticBeanstalkConfigurationTemplate:     "name",
	elasticBeanstalkEnvironment:               "id",
	elasticLoadBalancingLoadBalancer:          "name",
	elasticLoadBalancingV2Listener:            "arn",
	elasticLoadBalancingV2ListenerRule:        "arn",
	elasticLoadBalancingV2LoadBalancer:        "arn",
	elasticLoadBalancingV2TargetGroup:         "arn",
	emrCluster:                                "name",
	emrSecurityConfiguration:                  "name",
	emrInstanceGroup:                          "name",
	fsxFileSystem:                             "id",
	gameLiftAlias:                             "id",
	gameLiftBuild:                             "id",
	gameLiftFleet:                             "id",
	gameLiftGameSessionQueue:                  "name",
	glueConnection:                            "catalog_id",
	glueCrawler:                               "name",
	glueDatabase:                              "name",
	glueJob:                                   "name",
	glueSecurityConfiguration:                 "name",
	glueTable:                                 "name",
	glueTrigger:                               "name",
	glueWorkflow:                              "name",
	glacierVault:                              "name",
	globalAcceleratorAccelerator:              "arn",
	globalAcceleratorEndpointGroup:            "arn",
	globalAcceleratorListener:                 "arn",
	guardDutyDetector:                         "id",
	iamAccessKey:                              "id",
	iamAccountAlias:                           "account_alias",
	iamGroup:                                  "name",
	iamGroupPolicy:                            "name",
	iamInstanceProfile:                        "name",
	iamOpenidConnectProvider:                  "arn",
	iamPolicy:                                 "name",
	iamRole:                                   "name",
	iamRolePolicy:                             "name",
	iamSamlProvider:                           "arn",
	iamServerCertificate:                      "name",
	iamServiceLinkedRole:                      "name",
	iamUser:                                   "name",
	iamUserPolicy:                             "names",
	iamUserSSHKey:                             "ssh_public_key_id",
	inspectorAssessmentTarget:                 "arn",
	inspectorAssessmentTemplate:               "arn",
	ioTCertificate:                            "id",
	ioTPolicy:                                 "name",
	iotRoleAlias:                              "alias",
	ioTThing:                                  "name",
	iotThingType:                              "name",
	ioTTopicRule:                              "name",
	kinesisStream:                             "arn",
	kinesisAnalyticsApplication:               "name",
	kinesisVideoStream:                        "name",
	kinesisFirehoseDeliveryStream:             "name",
	kmsAlias:                                  "name",
	kmsGrant:                                  "grant_id",
	kmsKey:                                    "key_id",
	lambdaAlias:                               "name",
	lambdaEventSourceMapping:                  "uuid",
	lambdaFunction:                            "function_name",
	lambdaLayerVersion:                        "arn",
	licenseManagerLicenseConfiguration:        "id",
	lightsailDomain:                           "domain_name",
	lightsailInstance:                         "name",
	lightsailKeyPair:                          "name",
	lightsailStaticIP:                         "name",
	mediaConvertQueue:                         "name",
	mediaLiveChannel:                          "channel_id",
	mediaStoreContainer:                       "name",
	mskCluster:                                "cluster_name",
	mskConfiguration:                          "name",
	neptuneDBCluster:                          "cluster_resource_id",
	neptuneDBClusterParameterGroup:            "name",
	neptuneDBClusterSnapshot:                  "db_cluster_snapshot_identifier",
	neptuneDBEventSubscription:                "name",
	neptuneDBInstance:                         "id",
	neptuneDBParameterGroup:                   "name",
	neptuneDBSubnetGroup:                      "name",
	opsWorksApp:                               "id",
	opsWorksInstance:                          "id",
	opsWorksLayer:                             "id",
	opsWorkRdsDbInstance:                      "rds_db_instance_arn",
	opsWorksStack:                             "id",
	opsWorksUserProfile:                       "user_arn",
	organizationsAccount:                      "id",
	organizationsOrganization:                 "id",
	organizationsOrganizationalUnit:           "id",
	organizationsPolicy:                       "id",
	pinpointApp:                               "application_id",
	qLDBLedger:                                "name",
	quickSightGroup:                           "group_name",
	quickSightUser:                            "user_name",
	rdsDBCluster:                              "cluster_identifier",
	rdsDBClusterEndpoint:                      "cluster_endpoint_identifier",
	rdsDBClusterParameterGroup:                "name",
	rdsDBClusterSnapshot:                      "db_cluster_snapshot_identifier",
	rdsDBInstance:                             "identifier",
	rdsDBParameterGroup:                       "name",
	rdsDBSecurityGroup:                        "name",
	rdsDBSnapshot:                             "db_snapshot_identifier",
	rdsDBSubnetGroup:                          "name",
	rdsEventSubscription:                      "name",
	rdsGlobalCluster:                          "global_cluster_identifier",
	rdsOptionGroup:                            "name",
	redshiftCluster:                           "cluster_identifier",
	redshiftClusterParameterGroup:             "name",
	redshiftClusterSecurityGroup:              "name",
	redshiftClusterSubnetGroup:                "name",
	redshiftEventSubscription:                 "name",
	redshiftSnapshotCopyGrant:                 "snapshot_copy_grant_name",
	redshiftSnapshotSchedule:                  "identifier",
	resourceGroupsGroup:                       "name",
	route53DelegationSet:                      "id",
	route53HealthCheck:                        "id",
	route53HostedZone:                         "zone_id",
	route53RecordSet:                          "name",
	route53QueryLog:                           "id",
	route53ResolverResolverEndpoint:           "id",
	route53ResolverResolverRule:               "id",
	route53ResolverResolverRuleAssociation:    "id",
	s3AccessPoint:                             "name",
	s3Bucket:                                  "bucket",
	sageMakerEndpoint:                         "name",
	sageMakerEndpointConfig:                   "name",
	sageMakerModel:                            "name",
	sageMakerNotebookInstance:                 "name",
	sageMakerNotebookInstanceLifecycleConfig:  "name",
	secretsManagerSecret:                      "name",
	secretsManagerSecretVersion:               "version_id",
	sesConfigurationSet:                       "name",
	sesDomainIdentity:                         "domain",
	sesEmailIdentity:                          "email",
	sesReceiptFilter:                          "name",
	sesReceiptRuleSet:                         "rule_set_name",
	sesTemplate:                               "name",
	sdbDomain:                                 "name",
	snsPlatformApplication:                    "arn",
	snsSubscription:                           "arn",
	snsTopic:                                  "arn",
	sqsQueue:                                  "id",
	stepFunctionsActivity:                     "name",
	stepFunctionsStateMachine:                 "name",
	shieldProtection:                          "id",
	ssmActivation:                             "id",
	ssmAssociation:                            "association_id",
	ssmDocument:                               "name",
	ssmMaintenanceWindow:                      "id",
	ssmMaintenanceWindowTarget:                "id",
	ssmMaintenanceWindowTask:                  "id",
	ssmParameter:                              "name",
	ssmPatchBaseline:                          "id",
	ssmPatchGroup:                             "patch_group",
	ssmResourceDataSync:                       "name",
	storageGatewayCachedISCSIVolume:           "volume_id",
	storageGatewayGateway:                     "gateway_id",
	storageGatewayNFSFileShare:                "fileshare_id",
	storageGatewaySMBFileShare:                "fileshare_id",
	swfDomain:                                 "name",
	transferServer:                            "id",
	transferUser:                              "arn",
	wafByteMatchSet:                           "id",
	wafGeoMatchSet:                            "id",
	wafIPSet:                                  "id",
	wafRateBasedRule:                          "id",
	wafRegexMatchSet:                          "id",
	wafRegexPatternSet:                        "id",
	wafRule:                                   "id",
	wafRuleGroup:                              "id",
	wafSizeConstraintSet:                      "id",
	wafSQLInjectionMatchSet:                   "id",
	wafWebACL:                                 "id",
	wafXSSMatchSet:                            "id",
	wafRegionalByteMatchSet:                   "id",
	wafRegionalGeoMatchSet:                    "id",
	wafRegionalIPSet:                          "id",
	wafRegionalRateBasedRule:                  "id",
	wafregionalRegexMatchSet:                  "id",
	wafRegionalRegexPatternSet:                "id",
	wafRegionalRule:                           "id",
	wafregionalRuleGroup:                      "id",
	wafRegionalSizeConstraintSet:              "id",
	wafRegionalSQLInjectionMatchSet:           "id",
	wafRegionalWebACL:                         "id",
	wafRegionalXSSMatchSet:                    "id",
	workLinkFleet:                             "display_name",
	workspacesIPGroup:                         "id",
	workSpacesDirectory:                       "id",
	xraySamplingRule:                          "rule_name",
}

func (rType resourceType) physicalResourceIDTerraform() (string, bool) {
	value, ok := terraformPhysicalResourceIDs[rType]
	return value, ok
}

var resourceBlacklistMap = map[string][]string{
	"terraform default resources": {
		"aws_default_network_acl",
		"aws_default_route_table",
		"aws_default_security_group",
		"aws_default_subnet",
		"aws_default_vpc",
		"aws_default_vpc_dhcp_options",
	},
	"terraform accepter resources": {
		"aws_dx_hosted_private_virtual_interface_accepter",
		"aws_dx_hosted_public_virtual_interface_accepter",
		"aws_dx_hosted_transit_virtual_interface_accepter",
		"aws_ec2_transit_gateway_vpc_attachment_accepter",
		"aws_guardduty_invite_accepter",
		"aws_ram_resource_share_accepter",
		"aws_vpc_peering_connection_accepter",
	},
	"no identifier found": {
		"AWS::ACMPCA::CertificateAuthorityActivation",
		"AWS::AmazonMQ::ConfigurationAssociation",
		"aws_api_gateway_account",
		"AWS::ApiGateway::Account",
		"aws_api_gateway_base_path_mapping",
		"AWS::ApiGateway::BasePathMapping",
		"aws_api_gateway_gateway_response",
		"AWS::ApiGateway::GatewayResponse",
		"aws_api_gateway_method",
		"AWS::ApiGateway::Method",
		"aws_api_gateway_authorizer",
		"AWS::ApiGateway::Authorizer",
		"aws_api_gateway_deployment",
		"AWS::ApiGateway::Deployment",
		"aws_api_gateway_documentation_part",
		"AWS::ApiGateway::DocumentationPart",
		"aws_api_gateway_documentation_version",
		"AWS::ApiGateway::DocumentationVersion",
		"aws_api_gateway_model",
		"AWS::ApiGateway::Model",
		"aws_api_gateway_request_validator",
		"AWS::ApiGateway::RequestValidator",
		"aws_api_gateway_resource",
		"AWS::ApiGateway::Resource",
		"aws_api_gateway_usage_plan_key",
		"AWS::ApiGateway::UsagePlanKey",
		"aws_api_gateway_stage",
		"AWS::ApiGateway::Stage",
		"AWS::ApiGatewayV2::ApiMapping",
		"AWS::ApiGatewayV2::Authorizer",
		"AWS::ApiGatewayV2::Deployment",
		"AWS::ApiGatewayV2::Integration",
		"AWS::ApiGatewayV2::IntegrationResponse",
		"AWS::ApiGatewayV2::Model",
		"AWS::ApiGatewayV2::Route",
		"AWS::ApiGatewayV2::RouteResponse",
		"AWS::ApiGatewayV2::Stage",
		"AWS::ApiGatewayV2::ApiGatewayManagedOverrides",
		"AWS::AppConfig::ConfigurationProfile",
		"AWS::AppConfig::Deployment",
		"AWS::AppConfig::Environment",
		"aws_appautoscaling_target",
		"AWS::ApplicationAutoScaling::ScalableTarget",
		"aws_appautoscaling_policy",
		"AWS::ApplicationAutoScaling::ScalingPolicy",
		"aws_appmesh_route",
		"AWS::AppMesh::Route",
		"aws_appmesh_virtual_node",
		"AWS::AppMesh::VirtualNode",
		"aws_appmesh_virtual_router",
		"AWS::AppMesh::VirtualRouter",
		"aws_appmesh_virtual_service",
		"AWS::AppMesh::VirtualService",
		"AWS::AppStream::StackFleetAssociation",
		"AWS::AppStream::StackUserAssociation",
		"AWS::AppStream::User",
		"AWS::AppSync::ApiCache",
		"aws_appsync_api_key",
		"AWS::AppSync::ApiKey",
		"aws_appsync_datasource",
		"AWS::AppSync::DataSource",
		"AWS::AppSync::FunctionConfiguration",
		"AWS::AppSync::GraphQLSchema",
		"aws_appsync_resolver",
		"AWS::AppSync::Resolver",
		"aws_autoscaling_lifecycle_hook",
		"AWS::AutoScaling::LifecycleHook",
		"aws_budgets_budget",
		"AWS::Budgets::Budget",
		"AWS::CodePipeline::CustomActionType",
		"AWS::CodeStar::GitHubRepository",
		"aws_codestarnotifications_notification_rule",
		"AWS::CodeStarNotifications::NotificationRule",
		"aws_cognito_identity_pool_roles_attachment",
		"AWS::Cognito::IdentityPoolRoleAttachment",
		"aws_cognito_user_pool_domain",
		"AWS::Cognito::UserPoolDomain",
		"AWS::Cognito::UserPoolRiskConfigurationAttachment",
		"AWS::Cognito::UserPoolUICustomizationAttachment",
		"AWS::Cognito::UserPoolUserToGroupAttachment",
		"AWS::CloudWatch::AnomalyDetector",
		"AWS::CloudWatch::CompositeAlarm",
		"AWS::Events::EventBusPolicy",
		"AWS::EC2::ClientVpnAuthorizationRule",
		"aws_vpn_connection_route",
		"AWS::EC2::ClientVpnRoute",
		"aws_ec2_client_vpn_network_association",
		"AWS::EC2::ClientVpnTargetNetworkAssociation",
		"AWS::EC2::GatewayRouteTableAssociation",
		"AWS::EC2::LocalGatewayRoute",
		"AWS::EC2::LocalGatewayRouteTableVPCAssociation",
		"aws_network_acl_rule",
		"AWS::EC2::NetworkAclEntry",
		"aws_security_group_rule",
		"AWS::EC2::SecurityGroupEgress",
		"aws_security_group_rule",
		"AWS::EC2::SecurityGroupIngress",
		"AWS::EC2::SubnetCidrBlock",
		"aws_ec2_transit_gateway_route_table_association",
		"AWS::EC2::TransitGatewayRouteTableAssociation",
		"aws_ec2_transit_gateway_route_table_propagation",
		"AWS::EC2::TransitGatewayRouteTablePropagation",
		"aws_volume_attachment",
		"AWS::EC2::VolumeAttachment",
		"aws_vpc_dhcp_options_association",
		"AWS::EC2::VPCDHCPOptionsAssociation",
		"AWS::EC2::VPCEndpointServicePermissions",
		"AWS::EC2::VPCGatewayAttachment",
		"AWS::EC2::VPNConnectionRoute",
		"AWS::EC2::VPNGatewayRoutePropagation",
		"AWS::ECS::PrimaryTaskSet",
		"AWS::ECS::TaskSet",
		"AWS::ElastiCache::SecurityGroupIngress",
		"AWS::EMR::InstanceFleetConfig",
		"AWS::EMR::InstanceGroupConfig",
		"AWS::EMR::Step",
		"aws_lb_listener_certificate",
		"AWS::ElasticLoadBalancingV2::ListenerCertificate",
		"aws_glue_classifier",
		"AWS::Glue::Classifier",
		"AWS::Glue::DataCatalogEncryptionSettings",
		"AWS::Glue::Partition",
		"AWS::GuardDuty::Filter",
		"aws_guardduty_ipset",
		"AWS::GuardDuty::IPSet",
		"AWS::GuardDuty::Master",
		"aws_guardduty_member",
		"AWS::GuardDuty::Member",
		"aws_guardduty_threatintelset",
		"AWS::GuardDuty::ThreatIntelSet",
		"aws_iam_user_group_membership",
		"AWS::IAM::UserToGroupAddition",
		"aws_inspector_resource_group",
		"AWS::Inspector::ResourceGroup",
		"aws_iot_policy_attachment",
		"AWS::IoT::PolicyPrincipalAttachment",
		"aws_iot_thing_principal_attachment",
		"AWS::IoT::ThingPrincipalAttachment",
		"AWS::IoTThingsGraph::FlowTemplate",
		"AWS::IoT1Click::Placement",
		"AWS::KinesisAnalytics::ApplicationOutput",
		"AWS::KinesisAnalytics::ApplicationReferenceDataSource",
		"AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption",
		"AWS::KinesisAnalyticsV2::ApplicationOutput",
		"AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource",
		"AWS::LakeFormation::DataLakeSettings",
		"AWS::LakeFormation::Permissions",
		"AWS::Lambda::EventInvokeConfig",
		"AWS::Lambda::EventSourceMapping",
		"AWS::Lambda::LayerVersionPermission",
		"AWS::Lambda::Version",
		"AWS::Lambda::Permission",
		"AWS::ManagedBlockchain::Member",
		"AWS::ManagedBlockchain::Node",
		"aws_opsworks_permission",
		"AWS::OpsWorks::ElasticLoadBalancerAttachment",
		"AWS::OpsWorksCM::Server",
		"aws_pinpoint_adm_channel",
		"AWS::Pinpoint::ADMChannel",
		"aws_pinpoint_apns_channel",
		"AWS::Pinpoint::APNSChannel",
		"aws_pinpoint_apns_sandbox_channel",
		"AWS::Pinpoint::APNSSandboxChannel",
		"aws_pinpoint_apns_voip_channel",
		"AWS::Pinpoint::APNSVoipChannel",
		"aws_pinpoint_apns_voip_sandbox_channel",
		"AWS::Pinpoint::APNSVoipSandboxChannel",
		"AWS::Pinpoint::ApplicationSettings",
		"aws_pinpoint_baidu_channel",
		"AWS::Pinpoint::BaiduChannel",
		"AWS::Pinpoint::Campaign",
		"aws_pinpoint_email_channel",
		"AWS::Pinpoint::EmailChannel",
		"aws_pinpoint_event_stream",
		"AWS::Pinpoint::EventStream",
		"aws_pinpoint_gcm_channel",
		"AWS::Pinpoint::GCMChannel",
		"AWS::Pinpoint::Segment",
		"aws_pinpoint_sms_channel",
		"AWS::Pinpoint::SMSChannel",
		"AWS::Pinpoint::VoiceChannel",
		"AWS::PinpointEmail::ConfigurationSet",
		"AWS::PinpointEmail::ConfigurationSetEventDestination",
		"AWS::PinpointEmail::DedicatedIpPool",
		"AWS::PinpointEmail::Identity",
		"aws_ram_resource_share",
		"aws_ram_principal_association",
		"aws_ram_resource_association",
		"AWS::RAM::ResourceShare",
		"AWS::RDS::DBSecurityGroupIngress",
		"AWS::Redshift::ClusterSecurityGroupIngress",
		"AWS::RoboMaker::RobotApplicationVersion",
		"AWS::RoboMaker::SimulationApplicationVersion",
		"AWS::Route53::RecordSetGroup",
		"AWS::EventSchemas::Schema",
		"AWS::FMS::NotificationChannel",
		"AWS::FMS::Policy",
		"aws_s3_bucket_policy",
		"AWS::S3::BucketPolicy",
		"AWS::SecretsManager::ResourcePolicy",
		"AWS::SecretsManager::RotationSchedule",
		"AWS::SecretsManager::SecretTargetAttachment",
		"AWS::SecurityHub::Hub",
		"AWS::ServiceDiscovery::Instance",
		"aws_ses_event_destination",
		"AWS::SES::ConfigurationSetEventDestination",
		"aws_sns_topic_policy",
		"AWS::SNS::TopicPolicy",
		"aws_sqs_queue_policy",
		"AWS::SQS::QueuePolicy",
		"aws_wafregional_web_acl_association",
		"AWS::WAFRegional::WebACLAssociation",
		"AWS::WAFv2::WebACLAssociation",
		"aws_acm_certificate_validation",
		"aws_api_gateway_integration",
		"aws_api_gateway_integration_response",
		"aws_api_gateway_method_response",
		"aws_api_gateway_method_settings",
		"aws_athena_database",
		"aws_autoscaling_attachment",
		"aws_autoscaling_notification",
		"aws_cloudwatch_event_permission",
		"aws_cloudwatch_log_destination_policy",
		"aws_codebuild_webhook",
		"aws_config_configuration_recorder_status",
		"aws_directory_service_conditional_forwarder",
		"aws_directory_service_log_subscription",
		"aws_dx_connection_association",
		"aws_ami_launch_permission",
		"aws_ebs_default_kms_key",
		"aws_ebs_encryption_by_default",
		"aws_ec2_transit_gateway_route",
		"aws_snapshot_create_volume_permission",
		"aws_spot_datafeed_subscription",
		"aws_ecr_lifecycle_policy",
		"aws_ecr_repository_policy",
		"aws_app_cookie_stickiness_policy",
		"aws_elb_attachment",
		"aws_lb_cookie_stickiness_policy",
		"aws_lb_ssl_negotiation_policy",
		"aws_load_balancer_backend_server_policy",
		"aws_load_balancer_listener_policy",
		"aws_load_balancer_policy",
		"aws_proxy_protocol_policy",
		"aws_lb_target_group_attachment",
		"aws_elasticsearch_domain_policy",
		"aws_fms_admin_account",
		"aws_glacier_vault_lock",
		"aws_iam_account_password_policy",
		"aws_iam_group_membership",
		"aws_iam_group_policy_attachment",
		"aws_iam_policy_attachment",
		"aws_iam_role_policy_attachment",
		"aws_iam_user_login_profile",
		"aws_iam_user_policy_attachment",
		"aws_kms_ciphertext",
		"aws_lambda_function_event_invoke_config^",
		"aws_lambda_permission",
		"aws_lambda_provisioned_concurrency_config",
		"aws_licensemanager_association",
		"aws_lightsail_static_ip_attachment",
		"aws_media_store_container_policy",
		"aws_db_instance_role_association",
		"aws_redshift_snapshot_schedule_association",
		"aws_route53_zone_association",
		"aws_s3_account_public_access_block",
		"aws_s3_bucket_analysis_configuration",
		"aws_s3_bucket_inventory",
		"aws_s3_bucket_metric",
		"aws_s3_bucket_notification",
		"aws_s3_bucket_public_access_block",
		"aws_organizations_policy_attachment",
		"aws_securityhub_account",
		"aws_securityhub_product_subscription",
		"aws_securityhub_standards_subscription",
		"aws_securityhub_member",
		"aws_ses_domain_identity_verification",
		"aws_ses_domain_dkim",
		"aws_ses_domain_mail_from",
		"aws_ses_receipt_rule",
		"aws_ses_identity_notification_topic",
		"aws_ses_identity_policy",
		"aws_servicequotas_service_quota",
		"aws_sns_sms_preferences",
		"aws_storagegateway_cache",
		"aws_storagegateway_upload_buffer",
		"aws_storagegateway_working_storage",
		"aws_transfer_ssh_key",
		"aws_main_route_table_association",
		"aws_route",
		"aws_network_interface_sg_attachment",
		"aws_vpc_endpoint_route_table_association",
		"aws_vpc_endpoint_service_allowed_principal",
		"aws_vpc_endpoint_subnet_association",
		"aws_vpc_peering_connection_options",
		"aws_vpn_gateway_attachment",
		"aws_vpn_gateway_route_propagation",
		"aws_worklink_website_certificate_authority_association",
		"aws_dx_bgp_peer",
		"AWS::NetworkManager::CustomerGatewayAssociation",
		"AWS::NetworkManager::LinkAssociation",
		"AWS::NetworkManager::TransitGatewayRegistration",
	},
	"not in the scope of this application": {
		"aws_dynamodb_table_item",
		"aws_s3_bucket_object",
		"aws_cloudwatch_log_stream",
		"aws_servicecatalog_portfolio",
		"aws_cloudformation_stack",
		"AWS::Logs::LogStream",
		"AWS::ServiceCatalog::AcceptedPortfolioShare",
		"AWS::ServiceCatalog::CloudFormationProduct",
		"AWS::ServiceCatalog::CloudFormationProvisionedProduct",
		"AWS::ServiceCatalog::LaunchNotificationConstraint",
		"AWS::ServiceCatalog::LaunchRoleConstraint",
		"AWS::ServiceCatalog::LaunchTemplateConstraint",
		"AWS::ServiceCatalog::Portfolio",
		"AWS::ServiceCatalog::PortfolioPrincipalAssociation",
		"AWS::ServiceCatalog::PortfolioProductAssociation",
		"AWS::ServiceCatalog::PortfolioShare",
		"AWS::ServiceCatalog::ResourceUpdateConstraint",
		"AWS::ServiceCatalog::StackSetConstraint",
		"AWS::ServiceCatalog::TagOption",
		"AWS::ServiceCatalog::TagOptionAssociation",
		"AWS::CloudFormation::CustomResource",
		"AWS::CloudFormation::Macro",
		"AWS::CloudFormation::Stack",
		"AWS::CloudFormation::WaitCondition",
		"AWS::CloudFormation::WaitConditionHandle",
		"aws_cloudformation_stack_set",
		"aws_cloudformation_stack_set_instance",
	},
	"Not in the v2 sdk yet": {
		"AWS::CodeStarConnections::Connection",
		"AWS::Chatbot::SlackChannelConfiguration",
		"AWS::ApiGatewayV2::VpcLink",
	},
	"No aws api": {
		"AWS::ACMPCA::Certificate",
		"AWS::Cassandra::Keyspace",
		"AWS::Cassandra::Table",
	},
}
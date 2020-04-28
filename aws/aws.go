package aws

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/urfave/cli/v2"
)

// Start is the entrypoint in the application
func Start(gitTag string, gitCommit string) {
	setGlobalLogLevel()

	logInfo("Startup")

	cliArgs := os.Args[1:]
	logDebug(cliArgs)

	// Initiate application state
	var state = state{
		Real: awsResourceMap{},
		Cfn:  extResourceMap{},
		Tf:   extResourceMap{},
	}

	// Read state from disk
	err := state.readFromDisk()
	if err != nil {
		logError("Could not read state from disk", err)
	}

	app := &cli.App{
		Name:                 "aws-resource-list",
		Usage:                "list all your aws resources",
		HelpName:             "aws-resource-list",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "refresh",
				Usage: "Refresh aws resources",
				Subcommands: []*cli.Command{
					{
						Name:  "real",
						Usage: "Refresh currently deployed aws resources",
						Action: func(c *cli.Context) error {
							config := awsStart()
							state.Real = getRealState(config)
							return nil
						},
					},
					{
						Name:  "cfn",
						Usage: "Refresh aws resources that are deployed through cloudformation",
						Action: func(c *cli.Context) error {
							config := awsStart()
							cfnState, err := getCloudForamtionState(config)
							if err != nil {
								logFatal("Could not fetch cloudformation resources:", err)
							}
							state.Cfn = cfnState
							return nil
						},
					},
					{
						Name:      "tf",
						Usage:     "Refresh aws resources that are deployed through terraform",
						ArgsUsage: "['terraform show -json' output file]",
						Action: func(c *cli.Context) error {
							tfjsonFile := c.Args().Get(0)
							err := getTerraformState(tfjsonFile, &state.Tf)
							if err != nil {
								logFatal("Could not fetch terraform resources:", err)
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "print",
				Usage: "Print aws resources",
				Subcommands: []*cli.Command{
					{
						Name:  "real",
						Usage: "Print currently deployed aws resources",
						Action: func(c *cli.Context) error {
							state.Real.print()
							return nil
						},
					},
					{
						Name:  "cfn",
						Usage: "Print aws resources that are deployed through cloudformation",
						Action: func(c *cli.Context) error {
							state.Cfn.print()
							return nil
						},
					},
					{
						Name:  "tf",
						Usage: "Print aws resources that are deployed through terraform",
						Action: func(c *cli.Context) error {
							state.Tf.print()
							return nil
						},
					},
				},
			},
			{
				Name:  "compare",
				Usage: "Print resources that exist in reality but not in IaC",
				Action: func(c *cli.Context) error {
					state.filter().print()
					return nil
				},
			},
			{
				Name:  "types",
				Usage: "Print resource types",
				Subcommands: []*cli.Command{
					{
						Name:  "cfn",
						Usage: "Print CloudFormation resource types that are implemented",
						Action: func(c *cli.Context) error {
							for key := range cloudformationTypeMap {
								fmt.Println(key)
							}
							return nil
						},
					},
					{
						Name:  "tf",
						Usage: "Print Terraform resource types that are implemented",
						Action: func(c *cli.Context) error {
							for key := range terraformTypeMap {
								fmt.Println(key)
							}
							return nil
						},
					},
					{
						Name:  "missing",
						Usage: "Print resources that are not implemented/blacklisted",
						Subcommands: []*cli.Command{
							{
								Name:  "cfn",
								Usage: "Print CloudFormation resources that are not implemented/blacklisted",
								Action: func(c *cli.Context) error {
									printMissingCloudformationResources()
									return nil
								},
							},
							{
								Name:  "tf",
								Usage: "Print Terraform resources that are not implemented/blacklisted",
								Action: func(c *cli.Context) error {
									printMissingTerraformResources()
									return nil
								},
							},
						},
					},
				},
			},
			{
				Name:  "version",
				Usage: "Show version information",
				Action: func(c *cli.Context) error {
					fmt.Println(gitTag, "@", gitCommit)
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		logFatal(err)
	}

	// Write state to disk
	err = state.writeToDisk()
	if err != nil {
		logFatal("Could not write state to disk", err)
	}
}

var accountID = ""

func getAccountID(config aws.Config) (err error) {
	svc := sts.New(config)

	resp, err := svc.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{}).Send(context.Background())

	if err != nil {
		return err
	}
	accountID = *resp.Account
	return nil
}

func getRealState(config aws.Config) (resources awsResourceMap) {
	serviceFunctions := []func(config aws.Config) (resources awsResourceMap){
		getAPIGateway,
		getAPIGatewayV2,
		getAccessAnalyzer,
		getAcm,
		getAcmpca,
		getAlexaForBusiness,
		getAmplify,
		getAppConfig,
		getAppMesh,
		getAppStream,
		getAppSync,
		getApplicationAutoScaling,
		getAthena,
		getAutoScaling,
		getAutoScalingPlans,
		getBackup,
		getBatch,
		getCloud9,
		getCloudHSMV2,
		getCloudTrail,
		getCloudWatch,
		getCloudWatchEvents,
		getCloudfront,
		getCloudwatchLogs,
		getCodeBuild,
		getCodeCommit,
		getCodeDeploy,
		getCodeGuruProfiler,
		getCodePipeline,
		getCognitoIdentity,
		getCognitoIdentityProvider,
		getConfigService,
		getCostAndUsageReportService,
		getCostExplorer,
		getDAX,
		getDLM,
		getDataPipeline,
		getDataSync,
		getDatabaseMigrationService,
		getDetective,
		getDeviceFarm,
		getDirectConnect,
		getDirectoryService,
		getDocDB,
		getDynamoDB,
		getEc2,
		getEcr,
		getEcs,
		getEfs,
		getEks,
		getElasticBeanstalk,
		getElasticLoadBalancing,
		getElasticLoadBalancingV2,
		getElasticSearchService,
		getElasticTranscoder,
		getElasticache,
		getEmr,
		getFirehose,
		getFsx,
		getGameLift,
		getGlacier,
		getGlobalAccelerator,
		getGlue,
		getGreengrass,
		getGroundStation,
		getGuardDuty,
		getIam,
		getImageBuilder,
		getInspector,
		getIoT,
		getIoT1ClickDevicesService,
		getIoT1ClickProjects,
		getIoTAnalytics,
		getIoTEvents,
		getKafka,
		getKinesis,
		getKinesisAnalytics,
		getKinesisAnalyticsV2,
		getKinesisVideo,
		getKms,
		getLakeFormation,
		getLambda,
		getLicenseManager,
		getLightsail,
		getMQ,
		getMacie,
		getMediaConvert,
		getMediaLive,
		getMediaStore,
		getNeptune,
		getNetworkManager,
		getOpsWorks,
		getOrganizations,
		getPinpoint,
		getQLDB,
		getQuickSight,
		getRds,
		getRedshift,
		getResourceGroups,
		getRoboMaker,
		getRoute53,
		getRoute53Resolver,
		getS3,
		getS3Control,
		getSWF,
		getSageMaker,
		getSchemas,
		getSecretsManager,
		getServiceDiscovery,
		getSes,
		getSfn,
		getShield,
		getSimpleDB,
		getSns,
		getSns,
		getSqs,
		getSsm,
		getStorageGateway,
		getTransfer,
		getWaf,
		getWafRegional,
		getWafv2,
		getWorkLink,
		getWorkSpaces,
		getXray,
	}

	resourceMapAllServices := []awsResourceMap{}
	resourceMapChannel := make(chan awsResourceMap)
	var resourceMapWaitGroup sync.WaitGroup
	parallelServices := 10

	servicesCount := len(serviceFunctions) - 1

	count := 1
	for index, serviceFunction := range serviceFunctions {
		resourceMapWaitGroup.Add(1)
		go func(serviceFunction func(config aws.Config) (resources awsResourceMap)) {
			functionName := runtime.FuncForPC(reflect.ValueOf(serviceFunction).Pointer()).Name()
			functionName = functionName[(strings.LastIndex(functionName, ".") + 1):]

			logDebug("Starting", functionName)
			resourceMapChannel <- serviceFunction(config)
			logDebug("Finished", functionName)

			resourceMapWaitGroup.Done()
		}(serviceFunction)

		if count == parallelServices || index == servicesCount {
			// Append the awsResourceMaps to the slice until all are listed
			first := true
			for rMap := range resourceMapChannel {
				resourceMapAllServices = append(resourceMapAllServices, rMap)
				if first {
					// Wait for all functions to finish
					go func() {
						resourceMapWaitGroup.Wait()
						logDebug("Got all services for this batch. Closing channel...")
						close(resourceMapChannel)
					}()
					first = false
				}
			}
			resourceMapChannel = make(chan awsResourceMap)
			count = 1
		} else {
			count++
		}
	}

	resources = reduce(resourceMapAllServices...)
	return
}

func awsStart() (config aws.Config) {
	config, err := external.LoadDefaultAWSConfig()
	if err != nil {
		logFatal("Could not load SDK config", err)
	}
	logInfo("Created SDK config")

	err = getAccountID(config)
	if err != nil {
		logFatal("Could not get caller identity", err)
	}
	return
}

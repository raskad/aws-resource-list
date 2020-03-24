package aws

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/urfave/cli/v2"
)

var accountID = ""

// Start is the entrypoint in the application
func Start(gitTag string, gitCommit string) {
	setGlobalLogLevel()

	logInfo("Startup")

	cliArgs := os.Args[1:]
	logDebug(cliArgs)

	// Initiate application state
	var state = state{
		real: resourceMap{},
		cfn:  resourceMap{},
		tf:   resourceMap{},
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
							state[real] = getRealState(config)
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
							state[cfn] = cfnState
							return nil
						},
					},
					{
						Name:      "tf",
						Usage:     "Refresh aws resources that are deployed through terraform",
						ArgsUsage: "['terraform show -json' output file]",
						Action: func(c *cli.Context) error {
							tfjsonFile := c.Args().Get(0)
							cfnState, err := getTerraformState(tfjsonFile)
							if err != nil {
								logFatal("Could not fetch terraform resources:", err)
							}
							state[tf] = cfnState
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
							state[real].print()
							return nil
						},
					},
					{
						Name:  "cfn",
						Usage: "Print aws resources that are deployed through cloudformation",
						Action: func(c *cli.Context) error {
							state[cfn].print()
							return nil
						},
					},
					{
						Name:  "tf",
						Usage: "Print aws resources that are deployed through terraform",
						Action: func(c *cli.Context) error {
							state[tf].print()
							return nil
						},
					},
				},
			},
			{
				Name:  "compare",
				Usage: "Print resources that exist in reality but not in IaC",
				Action: func(c *cli.Context) error {
					state.filter(real, cfn).print()
					state.filter(real, tf).print()
					return nil
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

func getAccountID(config aws.Config) (err error) {
	svc := sts.New(config)

	resp, err := svc.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{}).Send(context.Background())

	if err != nil {
		return err
	}
	accountID = *resp.Account
	return nil
}

func run(f func(config aws.Config) (resources resourceMap), config aws.Config, c chan resourceMap, wg *sync.WaitGroup) {
	wg.Add(1)
	c <- f(config)
	wg.Done()
}

func getRealState(config aws.Config) (resources resourceMap) {
	r := []resourceMap{}
	c := make(chan resourceMap)
	var wg sync.WaitGroup

	go run(getAccessAnalyzer, config, c, &wg)
	go run(getAcm, config, c, &wg)
	go run(getAcmpca, config, c, &wg)
	go run(getAlexaForBusiness, config, c, &wg)
	go run(getAmplify, config, c, &wg)
	go run(getAPIGateway, config, c, &wg)
	go run(getAPIGatewayV2, config, c, &wg)
	go run(getAppConfig, config, c, &wg)
	go run(getAppMesh, config, c, &wg)
	go run(getAppStream, config, c, &wg)
	go run(getAppSync, config, c, &wg)
	go run(getAthena, config, c, &wg)
	go run(getAutoScaling, config, c, &wg)
	go run(getAutoScalingPlans, config, c, &wg)
	go run(getBackup, config, c, &wg)
	go run(getBatch, config, c, &wg)
	go run(getCloud9, config, c, &wg)
	go run(getCloudfront, config, c, &wg)
	go run(getCloudTrail, config, c, &wg)
	go run(getCloudWatch, config, c, &wg)
	go run(getCloudWatchEvents, config, c, &wg)
	go run(getCodeBuild, config, c, &wg)
	go run(getCodeCommit, config, c, &wg)
	go run(getCodeDeploy, config, c, &wg)
	go run(getCodePipeline, config, c, &wg)
	go run(getCognitoIdentity, config, c, &wg)
	go run(getCognitoIdentityProvider, config, c, &wg)
	go run(getConfig, config, c, &wg)
	go run(getDataPipeline, config, c, &wg)
	go run(getDAX, config, c, &wg)
	go run(getDirectoryService, config, c, &wg)
	go run(getDLM, config, c, &wg)
	go run(getDms, config, c, &wg)
	go run(getDocDB, config, c, &wg)
	go run(getDynamoDB, config, c, &wg)
	go run(getEc2, config, c, &wg)
	go run(getEcr, config, c, &wg)
	go run(getEcs, config, c, &wg)
	go run(getEfs, config, c, &wg)
	go run(getEks, config, c, &wg)
	go run(getElasticache, config, c, &wg)
	go run(getElasticsearch, config, c, &wg)
	go run(getElasticBeanstalk, config, c, &wg)
	go run(getElasticLoadBalancing, config, c, &wg)
	go run(getElasticLoadBalancingV2, config, c, &wg)
	go run(getEmr, config, c, &wg)
	go run(getFirehose, config, c, &wg)
	go run(getFsx, config, c, &wg)
	go run(getGameLift, config, c, &wg)
	go run(getGlue, config, c, &wg)
	go run(getGreengrass, config, c, &wg)
	go run(getGroundStation, config, c, &wg)
	go run(getGuardDuty, config, c, &wg)
	go run(getIam, config, c, &wg)
	go run(getInspector, config, c, &wg)
	go run(getIoT, config, c, &wg)
	go run(getIoT1ClickDevicesService, config, c, &wg)
	go run(getIoT1ClickProjects, config, c, &wg)
	go run(getIoTAnalytics, config, c, &wg)
	go run(getIoTEvents, config, c, &wg)
	go run(getKinesis, config, c, &wg)
	go run(getKinesisAnalytics, config, c, &wg)
	go run(getKinesisAnalyticsV2, config, c, &wg)
	go run(getKms, config, c, &wg)
	go run(getLakeFormation, config, c, &wg)
	go run(getLambda, config, c, &wg)
	go run(getCloudwatchLogs, config, c, &wg)
	go run(getMq, config, c, &wg)
	go run(getMediaConvert, config, c, &wg)
	go run(getMediaLive, config, c, &wg)
	go run(getMediaStore, config, c, &wg)
	go run(getMsk, config, c, &wg)
	go run(getNeptune, config, c, &wg)
	go run(getOpsWorks, config, c, &wg)
	go run(getPinpoint, config, c, &wg)
	go run(getQLDB, config, c, &wg)
	go run(getRds, config, c, &wg)
	go run(getRedshift, config, c, &wg)
	go run(getRoboMaker, config, c, &wg)
	go run(getRoute53, config, c, &wg)
	go run(getRoute53Resolver, config, c, &wg)
	go run(getS3, config, c, &wg)
	go run(getS3Control, config, c, &wg)
	go run(getSageMaker, config, c, &wg)
	go run(getSchemas, config, c, &wg)
	go run(getSdb, config, c, &wg)
	go run(getSecretsManager, config, c, &wg)
	go run(getServiceDiscovery, config, c, &wg)
	go run(getSes, config, c, &wg)
	go run(getSfn, config, c, &wg)
	go run(getSns, config, c, &wg)
	go run(getSns, config, c, &wg)
	go run(getSqs, config, c, &wg)
	go run(getSsm, config, c, &wg)
	go run(getTransfer, config, c, &wg)
	go run(getWaf, config, c, &wg)
	go run(getWafRegional, config, c, &wg)
	go run(getWafv2, config, c, &wg)
	go run(getWorkSpaces, config, c, &wg)

	// Append the resourceMaps to the slice until all are listed
	first := true
	for rMap := range c {
		r = append(r, rMap)
		if first {
			// Wait for all functions to finish
			go func() {
				wg.Wait()
				close(c)
			}()
			first = false
		}
	}

	resources = reduce(r...)
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

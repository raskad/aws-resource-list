package aws

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
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
	}

	// Read state from disk
	err := state.readFromDisk()
	if err != nil {
		logError("Could not read state from disk", err)
	}

	if len(cliArgs) == 0 {
		cliArgs = append(cliArgs, "help")
	}

	switch cliArgs[0] {
	case "refresh":
		switch cliArgs[1] {
		case "cfn":
			session := awsStart()
			cfnState, err := getCloudForamtionState(session)
			if err != nil {
				logFatal("Could not fetch cloudformation resources:", err)
			}
			state[cfn] = cfnState
		case "real":
			session := awsStart()
			state[real] = getRealState(session)
		default:
			logFatal("Cli argument", cliArgs[1], "invalid")
		}
	case "print":
		switch cliArgs[1] {
		case "real":
			state[real].print()
		case "cfn":
			state[cfn].print()
		case "tf":
			state[tf].print()
		default:
			logFatal("Cli argument", cliArgs[1], "invalid")
		}
	case "compare":
		state.filter(cfn, real).print()
	case "version":
		fmt.Println(gitTag, "@", gitCommit)
	default:
		fmt.Println("Refresh aws resources: aws-resource-list refresh [cfn/real]")
		fmt.Println("Print resources: aws-resource-list print [cfn/real]")
		fmt.Println("Print resources that exist in reality but not in IaC : aws-resource-list compare")
	}

	// Write state to disk
	err = state.writeToDisk()
	if err != nil {
		logFatal("Could not write state to disk", err)
	}
}

func getAccountID(session *session.Session) (err error) {
	client := sts.New(session)
	c, err := client.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return err
	}
	accountID = *c.Account
	return nil
}

func run(f func(session *session.Session) (resources resourceMap), session *session.Session, c chan resourceMap, wg *sync.WaitGroup) {
	wg.Add(1)
	c <- f(session)
	wg.Done()
}

func getRealState(session *session.Session) (resources resourceMap) {
	r := []resourceMap{}
	c := make(chan resourceMap)
	var wg sync.WaitGroup

	go run(getAccessAnalyzer, session, c, &wg)
	go run(getAcm, session, c, &wg)
	go run(getAcmpca, session, c, &wg)
	go run(getAlexaForBusiness, session, c, &wg)
	go run(getAmplify, session, c, &wg)
	go run(getAPIGateway, session, c, &wg)
	go run(getAPIGatewayV2, session, c, &wg)
	go run(getAppConfig, session, c, &wg)
	go run(getAppMesh, session, c, &wg)
	go run(getAppStream, session, c, &wg)
	go run(getAppSync, session, c, &wg)
	go run(getAthena, session, c, &wg)
	go run(getAutoScaling, session, c, &wg)
	go run(getAutoScalingPlans, session, c, &wg)
	go run(getBackup, session, c, &wg)
	go run(getBatch, session, c, &wg)
	go run(getCloud9, session, c, &wg)
	go run(getCloudfront, session, c, &wg)
	go run(getCloudTrail, session, c, &wg)
	go run(getCloudWatch, session, c, &wg)
	go run(getCloudWatchEvents, session, c, &wg)
	go run(getCodeBuild, session, c, &wg)
	go run(getCodeCommit, session, c, &wg)
	go run(getCodeDeploy, session, c, &wg)
	go run(getCodePipeline, session, c, &wg)
	go run(getCognitoIdentity, session, c, &wg)
	go run(getCognitoIdentityProvider, session, c, &wg)
	go run(getConfig, session, c, &wg)
	go run(getDataPipeline, session, c, &wg)
	go run(getDAX, session, c, &wg)
	go run(getDirectoryService, session, c, &wg)
	go run(getDLM, session, c, &wg)
	go run(getDms, session, c, &wg)
	go run(getDocDB, session, c, &wg)
	go run(getDynamoDB, session, c, &wg)
	go run(getEc2, session, c, &wg)
	go run(getEcr, session, c, &wg)
	go run(getEcs, session, c, &wg)
	go run(getEfs, session, c, &wg)
	go run(getEks, session, c, &wg)
	go run(getElasticache, session, c, &wg)
	go run(getElasticsearch, session, c, &wg)
	go run(getElasticBeanstalk, session, c, &wg)
	go run(getElasticLoadBalancingV2, session, c, &wg)
	go run(getElb, session, c, &wg)
	go run(getEmr, session, c, &wg)
	go run(getFirehose, session, c, &wg)
	go run(getFsx, session, c, &wg)
	go run(getGameLift, session, c, &wg)
	go run(getGlue, session, c, &wg)
	go run(getGuardDuty, session, c, &wg)
	go run(getIam, session, c, &wg)
	go run(getInspector, session, c, &wg)
	go run(getIoT, session, c, &wg)
	go run(getIoT1ClickDevicesService, session, c, &wg)
	go run(getIoT1ClickProjects, session, c, &wg)
	go run(getIoTAnalytics, session, c, &wg)
	go run(getIoTEvents, session, c, &wg)
	go run(getKinesis, session, c, &wg)
	go run(getKinesisAnalytics, session, c, &wg)
	go run(getKinesisAnalyticsV2, session, c, &wg)
	go run(getKms, session, c, &wg)
	go run(getLakeFormation, session, c, &wg)
	go run(getLambda, session, c, &wg)
	go run(getCloudwatchLogs, session, c, &wg)
	go run(getMq, session, c, &wg)
	go run(getMediaConvert, session, c, &wg)
	go run(getMediaLive, session, c, &wg)
	go run(getMediaStore, session, c, &wg)
	go run(getMsk, session, c, &wg)
	go run(getNeptune, session, c, &wg)
	go run(getOpsWorks, session, c, &wg)
	go run(getPinpoint, session, c, &wg)
	go run(getQLDB, session, c, &wg)
	go run(getRds, session, c, &wg)
	go run(getRedshift, session, c, &wg)
	go run(getRoboMaker, session, c, &wg)
	go run(getRoute53, session, c, &wg)
	go run(getRoute53Resolver, session, c, &wg)
	go run(getS3, session, c, &wg)
	go run(getS3Control, session, c, &wg)
	go run(getSageMaker, session, c, &wg)
	go run(getSchemas, session, c, &wg)
	go run(getSdb, session, c, &wg)
	go run(getSecretsManager, session, c, &wg)
	go run(getServiceDiscovery, session, c, &wg)
	go run(getSes, session, c, &wg)
	go run(getSns, session, c, &wg)
	go run(getSns, session, c, &wg)
	go run(getSqs, session, c, &wg)
	go run(getSsm, session, c, &wg)
	go run(getTransfer, session, c, &wg)
	go run(getWaf, session, c, &wg)
	go run(getWafRegional, session, c, &wg)
	go run(getWafv2, session, c, &wg)
	go run(getWorkSpaces, session, c, &wg)

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

func awsStart() (sess *session.Session) {
	sess, err := session.NewSession()
	if err != nil {
		logFatal("Could not create aws session", err)
	}
	logInfo("Created aws session")

	err = getAccountID(sess)
	if err != nil {
		logFatal("Could not get caller identity", err)
	}
	return
}

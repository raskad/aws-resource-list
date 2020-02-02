package aws

import (
	"fmt"
	"os"

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

func getRealState(session *session.Session) (resources resourceMap) {
	resources = reduce(
		getAccessAnalyzer(session),
		getAcm(session),
		getAutoScaling(session),
		getAutoScalingPlans(session),
		getBackup(session),
		getCloudfront(session),
		getCloudTrail(session),
		getCloudWatch(session),
		getCloudWatchEvents(session),
		getConfig(session),
		getDynamoDB(session),
		getEc2(session),
		getEcr(session),
		getEcs(session),
		getEfs(session),
		getEks(session),
		getElasticache(session),
		getElasticsearch(session),
		getElasticLoadBalancingV2(session),
		getFirehose(session),
		getFsx(session),
		getGlue(session),
		getGuardDuty(session),
		getIam(session),
		getKinesis(session),
		getKinesisAnalytics(session),
		getKinesisAnalyticsV2(session),
		getKms(session),
		getLambda(session),
		getCloudwatchLogs(session),
		getMsk(session),
		getRds(session),
		getRedshift(session),
		getRoute53(session),
		getRoute53Resolver(session),
		getS3(session),
		getS3Control(session),
		getSecretsManager(session),
		getSes(session),
		getSns(session),
		getSqs(session),
		getSsm(session),
		getWaf(session),
		getWafRegional(session),
		getWafv2(session),
	)
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

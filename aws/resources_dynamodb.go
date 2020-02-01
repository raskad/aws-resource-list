package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getDynamoDB(session *session.Session) (resources resourceMap) {
	client := dynamodb.New(session)
	resourcesSliceErrorMap := resourceSliceErrorMap{
		dynamoDBTable: getDynamoDBTable(client),
	}
	resources = resourcesSliceErrorMap.unwrap()
	return
}

func getDynamoDBTable(client *dynamodb.DynamoDB) (r resourceSliceError) {
	r.err = client.ListTablesPages(&dynamodb.ListTablesInput{}, func(page *dynamodb.ListTablesOutput, lastPage bool) bool {
		logDebug("List DynamoDBTable resources page")
		for _, resource := range page.TableNames {
			logDebug("Got DynamoDBTable resource with PhysicalResourceId", *resource)
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

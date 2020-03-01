package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getDynamoDB(session *session.Session) (resources resourceMap) {
	client := dynamodb.New(session)
	resources = reduce(
		getDynamoDBTable(client).unwrap(dynamoDBTable),
	)
	return
}

func getDynamoDBTable(client *dynamodb.DynamoDB) (r resourceSliceError) {
	r.err = client.ListTablesPages(&dynamodb.ListTablesInput{}, func(page *dynamodb.ListTablesOutput, lastPage bool) bool {
		for _, resource := range page.TableNames {
			r.resources = append(r.resources, *resource)
		}
		return true
	})
	return
}

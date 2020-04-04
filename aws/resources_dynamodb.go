package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func getDynamoDB(config aws.Config) (resources awsResourceMap) {
	client := dynamodb.New(config)

	dynamoDBTableNames := getDynamoDBTableNames(client)
	dynamoDBGlobalTableNames := getDynamoDBGlobalTableNames(client)

	resources = awsResourceMap{
		dynamoDBTable:       dynamoDBTableNames,
		dynamoDBGlobalTable: dynamoDBGlobalTableNames,
	}
	return
}

func getDynamoDBTableNames(client *dynamodb.Client) (resources []string) {
	req := client.ListTablesRequest(&dynamodb.ListTablesInput{})
	p := dynamodb.NewListTablesPaginator(req)
	for p.Next(context.Background()) {
		if p.Err() != nil {
			logErr(p.Err())
			return
		}
		page := p.CurrentPage()
		resources = append(resources, page.TableNames...)
	}
	return
}

func getDynamoDBGlobalTableNames(client *dynamodb.Client) (resources []string) {
	input := dynamodb.ListGlobalTablesInput{}
	for {
		page, err := client.ListGlobalTablesRequest(&input).Send(context.Background())
		if err != nil {
			logErr(err)
			return
		}
		for _, resource := range page.GlobalTables {
			resources = append(resources, *resource.GlobalTableName)
		}
		if page.LastEvaluatedGlobalTableName == nil {
			return
		}
		input.ExclusiveStartGlobalTableName = page.LastEvaluatedGlobalTableName
	}
}

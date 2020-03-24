package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func getDynamoDB(config aws.Config) (resources resourceMap) {
	client := dynamodb.New(config)

	dynamoDBTableNames := getDynamoDBTableNames(client)

	resources = resourceMap{
		dynamoDBTable: dynamoDBTableNames,
	}
	return
}

func getDynamoDBTableNames(client *dynamodb.Client) (resources []string) {
	req := client.ListTablesRequest(&dynamodb.ListTablesInput{})
	p := dynamodb.NewListTablesPaginator(req)
	for p.Next(context.Background()) {
		logErr(p.Err())
		page := p.CurrentPage()
		for _, resource := range page.TableNames {
			resources = append(resources, resource)
		}
	}
	return
}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func getDynamoDB(config aws.Config) (resources resourceMap) {
	client := dynamodb.New(config)
	resources = reduce(
		getDynamoDBTable(client).unwrap(dynamoDBTable),
	)
	return
}

func getDynamoDBTable(client *dynamodb.Client) (r resourceSliceError) {
	req := client.ListTablesRequest(&dynamodb.ListTablesInput{})
	p := dynamodb.NewListTablesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, resource := range page.TableNames {
			r.resources = append(r.resources, resource)
		}
	}
	r.err = p.Err()
	return
}

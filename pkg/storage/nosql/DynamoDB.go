package nosql

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	Client *dynamodb.Client
}

func (d *DynamoDBClient) ListTables(ctx context.Context) (*dynamodb.ListTablesOutput, error) {
	return d.Client.ListTables(ctx, &dynamodb.ListTablesInput{})
}

func (d *DynamoDBClient) CreateTable(ctx context.Context, input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	return d.Client.CreateTable(ctx, input)
}

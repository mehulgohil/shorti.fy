package config

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mehulgohil/shorti.fy/redirect/infrastructures"
	"sync"
)

var (
	dynamoDBObj  *DBClientHandler
	dynamoDBOnce sync.Once
)

type IDynamoDB interface {
	InitLocalDBConnection()
	InitTables()
}

type DBClientHandler struct {
	DBClient *infrastructures.DynamoDBClient
}

// InitLocalDBConnection initialize dynamodb connection
func (d *DBClientHandler) InitLocalDBConnection() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: EnvVariables.DynamoDBURL}, nil
			})),
	)

	if err != nil {
		panic(fmt.Sprintf("unable to connect DB, %v", err))
	}

	// Using the Config value, create the DynamoDB client
	d.DBClient = &infrastructures.DynamoDBClient{
		Client: dynamodb.NewFromConfig(cfg),
	}
}

func (d *DBClientHandler) InitTables() {
	// making sure the URL table exists
	// if not, we create a new table
	if d.createTableIfNotExist("URL") {
		ZapLogger.Info("Successfully initialized new URL table")
	} else {
		ZapLogger.Info("URL table already exist")
	}
}

func DynamoDB() IDynamoDB {
	if dynamoDBObj == nil {
		dynamoDBOnce.Do(func() {
			dynamoDBObj = &DBClientHandler{}
		})
	}
	return dynamoDBObj
}

func (d *DBClientHandler) createTableIfNotExist(tableName string) bool {
	if d.tableExists(tableName) {
		return false
	}
	_, err := d.DBClient.Client.CreateTable(context.TODO(), d.buildCreateTableInput(tableName))
	if err != nil {
		panic(fmt.Sprintf("create table failed, %v", err))
	}
	return true
}

func (d *DBClientHandler) tableExists(name string) bool {
	tables, err := d.DBClient.Client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		panic(fmt.Sprintf("unable to list tables in DB, %v", err))
	}
	for _, n := range tables.TableNames {
		if n == name {
			return true
		}
	}
	return false
}

func (d *DBClientHandler) buildCreateTableInput(tableName string) *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("HashKey"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("HashKey"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	}
}

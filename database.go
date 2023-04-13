package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mehulgohil/shorti.fy/pkg/storage/nosql"
	"sync"
)

var (
	dynamoDBObj  *db
	dynamoDBOnce sync.Once
)

type IDynamoDB interface {
	InitDBConnection() *nosql.DynamoDBClient
}

type db struct{}

// InitDBConnection initialize dynamodb connection
func (d *db) InitDBConnection() *nosql.DynamoDBClient {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
	)

	if err != nil {
		panic(fmt.Sprintf("unable to connect DB, %v", err))
	}

	// Using the Config value, create the DynamoDB client
	return &nosql.DynamoDBClient{
		Client: dynamodb.NewFromConfig(cfg),
	}
}

func DynamoDB() IDynamoDB {
	if dynamoDBObj == nil {
		dynamoDBOnce.Do(func() {
			dynamoDBObj = &db{}
		})
	}
	return dynamoDBObj
}

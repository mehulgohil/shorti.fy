package interfaces

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type IDataAccessLayer interface {
	ListTables(ctx context.Context) (*dynamodb.ListTablesOutput, error)
}

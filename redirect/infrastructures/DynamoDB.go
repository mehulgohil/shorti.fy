package infrastructures

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mehulgohil/shorti.fy/redirect/models"
)

type DynamoDBClient struct {
	Client *dynamodb.Client
}

func (d *DynamoDBClient) SaveItem(input *dynamodb.UpdateItemInput) error {
	_, err := d.Client.UpdateItem(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDBClient) GetItem(hashKey string) (models.URLTable, error) {
	tableItem := models.URLTable{}
	item, err := d.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String("URL"),
		Key: map[string]types.AttributeValue{
			"HashKey": &types.AttributeValueMemberS{Value: hashKey},
		},
	})
	if err != nil {
		return tableItem, err
	}

	err = attributevalue.UnmarshalMap(item.Item, &tableItem)
	if err != nil {
		return tableItem, err
	}

	return tableItem, nil
}

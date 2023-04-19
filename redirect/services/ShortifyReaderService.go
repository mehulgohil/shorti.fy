package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kataras/iris/v12/x/errors"
	"github.com/mehulgohil/shorti.fy/redirect/interfaces"
	"github.com/mehulgohil/shorti.fy/redirect/models"
)

type ShortifyReaderService struct {
	interfaces.IDataAccessLayer
}

// Reader get long url from db
func (s *ShortifyReaderService) Reader(shortURLHash string) (string, error) {

	// get the long url from db
	item, err := s.getItemFromDB(shortURLHash)
	if err != nil {
		return "", err
	}

	if item.LongURL == "" {
		return "", errors.New("long url not found, please check the short url and try again")
	}

	// increment the hitcount and update the item
	go func() {
		item.HitCount = item.HitCount + 1
		_ = s.upsertItemToTable(item)
	}()

	return item.LongURL, nil
}

func (s *ShortifyReaderService) getItemFromDB(hashKey string) (models.URLTable, error) {
	tableItem := models.URLTable{}
	item, err := s.GetItem(context.Background(), &dynamodb.GetItemInput{
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

func (s *ShortifyReaderService) upsertItemToTable(item models.URLTable) error {
	marshedData, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = s.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("URL"),
		Item:      marshedData,
	})
	if err != nil {
		return err
	}

	return nil
}

package services

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kataras/iris/v12/x/errors"
	"github.com/mehulgohil/shorti.fy/interfaces"
	"github.com/mehulgohil/shorti.fy/models"
	"math/rand"
	"strconv"
	"time"
)

type ShortifyService struct {
	interfaces.IEncodingAlgorithm
	interfaces.IHashingAlgorithm
	interfaces.IDataAccessLayer
}

// Reader get long url from db
func (s *ShortifyService) Reader(shortURLHash string) (string, error) {

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

// Writer shortens the long url and returns a short url
func (s *ShortifyService) Writer(longURL string, userEmail string) (string, error) {
	encodedString, err := s.getUniqueHash(longURL + userEmail)
	if err != nil {
		return "", err
	}

	// creating item struct
	err = s.upsertItemToTable(models.URLTable{
		HashKey:        encodedString,
		LongURL:        longURL,
		CreatedAt:      time.Now(),
		ExpirationDate: time.Now().AddDate(1, 0, 0), //setting expiration date as 1 year from now
		HitCount:       0,                           // initialize hitcount to 0
		CreatedBy:      userEmail,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://localhost:8080/%s", encodedString), nil
}

func (s *ShortifyService) getUniqueHash(str string) (string, error) {
	for true {
		encodedString := s.Encode(s.Hash(str))
		if len(encodedString) > 7 {
			// selecting only the first 7 characters
			encodedString = encodedString[:7]
		}

		db, err := s.getItemFromDB(encodedString)
		if err != nil {
			return "", err
		}
		if db.HashKey == "" {
			return encodedString, nil
		}
		// append a random integer at end to avoid collision
		str = str + strconv.Itoa(rand.Int())
	}
	return "", nil
}

func (s *ShortifyService) upsertItemToTable(item models.URLTable) error {
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

// ================= Utility Functions
func (s *ShortifyService) getItemFromDB(hashKey string) (models.URLTable, error) {
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

package interfaces

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mehulgohil/shorti.fy/redirect/models"
)

type IDataAccessLayer interface {
	GetItem(hashKey string) (models.URLTable, error)
	SaveItem(input *dynamodb.UpdateItemInput) error
}

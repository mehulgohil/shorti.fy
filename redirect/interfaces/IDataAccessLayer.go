package interfaces

import (
	"github.com/mehulgohil/shorti.fy/redirect/models"
)

type IDataAccessLayer interface {
	GetItem(hashKey string) (models.URLTable, error)
	SaveItem(item models.URLTable) error
}

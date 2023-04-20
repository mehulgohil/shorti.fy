package services

import (
	"github.com/kataras/iris/v12/x/errors"
	"github.com/mehulgohil/shorti.fy/redirect/interfaces"
)

type ShortifyReaderService struct {
	interfaces.IDataAccessLayer
	interfaces.IRedisLayer
}

// Reader get long url from db
func (s *ShortifyReaderService) Reader(shortURLHash string) (string, error) {

	// get the long url from db
	item, err := s.GetItem(shortURLHash)
	if err != nil {
		return "", err
	}

	if item.LongURL == "" {
		return "", errors.New("long url not found, please check the short url and try again")
	}

	// increment the hitcount and update the item
	go func() {
		item.HitCount = item.HitCount + 1
		_ = s.SaveItem(item)
	}()

	return item.LongURL, nil
}

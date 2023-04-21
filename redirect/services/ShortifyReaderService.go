package services

import (
	"github.com/kataras/iris/v12/x/errors"
	"github.com/mehulgohil/shorti.fy/redirect/interfaces"
	"time"
)

const (
	twoMonthDuration        = time.Hour * 1440
	longUrlNotFoundErrorMsg = "long url not found, please check the short url and try again"
)

type ShortifyReaderService struct {
	interfaces.IDataAccessLayer
	interfaces.IRedisLayer
}

// Reader get long url from db
func (s *ShortifyReaderService) Reader(shortURLHash string) (string, error) {
	// increment the hitcount and update the item
	defer func() {
		go func() {
			_ = s.incrementHitCount(shortURLHash)
		}()
	}()

	// checking and getting value from redis
	cacheValue, err := s.GetKeyValue(shortURLHash)
	if err == nil {
		return cacheValue, nil
	}

	// get the long url from db
	item, err := s.GetItem(shortURLHash)
	if err != nil {
		return "", err
	}

	if item.LongURL == "" {
		return "", errors.New(longUrlNotFoundErrorMsg)
	}

	// caching data into redis with expiration of 2 months
	go func() {
		_ = s.SetKeyValue(shortURLHash, item.LongURL, twoMonthDuration)
	}()

	return item.LongURL, nil
}

func (s *ShortifyReaderService) incrementHitCount(shortURLHash string) error {
	item, err := s.GetItem(shortURLHash)
	if err != nil {
		return err
	}
	item.HitCount = item.HitCount + 1

	err = s.SaveItem(item)
	if err != nil {
		return err
	}

	return nil
}

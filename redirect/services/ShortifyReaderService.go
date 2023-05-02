package services

import (
	"fmt"
	"github.com/kataras/iris/v12/x/errors"
	"github.com/mehulgohil/shorti.fy/redirect/interfaces"
	"github.com/mehulgohil/shorti.fy/redirect/models"
	"go.uber.org/zap"
	"time"
)

const (
	twoMonthDuration        = time.Hour * 1440
	longUrlNotFoundErrorMsg = "long url not found, please check the short url and try again"
)

type ShortifyReaderService struct {
	interfaces.IDataAccessLayer
	interfaces.IRedisLayer
	Logger *zap.Logger
}

// Reader get long url from db
func (s *ShortifyReaderService) Reader(shortURLHash string) (string, error) {
	var successfullyRedirected = false

	// increment the hitcount and update the item
	defer func() {
		go func() {
			if successfullyRedirected {
				err := s.incrementHitCount(shortURLHash)
				if err != nil {
					s.Logger.Error(fmt.Sprintf("error incrementing hitcount - %s", err.Error()))
				}
			}
		}()
	}()

	// checking and getting value from redis
	cacheValue, err := s.GetKeyValue(shortURLHash)
	if err == nil {
		successfullyRedirected = true
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
	go s.cacheData(shortURLHash, item)

	successfullyRedirected = true
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

func (s *ShortifyReaderService) cacheData(shortURLHash string, item models.URLTable) {
	err := s.SetKeyValue(shortURLHash, item.LongURL, twoMonthDuration)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("error caching value to redis - %s", err.Error()))
	}
}

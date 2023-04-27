package services

import (
	"fmt"
	"github.com/mehulgohil/shorti.fy/writer/config"
	"github.com/mehulgohil/shorti.fy/writer/interfaces"
	"github.com/mehulgohil/shorti.fy/writer/models"
	"math/rand"
	"strconv"
	"time"
)

type ShortifyWriterService struct {
	interfaces.IEncodingAlgorithm
	interfaces.IHashingAlgorithm
	interfaces.IDataAccessLayer
	EnvVariables config.EnvConfig
}

// Writer shortens the long url and returns a short url
func (s *ShortifyWriterService) Writer(longURL string, userEmail string) (string, error) {
	encodedString, err := s.getUniqueHash(longURL + userEmail)
	if err != nil {
		return "", err
	}

	// creating item struct
	err = s.SaveItem(models.URLTable{
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

	return fmt.Sprintf("%s/v1/%s", s.EnvVariables.APPDomain, encodedString), nil
}

func (s *ShortifyWriterService) getUniqueHash(str string) (string, error) {
	for true {
		encodedString := s.Encode(s.Hash(str))
		if len(encodedString) > 7 {
			// selecting only the first 7 characters
			encodedString = encodedString[:7]
		}

		db, err := s.GetItem(encodedString)
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

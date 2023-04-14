package services

import (
	"fmt"
	"github.com/mehulgohil/shorti.fy/interfaces"
)

type ShortifyService struct {
	interfaces.IEncodingAlgorithm
	interfaces.IHashingAlgorithm
	interfaces.IDataAccessLayer
}

// temporary creating map to store the hashValue
// will remove once we have DB layer implemented
var hashMap = map[string]string{}

// Reader get long url from db
func (s *ShortifyService) Reader(shortURL string) (string, error) {
	// TODO: to fetch the original url from DB

	return hashMap[shortURL], nil
}

// Writer shortens the long url and returns a short url
func (s *ShortifyService) Writer(longURL string, userEmail string) (string, error) {
	encodedString := s.Encode(s.Hash(longURL + userEmail))

	if len(encodedString) > 7 {
		// selecting only the first 7 characters
		encodedString = encodedString[:7]
	}

	// TODO: to save the original longURL to DB
	hashMap[encodedString] = longURL

	return fmt.Sprintf("http://localhost:8080/%s", encodedString), nil
}

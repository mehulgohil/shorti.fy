package services

import (
	"fmt"
	"github.com/mehulgohil/shorti.fy/interfaces"
)

type ShortifyService struct {
	interfaces.IEncodingAlgorithm
	interfaces.IHashingAlgorithm
}

// temporary creating map to store the hashValue
// will remove once we have DB layer implemented
var hashMap = map[string]string{}

// Reader get long url from db
func (s *ShortifyService) Reader(url string) (string, error) {
	// TODO: to fetch the original url from DB

	return hashMap[url], nil
}

// Writer shortens the long url and returns a short url
func (s *ShortifyService) Writer(url string, userEmail string) (string, error) {
	encodedString := s.Encode(s.Hash(url + userEmail))

	if len(encodedString) > 7 {
		// selecting only the first 7 characters
		encodedString = encodedString[:7]
	}

	// TODO: to save the original url to DB
	hashMap[encodedString] = url

	return fmt.Sprintf("http://localhost:8080/%s", encodedString), nil
}

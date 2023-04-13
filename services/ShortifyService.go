package services

import (
	"fmt"
	"github.com/mehulgohil/shorti.fy/interfaces"
)

type ShortifyService struct {
	interfaces.EncodingAlgorithm
}

// temporary creating map to store the hashValue
// will remove once we have DB layer implemented
var hashMap = map[string]string{}

func (s *ShortifyService) Reader(url string) (string, error) {
	// TODO: to fetch the original url from DB

	return hashMap[url], nil
}

func (s *ShortifyService) Writer(url string) (string, error) {
	encodedString := s.Encode(url)
	if len(encodedString) > 7 {
		encodedString = encodedString[:7]
	}

	// TODO: to save the original url to DB
	hashMap[encodedString] = url

	return fmt.Sprintf("http://localhost:8080/%s", encodedString), nil
}

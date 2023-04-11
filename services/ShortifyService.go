package services

type ShortifyService struct{}

func (s *ShortifyService) Reader(url string) (string, error) {
	// we fetch the original url from DB

	return "https://originalurl", nil
}

func (s *ShortifyService) Writer(url string) (string, error) {
	// shorten the url and save it in DB
	// return the new short url

	return "https://newshorturl", nil
}

package interfaces

type IShortifyService interface {
	Reader(shortURL string) (string, error)
	Writer(longURL string, userEmail string) (string, error)
}

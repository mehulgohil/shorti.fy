package interfaces

type IShortifyService interface {
	Reader(shortURLHash string) (string, error)
	Writer(longURL string, userEmail string) (string, error)
}

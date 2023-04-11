package interfaces

type IShortifyService interface {
	Reader(url string) (string, error)
	Writer(url string) (string, error)
}

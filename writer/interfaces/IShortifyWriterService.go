package interfaces

type IShortifyWriterService interface {
	Writer(longURL string, userEmail string) (string, error)
}

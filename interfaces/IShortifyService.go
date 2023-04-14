package interfaces

type IShortifyReaderService interface {
	Reader(shortURLHash string) (string, error)
}

type IShortifyWriterService interface {
	Writer(longURL string, userEmail string) (string, error)
}

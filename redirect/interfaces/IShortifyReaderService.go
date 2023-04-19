package interfaces

type IShortifyReaderService interface {
	Reader(shortURLHash string) (string, error)
}

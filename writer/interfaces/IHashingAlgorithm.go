package interfaces

type IHashingAlgorithm interface {
	Hash(input string) string
}

package interfaces

type EncodingAlgorithm interface {
	Encode(input string) string
}

package algorithm

import "github.com/jxskiss/base62"

type Base62Algorithm struct{}

func (b *Base62Algorithm) Encode(input string) string {
	return string(base62.Encode([]byte(input)))
}

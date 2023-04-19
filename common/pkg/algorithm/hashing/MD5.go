package hashing

import "crypto/md5"

type MD5Hash struct{}

func (h *MD5Hash) Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return string(hash[:])
}

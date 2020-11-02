package utils

import "crypto/sha1"

func SHA1(byteSlice []byte) []byte {
	crypt := sha1.New()
	crypt.Write(byteSlice)
	return crypt.Sum(nil)
}

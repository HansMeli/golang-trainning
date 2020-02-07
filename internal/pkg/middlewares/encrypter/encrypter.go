package middlewares

import (
	"crypto/md5"
	"fmt"
)

func DigestString(t string, publicKey string, privateKey string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(t+privateKey+publicKey)))
}
